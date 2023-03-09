package tests

import (
	"context"
	"finize-functions.app/data/model"
	"finize-functions.app/functions"
	"finize-functions.app/service"
	"finize-functions.app/tests/fake"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestOnUserUpdated(t *testing.T) {
	database := fake.NewFirestoreService[model.RecurringTask](context.Background())
	taskService := testFactory.TaskService()

	oldUser := fake.NewUser(fake.NewUserEvent("userId1", "name", "email"))
	newUser := fake.NewUser(fake.NewUserEvent("userId1", "name", "email"))
	newUser.Profile.Currency = "CURR2"

	transaction := fake.NewTransaction(fake.NewTransactionEvent("id", "name", 50.0, 50.0,
		25.0, time.Now(), nil, nil, nil))
	transaction.AmountLocal.Currency = "CURR1"

	_, _ = database.Create(service.TasksDB(), nil, fake.MapTo[map[string]interface{}](
		fake.NewRecurringTask("userId1", model.CreateTransaction,
			3, model.Weekly, "zone", nil, fake.MapTo[map[string]interface{}](transaction)),
	))
	_, _ = database.Create(service.TasksDB(), nil, fake.MapTo[map[string]interface{}](
		fake.NewRecurringTask("userId1", model.CreateTransaction,
			3, model.Weekly, "zone", nil, fake.MapTo[map[string]interface{}](transaction)),
	))

	_ = functions.OnUserUpdated(testFactory, oldUser, newUser, []string{"profile.currency"})

	tasks, _ := taskService.FindByUser("userId1")
	assert.True(t, tasks != nil)
	for _, task := range tasks {
		assert.Equal(t, newUser.Profile.Currency, fake.MapTo[model.Transaction](task.Data).AmountLocal.Currency)
	}
}
