package service

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/data/model"
	"finize-functions.app/util"
	"fmt"
)

type TransactionService struct {
	db     FirestoreService[model.Transaction]
	userID string
}

func transactionsDB(userID string) string {
	return fmt.Sprintf("user-transactions/%s/transactions", userID)
}

func transactionDoc(userID string, id string) string {
	return fmt.Sprintf("%s/%s", transactionsDB(userID), id)
}

func NewTransactionService(userId string) TransactionService {
	return TransactionService{db: NewFirestoreService[model.Transaction](), userID: userId}
}

func (service *TransactionService) Doc(id string) *firestore.DocumentRef {
	return service.db.Doc(transactionDoc(service.userID, id))
}

func (service *TransactionService) FindByID(id string) (*model.Transaction, error) {
	return service.db.Find(transactionDoc(service.userID, id), nil)
}

func (service *TransactionService) FindByIDWith(id string, tx *firestore.Transaction) (*model.Transaction, error) {
	return service.db.Find(transactionDoc(service.userID, id), tx)
}

func (service *TransactionService) Create(transaction model.Transaction) (string, error) {
	data, _ := util.MapTo[map[string]interface{}](transaction)
	return service.db.Create(transactionsDB(service.userID), data)
}

func (service *TransactionService) Update(id string, doc map[string]interface{}) (bool, error) {
	return service.db.Update(transactionDoc(service.userID, id), doc)
}

func (service *TransactionService) Delete(id string) (bool, error) {
	return service.db.Delete(transactionDoc(service.userID, id))
}
