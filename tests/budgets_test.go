package tests

import (
	"finize-functions.app/functions"
	"finize-functions.app/tests/fake"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOnBudgetDeleted(t *testing.T) {
	id1, _ := accountService.Create(fake.NewAccount(fake.NewAccountEvent("id1", "name", 50.0, util.Pointer("budget"))))
	id2, _ := accountService.Create(fake.NewAccount(fake.NewAccountEvent("id2", "name", 50.0, util.Pointer("budget"))))
	id3, _ := accountService.Create(fake.NewAccount(fake.NewAccountEvent("id3", "name", 50.0, nil)))

	_ = functions.OnBudgetDeleted(testFactory, fake.NewBudget(fake.NewBudgetEvent("budget", "name", 22.0)))

	a1, _ := accountService.FindByID(*id1, nil)
	a2, _ := accountService.FindByID(*id2, nil)
	a3, _ := accountService.FindByID(*id3, nil)
	assert.True(t, a1.Budget == nil)
	assert.True(t, a2.Budget == nil)
	assert.True(t, a3.Budget == nil)
}
