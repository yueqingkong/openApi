package okex

import (
	"github.com/yueqingkong/openApi/conset"
	"github.com/yueqingkong/openApi/db"
	"github.com/yueqingkong/openApi/util"
	"log"
	"strconv"
	"time"
)

type Base struct {
	*Api
}

func (self *Base) Init(strings []string) {
	if len(strings) >= 3 {
		self.Api = &Api{}
		self.InitKeys(strings[0], strings[1], strings[2])
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

func (self *Base) instId(symbol conset.SYMBOL, period conset.PERIOD) string {
	var s string
	switch period {
	case conset.SPOT:
		switch symbol {
		case conset.BTC_USD, conset.BTC_USDT:
			s = "BTC-USDT"
		case conset.ETH_USD, conset.ETH_USDT:
			s = "ETH-USDT"
		}
	case conset.SWAP:
		switch symbol {
		case conset.BTC_USD:
			s = "BTC-USD-SWAP"
		case conset.BTC_USDT:
			s = "BTC-USDT-SWAP"
		case conset.ETH_USD:
			s = "ETH-USD-SWAP"
		case conset.ETH_USDT:
			s = "ETH-USDT-SWAP"
		}
	case conset.WEEK:
		switch symbol {
		case conset.BTC_USD:

		}
	case conset.WEEK_NEXT:
		switch symbol {
		case conset.BTC_USD:

		}
	case conset.QUARTER:
		switch symbol {
		case conset.BTC_USD:

		}
	case conset.QUARTER_NEXT:
		switch symbol {
		case conset.BTC_USD:

		}
	}
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
	var s string
	switch symbol {
	case conset.BTC_USD:
		s = "btc_usd"
	case conset.BTC_USDT:
		s = "btc_usdt"
	case conset.ETH_USD:
		s = "eth_usd"
	case conset.ETH_USDT:
		s = "eth_usdt"
	}
	return s
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

func (self *Base) Side(period conset.PERIOD, direct conset.OPERATION) (string, string) {
	var side string
	var posSide string

	switch period {
	case conset.SPOT:
		if direct == conset.BUY_HIGH || direct == conset.SELL_HIGH {
			side = conset.BUY
		} else {
			side = conset.SELL
		}
	case conset.SWAP:
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
	}
	return side, posSide
}

func (self *Base) Pull(symbol conset.SYMBOL, period conset.PERIOD, times conset.TIMES, start time.Time) bool {
	candles := self.Candles(self.instId(symbol, period), self.Times(times), self.before(start))

	if len(candles) == 0 {
		log.Printf("Pull : 同步完成")
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

		if k != 0 { // 最近时间一条有效的K线不保存
			coin := &db.Coin{}
			if err := coin.Create(self.Plat(), symbol, period, times, float32(open), float32(close), float32(high), float32(low), float32(volume), timetamp); err != nil {
				log.Printf("Create err: %+v", err)
			}
		}
	}
	return false
}

func (self *Base) Price(symbol conset.SYMBOL, period conset.PERIOD) float32 {
	tickers := self.Ticker(self.instId(symbol, period))
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
	_, poside := self.Side(period, direct)

	bs := self.setLeverage(self.instId(symbol, period), level, self.TdMode(period), poside)
	if len(bs) == 0 {
		return false
	}

	return true
}

func (self *Base) Orders(symbol conset.SYMBOL, period conset.PERIOD, direct conset.OPERATION, price, sz float32) bool {
	side, poside := self.Side(period, direct)
	price = priceLimit(direct, price)

	orders := self.Order(self.instId(symbol, period), self.TdMode(period), side, poside, price, sz)
	if len(orders) == 0 {
		return false
	}

	return orders[0].SCode == "0"
}

// limit 成交价格
func priceLimit(direct conset.OPERATION, price float32) float32 {
	rate := float32(0.001)
	switch direct {
	case conset.BUY_HIGH, conset.SELL_LOW:
		price = price * (1.0 + rate)
	case conset.SELL_HIGH, conset.BUY_LOW:
		price = price * (1.0 - rate)
	}
	return price
}
