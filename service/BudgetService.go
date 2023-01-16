package service

import (
	"finize-functions.app/data/model"
	"finize-functions.app/util"
	"fmt"
)

type BudgetService struct {
	db FirestoreService[model.Budget]
}

func budgetsDB(userID string) string {
	return fmt.Sprintf("user-budgets/%s/budgets", userID)
}

func budgetDoc(userID string, id string) string {
	return fmt.Sprintf("%s/%s", budgetsDB(userID), id)
}

func (service *BudgetService) FindByID(userID string, id string) (*model.Budget, error) {
	return service.db.Find(budgetDoc(userID, id))
}

func (service *BudgetService) Create(userID string, budget model.Budget) (string, error) {
	data, _ := util.MapTo[map[string]interface{}](budget)
	return service.db.Create(budgetsDB(userID), data)
}

func (service *BudgetService) Update(userID string, id string, doc map[string]interface{}) (bool, error) {
	return service.db.Update(budgetDoc(userID, id), doc)
}

func (service *BudgetService) Delete(userID string, id string) (bool, error) {
	return service.db.Delete(budgetDoc(userID, id))
}
