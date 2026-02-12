package okex

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/yueqingkong/openApi/conset"
	"github.com/yueqingkong/openApi/db"
	"github.com/yueqingkong/openApi/util"
)

type Base struct {
	*Api
	*Ws
}

var (
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

// base 交易货币
// quote 计价货币(USD, USDT)
//func (self *Base) instId(base conset.CCY, quote conset.CCY) string {
//	return fmt.Sprintf("%s-%s", strings.ToUpper(ccyStringMap[base]), strings.ToUpper(ccyStringMap[quote]))
//}

// base 交易货币
// quote 计价货币(USD, USDT)
func (self *Base) InstId(base conset.CCY, quote conset.CCY, period conset.PERIOD) string {
	if base == "" || quote == "" || period == 0 {
		return ""
	}

	switch period {
	case conset.SPOT, conset.MARGIN:
		return fmt.Sprintf("%s-%s", strings.ToUpper(string(base)), strings.ToUpper(string(quote)))
	case conset.SWAP:
		return fmt.Sprintf("%s-%s-%s", strings.ToUpper(string(base)), strings.ToUpper(string(quote)), strings.ToUpper(db.Period(period)))
	}
	return ""
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

func TdMode(period conset.PERIOD) string {
	var s string
	switch period {
	case conset.SPOT:
		s = "cash"
	case conset.SWAP, conset.MARGIN:
		s = "isolated"
	default:
		s = "isolated"
	}
	return s
}

func Side(period conset.PERIOD, direct conset.OPERATION) (string, string) {
	var side string
	var posSide string

	switch direct {
	case conset.BUY_HIGH:
		side = conset.BUY
	case conset.BUY_LOW:
		side = conset.SELL
	case conset.SELL_HIGH:
		side = conset.SELL
	case conset.SELL_LOW:
		side = conset.BUY
	}

	if period == conset.SWAP {
		switch direct {
		case conset.BUY_HIGH:
			posSide = conset.LONG
		case conset.BUY_LOW:
			posSide = conset.SHORT
		case conset.SELL_HIGH:
			posSide = conset.LONG
		case conset.SELL_LOW:
			posSide = conset.SHORT
		}
	}
	return side, posSide
}

func (self *Base) Pull(base conset.CCY, quote conset.CCY, period conset.PERIOD, times conset.TIMES, start time.Time) bool {
	candles := self.Candles(self.InstId(base, quote, period), db.Times(times), self.before(start), 300)

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
			coin := &db.Coin{
				Plat:       db.Plat(self.Plat()),
				Symbol:     db.Symbol(base, quote),
				Times:      db.Times(times),
				Period:     db.Period(period),
				Timestamp:  timetamp,
				Open:       float32(open),
				Close:      float32(close),
				High:       float32(high),
				Low:        float32(low),
				Volume:     float32(volume),
				CreateTime: time.Unix(timetamp/1000, 0)}
			if err := coin.Create(); err != nil {
				log.Printf("Create err: %+v", err)
			}
		}
	}
	return false
}

