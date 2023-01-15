package model

import (
	"finize-functions.app/data/model"
	"finize-functions.app/tests/fake"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountEventMapParsing(t *testing.T) {
	want := fake.NewAccountEvent("id", "name", 50.0, nil)
	got, _ := util.MapTo[model.AccountEvent](fake.NewAccountEventMap(want))
	assert.Equal(t, want, got)

	want = fake.NewAccountEvent("id", "name", 50.0, util.Pointer("budget"))
	got, _ = util.MapTo[model.AccountEvent](fake.NewAccountEventMap(want))
	assert.Equal(t, want, got)
}

func TestAccountFromEvent(t *testing.T) {
	got := fake.NewAccountEvent("id", "name", 50.0, nil)
	want, _ := util.MapTo[model.Account](got)

	assert.Equal(t, "id", want.ID)
	assert.Equal(t, "name", want.Name)
	assert.Equal(t, 50.0, want.Balance)
	assert.Equal(t, "type", want.Type)
	assert.Equal(t, 50.0, want.Balance)
	assert.Equal(t, "CURR", want.Currency)
	assert.Equal(t, nil, want.Budget)

	got = fake.NewAccountEvent("id", "name", 50.0, util.Pointer("budget"))
	want, _ = util.MapTo[model.Account](got)
	assert.Equal(t, "budget", want.Budget)
}
