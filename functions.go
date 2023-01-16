package finize_functions

import (
	"context"
	"finize-functions.app/data"
	"finize-functions.app/data/model"
)

func OnTransactionCreated(ctx context.Context, e data.FirestoreEvent[model.TransactionEvent]) error {
	return nil
}

func OnTransactionUpdated(ctx context.Context, e data.FirestoreEvent[model.TransactionEvent]) error {
	return nil
}

func GetExchangeRate(ctx context.Context) error {
	return nil
}
