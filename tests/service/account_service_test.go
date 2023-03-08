package service

import (
	"finize-functions.app/tests/fake"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	testAccount := fake.NewAccount(fake.NewAccountEvent("id", "name", 50.0, nil))
	id, err := accountService.Create(testAccount)
	assert.Nil(t, err)
	assert.True(t, !util.NullOrEmpty(id))

	testAccount.Budget = util.Pointer("test-budget")
	id, err = accountService.Create(testAccount)
	assert.Nil(t, err)
	assert.True(t, !util.NullOrEmpty(id))
}

func TestGetAccount(t *testing.T) {
	testAccount := fake.NewAccount(fake.NewAccountEvent("id", "name", 50.0, nil))
	id, _ := accountService.Create(testAccount)
	testAccount.ID = *id

	account, _ := accountService.FindByID(testAccount.ID, nil)
	assert.Equal(t, testAccount, *account)

	testAccount = fake.NewAccount(fake.NewAccountEvent("id", "name", 50.0, util.Pointer("test-budget")))
	id, _ = accountService.Create(testAccount)
	testAccount.ID = *id

	account, _ = accountService.FindByID(testAccount.ID, nil)
	assert.Equal(t, testAccount, *account)
}

func TestGetAccountsByBudget(t *testing.T) {
	testBudget := fake.NewBudget(fake.NewBudgetEvent("", "name", 50.0))
	budgetID, _ := budgetService.Create(testBudget)

	_, _ = accountService.Create(fake.NewAccount(fake.NewAccountEvent("id1", "name", 50.0, budgetID)))
	_, _ = accountService.Create(fake.NewAccount(fake.NewAccountEvent("id2", "name", 50.0, budgetID)))
	_, _ = accountService.Create(fake.NewAccount(fake.NewAccountEvent("id3", "name", 50.0, nil)))

	accounts, _ := accountService.FindByBudget(*budgetID)
	assert.Equal(t, 2, len(accounts))
}
