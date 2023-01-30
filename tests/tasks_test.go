package tests

import (
	"context"
	"finize-functions.app/data/model"
	"finize-functions.app/functions"
	services "finize-functions.app/service"
	"finize-functions.app/tests/fake"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestRecurringTasks(t *testing.T) {
	database := fake.NewFirestoreService[model.RecurringTask](context.Background())
	transactions := fake.NewFirestoreService[model.Transaction](context.Background())
	ClearDatabase()

	transaction := fake.MapTo[map[string]interface{}](fake.NewTransaction(fake.NewTransactionEvent("", "name",
		7, 8, 9, time.Now(), util.Pointer("toAccountID"), nil, nil)))

	testTasks := []model.RecurringTask{
		fake.NewRecurringTask("user1", model.CreateTransaction,
			1, model.Weekly, "Etc/GMT-5", util.Pointer(fake.NewClock(2023, time.January, 15).Now()), transaction),
		fake.NewRecurringTask("user11", model.CreateTransaction,
			1, model.Weekly, "Etc/GMT-5", util.Pointer(fake.NewClock(2023, time.January, 8).Now()), transaction),
		fake.NewRecurringTask("user2", model.CreateTransaction,
			15, model.Monthly, "Etc/GMT-5", util.Pointer(fake.NewClock(2022, time.December, 15).Now()), transaction),
		fake.NewRecurringTask("user22", model.CreateTransaction,
			15, model.Monthly, "Etc/GMT-5", util.Pointer(fake.NewClock(2023, time.January, 15).Now()), transaction),
		fake.NewRecurringTask("user3", model.CreateTransaction,
			16, model.Monthly, "Etc/GMT", nil, transaction),
		fake.NewRecurringTask("user4", model.CreateTransaction,
			13, model.Yearly, "Etc/GMT+5", nil, transaction),
		fake.NewRecurringTask("user5", model.CreateTransaction,
			15, model.Yearly, "Etc/GMT-5", nil, transaction),
		fake.NewRecurringTask("user6", model.CreateTransaction,
			0, model.Weekly, "Etc/GMT+7", nil, transaction),
		fake.NewRecurringTask("user7", model.CreateTransaction,
			15, model.Yearly, "Etc/GMT+5", nil, transaction),
	}
	for i, testTask := range testTasks {
		_, _ = database.Create("tasks", util.Pointer("task"+strconv.Itoa(i)), fake.MapTo[map[string]interface{}](testTask))
	}

	err := functions.ProcessRecurringTasks(testFactory, fake.NewClock(2023, time.January, 15))

	// Already Processed
	user1, _ := transactions.GetAll(services.TransactionsDB("user1"))
	user22, _ := transactions.GetAll(services.TransactionsDB("user22"))
	assert.Equal(t, 0, len(user1))
	assert.Equal(t, 0, len(user22))

	// Need Processing
	user2, _ := transactions.GetAll(services.TransactionsDB("user2"))
	user5, _ := transactions.GetAll(services.TransactionsDB("user5"))
	assert.Equal(t, 1, len(user2))
	assert.Equal(t, 1, len(user5))
	assert.True(t, !util.NullOrEmpty(&user2[0].ID))
	assert.True(t, !util.NullOrEmpty(&user5[0].ID))

	// Ignore Processing
	user3, _ := transactions.GetAll(services.TransactionsDB("user3"))
	user4, _ := transactions.GetAll(services.TransactionsDB("user4"))
	user6, _ := transactions.GetAll(services.TransactionsDB("user6"))
	user7, _ := transactions.GetAll(services.TransactionsDB("user7"))
	user11, _ := transactions.GetAll(services.TransactionsDB("user11"))
	assert.Equal(t, 0, len(user3))
	assert.Equal(t, 0, len(user4))
	assert.Equal(t, 0, len(user6))
	assert.Equal(t, 0, len(user7))
	assert.Equal(t, 0, len(user11))

	assert.True(t, err == nil)
}
