package db

import (
	"errors"
	"github.com/yueqingkong/openApi/conset"
	"github.com/yueqingkong/openApi/util"
	"log"
	"time"
	"xorm.io/builder"
)

// Coin K线数据
type Coin struct {
	Id         int64
	Plat       string    `xorm:"plat varchar(255) unique(pl-time) index(pl-sy-t) index(p-s-t-p-c)"`
	Symbol     string    `xorm:"symbol varchar(255) unique(pl-time) index(pl-sy-t) index(p-s-t-p-c)"`
	Times      string    `xorm:"times varchar(255) unique(pl-time) index(pl-sy-t) index(p-s-t-p-c)"`         // 时间间隔
	Period     string    `json:"period" xorm:"varchar(255) unique(pl-time) index(pl-sy-t) index(p-s-t-p-c)"` // 合约类型 spot,week
	Open       float32   `xorm:"float"`
	Close      float32   `xorm:"float"`
	High       float32   `xorm:"float"`
	Low        float32   `xorm:"float"`
	Volume     float32   `xorm:"float"`
	Timestamp  int64     `json:"time_stamp" xorm:"bigint time_stamp index unique(pl-time)"` // 毫秒
	CreateTime time.Time `json:"create_time" xorm:"DATETIME create_time index(p-s-t-p-c)"`
}

func dPlat(p conset.PLAT) string {
	var s string
	switch p {
	case conset.OKEX:
		s = "okex"
	case conset.QKL123:
		s = "qkl123"
	}
	return s
}

