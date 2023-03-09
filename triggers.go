package finize_functions

import (
	"cloud.google.com/go/functions/metadata"
	"context"
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/functions"
	"finize-functions.app/service"
	"finize-functions.app/util"
	"log"
)

func init() {
	if err := service.InitFirestore(context.Background()); err != nil {
		log.Fatalf("Failed to initialize Firestore %v", err)
	}
}

//goland:noinspection GoUnusedExportedFunction,GoUnusedParameter
func OnTransactionCreated(ctx context.Context, e data.FirestoreEvent[model.TransactionEvent]) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		log.Fatalf("Failed to get metadata %v", err)
	}

	factory := service.NewServiceFactory(ctx, meta.EventID, e.UserID())
	if factory.EventService().IsProcessed() {
		return nil
	}

	transaction, err := util.MapTo[model.Transaction](e.Value.Data)
	if err != nil {
		log.Fatalf("Failed to parse transaction %v", err)
	}
	return functions.OnTransactionCreated(factory, transaction)
}

//goland:noinspection GoUnusedExportedFunction,GoUnusedParameter
func OnTransactionUpdated(ctx context.Context, e data.FirestoreEvent[model.TransactionEvent]) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		log.Fatalf("Failed to get metadata %v", err)
	}

	factory := service.NewServiceFactory(ctx, meta.EventID, e.UserID())
	if factory.EventService().IsProcessed() {
		return nil
	}

	transactionOld, err := util.MapTo[model.Transaction](e.OldValue.Data)
	transactionNew, err2 := util.MapTo[model.Transaction](e.Value.Data)
	if err != nil || err2 != nil {
		log.Fatalf("Failed to parse transaction %v", err)
	}
	return functions.OnTransactionUpdated(factory, transactionOld, transactionNew, e.UpdateMask.Fields)
}

//goland:noinspection GoUnusedExportedFunction,GoUnusedParameter
func OnTransactionDeleted(ctx context.Context, e data.FirestoreEvent[model.TransactionEvent]) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		log.Fatalf("Failed to get metadata %v", err)
	}

	factory := service.NewServiceFactory(ctx, meta.EventID, e.UserID())
	if factory.EventService().IsProcessed() {
		return nil
	}

	transaction, err := util.MapTo[model.Transaction](e.Value.Data)
	if err != nil {
		log.Fatalf("Failed to parse transaction %v", err)
	}
	return functions.OnTransactionDeleted(factory, transaction)
}

//goland:noinspection GoUnusedExportedFunction,GoUnusedParameter
func OnBudgetDeleted(ctx context.Context, e data.FirestoreEvent[model.BudgetEvent]) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		log.Fatalf("Failed to get metadata %v", err)
	}

	factory := service.NewServiceFactory(ctx, meta.EventID, e.UserID())
	if factory.EventService().IsProcessed() {
		return nil
	}

	budget, err := util.MapTo[model.Budget](e.Value.Data)
	if err != nil {
		log.Fatalf("Failed to parse budget %v", err)
	}
	return functions.OnBudgetDeleted(factory, budget)
}

//goland:noinspection GoUnusedExportedFunction,GoUnusedParameter
func OnUserUpdated(ctx context.Context, e data.FirestoreEvent[model.UserEvent]) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		log.Fatalf("Failed to get metadata %v", err)
	}

	factory := service.NewServiceFactory(ctx, meta.EventID, e.UserID())
	if factory.EventService().IsProcessed() {
		return nil
	}

	userOld, err := util.MapTo[model.User](e.OldValue.Data)
	userNew, err2 := util.MapTo[model.User](e.Value.Data)
	if err != nil || err2 != nil {
		log.Fatalf("Failed to parse transaction %v", err)
	}
	return functions.OnUserUpdated(factory, userOld, userNew, e.UpdateMask.Fields)
}

//goland:noinspection GoUnusedExportedFunction,GoUnusedParameter
func OnAccountUpdated(ctx context.Context, e data.FirestoreEvent[model.AccountEvent]) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		log.Fatalf("Failed to get metadata %v", err)
	}

	factory := service.NewServiceFactory(ctx, meta.EventID, e.UserID())
	if factory.EventService().IsProcessed() {
		return nil
	}

	accountOld, err := util.MapTo[model.Account](e.OldValue.Data)
	accountNew, err2 := util.MapTo[model.Account](e.Value.Data)
	if err != nil || err2 != nil {
		log.Fatalf("Failed to parse transaction %v", err)
	}
	return functions.OnAccountUpdated(factory, accountOld, accountNew, e.UpdateMask.Fields)
}
