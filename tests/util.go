package tests

import (
	"finize-functions.app/util"
	"fmt"
	"os"
)

func ResetFirestore() {
	_, err := util.GetApiCall("http://localhost:8080/emulator/v1/projects/firestore-emulator-example/databases/(default)/documents")
	if err != nil {
		fmt.Printf("Failed to reset Firestore\n")
	}
}

func ExitOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %v", msg, err)
		os.Exit(0)
	}
}