func dSymbol(symbol conset.SYMBOL) string {
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

func dPeriod(period conset.PERIOD) string {
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

func dTimes(times conset.TIMES) string {
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

func (self *Coin) Create(pt conset.PLAT, symbol conset.SYMBOL, period conset.PERIOD, times conset.TIMES, open, close, high, low, volume float32, timetamp int64) error {
	coin := &Coin{
		Plat:       dPlat(pt),
		Symbol:     dSymbol(symbol),
		Period:     dPeriod(period),
		Times:      dTimes(times),
		Open:       open,
		Close:      close,
		High:       high,
		Low:        low,
		Volume:     volume,
		Timestamp:  timetamp,
		CreateTime: util.SecondsTime(timetamp / 1000),
	}

	_, err := Engine().InsertOne(coin)
	return err
}

func (self *Coin) LastTime(pt conset.PLAT, symbol conset.SYMBOL, period conset.PERIOD, times conset.TIMES) (bool, time.Time) {
	var startTime time.Time

	coin := &Coin{Plat: dPlat(pt), Symbol: dSymbol(symbol), Period: dPeriod(period), Times: dTimes(times)}
	if nil == coin.last() {
		startTime = coin.CreateTime
	}

	// 最后一条记录是昨天的
	diffHours := time.Now().Sub(startTime).Hours()

	// 是否最新的数据
	if times == conset.MIN_15 {
		if diffHours < 0.5 {
			return false, time.Time{}
		}
	} else if times == conset.MIN_30 {
		if diffHours < 1 {
			return false, time.Time{}
		}
	} else if times == conset.H_1 {
		if diffHours < 2 {
			return false, time.Time{}
		}
	} else if times == conset.H_2 {
		if diffHours < 4 {
			return false, time.Time{}
		}
	} else if times == conset.H_4 {
		if diffHours < 8 {
			return false, time.Time{}
		}
	} else if times == conset.H_6 {
		if diffHours < 12 {
			return false, time.Time{}
		}
	} else if times == conset.H_12 {
		if diffHours < 24 {
			return false, time.Time{}
		}
	} else if times == conset.D_1 {
		if diffHours < 48 {
			return false, time.Time{}
		}
	}

	// 避免重复返回最后一条的k线数据，加30s
	if !startTime.IsZero() {
		startTime = startTime.Add(time.Duration(30) * time.Second)
	}
	return true, startTime
}

func (self *Coin) last() error {
	if b, err := Engine().Desc("create_time").Get(self); err != nil || !b {
		return errors.New("get")
	}

	return nil
}

func (self *Coin) All(pt conset.PLAT, symbol conset.SYMBOL, period conset.PERIOD, times conset.TIMES, start time.Time) ([]*Coin, error) {
	coins := make([]*Coin, 0)

	sql, args, _ := builder.ToSQL(builder.Gte{"create_time": start})
	if err := Engine().Where(sql, args...).Asc("create_time").Find(&coins, &Coin{Plat: dPlat(pt), Symbol: dSymbol(symbol), Period: dPeriod(period), Times: dTimes(times)}); err != nil {
		return nil, err
	}

	return coins, nil
}

func (self *Coin) Lasts(pt conset.PLAT, symbol conset.SYMBOL, period conset.PERIOD, times conset.TIMES, limit int, end time.Time) ([]Coin, error) {
	coins := make([]Coin, 0)

	sql, args, _ := builder.ToSQL(builder.Lt{"create_time": end})
	if err := Engine().Where(sql, args...).Desc("create_time").Limit(limit).Find(&coins, &Coin{Plat: dPlat(pt), Symbol: dSymbol(symbol), Period: dPeriod(period), Times: dTimes(times)}); err != nil {
		return nil, err
	}

	index := make([]Coin, 0)
	for i := len(coins) - 1; i >= 0; i-- {
		index = append(index, coins[i])
	}

	return index, nil
}

// 移动平均线
// N日移动平均线=N日收市价之和/N
func (self *Coin) MA(pt conset.PLAT, symbol conset.SYMBOL, period conset.PERIOD, times conset.TIMES, limit int, end time.Time) float32 {
	if coins, err := self.Lasts(pt, symbol, period, times, limit, end); err != nil {
		return 0.0
	} else {
		var total float32
		for _, value := range coins {
			total += value.Close
		}

		length := len(coins)
		return total / float32(length)
	}
}

// 平滑移动平均线
// EMA(12) = [2/(12+1)]*今日收盘价+[11/(12+1)]*作日EMA(12)
func (self *Coin) EMA(pt conset.PLAT, symbol conset.SYMBOL, period conset.PERIOD, times conset.TIMES, limit int, end time.Time) float32 {
	if coins, err := self.Lasts(pt, symbol, period, times, limit, end); err != nil {
		return 0.0
	} else {
		log.Print(coins)

		factors := 2.0 / (float32(limit) + 1.0)
		log.Printf("factors : %f", factors)

		var value float32
		for i := 0; i < len(coins); i++ {
			coin := coins[i]
			if i == 0 {
				value = coin.Close
			} else {
				value = coin.Close*factors + value*(1.0-factors)
			}
		}

		return value
	}
}

// 通道
// N日移动平均线=N日收市价之和/N
func (self *Coin) Chanel(pt conset.PLAT, symbol conset.SYMBOL, period conset.PERIOD, times conset.TIMES, limit int, end time.Time) (float32, float32) {
	if coins, err := self.Lasts(pt, symbol, period, times, limit, end); err != nil {
		return 0.0, 0.0
	} else {
		var low float32
		var high float32

		for k, v := range coins {
			if k == 0 {
				low = v.Close
				high = v.Close
			} else {
				if v.Low < low {
					low = v.Low
				}
				if v.High > high {
					high = v.High
				}
			}
		}

		return low, high
	}
}

//  平均波动幅度
//  1、当前交易日的最高价与最低价间的波幅
//  2、前一交易日收盘价与当个交易日最高价间的波幅
//  3、前一交易日收盘价与当个交易日最低价间的波幅
func (self *Coin) ATR(pt conset.PLAT, symbol conset.SYMBOL, period conset.PERIOD, times conset.TIMES, limit int, end time.Time) float32 {
	if coins, err := self.Lasts(pt, symbol, period, times, limit, end); err != nil {
		return 0.0
	} else {
		var totalRange float32
		for k, value := range coins {
			var dayRange float32
			if k == 0 {
				dayRange = value.High - value.Low
			} else {
				last := coins[k-1]
				lastHigh := util.Abs(last.Close - value.High)
				lastLow := util.Abs(last.Close - value.Low)
				todayRange := value.High - value.Low

				dayRange = util.Max(todayRange, lastHigh, lastLow)
			}
			totalRange += dayRange
		}

		length := len(coins)
		atr := totalRange / float32(length)
		return atr
	}
}