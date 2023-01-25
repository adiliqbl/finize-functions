package finize_functions

import (
	"context"
	"encoding/json"
	"finize-functions.app/functions"
	"finize-functions.app/service"
	"fmt"
	"log"
	"net/http"
)

func init() {
	if err := service.InitFirestore(context.Background()); err != nil {
		log.Fatalf("Failed to initialize Firestore %v", err)
	}
}

//goland:noinspection GoUnusedExportedFunction
func GetExchangeRate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fromIso := r.URL.Query().Get("from")
	toIso := r.URL.Query().Get("to")

	if fromIso == "" || toIso == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("From/To Iso Missing"))
		return
	}

	factory := service.NewServiceFactory(context.Background(), "", "")
	if rate, err := functions.GetExchangeRate(factory, fromIso, toIso); err != nil {
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
