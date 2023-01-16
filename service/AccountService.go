package service

import (
	"finize-functions.app/data/model"
	"finize-functions.app/util"
	"fmt"
)

type AccountService struct {
	db FirestoreService[model.Account]
}

func accountsDB(userID string) string {
	return fmt.Sprintf("user-accounts/%s/accounts", userID)
}

func accountDoc(userID string, id string) string {
	return fmt.Sprintf("%s/%s", accountsDB(userID), id)
}

func NewAccountService() AccountService {
	return AccountService{db: newFirestoreService[model.Account]()}
}

func (service *AccountService) FindByID(userID string, id string) (*model.Account, error) {
	return service.db.Find(accountDoc(userID, id))
}

func (service *AccountService) Create(userID string, account model.Account) (string, error) {
	data, _ := util.MapTo[map[string]interface{}](account)
	return service.db.Create(accountsDB(userID), data)
}

func (service *AccountService) Update(userID string, id string, doc map[string]interface{}) (bool, error) {
	return service.db.Update(accountDoc(userID, id), doc)
}

func (service *AccountService) Delete(userID string, id string) (bool, error) {
	return service.db.Delete(accountDoc(userID, id))
}
