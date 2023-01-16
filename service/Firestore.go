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
	Find(path string) (*T, error)
	Create(collection string, doc map[string]interface{}) (string, error)
	Update(path string, doc map[string]interface{}) (bool, error)
	Delete(path string) (bool, error)
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

func newFirestoreService[T any]() FirestoreService[T] {
	return &Firestore[T]{client: firestoreDatabase, ctx: context.Background()}
}

func (store *Firestore[T]) Find(path string) (*T, error) {
	snap, err := store.client.Doc(path).Get(store.ctx)
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
