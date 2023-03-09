package service

import (
	"context"
	"finize-functions.app/data/model"
	"finize-functions.app/service"
	"finize-functions.app/tests"
	"finize-functions.app/tests/fake"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTasksPagination(t *testing.T) {
	database := fake.NewFirestoreService[model.RecurringTask](context.Background())
	tests.ClearDatabase()

	for i := 0; i < 30; i++ {
		testTask := fake.NewRecurringTask("userId", model.CreateTransaction,
			3, model.Weekly, "zone", nil, map[string]interface{}{"test": "value"})
		_, _ = database.Create(service.TasksDB(), nil, fake.MapTo[map[string]interface{}](testTask))
	}

	tasks, err := taskService.Paginate(0, 500)

	assert.True(t, err == nil)
	assert.True(t, tasks != nil)
	assert.Equal(t, 30, len(tasks))
}

func TestTasksByUser(t *testing.T) {
	database := fake.NewFirestoreService[model.RecurringTask](context.Background())
	tests.ClearDatabase()

	_, _ = database.Create(service.TasksDB(), nil, fake.MapTo[map[string]interface{}](
		fake.NewRecurringTask("userId1", model.CreateTransaction,
			3, model.Weekly, "zone", nil, map[string]interface{}{"test": "value"}),
	))
	_, _ = database.Create(service.TasksDB(), nil, fake.MapTo[map[string]interface{}](
		fake.NewRecurringTask("userId1", model.CreateTransaction,
			3, model.Weekly, "zone", nil, map[string]interface{}{"test": "value"}),
	))
	_, _ = database.Create(service.TasksDB(), nil, fake.MapTo[map[string]interface{}](
		fake.NewRecurringTask("userId1", model.CreateTransaction,
			3, model.Weekly, "zone", nil, map[string]interface{}{"test": "value"}),
	))
	_, _ = database.Create(service.TasksDB(), nil, fake.MapTo[map[string]interface{}](
		fake.NewRecurringTask("userId2", model.CreateTransaction,
			3, model.Weekly, "zone", nil, map[string]interface{}{"test": "value"}),
	))

	tasks1, err1 := taskService.FindByUser("userId1")
	tasks2, err2 := taskService.FindByUser("userId2")

	assert.True(t, err1 == nil)
	assert.True(t, err2 == nil)
	assert.Equal(t, 3, len(tasks1))
	assert.Equal(t, 1, len(tasks2))
}

func TestTasksByAccount(t *testing.T) {
	database := fake.NewFirestoreService[model.RecurringTask](context.Background())
	tests.ClearDatabase()

	transaction := fake.NewTransaction(fake.NewTransactionEvent("id", "name", 50.0, 50.0,
		25.0, time.Now(), nil, nil, nil))

	transaction.AccountTo = util.Pointer("account1")
	_, _ = database.Create(service.TasksDB(), nil, fake.MapTo[map[string]interface{}](
		fake.NewRecurringTask("userId1", model.CreateTransaction,
			3, model.Weekly, "zone", nil, fake.MapTo[map[string]interface{}](transaction)),
	))

	transaction.AccountTo = nil
	transaction.AccountFrom = util.Pointer("account1")
	_, _ = database.Create(service.TasksDB(), nil, fake.MapTo[map[string]interface{}](
		fake.NewRecurringTask("userId1", model.CreateTransaction,
			3, model.Weekly, "zone", nil, fake.MapTo[map[string]interface{}](transaction)),
	))

	transaction.AccountTo = util.Pointer("account0")
	transaction.AccountFrom = util.Pointer("account1")
	_, _ = database.Create(service.TasksDB(), nil, fake.MapTo[map[string]interface{}](
		fake.NewRecurringTask("userId1", model.CreateTransaction,
			3, model.Weekly, "zone", nil, fake.MapTo[map[string]interface{}](transaction)),
	))

	transaction.AccountTo = nil
	transaction.AccountFrom = util.Pointer("account2")
	_, _ = database.Create(service.TasksDB(), nil, fake.MapTo[map[string]interface{}](
		fake.NewRecurringTask("userId1", model.CreateTransaction,
			3, model.Weekly, "zone", nil, fake.MapTo[map[string]interface{}](transaction)),
	))

	tasks1, err1 := taskService.FindByAccount("account1")
	tasks2, err2 := taskService.FindByAccount("account3")

	assert.True(t, err1 == nil)
	assert.True(t, err2 == nil)
	assert.Equal(t, 3, len(tasks1))
	assert.Equal(t, 0, len(tasks2))
}
