package okex

import (
	"fmt"
	"github.com/yueqingkong/openApi/conset"
	"log"
	"testing"
)

// go test -v -run TestUsdCny
func TestUsdCny(t *testing.T) {
	base := &Base{}
	rate := base.UsdCny()
	log.Println(rate)
}

// go test -v -run TestPull
func TestPull(t *testing.T) {
	base := &Base{}
	base.Init([]string{"", "", ""})
	inst := base.instIds(conset.BTC, conset.USD, conset.SWAP)
	candles := base.Candles(inst, "15m", "1675068300000")
	t.Log(candles)
}

// go test -v -run TestPrice
func TestPrice(t *testing.T) {
	base := &Base{}
	base.Init([]string{"", "", ""})

	pr := base.Price(conset.BTC, conset.USD, conset.SWAP)
	t.Log(pr)
}

// go test -v -run TestFundingRate
func TestFundingRate(t *testing.T) {
	base := &Base{}
	base.Init([]string{"", "", ""})

	rate, nextRate := base.FundingRate(conset.BTC, conset.USD)
	t.Log(fmt.Sprintf("%.5f",rate), nextRate)
}

// go test -v -run TestOrders
func TestOrders(t *testing.T) {
	base := &Base{}
	base.Init([]string{"", "", ""})

	b, orderid := base.Orders(conset.BTC, conset.USDT, conset.SWAP, conset.SELL_HIGH, 23300, 1.0)
	t.Log(b, orderid)
}

// go test -v -run TestOrderInfos
func TestOrderInfos(t *testing.T) {
	base := &Base{}
	base.Init([]string{"", "", ""})

	b, info := base.OrderInfos(conset.BCH, conset.USDT, conset.SWAP, "")
	t.Log(b, info)
}

// go test -v -run TestSubscribeTickers
func TestSubscribeTickers(t *testing.T) {
	base := &Base{}
	base.Init([]string{"", "", ""})

	ccies := [2]conset.CCY{conset.BTC, conset.USDT}
	base.SubscribeTickers([][2]conset.CCY{ccies}, conset.SWAP, func(bas conset.CCY, quote conset.CCY, f float32) {
		t.Log(base.instIds(bas, quote, conset.SWAP), f)
	})

	c := make(chan bool)
	<-c
}
