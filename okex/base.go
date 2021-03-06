package okex

import (
	"github.com/yueqingkong/openApi/conset"
	"github.com/yueqingkong/openApi/db"
	"github.com/yueqingkong/openApi/util"
	"log"
	"strconv"
	"strings"
	"time"
)

type Base struct {
	*Api
	*Ws
}

func (self *Base) Init(strings []string) {
	if len(strings) >= 3 {
		self.Api = &Api{}
		self.Ws = &Ws{}

		self.InitApiKeys(strings[0], strings[1], strings[2])
		self.InitWsKeys(strings[0], strings[1], strings[2])
	}
}

func (self *Base) ccy(symbol conset.SYMBOL) string {
	var s string

	switch symbol {
	case conset.BTC_USD, conset.BTC_USDT:
		s = "BTC"
	case conset.ETH_USD, conset.ETH_USDT:
		s = "ETH"
	}
	return s
}

func (self *Base) instId(symbol conset.SYMBOL) string {
	var s string
	s = self.Symbol(symbol)
	s = strings.ReplaceAll(s, "_", "-")
	s = strings.ToUpper(s)
	s += "-SWAP"
	return s
}

func (self *Base) before(start time.Time) string {
	var s string
	if !start.IsZero() {
		s = util.UnixMillis(start)
	}
	return s
}

func (self *Base) Plat() conset.PLAT {
	return conset.OKEX
}

func (self *Base) Symbol(symbol conset.SYMBOL) string {
	//var s string
	//switch symbol {
	//case conset.BTC_USD:
	//	s = "btc_usd"
	//case conset.BTC_USDT:
	//	s = "btc_usdt"
	//case conset.ETH_USD:
	//	s = "eth_usd"
	//case conset.ETH_USDT:
	//	s = "eth_usdt"
	//}
	return db.SymbolToString(symbol)
}

func (self *Base) Period(period conset.PERIOD) string {
	var s string
	switch period {
	case conset.SPOT:
		s = "spot"
	case conset.SWAP:
		s = "swap"
	case conset.WEEK:
		s = "week"
	case conset.WEEK_NEXT:
		s = "week_next"
	case conset.QUARTER:
		s = "quarter"
	case conset.QUARTER_NEXT:
		s = "quarter_next"
	}
	return s
}

func (self *Base) Times(times conset.TIMES) string {
	var s string
	switch times {
	case conset.MIN_15:
		s = "15m"
	case conset.MIN_30:
		s = "30m"
	case conset.H_1:
		s = "1H"
	case conset.H_6:
		s = "6H"
	case conset.H_12:
		s = "12H"
	case conset.D_1:
		s = "1D"
	}
	return s
}

func (self *Base) TdMode(period conset.PERIOD) string {
	var s string
	switch period {
	case conset.SPOT:
		s = "cash"
	case conset.SWAP:
		s = "isolated"
	default:
		s = "isolated"
	}
	return s
}

func (self *Base) Side(direct conset.OPERATION) (string, string) {
	var side string
	var posSide string

	switch direct {
	case conset.BUY_HIGH:
		side = conset.BUY
		posSide = conset.LONG
	case conset.BUY_LOW:
		side = conset.SELL
		posSide = conset.SHORT
	case conset.SELL_HIGH:
		side = conset.SELL
		posSide = conset.LONG
	case conset.SELL_LOW:
		side = conset.BUY
		posSide = conset.SHORT
	}
	return side, posSide
}

func (self *Base) Pull(symbol conset.SYMBOL, times conset.TIMES, start time.Time) bool {
	candles := self.Candles(self.instId(symbol), self.Times(times), self.before(start))

	if len(candles) == 0 {
		log.Printf("Pull : ????????????")
		return true
	}

	for k, value := range candles {
		var arr = value
		var timetamp = util.Int64(arr[0])
		var open, _ = strconv.ParseFloat(arr[1], 32)
		var close, _ = strconv.ParseFloat(arr[4], 32)
		var high, _ = strconv.ParseFloat(arr[2], 32)
		var low, _ = strconv.ParseFloat(arr[3], 32)
		var volume, _ = strconv.ParseFloat(arr[5], 32)

		if k != 0 { // ???????????????????????????K????????????
			coin := &db.Coin{}
			if err := coin.Create(self.Plat(), symbol, times, float32(open), float32(close), float32(high), float32(low), float32(volume), timetamp); err != nil {
				log.Printf("Create err: %+v", err)
			}
		}
	}
	return false
}

func (self *Base) Price(symbol conset.SYMBOL) float32 {
	tickers := self.Ticker(self.instId(symbol))
	if len(tickers) == 0 {
		return 0.0
	}

	return util.Float32(tickers[0].Last)
}

func (self *Base) Balance(symbol conset.SYMBOL) float32 {
	bs := self.balance(self.ccy(symbol))
	if len(bs) == 0 {
		return 0.0
	}

	return util.Float32(bs[0].TotalEq)
}

func (self *Base) SetLeverage(symbol conset.SYMBOL, period conset.PERIOD, direct conset.OPERATION, level string) bool {
	_, poside := self.Side(direct)

	bs := self.setLeverage(self.instId(symbol), level, self.TdMode(period), poside)
	if len(bs) == 0 {
		return false
	}

	return true
}

func (self *Base) Orders(symbol conset.SYMBOL, period conset.PERIOD, direct conset.OPERATION, price, sz float32) bool {
	side, poside := self.Side(direct)
	price = priceLimit(direct, price)

	orders := self.Order(self.instId(symbol), self.TdMode(period), side, poside, price, sz)
	if len(orders) == 0 {
		return false
	}

	return orders[0].SCode == "0"
}

// limit ????????????
func priceLimit(direct conset.OPERATION, price float32) float32 {
	rate := float32(0.01)
	switch direct {
	case conset.BUY_HIGH, conset.SELL_LOW:
		price = price * (1.0 + rate)
	case conset.SELL_HIGH, conset.BUY_LOW:
		price = price * (1.0 - rate)
	}
	return price
}

func (self *Base) SubscribeTickers(symbols []conset.SYMBOL, f func(conset.SYMBOL, float32)) {
	instIds := make([]string, 0)
	for _, symbol := range symbols {
		instIds = append(instIds, self.instId(symbol))
	}

	self.WsTickers(instIds, f)
}
