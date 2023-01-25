package tests

import (
	"context"
	"finize-functions.app/functions"
	"finize-functions.app/tests/fake"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetExchangeRate(t *testing.T) {
	fake.MockForexService.On("GetRates", "USD").Return(map[string]float64{
		"PKR": 200,
		"USD": 1,
		"one": 1,
		"two": 2,
	})

	_, _ = database.Doc("exchange-rates/USD").Delete(context.Background())
	rate := exchangeRateService.GetRate("USD", "PKR")
	assert.True(t, rate == nil)

	rate, _ = functions.GetExchangeRate(testFactory, "USD", "PKR")
	assert.True(t, rate != nil)
	assert.Equal(t, 200.0, rate.Rate)

	usdToPkr := exchangeRateService.GetRate("USD", "PKR")
	usdToTwo := exchangeRateService.GetRate("USD", "two")
	usdToThree := exchangeRateService.GetRate("USD", "three")

	assert.True(t, usdToPkr != nil)
	assert.True(t, usdToPkr != nil)
	assert.True(t, usdToThree == nil)
	assert.Equal(t, 200.0, usdToPkr.Rate)
	assert.Equal(t, 2.0, usdToTwo.Rate)
}
