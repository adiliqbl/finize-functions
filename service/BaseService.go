package service

import (
	"cloud.google.com/go/firestore"
)

type BaseService[T any] interface {
	FindByID(id string, tx *firestore.Transaction) (*T, error)
	Create(doc T) (*string, error)
	Update(id string, doc map[string]interface{}) (bool, error)
	Delete(id string) (bool, error)
}
