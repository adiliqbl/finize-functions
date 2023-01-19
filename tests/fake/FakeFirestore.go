package fake

import (
	"cloud.google.com/go/firestore"
	"context"
	"finize-functions.app/service"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

var testFirestoreDatabase *firestore.Client

func InitTestFirestore() {
	if testFirestoreDatabase != nil {
		return
	}

	os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8080")

	// Use the application default credentials.
	conf := &firebase.Config{ProjectID: "demo-project"}

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, conf,
		option.WithEndpoint("http://localhost:8080"),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
		os.Exit(0)
		return
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("app.firestoreDB: %v", err)
		os.Exit(0)
		return
	}

	testFirestoreDatabase = client
	return
}

func NewFirestoreService[T any](ctx context.Context) service.FirestoreService[T] {
	return service.NewFirestoreService[T](ctx, testFirestoreDatabase)
}
