package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExchangeRateService(t *testing.T) {
	_, _ = database.Doc("exchange-rates/PKR").Delete(context.Background())
	_, _ = database.Doc("exchange-rates/GBP").Delete(context.Background())
	_, _ = database.Doc("exchange-rates/USD").Delete(context.Background())

	_ = exchangeRateService.SetRates("GBP", map[string]float64{
		"USD": 20,
	})
	_ = exchangeRateService.SetRates("PKR", map[string]float64{
		"USD":   1 / 200.0,
		"GBP":   1 / 220.0,
		"one":   10,
		"two":   10,
		"three": 10,
	})

	pkrToUsd := exchangeRateService.GetRate("PKR", "USD")
	usdToPkr := exchangeRateService.GetRate("USD", "PKR")
	gbpToUsd := exchangeRateService.GetRate("GBP", "USD")
	gbpToPkr := exchangeRateService.GetRate("GBP", "PKR")
	pkrToThree := exchangeRateService.GetRate("PKR", "three")
	notFound1 := exchangeRateService.GetRate("CNY", "GBP")
	notFound2 := exchangeRateService.GetRate("EUR", "GBP")

	assert.True(t, pkrToUsd != nil)
	assert.True(t, usdToPkr != nil)
	assert.True(t, gbpToUsd != nil)
	assert.True(t, gbpToPkr != nil)
	assert.True(t, pkrToThree != nil)
	assert.True(t, notFound1 == nil)
	assert.True(t, notFound2 == nil)
	assert.Equal(t, 200.0, usdToPkr.Rate)
	assert.Equal(t, 1/200.0, pkrToUsd.Rate)
	assert.Equal(t, 20.0, gbpToUsd.Rate) // Update GBP<>PKR without removing USD rate
	assert.Equal(t, 220.0, gbpToPkr.Rate)
	assert.Equal(t, 10.0, pkrToThree.Rate)
}