func (self *Base) Price(base conset.CCY, quote conset.CCY, period conset.PERIOD) float32 {
	tickers := self.Ticker(self.InstId(base, quote, period))
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

func (self *Base) SupportCoin() *SupportCoinBody {
	return self.Api.SupportCoin()
}

// times: [5m/1H/1D]
func (self *Base) TakerVolume(base conset.CCY, period conset.PERIOD, start, end int64, times conset.TIMES) [][]string {
	instType := "CONTRACTS"
	if period == conset.SPOT {
		instType = "SPOT"
	}

	return self.Api.TakerVolume(string(base), instType, strconv.FormatInt(start, 10), strconv.FormatInt(end, 10), db.Times(times))
}

// times: [5m/1H/1D]
func (self *Base) LoanRatio(base conset.CCY, start, end int64, times conset.TIMES) [][]string {
	return self.Api.LoanRatio(string(base), strconv.FormatInt(start, 10), strconv.FormatInt(end, 10), db.Times(times))
}

// times: [5m/1H/1D]
func (self *Base) SwapAccountRatio(base conset.CCY, start, end int64, times conset.TIMES) [][]string {
	return self.Api.SwapAccountRatio(string(base), strconv.FormatInt(start, 10), strconv.FormatInt(end, 10), db.Times(times))
}

// times: [5m/1H/1D]
func (self *Base) InterestVolume(base conset.CCY, start, end int64, times conset.TIMES) [][]string {
	return self.Api.InterestVolume(string(base), strconv.FormatInt(start, 10), strconv.FormatInt(end, 10), db.Times(times))
}

func (self *Base) TopSentimentIndex(base conset.CCY, start, num int64, times conset.TIMES) [][]string {
	return self.Api.TopSentimentIndex(string(base), strconv.FormatInt(start, 10), strconv.FormatInt(num, 10), db.Times(times))
}

func (self *Base) TopAverageIndex(base conset.CCY, start, num int64, times conset.TIMES) [][]string {
	return self.Api.TopAverageIndex(string(base), strconv.FormatInt(start, 10), strconv.FormatInt(num, 10), db.Times(times))
}

func (self *Base) FundingRate(base conset.CCY, quote conset.CCY) (float32, float32) {
	rates := self.Api.FundingRate(self.InstId(base, quote, conset.SWAP))
	if len(rates) == 0 {
		return 0, 0
	}

	return util.Float32(rates[0].FundingRate), util.Float32(rates[0].NextFundingRate)
}

func (self *Base) Balance(c conset.CCY) []*Balance {
	return self.Api.balance(strings.ToUpper(string(c)))
}

func (self *Base) PullInstrument(period conset.PERIOD) {
	instruments := self.Instrument(period, "", "")
	for _, ins := range instruments {
		instrument := &db.Instrument{}
		instrument.CreateOrUpdate(self.Plat(), period, ins.InstID, ins.Uly, ins.InstFamily, ins.SettleCcy, ins.CtVal, ins.State)
	}
}

func (self *Base) Instrument(period conset.PERIOD, base conset.CCY, quote conset.CCY) []*Instrument {
	return self.Api.Instruments(strings.ToUpper(db.Period(period)), "", self.InstId(base, quote, period))
}

func (self *Base) OrderInfo(base conset.CCY, quote conset.CCY, period conset.PERIOD, orderId string) (bool, *OrderInfo) {
	infos := self.Api.OrderInfo(self.InstId(base, quote, period), orderId)
	if len(infos) == 0 {
		return false, nil
	}

	return true, infos[0]
}

func (self *Base) Order(base conset.CCY, quote conset.CCY, period conset.PERIOD, op conset.OPERATION, price, sz, rate float32) (bool, *OrderRes) {
	orders := self.Api.Order((&OrderParam{}).Format(base, quote, period, op, price, sz, rate))
	if len(orders) == 0 {
		return false, nil
	}

	return orders[0].SCode == "0", orders[0]
}

func (param *OrderParam) Format(bs conset.CCY, quote conset.CCY, period conset.PERIOD, op conset.OPERATION, price, sz, rate float32) *OrderParam {
	side, poside := Side(period, op)
	price = priceLimit(op, price, rate)

	param.InstId = (&Base{}).InstId(bs, quote, period)
	param.TdMode = TdMode(period)
	param.Side = side
	param.PosSide = poside
	param.OrdType = "limit"
	param.Px = price
	param.Sz = sz
	return param
}

func (self *Base) BatchOrder(params []*OrderParam) (bool, []*OrderRes) {
	orders := self.Api.BatchOrder(params)
	if len(orders) == 0 {
		return false, nil
	}

	return orders[0].SCode == "0", orders
}

func (self *Base) SavingsBalance(base conset.CCY) []*SavingsBalance {
	return self.Api.SavingsBalance(strings.ToUpper(string(base)))
}

// side: purchase：申购 redempt：赎回
func (self *Base) SavingsPurchaseRedempt(base conset.CCY, amt float32, side string, rate float32) []*SavingsPurchaseRedempt {
	amtStr := util.FloatString(amt)
	rateStr := ""
	if rate > 0 {
		rateStr = util.FloatString(rate * 100)
	}

	return self.Api.SavingsPurchaseRedempt(strings.ToUpper(string(base)), amtStr, side, rateStr)
}

// 资金划转
func (self *Base) Transfer(base conset.CCY, amt float32, typ, from, to, fromSubAccount, toSubAccount string) []*Transfer {
	amtStr := util.FloatString(amt)
	return self.Api.Transfer(typ, strings.ToUpper(string(base)), amtStr, from, to, fromSubAccount, toSubAccount)
}

// 获取资金账户余额
func (self *Base) AssetBalance(base conset.CCY) []*AssetBalance {
	return self.Api.AssetBalance(strings.ToUpper(string(base)))
}

// limit 成交价格
func priceLimit(direct conset.OPERATION, price, rate float32) float32 {
	if rate == 0 {
		rate = float32(0.01)
	}

	switch direct {
	case conset.BUY_HIGH, conset.SELL_LOW:
		price = price * (1.0 + rate)
	case conset.SELL_HIGH, conset.BUY_LOW:
		price = price * (1.0 - rate)
	}
	return price
}

func (self *Base) SubscribeTickers(channel string, instIds []string, f func(conset.CCY, conset.CCY, interface{})) {
	self.WsTickers(channel, instIds, f)
}
