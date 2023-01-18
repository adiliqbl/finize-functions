package service

import (
	"cloud.google.com/go/firestore"
	"context"
	"finize-functions.app/config"
	"finize-functions.app/util"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

var firestoreDatabase *firestore.Client

type Firestore[T any] struct {
	client *firestore.Client
	ctx    context.Context
}

type FirestoreService[T any] interface {
	Find(path string, tx *firestore.Transaction) (*T, error)
	Create(collection string, doc map[string]interface{}) (string, error)
	Update(path string, doc map[string]interface{}) (bool, error)
	Delete(path string) (bool, error)

	Doc(path string) *firestore.DocumentRef
	Collection(collection string) *firestore.CollectionRef
	Batch() *firestore.BulkWriter
	Transaction(run func(tx *firestore.Transaction) error) error
}

func InitFirestore(ctx context.Context) error {
	// Use the application default credentials.
	conf := &firebase.Config{ProjectID: config.ProjectIdD}

	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
		return err
	}

	firestoreDatabase = client
	return nil
}

func newFirestoreService[T any](ctx context.Context) FirestoreService[T] {
	return &Firestore[T]{client: firestoreDatabase, ctx: ctx}
}

func (store *Firestore[T]) Doc(path string) *firestore.DocumentRef {
	return store.client.Doc(path)
}

func (store *Firestore[T]) Collection(path string) *firestore.CollectionRef {
	return store.client.Collection(path)
}

func (store *Firestore[T]) Batch() *firestore.BulkWriter {
	return store.client.BulkWriter(store.ctx)
}

func (store *Firestore[T]) Transaction(run func(tx *firestore.Transaction) error) error {
	return store.client.RunTransaction(store.ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		return run(tx)
	})
}

func (store *Firestore[T]) Find(path string, tx *firestore.Transaction) (*T, error) {
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

func (store *Firestore[T]) Create(collection string, doc map[string]interface{}) (string, error) {
	ref := store.client.Collection(collection).NewDoc()
	doc["id"] = ref.ID

	_, err := ref.Set(store.ctx, doc)
	if err != nil {
		log.Fatalf("Firestore.Create: %v", err)
		return "", err
	}

	return ref.ID, nil
}

func (store *Firestore[T]) Update(path string, doc map[string]interface{}) (bool, error) {
	_, err := store.client.Doc(path).Set(store.ctx, doc, firestore.MergeAll)
	if err != nil {
		log.Fatalf("Firestore.Create: %v", err)
		return false, err
	}

	return true, nil
}

func (store *Firestore[T]) Delete(path string) (bool, error) {
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
