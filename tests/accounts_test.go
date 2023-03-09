package tests

import (
	"context"
	"finize-functions.app/data/model"
	"finize-functions.app/functions"
	"finize-functions.app/service"
	"finize-functions.app/tests/fake"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestOnAccountUpdated(t *testing.T) {
	database := fake.NewFirestoreService[model.RecurringTask](context.Background())
	taskService := testFactory.TaskService()

	oldAccount := fake.NewAccount(fake.NewAccountEvent("account1", "name", 50.0, nil))
	oldAccount.Currency = "CURR"
	newAccount := fake.NewAccount(fake.NewAccountEvent("account1", "name", 50.0, nil))
	newAccount.Currency = "CURR2"

	transaction := fake.NewTransaction(fake.NewTransactionEvent("id", "name", 50.0, 50.0,
		25.0, time.Now(), nil, nil, nil))

	transaction.AccountTo = util.Pointer("account1")
	t1, _ := database.Create(service.TasksDB(), nil, fake.MapTo[map[string]interface{}](
		fake.NewRecurringTask("userId1", model.CreateTransaction,
			3, model.Weekly, "zone", nil, fake.MapTo[map[string]interface{}](transaction)),
	))

	transaction.AccountTo = nil
	transaction.AccountFrom = util.Pointer("account1")
	t2, _ := database.Create(service.TasksDB(), nil, fake.MapTo[map[string]interface{}](
		fake.NewRecurringTask("userId1", model.CreateTransaction,
			3, model.Weekly, "zone", nil, fake.MapTo[map[string]interface{}](transaction)),
	))

	_ = functions.OnAccountUpdated(testFactory, oldAccount, newAccount, []string{"currency"})

	tasks, _ := taskService.FindByAccount("account1")
	assert.True(t, tasks != nil)
	for _, task := range tasks {
		mTransaction := fake.MapTo[model.Transaction](task.Data)
		if task.ID == *t1 {
			assert.Equal(t, newAccount.Currency, mTransaction.AmountTo.Currency)
		} else if task.ID == *t2 {
			assert.Equal(t, newAccount.Currency, mTransaction.AmountFrom.Currency)
		}
	}
}
