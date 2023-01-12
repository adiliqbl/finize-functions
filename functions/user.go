package functions

import (
	"context"
	"finize-functions/data"
	"finize-functions/data/model"
	"log"
	"strings"
)

func OnUserCreated(ctx context.Context, e data.FirestoreEvent[model.User]) error {
	fullPath := strings.Split(e.Value.Name, "/documents/")[1]
	pathParts := strings.Split(fullPath, "/")
	collection := pathParts[0]
	doc := strings.Join(pathParts[1:], "/")

	log.Printf(collection + "/" + doc)

	curValue := e.Value.Data.ID.StringValue
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
