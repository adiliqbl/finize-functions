package tests

import (
	"finize-functions.app/util"
	"fmt"
	"log"
	"net/http"
	"os"
)

func ExitOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %v", msg, err)
		os.Exit(0)
	}
}

func ClearDatabase() {
	req, err := http.NewRequest("DELETE", "http://localhost:8080/emulator/v1/projects/demo-project/databases/(default)/documents", nil)
	_, err = util.MakeApiCall(req)
	if err != nil {
		log.Fatalf("Failed to clear database: %v", err)
	}
}
