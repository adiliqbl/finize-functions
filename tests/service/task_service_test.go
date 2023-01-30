package service

import (
	"context"
	"finize-functions.app/data/model"
	"finize-functions.app/tests"
	"finize-functions.app/tests/fake"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTasksPagination(t *testing.T) {
	database := fake.NewFirestoreService[model.RecurringTask](context.Background())
	tests.ClearDatabase()

	for i := 0; i < 30; i++ {
		testTask := fake.NewRecurringTask("userId", model.CreateTransaction,
			3, model.Weekly, "zone", nil, map[string]interface{}{"test": "value"})
		_, _ = database.Create("tasks", nil, fake.MapTo[map[string]interface{}](testTask))
	}

	tasks, err := taskService.Paginate(0, 500)

	assert.True(t, err == nil)
	assert.True(t, tasks != nil)
	assert.Equal(t, 30, len(tasks))
}
