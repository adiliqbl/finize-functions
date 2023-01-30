package finize_functions

import (
	"context"
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
func ProcessRecurringTasks(ctx context.Context) error {
	factory := service.NewServiceFactory(ctx, "", "")
	return functions.ProcessRecurringTasks(factory, util.NewClock())
}
