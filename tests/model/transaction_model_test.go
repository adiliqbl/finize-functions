package model

import (
	"finize-functions.app/data/model"
	"finize-functions.app/tests/fake"
	"finize-functions.app/util"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTransactionEventMapParsing(t *testing.T) {
	parsed := fake.NewTransactionEvent("id", "name", 50.0, 50.0, 25.0, time.Now(), util.Pointer("to"), nil, util.Pointer("budget"))
	event, _ := util.MapTo[model.TransactionEvent](fake.NewTransactionEventMap(parsed))
	want, _ := util.MapTo[map[string]interface{}](parsed)
	got, _ := util.MapTo[map[string]interface{}](event)
	assert.EqualValues(t, want, got)

	parsed = fake.NewTransactionEvent("id", "name", 50.0, 50.0, 0.0, time.Now(), util.Pointer("to"), util.Pointer("from"), nil)
	event, _ = util.MapTo[model.TransactionEvent](fake.NewTransactionEventMap(parsed))
	want, _ = util.MapTo[map[string]interface{}](parsed)
	got, _ = util.MapTo[map[string]interface{}](event)
	assert.Equal(t, want, got)
}

func TestTransactionFromEvent(t *testing.T) {
	date := time.Now()
	got := fake.NewTransactionEvent("id", "name", 50.0, 50.0, 25.0, date, util.Pointer("to"), nil, util.Pointer("budget"))
	want, _ := util.MapTo[model.Transaction](got)

	assert.Equal(t, "id", want.ID)
	assert.Equal(t, "name", want.Name)
	assert.Equal(t, 50.0, want.Amount.Amount)
	assert.Equal(t, "CURR", want.Amount.Currency)
	assert.Equal(t, 25.0, want.AmountTo.Amount)
	assert.Equal(t, "CURR", want.AmountTo.Currency)
	assert.Equal(t, "to", util.ValueOrNull(want.AccountTo))
	assert.True(t, want.AccountFrom == nil)
	assert.Equal(t, "budget", util.ValueOrNull(want.Budget))
	assert.Equal(t, []string{"One", "Two"}, want.Categories)

	got = fake.NewTransactionEvent("id", "name", 50.0, 50.0, 0.0, date, util.Pointer("to"), util.Pointer("from"), nil)
	want, _ = util.MapTo[model.Transaction](got)

	assert.True(t, want.AmountTo != nil)
	assert.True(t, want.AmountFrom != nil)
	assert.Equal(t, "to", util.ValueOrNull(want.AccountTo))
	assert.Equal(t, "from", util.ValueOrNull(want.AccountFrom))
	assert.True(t, want.Budget == nil)
}
