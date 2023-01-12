package service

import (
	"cloud.google.com/go/firestore"
	"context"
	"finize-functions/config"
	"finize-functions/util"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Firestore[T any] struct {
	client *firestore.Client
	ctx    context.Context
}

func NewFirestore[T any]() (*Firestore[T], error) {
	store := new(Firestore[T])

	// Use the application default credentials.
	conf := &firebase.Config{ProjectID: config.ProjectIdD}

	// Use context.Background() because the app/client should persist across
	// invocations.
	ctx := context.Background()
	store.ctx = ctx

	app, err := firebase.NewApp(store.ctx, conf)
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}

	store.client, err = app.Firestore(store.ctx)
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}

	return store, nil
}

func (store *Firestore[T]) FindByID(path string) (*T, error) {
	snap, err := store.client.Doc(path).Get(store.ctx)
	if err != nil {
		// Translate firestorm not found to application specific not found.
		if status.Code(err) == codes.NotFound {
			err = status.Error(codes.NotFound, "Doc by "+path+" not found")
		}
		return nil, err
	}

	return util.FromDocument[T](snap.Data())
}

func (store *Firestore[T]) Create(collection string, data T) (interface{}, error) {
	doc, err := util.ToDocument(data)
	if err != nil {
		return nil, err
	}

	ref := store.client.Collection(collection).NewDoc()
	doc["id"] = ref.ID

	_, err = ref.Set(store.ctx, doc)
	if err != nil {
		log.Fatalf("Firestore.Create: %v", err)
		return nil, err
	}

	return ref.ID, nil
}

func (store *Firestore[T]) Update(path string, data T) (bool, error) {
	doc, err := util.ToDocument(data)
	if err != nil {
		return false, err
	}

	_, err = store.client.Doc(path).Set(store.ctx, doc, firestore.MergeAll)
	if err != nil {
		log.Fatalf("Firestore.Create: %v", err)
		return false, err
	}

	return true, nil
}

func (store *Firestore[T]) DeleteByID(path string) (bool, error) {
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
