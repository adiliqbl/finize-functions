package tests

import (
	"finize-functions.app/functions"
	"finize-functions.app/tests/fake"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestOnTransactionCreated(t *testing.T) {
	fromAccountID, _ := accountService.Create(fake.NewAccount(fake.NewAccountEvent("", "name", 50.0, nil)))
	toAccountID, _ := accountService.Create(fake.NewAccount(fake.NewAccountEvent("", "name", 50.0, nil)))
	transactionID, _ := transactionService.Create(fake.NewTransaction(fake.NewTransactionEvent("", "name",
		7, 8, 9, time.Now(), toAccountID, fromAccountID, nil)))
	transaction, _ := transactionService.FindByID(*transactionID, nil)

	err := functions.OnTransactionCreated(testFactory, *transaction)
	assert.Nil(t, err)

	accountTo, _ := accountService.FindByID(*toAccountID, nil)
	accountFrom, _ := accountService.FindByID(*fromAccountID, nil)

	assert.Equal(t, 50.0+9.0, accountTo.Balance)
	assert.Equal(t, 50.0-9.0, accountFrom.Balance)
}

func TestOnTransactionAmountUpdated(t *testing.T) {
	accountID, _ := accountService.Create(fake.NewAccount(fake.NewAccountEvent("", "name", 50.0, nil)))
	transactionID, _ := transactionService.Create(fake.NewTransaction(fake.NewTransactionEvent("", "name",
		7, 8, 9, time.Now(), accountID, nil, nil)))

	oldTransaction := fake.NewTransaction(fake.NewTransactionEvent(*transactionID, "name",
		7, 8, 9, time.Now(), accountID, nil, nil))
	newTransaction := fake.NewTransaction(fake.NewTransactionEvent(*transactionID, "name",
		7, 8, 15, time.Now(), accountID, nil, nil))

	err := functions.OnTransactionCreated(testFactory, oldTransaction)
	assert.Nil(t, err)

	account, _ := accountService.FindByID(*accountID, nil)
	assert.Equal(t, 50.0+9.0, account.Balance)

	_ = functions.OnTransactionUpdated(testFactory, oldTransaction, newTransaction, []string{"amountTo.amount"})

	account, _ = accountService.FindByID(*accountID, nil)
	assert.Equal(t, 50.0+15.0, account.Balance)
}

func TestOnTransactionAccountUpdated(t *testing.T) {
	accountID, _ := accountService.Create(fake.NewAccount(fake.NewAccountEvent("", "name", 50.0, nil)))
	transactionID, _ := transactionService.Create(fake.NewTransaction(fake.NewTransactionEvent("", "name",
		7, 8, 9, time.Now(), accountID, nil, nil)))

	oldTransaction := fake.NewTransaction(fake.NewTransactionEvent(*transactionID, "name",
		7, 8, 9, time.Now(), accountID, nil, nil))
	newTransaction := fake.NewTransaction(fake.NewTransactionEvent(*transactionID, "name",
		7, 8, 12, time.Now(), nil, accountID, nil))

	err := functions.OnTransactionCreated(testFactory, oldTransaction)
	assert.Nil(t, err)

	account, _ := accountService.FindByID(*accountID, nil)
	assert.Equal(t, 50.0+9.0, account.Balance)

	_ = functions.OnTransactionUpdated(testFactory, oldTransaction, newTransaction, []string{"accountTo", "accountFrom", "amountFrom", "amountTo"})

	account, _ = accountService.FindByID(*accountID, nil)
	assert.Equal(t, 50.0-12.0, account.Balance)
}

func TestOnTransactionAccountSwitched(t *testing.T) {
	oldAccountID, _ := accountService.Create(fake.NewAccount(fake.NewAccountEvent("", "name", 50.0, nil)))
	newAccountID, _ := accountService.Create(fake.NewAccount(fake.NewAccountEvent("", "name", 70.0, nil)))
	transactionID, _ := transactionService.Create(fake.NewTransaction(fake.NewTransactionEvent("", "name",
		7, 8, 9, time.Now(), oldAccountID, nil, nil)))

	oldTransaction := fake.NewTransaction(fake.NewTransactionEvent(*transactionID, "name",
		7, 8, 9, time.Now(), oldAccountID, nil, nil))
	newTransaction := fake.NewTransaction(fake.NewTransactionEvent(*transactionID, "name",
		7, 8, 15, time.Now(), newAccountID, nil, nil))

	err := functions.OnTransactionCreated(testFactory, oldTransaction)
	assert.Nil(t, err)

	oldAccount, _ := accountService.FindByID(*oldAccountID, nil)
	newAccount, _ := accountService.FindByID(*newAccountID, nil)
	assert.Equal(t, 50.0+9.0, oldAccount.Balance)
	assert.Equal(t, 70.0, newAccount.Balance)

	_ = functions.OnTransactionUpdated(testFactory, oldTransaction, newTransaction, []string{"accountTo"})

	oldAccount, _ = accountService.FindByID(*oldAccountID, nil)
	newAccount, _ = accountService.FindByID(*newAccountID, nil)
	assert.Equal(t, 50.0, oldAccount.Balance)
	assert.Equal(t, 70.0+15.0, newAccount.Balance)
}

func TestOnTransactionAccountRemoved(t *testing.T) {
	toAccountID, _ := accountService.Create(fake.NewAccount(fake.NewAccountEvent("", "name", 50.0, nil)))
	fromAccountID, _ := accountService.Create(fake.NewAccount(fake.NewAccountEvent("", "name", 70.0, nil)))
	transactionID, _ := transactionService.Create(fake.NewTransaction(fake.NewTransactionEvent("", "name",
		7, 8, 9, time.Now(), toAccountID, fromAccountID, nil)))

	oldTransaction := fake.NewTransaction(fake.NewTransactionEvent(*transactionID, "name",
		7, 8, 9, time.Now(), toAccountID, fromAccountID, nil))
	newTransaction := fake.NewTransaction(fake.NewTransactionEvent(*transactionID, "name",
		7, 8, 15, time.Now(), toAccountID, nil, nil))

	err := functions.OnTransactionCreated(testFactory, oldTransaction)
	assert.Nil(t, err)

	toAccount, _ := accountService.FindByID(*toAccountID, nil)
	fromAccount, _ := accountService.FindByID(*fromAccountID, nil)
	assert.Equal(t, 50.0+9.0, toAccount.Balance)
	assert.Equal(t, 70-9.0, fromAccount.Balance)

	_ = functions.OnTransactionUpdated(testFactory, oldTransaction, newTransaction, []string{"accountFrom", "accountFrom.amount", "accountTo.amount"})

	toAccount, _ = accountService.FindByID(*toAccountID, nil)
	fromAccount, _ = accountService.FindByID(*fromAccountID, nil)
	assert.Equal(t, 50.0+15.0, toAccount.Balance)
	assert.Equal(t, 70.0, fromAccount.Balance)
}

func TestOnTransactionDeleted(t *testing.T) {
	toAccountID, _ := accountService.Create(fake.NewAccount(fake.NewAccountEvent("", "name", 50.0, nil)))
	fromAccountID, _ := accountService.Create(fake.NewAccount(fake.NewAccountEvent("", "name", 70.0, nil)))
	transactionID, _ := transactionService.Create(fake.NewTransaction(fake.NewTransactionEvent("", "name",
		7, 8, 9, time.Now(), toAccountID, fromAccountID, nil)))

	transaction := fake.NewTransaction(fake.NewTransactionEvent(*transactionID, "name",
		7, 8, 9, time.Now(), toAccountID, fromAccountID, nil))

	err := functions.OnTransactionCreated(testFactory, transaction)
	assert.Nil(t, err)

	toAccount, _ := accountService.FindByID(*toAccountID, nil)
	fromAccount, _ := accountService.FindByID(*fromAccountID, nil)
	assert.Equal(t, 50.0+9.0, toAccount.Balance)
	assert.Equal(t, 70-9.0, fromAccount.Balance)

	_ = functions.OnTransactionDeleted(testFactory, transaction)

	toAccount, _ = accountService.FindByID(*toAccountID, nil)
	fromAccount, _ = accountService.FindByID(*fromAccountID, nil)
	assert.Equal(t, 50.0, toAccount.Balance)
	assert.Equal(t, 70.0, fromAccount.Balance)
}
