package service

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/data/model"
	"finize-functions.app/util"
	"fmt"
)

type AccountService interface {
	BaseService[model.Account]
}

type accountServiceImpl struct {
	db     FirestoreService[model.Account]
	userID string
}

func accountsDB(userID string) string {
	return fmt.Sprintf("user-accounts/%s/accounts", userID)
}

func accountDoc(userID string, id string) string {
	return fmt.Sprintf("%s/%s", accountsDB(userID), id)
}

func NewAccountService(db FirestoreService[model.Account], userID string) AccountService {
	return &accountServiceImpl{db: db, userID: userID}
}

func (service *accountServiceImpl) Doc(id string) *firestore.DocumentRef {
	return service.db.Doc(accountDoc(service.userID, id))
}

func (service *accountServiceImpl) FindByID(id string) (*model.Account, error) {
	return service.db.Find(accountDoc(service.userID, id), nil)
}

func (service *accountServiceImpl) FindByIDWith(id string, tx *firestore.Transaction) (*model.Account, error) {
	return service.db.Find(accountDoc(service.userID, id), tx)
}

func (service *accountServiceImpl) Create(account model.Account) (string, error) {
	data, _ := util.MapTo[map[string]interface{}](account)
	return service.db.Create(accountsDB(service.userID), data)
}

func (service *accountServiceImpl) Update(id string, doc map[string]interface{}) (bool, error) {
	return service.db.Update(accountDoc(service.userID, id), doc)
}

func (service *accountServiceImpl) Delete(id string) (bool, error) {
	return service.db.Delete(accountDoc(service.userID, id))
}
