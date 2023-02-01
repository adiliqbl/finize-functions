package functions

import (
	"finize-functions.app/data/model"
	services "finize-functions.app/service"
	"log"
	"time"
)

func GetExchangeRate(factory services.Factory, iso string, toIso string) (*model.ExchangeRate, error) {
	if iso == toIso {
		return &model.ExchangeRate{Rate: 1.0, Date: time.Now().UTC()}, nil
	}

	exchange := factory.ExchangeRateService()

	if rate := exchange.GetRate(iso, toIso); rate != nil {
		if rate.Date.Sub(time.Now().UTC()).Hours()*24 <= 48 { // If less than 2 days, return cached
			return rate, nil
		}
	}

	rates := factory.ForexService().GetRates(iso)
	if rates == nil {
		log.Fatalf("Failed to get Forex rates for %s", iso)
	}

	delete(rates, iso)
	err := exchange.SetRates(iso, rates)
	if rate, ok := rates[toIso]; ok {
		return &model.ExchangeRate{Rate: rate, Date: time.Now().UTC()}, nil
	} else {
		return nil, err
	}
}
