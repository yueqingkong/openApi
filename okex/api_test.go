package okex

import (
	"log"
	"testing"
)

// go test -v -run TestExchangeRate
func TestExchangeRate(t *testing.T) {
	api := &Api{}
	rate := api.ExchangeRate()
	log.Println(rate[0].UsdCny)
}
