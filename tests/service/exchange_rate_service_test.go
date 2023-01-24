package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExchangeRateService(t *testing.T) {
	err := exchangeRateService.SetRate("PKR", "USD", 120)
	assert.Nil(t, err)

	toUSD := exchangeRateService.GetRate("PKR", "USD")
	toPKR := exchangeRateService.GetRate("USD", "PKR")
	fromGBP := exchangeRateService.GetRate("GBP", "PKR")
	toGBP := exchangeRateService.GetRate("PKR", "GBP")

	assert.True(t, toUSD != nil)
	assert.True(t, toPKR != nil)
	assert.True(t, fromGBP == nil)
	assert.True(t, toGBP == nil)
	assert.Equal(t, 120.0, *toUSD)
	assert.Equal(t, 1/120.0, *toPKR)
}
