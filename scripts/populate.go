package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var store *firestore.Client

func setupFirestore() *firestore.Client {
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("%v", err)
	}

	return client
}

func populateTransactionCategories(reset bool) {
	jsonFile, err := os.Open("transaction_categories.json")
	if err != nil {
		log.Fatalf("%v", err)
	}
	byteValue, _ := io.ReadAll(jsonFile)
	var data map[string]interface{}
	json.Unmarshal(byteValue, &data)

	categories := data["categories"].([]interface{})

	// Remove all categories
	if reset {
		refs := store.Collection("transaction-categories").DocumentRefs(context.Background())
		docs, _ := refs.GetAll()
		for _, doc := range docs {
			doc.Delete(context.Background())
		}
	}

	// Add categories from json
	for _, category := range categories {
		category := category.(map[string]interface{})
		name := category["name"].(string)
		path := fmt.Sprintf("transaction-categories/%s", name)
		_, err = store.Doc(path).Set(context.Background(), category)

		if err != nil {
			log.Fatalf("%v", err)
		}
	}

	defer jsonFile.Close()
}

func main() {
	store = setupFirestore()
	defer store.Close()

	populateTransactionCategories(false)
}
