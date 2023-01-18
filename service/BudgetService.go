package service

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/data/model"
	"finize-functions.app/util"
	"fmt"
)

type BudgetService interface {
	BaseService[model.Budget]
}

type budgetServiceImpl struct {
	db     FirestoreService[model.Budget]
	userID string
}

func budgetsDB(userID string) string {
	return fmt.Sprintf("user-budgets/%s/budgets", userID)
}

func budgetDoc(userID string, id string) string {
	return fmt.Sprintf("%s/%s", budgetsDB(userID), id)
}

func (service *budgetServiceImpl) Doc(id string) *firestore.DocumentRef {
	return service.db.Doc(budgetDoc(service.userID, id))
}

func (service *budgetServiceImpl) FindByID(id string) (*model.Budget, error) {
	return service.db.Find(budgetDoc(service.userID, id), nil)
}

func (service *budgetServiceImpl) FindByIDWith(id string, tx *firestore.Transaction) (*model.Budget, error) {
	return service.db.Find(budgetDoc(service.userID, id), tx)
}

func (service *budgetServiceImpl) Create(budget model.Budget) (string, error) {
	data, _ := util.MapTo[map[string]interface{}](budget)
	return service.db.Create(budgetsDB(service.userID), data)
}

func (service *budgetServiceImpl) Update(id string, doc map[string]interface{}) (bool, error) {
	return service.db.Update(budgetDoc(service.userID, id), doc)
}

func (service *budgetServiceImpl) Delete(id string) (bool, error) {
	return service.db.Delete(budgetDoc(service.userID, id))
}
