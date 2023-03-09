package functions

import (
	"cloud.google.com/go/firestore"
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/service"
	"finize-functions.app/util"
	"golang.org/x/exp/slices"
	"log"
)

func OnTransactionCreated(factory service.Factory, transaction model.Transaction) error {
	accounts := factory.AccountService()
	transactions := factory.TransactionService()

	return factory.Firestore().Transaction(func(tx *firestore.Transaction) []data.DatabaseOperation {
		var ops []data.DatabaseOperation

		ops = append(ops, data.DatabaseOperation{
			Ref: transactions.Doc(transaction.ID),
			Data: []firestore.Update{{
				Path:  model.FieldKeywords,
				Value: util.GenerateKeywords(transaction.Name),
			}},
		})

		if !util.NullOrEmpty(transaction.AccountFrom) {
			accountFrom, err := accounts.FindByID(*transaction.AccountFrom, tx)
			if err != nil {
				log.Fatalf("AccountService.FindByID: %v", err)
			}
			accountFrom.Balance = accountFrom.Balance - transaction.AmountFrom.Amount

			ops = append(ops, data.DatabaseOperation{
				Ref: accounts.Doc(accountFrom.ID),
				Data: []firestore.Update{{
					Path:  model.FieldBalance,
					Value: accountFrom.Balance,
				}},
			})
		}

		if !util.NullOrEmpty(transaction.AccountTo) {
			accountTo, err := accounts.FindByID(*transaction.AccountTo, tx)
			if err != nil {
				log.Fatalf("AccountService.FindByID: %v", err)
			}
			accountTo.Balance = accountTo.Balance + transaction.AmountTo.Amount

			ops = append(ops, data.DatabaseOperation{
				Ref: accounts.Doc(accountTo.ID),
				Data: []firestore.Update{{
					Path:  model.FieldBalance,
					Value: accountTo.Balance,
				}},
			})
		}

		return ops
	})
}

func OnTransactionUpdated(factory service.Factory, oldTransaction model.Transaction, transaction model.Transaction, fields []string) error {
	updateAccount := slices.Contains(fields, model.FieldAccountFrom) ||
		slices.Contains(fields, model.FieldAccountTo) ||
		slices.Contains(fields, model.FieldAmountTo) ||
		slices.Contains(fields, model.FieldAmountTo+"."+model.FieldAmount) ||
		slices.Contains(fields, model.FieldAmountFrom) ||
		slices.Contains(fields, model.FieldAmountFrom+"."+model.FieldAmount)

	return factory.Firestore().Transaction(func(tx *firestore.Transaction) []data.DatabaseOperation {
		var ops []data.DatabaseOperation

		if slices.Contains(fields, model.FieldName) {
			ops = append(ops, data.DatabaseOperation{
				Ref: factory.TransactionService().Doc(transaction.ID),
				Data: []firestore.Update{{
					Path:  model.FieldKeywords,
					Value: util.GenerateKeywords(transaction.Name),
				}},
			})
		}

		if updateAccount {
			accountOps := updateAccountBalances(factory, oldTransaction, transaction, tx)
			ops = append(ops, accountOps...)
		}

		return ops
	})
}

func OnTransactionDeleted(factory service.Factory, transaction model.Transaction) error {
	accounts := factory.AccountService()

	return factory.Firestore().Transaction(func(tx *firestore.Transaction) []data.DatabaseOperation {
		var ops []data.DatabaseOperation

		if transaction.AccountFrom != nil {
			account, err := accounts.FindByID(*transaction.AccountFrom, tx)
			if err != nil {
				log.Fatalf("AccountService.FindByID: %v", err)
			}
			account.Balance = account.Balance + transaction.AmountFrom.Amount

			ops = append(ops, data.DatabaseOperation{
				Ref: accounts.Doc(account.ID),
				Data: []firestore.Update{{
					Path:  model.FieldBalance,
					Value: account.Balance,
				}},
			})
		}

		if transaction.AccountTo != nil {
			account, err := accounts.FindByID(*transaction.AccountTo, tx)
			if err != nil {
				log.Fatalf("AccountService.FindByID: %v", err)
			}
			account.Balance = account.Balance - transaction.AmountTo.Amount

			ops = append(ops, data.DatabaseOperation{
				Ref: accounts.Doc(account.ID),
				Data: []firestore.Update{{
					Path:  model.FieldBalance,
					Value: account.Balance,
				}},
			})
		}

		return ops
	})
}

func updateAccountBalances(factory service.Factory, oldTransaction model.Transaction, transaction model.Transaction, tx *firestore.Transaction) []data.DatabaseOperation {
	accounts := factory.AccountService()
	mAccounts := map[string]model.Account{}

	var ops []data.DatabaseOperation

	if oldTransaction.AccountFrom != nil {
		account, err := accounts.FindByID(*oldTransaction.AccountFrom, tx)
		if err != nil {
			log.Fatalf("AccountService.FindByID: %v", err)
		}
		account.Balance = account.Balance + oldTransaction.AmountFrom.Amount
		mAccounts[account.ID] = *account
	}

	if oldTransaction.AccountTo != nil {
		account, err := accounts.FindByID(*oldTransaction.AccountTo, tx)
		if err != nil {
			log.Fatalf("AccountService.FindByID: %v", err)
		}
		account.Balance = account.Balance - oldTransaction.AmountTo.Amount
		mAccounts[account.ID] = *account
	}

	if transaction.AccountFrom != nil {
		var account model.Account

		if mAccount, ok := mAccounts[*transaction.AccountFrom]; ok {
			account = mAccount
		} else {
			apiAccount, err := accounts.FindByID(*transaction.AccountFrom, tx)
			if err != nil {
				log.Fatalf("AccountService.FindByID: %v", err)
			}
			account = *apiAccount
		}

		account.Balance = account.Balance - transaction.AmountFrom.Amount
		mAccounts[account.ID] = account
	}

	if transaction.AccountTo != nil {
		var account model.Account

		if mAccount, ok := mAccounts[*transaction.AccountTo]; ok {
			account = mAccount
		} else {
			apiAccount, err := accounts.FindByID(*transaction.AccountTo, tx)
			if err != nil {
				log.Fatalf("AccountService.FindByID: %v", err)
			}
			account = *apiAccount
		}

		account.Balance = account.Balance + transaction.AmountTo.Amount
		mAccounts[account.ID] = account
	}

	for _, account := range mAccounts {
		ops = append(ops, data.DatabaseOperation{
			Ref: accounts.Doc(account.ID),
			Data: []firestore.Update{{
				Path:  model.FieldBalance,
				Value: account.Balance,
			}},
		})
	}

	return ops
}
