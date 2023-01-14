package functions

import (
	"context"
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"log"
	"strings"
)

func OnUserCreated(ctx context.Context, e data.FirestoreEvent[model.User]) error {
	collection, doc, _ := e.Path()

	log.Printf(collection + "/" + doc)

	curValue := e.Value.Data.ID
	newValue := strings.ToUpper(curValue)
	if curValue == newValue {
		log.Printf("%q is already upper case: skipping", curValue)
		return nil
	}
	log.Printf("Replacing value: %q -> %q", curValue, newValue)

	//data := map[string]string{"original": newValue}
	//_, err := client.Collection(collection).Doc(doc).Set(ctx, data)
	//if err != nil {
	//	return fmt.Errorf("Set: %v", err)
	//}
	return nil
}
