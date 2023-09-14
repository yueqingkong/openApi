package okex

import (
	"encoding/json"
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

// go test -v -run TestInstrument
func TestInstrument(t *testing.T) {
	base := &Base{}
	ins := base.Instrument(conset.MARGIN, "", "")
	log.Println(ins)
}

// go test -v -run TestSupportCoin
func TestSupportCoin(t *testing.T) {
	base := &Base{}
	body := base.SupportCoin()
	t.Log(body)
}

// go test -v -run TestPull
func TestPull(t *testing.T) {
	base := &Base{}
	base.Init([]string{"", "", ""})
	inst := base.InstId(conset.BTC, conset.USD, conset.SWAP)
	candles := base.Candles(inst, "15m", "1675068300000")
	t.Log(candles)
}

// go test -v -run TestPrice
func TestPrice(t *testing.T) {
	base := &Base{}
	base.Init([]string{"", "", ""})

	pr := base.Price(conset.ETH, conset.USDT, conset.MARGIN)
	t.Log(pr)
}

// go test -v -run TestFundingRate
func TestFundingRate(t *testing.T) {
	base := &Base{}
	base.Init([]string{"", "", ""})

	rate, nextRate := base.FundingRate(conset.BTC, conset.USD)
	t.Log(fmt.Sprintf("%.5f", rate), nextRate)
}

// go test -v -run TestTakerVolume
func TestTakerVolume(t *testing.T) {
	base := &Base{}
	base.Init([]string{"", "", ""})

	res := base.TakerVolume(conset.BTC, conset.SPOT, 1679155200000, 1694448000000, conset.D_1)
	t.Log(fmt.Sprintf("%v", res))
}

// go test -v -run TestLoanRatio
func TestLoanRatio(t *testing.T) {
	base := &Base{}
	base.Init([]string{"", "", ""})

	res := base.LoanRatio(conset.BTC, 1679155200000, 1694448000000, conset.D_1)
	t.Log(fmt.Sprintf("%v", res))
}

// go test -v -run TestSwapAccountRatio
func TestSwapAccountRatio(t *testing.T) {
	base := &Base{}
	base.Init([]string{"", "", ""})

	res := base.SwapAccountRatio(conset.BTC, 1679155200000, 1694448000000, conset.D_1)
	t.Log(fmt.Sprintf("%v", res))
}

// go test -v -run TestInterestVolume
func TestInterestVolume(t *testing.T) {
	base := &Base{}
	base.Init([]string{"", "", ""})

	res := base.InterestVolume(conset.BTC, 0, 1694448000000, conset.D_1)
	t.Log(fmt.Sprintf("%v", res))
}

// go test -v -run TestODInfo
func TestODInfo(t *testing.T) {
	base := &Base{}
	base.Init([]string{"", "", ""})

	b, info := base.OrderInfo(conset.ETH, conset.USDT, conset.MARGIN, "611746420433305600")
	t.Log(b, info)
}

// go test -v -run TestOrder
func TestOrder(t *testing.T) {
	base := &Base{}
	base.Init([]string{"", "", ""})

	b, orderid := base.Order(conset.ETH, conset.USDT, conset.SPOT, conset.SELL_HIGH, 1841, 0.01)
	t.Log(b, orderid)
}

// go test -v -run TestSubscribeTickers
func TestSubscribeTickers(t *testing.T) {
	base := &Base{}
	base.Init([]string{"", "", ""})

	base.SubscribeTickers(CHANNEL_FUNDING_RATE, []string{"BTC-USDT-SWAP"}, func(bas conset.CCY, quote conset.CCY, i interface{}) {
		bytes, _ := json.Marshal(i)
		body := &FundingRateBody{}
		json.Unmarshal(bytes, body)

		t.Log(base.InstId(bas, quote, conset.SWAP), body)
	})

	c := make(chan bool)
	<-c
}
