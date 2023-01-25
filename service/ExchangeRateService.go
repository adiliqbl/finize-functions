package service

import (
	"finize-functions.app/data"
	"finize-functions.app/data/model"
	"finize-functions.app/util"
	"fmt"
	"log"
	"time"
)

type ExchangeRateService interface {
	GetRate(fromIso string, toIso string) *model.ExchangeRate
	SetRates(iso string, rates map[string]float64) error
}

type exchangeRateServiceImpl struct {
	db FirestoreService[map[string]model.ExchangeRate]
}

func exchangeRateDB() string {
	return "exchange-rates"
}

func exchangeRateDoc(fromIso string) string {
	return fmt.Sprintf("%s/%s", exchangeRateDB(), fromIso)
}

func NewExchangeRateService(db FirestoreService[map[string]model.ExchangeRate]) ExchangeRateService {
	return &exchangeRateServiceImpl{db: db}
}

func (service *exchangeRateServiceImpl) GetRate(fromIso string, toIso string) *model.ExchangeRate {
	if fromIso == toIso {
		return &model.ExchangeRate{Rate: 1.0, Date: time.Now().UTC()}
	}

	rates, err := service.db.Find(exchangeRateDoc(fromIso), nil)
	if err != nil || rates == nil {
		return nil
	}

	if rate, ok := (*rates)[toIso]; ok {
		return &rate
	} else {
		return nil
	}
}

func (service *exchangeRateServiceImpl) SetRates(fromIso string, rates map[string]float64) error {
	return service.db.Batch(func() []data.DatabaseOperation {
		var ops []data.DatabaseOperation

		date := time.Now().UTC()
		fromIsoMap := map[string]model.ExchangeRate{}
		for iso, rate := range rates {
			fromIsoMap[iso] = model.ExchangeRate{Rate: rate, Date: date}

			isoRate, err := util.MapTo[map[string]interface{}](model.ExchangeRate{
				Rate: 1 / rate,
				Date: date,
			})
			if err != nil {
				log.Fatalf("ExchangeRates.SetRates – Failed to convert: %v", err)
			}

			ops = append(ops, data.DatabaseOperation{
				Ref: service.db.Doc(exchangeRateDoc(iso)),
				Data: map[string]interface{}{
					fromIso: isoRate,
				},
			})
		}

		fromRate, err := util.MapTo[map[string]interface{}](fromIsoMap)
		if err != nil {
			log.Fatalf("ExchangeRates.SetRates – Failed to convert: %v", err)
		}

		ops = append(ops, data.DatabaseOperation{
			Ref:  service.db.Doc(exchangeRateDoc(fromIso)),
			Data: fromRate,
		})

		return ops
	})
}
