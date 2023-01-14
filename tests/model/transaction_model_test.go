package model

import (
	"finize-functions.app/data/model"
	"finize-functions.app/tests/fakedata"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTransactionEventMapParsing(t *testing.T) {
	parsed := fakedata.NewTransactionEvent("id", "name", 50.0, util.Pointer(25.0), time.Now(), util.Pointer("to"), nil, util.Pointer("budget"))
	event, _ := util.MapTo[model.TransactionEvent](fakedata.NewTransactionEventMap(parsed))
	want, _ := util.MapTo[map[string]interface{}](parsed)
	got, _ := util.MapTo[map[string]interface{}](event)
	assert.EqualValues(t, want, got)

	parsed = fakedata.NewTransactionEvent("id", "name", 50.0, nil, time.Now(), util.Pointer("to"), util.Pointer("from"), nil)
	event, _ = util.MapTo[model.TransactionEvent](fakedata.NewTransactionEventMap(parsed))
	want, _ = util.MapTo[map[string]interface{}](parsed)
	got, _ = util.MapTo[map[string]interface{}](event)
	assert.Equal(t, want, got)
}

func TestTransactionFromEvent(t *testing.T) {
	date := time.Now()
	got := fakedata.NewTransactionEvent("id", "name", 50.0, util.Pointer(25.0), date, util.Pointer("to"), nil, util.Pointer("budget"))
	want, _ := util.MapTo[model.Transaction](got)

	assert.Equal(t, "id", want.ID)
	assert.Equal(t, "name", want.Name)
	assert.Equal(t, 50.0, want.Amount)
	assert.Equal(t, 25.0, *want.AmountValue)
	assert.Equal(t, "to", *want.AccountTo)
	assert.True(t, want.AccountFrom == nil)
	assert.Equal(t, "budget", *want.Budget)
	assert.Equal(t, "CURR", want.Currency)
	assert.Equal(t, []string{"One", "Two"}, want.Category)

	got = fakedata.NewTransactionEvent("id", "name", 50.0, nil, date, util.Pointer("to"), util.Pointer("from"), nil)
	want, _ = util.MapTo[model.Transaction](got)

	assert.True(t, want.AmountValue == nil)
	assert.Equal(t, "to", *want.AccountTo)
	assert.Equal(t, "from", *want.AccountFrom)
	assert.True(t, want.Budget == nil)
}
