package service

import (
	"finize-functions.app/data/model"
	"finize-functions.app/util"
	"fmt"
	"time"
)

type ExchangeRateService interface {
	GetRate(fromIso string, toIso string) *model.ExchangeRate
	SetRate(fromIso string, toIso string, rate float64) error
}

type exchangeRateServiceImpl struct {
	db FirestoreService[map[string]interface{}]
}

func exchangeRateDB() string {
	return "exchange-rates"
}

func exchangeRateDoc(fromIso string) string {
	return fmt.Sprintf("%s/%s", exchangeRateDB(), fromIso)
}

func NewExchangeRateService(db FirestoreService[map[string]interface{}]) ExchangeRateService {
	return &exchangeRateServiceImpl{db: db}
}

func (service *exchangeRateServiceImpl) GetRate(fromIso string, toIso string) *model.ExchangeRate {
	rates, err := service.db.Find(exchangeRateDoc(fromIso), nil)
	if err != nil || rates == nil {
		return nil
	}

	if rate, ok := (*rates)[toIso]; ok {
		if exRate, err := util.MapTo[model.ExchangeRate](rate); err == nil {
			return &exRate
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (service *exchangeRateServiceImpl) SetRate(fromIso string, toIso string, rate float64) error {
	fromRate, err := util.MapTo[map[string]interface{}](model.ExchangeRate{
		Rate: rate,
		Date: time.Now().UTC(),
	})
	if err != nil {
		return err
	}

	fromRates, err := service.db.Find(fromIso, nil)
	if err != nil || fromRates == nil {
		_, err = service.db.Create(exchangeRateDB(), &fromIso, map[string]interface{}{toIso: fromRate})
		if err != nil {
			return err
		}
	} else {
		_, err = service.db.Update(exchangeRateDoc(fromIso), map[string]interface{}{toIso: fromRate})
		if err != nil {
			return err
		}
	}

	toRate, err := util.MapTo[map[string]interface{}](model.ExchangeRate{
		Rate: 1 / rate,
		Date: time.Now().UTC(),
	})
	if err != nil {
		return err
	}

	toRates, err := service.db.Find(toIso, nil)
	if err != nil || toRates == nil {
		_, err = service.db.Create(exchangeRateDB(), &toIso, map[string]interface{}{fromIso: toRate})
		if err != nil {
			return err
		}
	} else {
		_, err = service.db.Update(exchangeRateDoc(toIso), map[string]interface{}{fromIso: toRate})
		if err != nil {
			return err
		}
	}

	return nil
}
