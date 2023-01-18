package fake

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/util"
	"github.com/stretchr/testify/mock"
)

type TestFirestore[T any] struct {
	mock.Mock
}

type TestCollectionRef struct {
	firestore.CollectionRef
}

type TestDocumentRef struct {
	firestore.DocumentRef
}

type TestTransaction struct {
	firestore.Transaction
}

type TestBatch struct {
	firestore.BulkWriter
}

func (mock *TestFirestore[T]) Doc(path string) *TestDocumentRef {
	args := mock.Called(path)
	return util.Pointer(args.Get(0).(TestDocumentRef))
}

func (mock *TestFirestore[T]) Collection(path string) *TestCollectionRef {
	args := mock.Called(path)
	return util.Pointer(args.Get(0).(TestCollectionRef))
}

func (mock *TestFirestore[T]) Batch() TestBatch {
	args := mock.Called()
	return args.Get(0).(TestBatch)
}

func (mock *TestFirestore[T]) Transaction(run func(tx *TestTransaction) error) error {
	args := mock.Called(run)
	return args.Error(0)
}

func (mock *TestFirestore[T]) Find(path string, tx *TestTransaction) (*T, error) {
	args := mock.Called(path, tx)
	arg := args.Get(0).(T)
	return &arg, args.Error(1)
}

func (mock *TestFirestore[T]) Create(collection string, doc map[string]interface{}) (string, error) {
	args := mock.Called(collection, doc)
	return args.Get(0).(string), args.Error(1)
}

func (mock *TestFirestore[T]) Update(path string, doc map[string]interface{}) (bool, error) {
	args := mock.Called(path, doc)
	return args.Get(0).(bool), args.Error(1)
}

func (mock *TestFirestore[T]) Delete(path string) (bool, error) {
	args := mock.Called(path)
	return args.Get(0).(bool), args.Error(1)
}
