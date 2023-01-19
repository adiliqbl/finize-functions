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
	assert.True(t, !util.NullOrEmpty(&id))

	testAccount.Budget = util.Pointer("test-budget")
	id, err = accountService.Create(testAccount)
	assert.Nil(t, err)
	assert.True(t, !util.NullOrEmpty(&id))
}

func TestGetAccount(t *testing.T) {
	testAccount := fake.NewAccount(fake.NewAccountEvent("id", "name", 50.0, nil))
	testAccount.ID, _ = accountService.Create(testAccount)

	account, _ := accountService.FindByID(testAccount.ID)
	assert.Equal(t, testAccount, *account)

	testAccount = fake.NewAccount(fake.NewAccountEvent("id", "name", 50.0, util.Pointer("test-budget")))
	testAccount.ID, _ = accountService.Create(testAccount)

	account, _ = accountService.FindByID(testAccount.ID)
	assert.Equal(t, testAccount, *account)
}
