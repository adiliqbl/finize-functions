package service

import (
	"finize-functions.app/tests/fake"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateBudget(t *testing.T) {
	id, err := budgetService.Create(fake.NewBudget(fake.NewBudgetEvent("", "name", 50.0)))

	assert.Nil(t, err)
	assert.True(t, !util.NullOrEmpty(id))
}

func TestGetBudget(t *testing.T) {
	testBudget := fake.NewBudget(fake.NewBudgetEvent("", "name", 50.0))
	id, _ := budgetService.Create(testBudget)
	testBudget.ID = *id

	budget, _ := budgetService.FindByID(testBudget.ID, nil)
	assert.Equal(t, testBudget, *budget)
}
