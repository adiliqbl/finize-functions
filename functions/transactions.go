package functions

import (
	"context"
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/service"
	"log"
	"strings"
)

func OnTransactionCreated(ctx context.Context, event data.FirestoreEvent[model.Transaction]) error {
	service.InitFirestore(context.Background())

	transactions := service.NewTransactionService()
	accounts := service.NewAccountService()
	budgets := service.NewBudgetService()

	var budget *model.Budget
	var accountTo *model.Account
	var accountFrom *model.Account

	var transaction = event.Value.Data

	if transaction.Budget != nil {

	}

	collection, doc, _ := e.Path()

	log.Printf(collection + "/" + doc)

	curValue := e.Value.Data.ID
	newValue := strings.ToUpper(curValue)
	if curValue == newValue {
		log.Printf("%q is already upper case: skipping", curValue)
		return nil
	}
	log.Printf("Replacing value: %q -> %q", curValue, newValue)

	return nil
}
