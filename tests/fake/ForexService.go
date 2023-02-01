package fake

import (
	"finize-functions.app/service"
	"github.com/stretchr/testify/mock"
)

var MockForexService *ForexService

type ForexService struct {
	mock.Mock
}

func NewForexService() service.ForexService {
	if MockForexService == nil {
		MockForexService = new(ForexService)
	}
	return MockForexService
}

func (mock *ForexService) GetRates(iso string) map[string]float64 {
	args := mock.Called(iso)
	return args.Get(0).(map[string]float64)
}
