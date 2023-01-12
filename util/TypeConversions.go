package util

import (
	"encoding/json"
	"log"
)

func FromDocument[T any](doc map[string]interface{}) (*T, error) {
	var docType *T
	marshal, err := json.Marshal(doc)
	if err != nil {
		log.Fatalf("FromDocument: %v", err)
	}

	err = json.Unmarshal(marshal, &docType)
	if err != nil {
		log.Fatalf("FromDocument: %v", err)
	}

	return docType, nil
}

func ToDocument(data interface{}) (map[string]interface{}, error) {
	var doc map[string]interface{}
	marshal, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("ToDocument: %v", err)
	}

	err = json.Unmarshal(marshal, &doc)
	if err != nil {
		log.Fatalf("ToDocument: %v", err.Error())
	}

	parseDocumentWithoutValues(doc)

	return doc, nil
}

func parseDocumentWithoutValues(doc map[string]interface{}) interface{} {
	for key, value := range doc {
		switch value.(type) {
		case map[string]interface{}:
			mapValue := value.(map[string]interface{})
			if len(mapValue) == 1 {
				doc[key] = parseFirebaseValue(mapValue)
			} else {
				doc[key] = parseDocumentWithoutValues(mapValue)
			}
		default:
			continue
		}
	}
	return doc
}

func parseFirebaseValue(doc map[string]interface{}) interface{} {
	var mValue interface{}
	for _, value := range doc {
		mValue = value
		break
	}
	return mValue
}
