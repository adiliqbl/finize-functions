package service

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/data/model"
	"finize-functions.app/util"
	"fmt"
)

type TransactionService interface {
	BaseService[model.Transaction]
	Doc(id string) *firestore.DocumentRef
}

type transactionServiceImpl struct {
	db     FirestoreService[model.Transaction]
	userID string
}

func TransactionsDB(userID string) string {
	return fmt.Sprintf("user-transactions/%s/transactions", userID)
}

func TransactionDoc(userID string, id string) string {
	return fmt.Sprintf("%s/%s", TransactionsDB(userID), id)
}

func NewTransactionService(db FirestoreService[model.Transaction], userID string) TransactionService {
	return &transactionServiceImpl{db: db, userID: userID}
}

func (service *transactionServiceImpl) Doc(id string) *firestore.DocumentRef {
	return service.db.Doc(TransactionDoc(service.userID, id))
}

func (service *transactionServiceImpl) FindByID(id string, tx *firestore.Transaction) (*model.Transaction, error) {
	return service.db.Find(TransactionDoc(service.userID, id), tx)
}

func (service *transactionServiceImpl) Create(transaction model.Transaction) (*string, error) {
	data, _ := util.MapTo[map[string]interface{}](transaction)
	return service.db.Create(TransactionsDB(service.userID), nil, data)
}

func (service *transactionServiceImpl) Update(id string, doc map[string]interface{}) (bool, error) {
	return service.db.Update(TransactionDoc(service.userID, id), doc)
}

func (service *transactionServiceImpl) Delete(id string) (bool, error) {
	return service.db.Delete(TransactionDoc(service.userID, id))
}
