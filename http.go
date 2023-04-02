package finize_functions

import (
	"context"
	"encoding/json"
	"finize-functions.app/functions"
	"finize-functions.app/service"
	"finize-functions.app/util"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func init() {
	if err := service.InitFirestore(context.Background()); err != nil {
		log.Fatalf("Failed to initialize Firestore %v", err)
	}
}

//goland:noinspection GoUnusedExportedFunction
func GetExchangeRate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	fromIso := query.Get("from")
	toIso := query.Get("to")

	refresh := false
	if isRefresh, err := strconv.ParseBool(query.Get("refresh")); err == nil {
		refresh = isRefresh
	}

	if fromIso == "" || toIso == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("From/To Iso Missing"))
		return
	}

	factory := service.NewServiceFactory(context.Background(), "", "")
	if rate, err := functions.GetExchangeRate(factory, fromIso, toIso, refresh); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Failed to get exchange rates: %v", err)))
	} else {
		json, err := json.Marshal(rate)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
}

//goland:noinspection GoUnusedExportedFunction
func ProcessRecurringTasks(w http.ResponseWriter, _ *http.Request) {
	factory := service.NewServiceFactory(context.Background(), "", "")
	err := functions.ProcessRecurringTasks(factory, util.NewClock())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Processing Failed: %v", err)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Processing Completed")
	}
}
