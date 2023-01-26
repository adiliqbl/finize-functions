package service

import (
	"cloud.google.com/go/firestore"
	"context"
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/util"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
)

var firestoreDatabase *firestore.Client

type firestoreDB[T any] struct {
	client  *firestore.Client
	ctx     context.Context
	eventID string
}

type FirestoreDB interface {
	Doc(path string) *firestore.DocumentRef
	Collection(collection string) *firestore.CollectionRef
	Batch(run func() []data.DatabaseOperation) error
	Transaction(run func(tx *firestore.Transaction) []data.DatabaseOperation) error
}

type FirestoreService[T any] interface {
	FirestoreDB
	GetAll(collection string) ([]T, error)
	Paginate(query firestore.Query, start uint, limit uint) ([]T, error)
	Find(path string, tx *firestore.Transaction) (*T, error)
	Create(collection string, id *string, doc map[string]interface{}) (*string, error)
	Update(path string, doc map[string]interface{}) (bool, error)
	Delete(path string) (bool, error)
}

func InitFirestore(ctx context.Context) error {
	if firestoreDatabase != nil {
		return nil
	}

	// Use the application default credentials.
	conf := &firebase.Config{ProjectID: os.Getenv("GOOGLE_CLOUD_PROJECT")}

	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("app.firestoreDB: %v", err)
		return err
	}

	firestoreDatabase = client
	return nil
}

func newFirestoreService[T any](ctx context.Context, eventID string) FirestoreService[T] {
	return NewFirestoreService[T](ctx, firestoreDatabase, eventID)
}

func NewFirestoreService[T any](ctx context.Context, db *firestore.Client, eventID string) FirestoreService[T] {
	return &firestoreDB[T]{client: db, ctx: ctx, eventID: eventID}
}

func newFirestoreDB(ctx context.Context, eventID string) FirestoreDB {
	return NewFirestoreDB(ctx, firestoreDatabase, eventID)
}

func NewFirestoreDB(ctx context.Context, db *firestore.Client, eventID string) FirestoreDB {
	return &firestoreDB[map[string]interface{}]{client: db, ctx: ctx, eventID: eventID}
}

func (store *firestoreDB[T]) Doc(path string) *firestore.DocumentRef {
	return store.client.Doc(path)
}

func (store *firestoreDB[T]) Collection(path string) *firestore.CollectionRef {
	return store.client.Collection(path)
}

func (store *firestoreDB[T]) Batch(run func() []data.DatabaseOperation) error {
	ops := run()

	if len(ops) == 0 {
		return nil
	}

	batch := store.client.BulkWriter(store.ctx)

	if store.eventID != "" {
		events := NewEventService(NewFirestoreService[model.Event](store.ctx, store.client, store.eventID), store.eventID)
		if err := events.SetProcessedBatch(batch); err != nil {
			return err
		}
	}

	return data.CommitBatch(batch, ops)
}

func (store *firestoreDB[T]) Transaction(run func(tx *firestore.Transaction) []data.DatabaseOperation) error {
	return store.client.RunTransaction(store.ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		ops := run(tx)

		if len(ops) == 0 {
			return nil
		}

		if store.eventID != "" {
			events := NewEventService(NewFirestoreService[model.Event](store.ctx, store.client, store.eventID), store.eventID)
			if err := events.SetProcessed(tx); err != nil {
				return err
			}
		}

		return data.CommitTransaction(tx, ops)
	})
}

func (store *firestoreDB[T]) runQuery(query firestore.Query) ([]T, error) {
	iterator := query.Documents(store.ctx)
	snapshots, err := iterator.GetAll()

	if err != nil {
		log.Printf("Failed to get all documents: %v", err)
	}

	var docs []T
	for _, snap := range snapshots {
		doc, err := util.MapTo[T](snap.Data())

		if err != nil {
			log.Printf("Failed to parse task: %v", err)
			return nil, err
		}

		docs = append(docs, doc)
	}

	return docs, nil
}

func (store *firestoreDB[T]) Paginate(query firestore.Query, start uint, limit uint) ([]T, error) {
	return store.runQuery(query.Offset(int(start)).Limit(int(limit)))
}

func (store *firestoreDB[T]) GetAll(collection string) ([]T, error) {
	return store.runQuery(store.client.Collection(collection).Query)
}

func (store *firestoreDB[T]) Find(path string, tx *firestore.Transaction) (*T, error) {
	var snap *firestore.DocumentSnapshot
	var err error
	if tx == nil {
		snap, err = store.client.Doc(path).Get(store.ctx)
	} else {
		snap, err = tx.Get(store.client.Doc(path))
	}

	if err != nil {
		// Translate firestorm not found to application specific not found.
		if status.Code(err) == codes.NotFound {
			err = status.Error(codes.NotFound, "Doc by "+path+" not found")
		}
		return nil, err
	}

	document, err := util.MapTo[T](snap.Data())
	return &document, err
}

func (store *firestoreDB[T]) Create(collection string, id *string, doc map[string]interface{}) (*string, error) {
	var ref *firestore.DocumentRef
	if util.NullOrEmpty(id) {
		ref = store.client.Collection(collection).NewDoc()
		doc["id"] = ref.ID
	} else {
		ref = store.client.Doc(collection + "/" + *id)
		doc["id"] = ref.ID
	}

	_, err := ref.Set(store.ctx, doc)
	if err != nil {
		log.Fatalf("firestoreDB.Create: %v", err)
		return nil, err
	}

	return util.Pointer(ref.ID), nil
}

func (store *firestoreDB[T]) Update(path string, doc map[string]interface{}) (bool, error) {
	_, err := store.client.Doc(path).Set(store.ctx, doc, firestore.MergeAll)
	if err != nil {
		log.Fatalf("firestoreDB.Create: %v", err)
		return false, err
	}

	return true, nil
}

func (store *firestoreDB[T]) Delete(path string) (bool, error) {
	_, err := store.client.Doc(path).Delete(store.ctx)
	if err != nil {
		// Translate firestorm not found to application specific not found.
		if status.Code(err) == codes.NotFound {
			err = status.Error(codes.NotFound, "Doc by "+path+" not found")
		}
		return false, err
	}

	return true, nil
}
