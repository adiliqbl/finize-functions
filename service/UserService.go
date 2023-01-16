package service

import (
	"finize-functions.app/data/model"
	"finize-functions.app/util"
	"fmt"
)

type UserService struct {
	db FirestoreService[model.User]
}

func usersDB() string {
	return "users"
}

func userDoc(id string) string {
	return fmt.Sprintf("%s/%s", usersDB(), id)
}

func NewUserService() UserService {
	return UserService{db: newFirestoreService[model.User]()}
}

func (service *UserService) FindByID(id string) (*model.User, error) {
	return service.db.Find(userDoc(id))
}

func (service *UserService) Create(user model.User) (string, error) {
	data, _ := util.MapTo[map[string]interface{}](user)
	return service.db.Create(usersDB(), data)
}

func (service *UserService) Update(id string, doc map[string]interface{}) (bool, error) {
	return service.db.Update(userDoc(id), doc)
}

func (service *UserService) Delete(id string) (bool, error) {
	return service.db.Delete(userDoc(id))
}
