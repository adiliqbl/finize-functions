package tests

import (
	"context"
	services "finize-functions.app/service"
	"finize-functions.app/tests/fake"
	"fmt"
	"os"
	"testing"
)

var testFactory services.Factory
var accountService services.AccountService
var transactionService services.TransactionService
var exchangeRateService services.ExchangeRateService

func setup() {
	setupFirestore()

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

func setupFirestore() {
	if testFactory != nil {
		return
	}

	fake.InitTestFirestore()

	factory := fake.NewServiceFactory(context.Background(), "")
	userID, err := factory.UserService().Create(fake.NewUser(fake.NewUserEvent("", "name", "email@test.com")))
	ExitOnError(err, "Failed to create user")

	testFactory = fake.NewServiceFactory(context.Background(), *userID)
	accountService = testFactory.AccountService()
	transactionService = testFactory.TransactionService()
	exchangeRateService = testFactory.ExchangeRateService()
	_ = testFactory.ForexService()

	ClearDatabase()
}
