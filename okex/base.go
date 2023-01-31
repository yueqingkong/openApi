package okex

import (
	"fmt"
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

var (
	ccyStringMap = map[conset.CCY]string{conset.USD: "usd", conset.USDT: "usdt", // 稳定币
		conset.BTC: "btc", conset.ETH: "eth", conset.LTC: "ltc", conset.DOT: "dot", conset.DOGE: "doge",
		conset.LUNA: "luna", conset.TONCOIN: "toncoin", conset.SHIBI: "shib", conset.MATIC: "matic", conset.CRO: "cro", conset.BCH: "bch", conset.FTM: "ftm",
		conset.XLM: "xlm", conset.AXS: "axs", conset.ONE: "one", conset.NEAR: "near", conset.ICP: "icp", conset.LEO: "leo", conset.IOTA: "iota",
		conset.ADA: "ada", conset.FIL: "fil", conset.ATOM: "atom", conset.XRP: "xrp", conset.LINK: "link", conset.EOS: "eos", conset.UNI: "uni",
		conset.CRV: "crv", conset.THETA: "theta", conset.ALGO: "algo", conset.ETC: "etc", conset.SAND: "sand", conset.SOL: "sol", conset.XTZ: "xtz",
		conset.DASH: "dash", conset.TRX: "trx", conset.XMR: "xmr", conset.MANA: "mana", conset.SUSHI: "sushi", conset.ZEC: "zec", conset.SNX: "snx",
		conset.AVAX: "avax", conset.WAVES: "waves", conset.AAVE: "aave", conset.BSV: "bsv", conset.XCH: "xch", conset.ENS: "ens", conset.COMP: "comp",
		conset.EGLD: "egld"}

	toCcyMap = map[string]conset.CCY{}
)

func (self *Base) Init(strings []string) {
	if len(strings) >= 3 {
		self.Api = &Api{}
		self.Ws = &Ws{}

		self.InitApiKeys(strings[0], strings[1], strings[2])
		self.InitWsKeys(strings[0], strings[1], strings[2])
	}
}

func (self *Base) ToCcy(ccy string) conset.CCY {
	if len(toCcyMap) == 0 {
		for k, v := range ccyStringMap {
			toCcyMap[v] = k
		}
	}

	return toCcyMap[ccy]
}

// base 交易货币
// quote 计价货币(USD, USDT)
//func (self *Base) instId(base conset.CCY, quote conset.CCY) string {
//	return fmt.Sprintf("%s-%s", strings.ToUpper(ccyStringMap[base]), strings.ToUpper(ccyStringMap[quote]))
//}

// base 交易货币
// quote 计价货币(USD, USDT)
func (self *Base) instIds(base conset.CCY, quote conset.CCY, period conset.PERIOD) string {
	if base == 0 || quote == 0 || period == 0 {
		return ""
	}

	return fmt.Sprintf("%s-%s-%s", strings.ToUpper(ccyStringMap[base]), strings.ToUpper(ccyStringMap[quote]), strings.ToUpper(self.Period(period)))
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

func TdMode(period conset.PERIOD) string {
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

func Side(direct conset.OPERATION) (string, string) {
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

func (self *Base) Pull(base conset.CCY, quote conset.CCY, period conset.PERIOD, times conset.TIMES, start time.Time) bool {
	candles := self.Candles(self.instIds(base, quote, period), self.Times(times), self.before(start))

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
			if err := coin.Create(self.Plat(), base, quote, times, float32(open), float32(close), float32(high), float32(low), float32(volume), timetamp); err != nil {
				log.Printf("Create err: %+v", err)
			}
		}
	}
	return false
}

func (self *Base) Price(base conset.CCY, quote conset.CCY, period conset.PERIOD) float32 {
	tickers := self.Ticker(self.instIds(base, quote, period))
	if len(tickers) == 0 {
		return 0.0
	}

	return util.Float32(tickers[0].Last)
}

func (self *Base) UsdCny() float32 {
	rates := self.ExchangeRate()
	if len(rates) == 0 {
		return 0.0
	}

	return util.Float32(rates[0].UsdCny)
}

func (self *Base) Balance(c conset.CCY) float32 {
	bs := self.balance(strings.ToUpper(ccyStringMap[c]))
	if len(bs) == 0 {
		return 0.0
	}

	return util.Float32(bs[0].TotalEq)
}

func (self *Base) PullInstrument(period conset.PERIOD) {
	instruments := self.Instrument(period, 0, 0)
	for _, ins := range instruments {
		instrument := &db.Instrument{}
		instrument.CreateOrUpdate(self.Plat(), period, ins.InstID, ins.Uly, ins.InstFamily, ins.SettleCcy, ins.CtVal, ins.State)
	}
}

func (self *Base) Instrument(period conset.PERIOD, base conset.CCY, quote conset.CCY) []*Instrument {
	return self.Instruments(strings.ToUpper(self.Period(period)), "", self.instIds(base, quote, period))
}

func (self *Base) Orders(base conset.CCY, quote conset.CCY, period conset.PERIOD, direct conset.OPERATION, price, sz float32) bool {
	side, poside := Side(direct)
	price = priceLimit(direct, price)

	orders := self.Order(self.instIds(base, quote, period), TdMode(period), side, poside, price, sz)
	if len(orders) == 0 {
		return false
	}

	return orders[0].SCode == "0"
}

// limit 成交价格
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

func (self *Base) SubscribeTickers(symbols [][2]conset.CCY, period conset.PERIOD, f func(conset.CCY, conset.CCY, float32)) {
	ids := make([]string, 0)
	for _, symbol := range symbols {
		inst := self.instIds(symbol[0], symbol[1], period)
		ids = append(ids, inst)
	}

	self.WsTickers(ids, f)
}
