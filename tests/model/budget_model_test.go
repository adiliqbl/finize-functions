package model

import (
	"finize-functions.app/data/model"
	"finize-functions.app/tests/fakedata"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBudgetEventMapParsing(t *testing.T) {
	want := fakedata.NewBudgetEvent("id", "name", 50.0, 20.0)
	got, _ := util.MapTo[model.BudgetEvent](fakedata.NewBudgetEventMap(want))

	assert.Equal(t, want, got)
}

func TestBudgetFromEvent(t *testing.T) {
	got := fakedata.NewBudgetEvent("id", "name", 50.0, 20.0)
	want, _ := util.MapTo[model.Budget](got)

	assert.Equal(t, "id", want.ID)
	assert.Equal(t, "name", want.Name)
	assert.Equal(t, 50.0, want.Limit)
	assert.Equal(t, 20.0, want.Spent)
}
