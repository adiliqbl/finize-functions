package finize_functions

import (
	"context"
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/functions"
	"finize-functions.app/service"
	"finize-functions.app/util"
	"log"
)

func OnTransactionCreated(ctx context.Context, e data.FirestoreEvent[model.TransactionEvent]) error {
	_ = service.InitFirestore(context.Background())
	transaction, err := util.MapTo[model.Transaction](e.Value.Data)
	if err != nil {
		log.Fatalf("Failed to parse transaction %v", e.Value)
	}
	return functions.OnTransactionCreated(e.UserID(), transaction)
}

func OnTransactionUpdated(ctx context.Context, e data.FirestoreEvent[model.TransactionEvent]) error {
	_ = service.InitFirestore(context.Background())
	return nil
}

func GetExchangeRate(ctx context.Context) error {
	_ = service.InitFirestore(context.Background())
	return nil
}
