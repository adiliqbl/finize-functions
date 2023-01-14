package util

import (
	"encoding/json"
	"log"
)

func MapTo[T any](doc interface{}) (T, error) {
	var docType T
	marshal, err := json.Marshal(doc)
	if err != nil {
		log.Fatalf("MapTo: %v", err)
	}

	err = json.Unmarshal(marshal, &docType)
	if err != nil {
		log.Fatalf("MapTo: %v", err)
	}

	return docType, nil
}
