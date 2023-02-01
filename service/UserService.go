package service

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/data/model"
	"finize-functions.app/util"
	"fmt"
)

type UserService interface {
	BaseService[model.User]
}

type userServiceImpl struct {
	db FirestoreService[model.User]
}

func UsersDB() string {
	return "users"
}

func UserDoc(id string) string {
	return fmt.Sprintf("%s/%s", UsersDB(), id)
}

func NewUserService(db FirestoreService[model.User]) UserService {
	return &userServiceImpl{db: db}
}

func (service *userServiceImpl) FindByID(id string, tx *firestore.Transaction) (*model.User, error) {
	return service.db.Find(UserDoc(id), tx)
}

func (service *userServiceImpl) Create(user model.User) (*string, error) {
	data, _ := util.MapTo[map[string]interface{}](user)
	return service.db.Create(UsersDB(), nil, data)
}

func (service *userServiceImpl) Update(id string, doc map[string]interface{}) (bool, error) {
	return service.db.Update(UserDoc(id), doc)
}

func (service *userServiceImpl) Delete(id string) (bool, error) {
	return service.db.Delete(UserDoc(id))
}
