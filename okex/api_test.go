package okex

import (
	"log"
	"testing"
)

func TestExchangeRate(t *testing.T) {
	api := &Api{}
	rate := api.ExchangeRate()
	log.Println(rate[0].UsdCny)
}
