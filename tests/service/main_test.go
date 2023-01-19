package service

import (
	"context"
	"finize-functions.app/tests/fake"
	"fmt"
	"os"
	"testing"
)

func setup() {
	err := fake.InitTestFirestore(context.Background())
	if err != nil {
		fmt.Printf("\033[1;33m%s\033[0m", "> Failed to initialize Firestore\n")
		os.Exit(0)
		return
	}
	fmt.Printf("\033[1;33m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	fmt.Printf("\033[1;33m%s\033[0m", "> Teardown completed\n")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
