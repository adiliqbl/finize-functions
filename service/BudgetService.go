package service

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/data/model"
	"finize-functions.app/util"
	"fmt"
)

type BudgetService struct {
	db     FirestoreService[model.Budget]
	userID string
}

func budgetsDB(userID string) string {
	return fmt.Sprintf("user-budgets/%s/budgets", userID)
}

func budgetDoc(userID string, id string) string {
	return fmt.Sprintf("%s/%s", budgetsDB(userID), id)
}

func NewBudgetService(userId string) BudgetService {
	return BudgetService{db: NewFirestoreService[model.Budget](), userID: userId}
}

func (service *BudgetService) Doc(id string) *firestore.DocumentRef {
	return service.db.Doc(budgetDoc(service.userID, id))
}

func (service *BudgetService) FindByID(id string) (*model.Budget, error) {
	return service.db.Find(budgetDoc(service.userID, id), nil)
}

func (service *BudgetService) FindByIDWith(id string, tx *firestore.Transaction) (*model.Budget, error) {
	return service.db.Find(budgetDoc(service.userID, id), tx)
}

func (service *BudgetService) Create(budget model.Budget) (string, error) {
	data, _ := util.MapTo[map[string]interface{}](budget)
	return service.db.Create(budgetsDB(service.userID), data)
}

func (service *BudgetService) Update(id string, doc map[string]interface{}) (bool, error) {
	return service.db.Update(budgetDoc(service.userID, id), doc)
}

func (service *BudgetService) Delete(id string) (bool, error) {
	return service.db.Delete(budgetDoc(service.userID, id))
}
