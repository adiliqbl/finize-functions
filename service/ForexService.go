package service

import (
	"finize-functions.app/util"
	"fmt"
	"log"
	"net/http"
	"os"
)

type ForexService interface {
	GetRates(iso string) map[string]float64
}

type forexServiceImpl struct{}

func NewForexService() ForexService {
	return &forexServiceImpl{}
}

func (service *forexServiceImpl) GetRates(iso string) map[string]float64 {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.apilayer.com/exchangerates_data/latest?base=%s", iso), nil)
	req.Header["apikey"] = []string{os.Getenv("EXCHANGE_RATES_API")}

	result, err := util.MakeApiCall(req)
	if err != nil {
		log.Printf("%v", err)
	}

	if rates, ok := result["rates"].(map[string]interface{}); ok {
		values := map[string]float64{}
		for k := range rates {
			if v, ok := rates[k].(float64); ok {
				values[k] = v
			}
		}

		return values
	} else {
		return nil
	}
}
