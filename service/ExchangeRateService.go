package service

import (
	"finize-functions.app/util"
	"fmt"
)

type ExchangeRateService interface {
	GetRate(fromIso string, toIso string) *float64
	SetRate(fromIso string, toIso string, rate float64) error
}

type exchangeRateServiceImpl struct {
	db FirestoreService[map[string]interface{}]
}

func exchangeRateDB() string {
	return "exchange-rates"
}

func exchangeRateDoc(iso string) string {
	return fmt.Sprintf("%s/%s", exchangeRateDB(), iso)
}

func NewExchangeRateService(db FirestoreService[map[string]interface{}]) ExchangeRateService {
	return &exchangeRateServiceImpl{db: db}
}

func (service *exchangeRateServiceImpl) GetRate(fromIso string, toIso string) *float64 {
	rates, err := service.db.Find(exchangeRateDoc(fromIso), nil)
	if err != nil || rates == nil {
		return nil
	}

	if rate, ok := (*rates)[toIso]; ok {
		return util.Pointer(rate.(float64))
	} else {
		return nil
	}
}

func (service *exchangeRateServiceImpl) SetRate(fromIso string, toIso string, rate float64) error {
	fromRates, err := service.db.Find(fromIso, nil)
	if err != nil || fromRates == nil {
		_, err = service.db.Create(exchangeRateDB(), &fromIso, map[string]interface{}{toIso: rate})
		if err != nil {
			return err
		}
	} else {
		_, err = service.db.Update(exchangeRateDoc(fromIso), map[string]interface{}{toIso: rate})
		if err != nil {
			return err
		}
	}

	toRates, err := service.db.Find(fromIso, nil)
	if err != nil || toRates == nil {
		_, err = service.db.Create(exchangeRateDB(), &toIso, map[string]interface{}{fromIso: 1 / rate})
		if err != nil {
			return err
		}
	} else {
		_, err = service.db.Update(exchangeRateDoc(toIso), map[string]interface{}{fromIso: 1 / rate})
		if err != nil {
			return err
		}
	}

	return nil
}
