package tests

import (
	"finize-functions.app/util"
	"fmt"
)

func ResetFirestore() {
	_, err := util.GetApiCall("http://localhost:8080/emulator/v1/projects/firestore-emulator-example/databases/(default)/documents")
	if err != nil {
		fmt.Printf("Failed to reset Firestore\n")
	}
}
