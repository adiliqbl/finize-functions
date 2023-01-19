package service

import (
	"finize-functions.app/tests/fake"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateTransaction(t *testing.T) {
	testTransaction := fake.NewTransaction(fake.NewTransactionEvent("id", "name", 50.0, 50.0,
		25.0, time.Now(), util.Pointer("test-account"), nil, nil))

	id, err := transactionService.Create(testTransaction)
	assert.Nil(t, err)
	assert.True(t, !util.NullOrEmpty(&id))

	testTransaction.AccountTo = nil
	testTransaction.AccountFrom = util.Pointer("test-account")
	testTransaction.Budget = util.Pointer("test-budget")
	id, err = transactionService.Create(testTransaction)
	assert.Nil(t, err)
	assert.True(t, !util.NullOrEmpty(&id))
}

func TestGetTransaction(t *testing.T) {
	testTransaction := fake.NewTransaction(fake.NewTransactionEvent("id", "name", 50.0, 50.0,
		25.0, time.Now(), util.Pointer("test-account"), nil, nil))

	testTransaction.ID, _ = transactionService.Create(testTransaction)

	transaction, _ := transactionService.FindByID(testTransaction.ID)
	assert.Equal(t, testTransaction, *transaction)

	testTransaction.AccountTo = nil
	testTransaction.AccountFrom = util.Pointer("test-account")
	testTransaction.Budget = util.Pointer("test-budget")
	testTransaction.ID, _ = transactionService.Create(testTransaction)

	transaction, _ = transactionService.FindByID(testTransaction.ID)
	assert.Equal(t, testTransaction, *transaction)
}
