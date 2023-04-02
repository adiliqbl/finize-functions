package tests

import (
	"context"
	"finize-functions.app/functions"
	"finize-functions.app/service"
	"finize-functions.app/tests/fake"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetExchangeRateWithCache(t *testing.T) {
	_, _ = testFactory.Firestore().Doc(service.ExchangeRateDoc("USD")).Delete(context.Background())

	fake.MockForexService.On("GetRates", "USD").Return(map[string]float64{
		"PKR": 200,
		"USD": 1,
		"one": 25,
	})
	_ = exchangeRateService.SetRates("USD", map[string]float64{
		"PKR": 100,
	})

	assert.Equal(t, 100.0, exchangeRateService.GetRate("USD", "PKR").Rate)

	rate, _ := functions.GetExchangeRate(testFactory, "USD", "PKR", false)
	assert.True(t, rate != nil)
	assert.Equal(t, 100.0, rate.Rate)

	assert.True(t, exchangeRateService.GetRate("USD", "one") == nil)
	assert.Equal(t, 100.0, exchangeRateService.GetRate("USD", "PKR").Rate)

	// Ignoring cached values
	rate, _ = functions.GetExchangeRate(testFactory, "USD", "PKR", true)
	assert.True(t, rate != nil)
	assert.Equal(t, 200.0, rate.Rate)

	assert.Equal(t, 25.0, exchangeRateService.GetRate("USD", "one").Rate)
	assert.Equal(t, 200.0, exchangeRateService.GetRate("USD", "PKR").Rate)
}

func TestGetExchangeRateWithoutCache(t *testing.T) {
	_, _ = testFactory.Firestore().Doc(service.ExchangeRateDoc("USD")).Delete(context.Background())

	fake.MockForexService.On("GetRates", "USD").Return(map[string]float64{
		"PKR": 200,
		"USD": 1,
		"one": 25,
	})
	_ = exchangeRateService.SetRates("USD", map[string]float64{
		"PKR": 100,
	})

	assert.True(t, exchangeRateService.GetRate("USD", "one") == nil)
	assert.True(t, exchangeRateService.GetRate("USD", "two") == nil)

	rate, _ := functions.GetExchangeRate(testFactory, "USD", "one", false)
	assert.True(t, rate != nil)
	assert.Equal(t, 25.0, rate.Rate)

	assert.Equal(t, 25.0, exchangeRateService.GetRate("USD", "one").Rate)
	assert.Equal(t, 200.0, exchangeRateService.GetRate("USD", "PKR").Rate)
}
