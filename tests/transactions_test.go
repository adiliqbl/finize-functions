package tests

import (
	"finize-functions.app/functions"
	"finize-functions.app/tests/fake"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestOnTransactionCreated(t *testing.T) {
	budgetID, _ := budgetService.Create(fake.NewBudget(fake.NewBudgetEvent("id", "name", 50.0, 0.0)))
	accountID, _ := accountService.Create(fake.NewAccount(fake.NewAccountEvent("id", "name", 50.0, &budgetID)))
	transactionID, _ := transactionService.Create(fake.NewTransaction(fake.NewTransactionEvent("id", "name",
		7, 8, 9, time.Now(), &accountID, nil, &budgetID)))
	transaction, _ := transactionService.FindByID(transactionID)

	err := functions.OnTransactionCreated(testFactory, *transaction)
	assert.Nil(t, err)

	budget, _ := budgetService.FindByID(budgetID)
	account, _ := accountService.FindByID(accountID)

	assert.Equal(t, 8.0, budget.Spent)
	assert.Equal(t, 50.0, budget.Limit)
	assert.Equal(t, 50.0-9.0, account.Balance)
}
