package fake

import (
	"finize-functions.app/util"
	"github.com/stretchr/testify/mock"
)

type TestService[T any] struct {
	mock.Mock
}

func (mock *TestService[T]) Doc(id string) *TestDocumentRef {
	args := mock.Called(id)
	return util.Pointer(args.Get(0).(TestDocumentRef))
}

func (mock *TestService[T]) FindByID(id string) (*T, error) {
	args := mock.Called(id)
	return util.Pointer(args.Get(0).(T)), args.Error(1)
}

func (mock *TestService[T]) FindByIDWith(id string, tx *TestTransaction) (*T, error) {
	args := mock.Called(id, tx)
	return util.Pointer(args.Get(0).(T)), args.Error(1)
}

func (mock *TestService[T]) Create(doc T) (string, error) {
	args := mock.Called(doc)
	return args.Get(0).(string), args.Error(1)
}

func (mock *TestService[T]) Update(id string, doc map[string]interface{}) (bool, error) {
	args := mock.Called(id, doc)
	return args.Get(0).(bool), args.Error(1)
}

func (mock *TestService[T]) Delete(id string) (bool, error) {
	args := mock.Called(id)
	return args.Get(0).(bool), args.Error(1)
}
