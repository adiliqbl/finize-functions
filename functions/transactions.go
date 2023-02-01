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

	return factory.Firestore().Transaction(func(tx *firestore.Transaction) []data.DatabaseOperation {
		var ops []data.DatabaseOperation

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
	if !(slices.Contains(fields, "accountFrom") ||
		slices.Contains(fields, "accountTo") ||
		slices.Contains(fields, "amountFrom") ||
		slices.Contains(fields, "amountTo") ||
		slices.Contains(fields, "amountFrom.amount") ||
		slices.Contains(fields, "amountTo.amount")) {
		return nil
	}

	accounts := factory.AccountService()

	return factory.Firestore().Transaction(func(tx *firestore.Transaction) []data.DatabaseOperation {
		mAccounts := map[string]model.Account{}

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

		var ops []data.DatabaseOperation

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
