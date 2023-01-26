package service

import (
	"context"
	"finize-functions.app/data/model"
	"finize-functions.app/tests/fake"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestTasksPagination(t *testing.T) {
	database := fake.NewFirestoreService[model.RecurringTask](context.Background())

	for i := 0; i <= 30; i++ {
		testTask := fake.NewRecurringTask("recurring-user"+strconv.Itoa(i), model.CreateTransaction,
			24*7, map[string]interface{}{"test": "value"})
		_, _ = database.Create("paginate-users", util.Pointer("user"+strconv.Itoa(i)), fake.MapTo[map[string]interface{}](testTask))
	}

	tasks, err := taskService.PaginateTasks(0, 500)

	assert.True(t, err == nil)
	assert.True(t, tasks != nil)
	assert.Equal(t, 30, len(tasks))
}
