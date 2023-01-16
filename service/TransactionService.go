package service

import (
	"finize-functions.app/data/model"
	"finize-functions.app/util"
	"fmt"
)

type TransactionService struct {
	db FirestoreService[model.Transaction]
}

func transactionsDB(userID string) string {
	return fmt.Sprintf("user-transactions/%s/transactions", userID)
}

func transactionDoc(userID string, id string) string {
	return fmt.Sprintf("%s/%s", transactionsDB(userID), id)
}

func (service *TransactionService) FindByID(userID string, id string) (*model.Transaction, error) {
	return service.db.Find(transactionDoc(userID, id))
}

func (service *TransactionService) Create(userID string, transaction model.Transaction) (string, error) {
	data, _ := util.MapTo[map[string]interface{}](transaction)
	return service.db.Create(transactionsDB(userID), data)
}

func (service *TransactionService) Update(userID string, id string, doc map[string]interface{}) (bool, error) {
	return service.db.Update(transactionDoc(userID, id), doc)
}

func (service *TransactionService) Delete(userID string, id string) (bool, error) {
	return service.db.Delete(transactionDoc(userID, id))
}
