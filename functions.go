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

//goland:noinspection GoUnusedExportedFunction,GoUnusedParameter
func OnTransactionCreated(ctx context.Context, e data.FirestoreEvent[model.TransactionEvent]) error {
	if err := service.InitFirestore(context.Background()); err != nil {
		log.Fatalf("Failed to initialize Firestore %v", e.Value)
	}
	transaction, err := util.MapTo[model.Transaction](e.Value.Data)
	if err != nil {
		log.Fatalf("Failed to parse transaction %v", e.Value)
	}
	return functions.OnTransactionCreated(service.NewServiceFactory(context.Background(), e.UserID()), transaction)
}

//goland:noinspection GoUnusedExportedFunction,GoUnusedParameter
func OnTransactionUpdated(ctx context.Context, e data.FirestoreEvent[model.TransactionEvent]) error {
	if err := service.InitFirestore(context.Background()); err != nil {
		log.Fatalf("Failed to initialize Firestore %v", e.Value)
	}
	transactionOld, err := util.MapTo[model.Transaction](e.OldValue.Data)
	transactionNew, err2 := util.MapTo[model.Transaction](e.Value.Data)
	if err != nil || err2 != nil {
		log.Fatalf("Failed to parse transaction %v", e.Value)
	}
	return functions.OnTransactionUpdated(service.NewServiceFactory(context.Background(), e.UserID()), transactionOld, transactionNew, e.UpdateMask.Fields)
}

//goland:noinspection GoUnusedExportedFunction,GoUnusedParameter
func OnTransactionDeleted(ctx context.Context, e data.FirestoreEvent[model.TransactionEvent]) error {
	_ = service.InitFirestore(context.Background())
	transaction, err := util.MapTo[model.Transaction](e.OldValue.Data)
	if err != nil {
		log.Fatalf("Failed to parse transaction %v", e.Value)
	}
	return functions.OnTransactionDeleted(service.NewServiceFactory(context.Background(), e.UserID()), transaction)
}

//goland:noinspection GoUnusedExportedFunction,GoUnusedParameter
func GetExchangeRate(ctx context.Context) error {
	_ = service.InitFirestore(context.Background())
	return nil
}
