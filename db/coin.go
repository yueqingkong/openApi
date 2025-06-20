package db

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/yueqingkong/openApi/conset"
	"github.com/yueqingkong/openApi/util"
	"xorm.io/builder"
)

// Coin K线数据
type Coin struct {
	Id         int64
	Plat       string    `xorm:"plat unique(p-p-s-t-c)"`
	Period     string    `xorm:"period unique(p-p-s-t-c)"` // spot/swap
	Symbol     string    `xorm:"symbol unique(p-p-s-t-c)"`
	Times      string    `xorm:"times unique(p-p-s-t-c)"` // 时间间隔
	Open       float32   `xorm:"float"`
	Close      float32   `xorm:"float"`
	High       float32   `xorm:"float"`
	Low        float32   `xorm:"float"`
	Volume     float32   `xorm:"float"`
	Timestamp  int64     `json:"time_stamp" xorm:"bigint unique(p-p-s-t-c)"` // 毫秒
	CreateTime time.Time `json:"create_time" xorm:"create_time"`
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"updated"`
}

func Plat(p conset.PLAT) string {
	var s string
	switch p {
	case conset.OKEX:
		s = "okex"
	case conset.QKL123:
		s = "qkl123"
	case conset.COIN_MARKET_CAP:
		s = "coin_market_cap"
	}
	return s
}

func Times(times conset.TIMES) string {
	var s string
	switch times {
	case conset.MIN_15:
		s = "15m"
	case conset.MIN_30:
		s = "30m"
	case conset.H_1:
		s = "1H"
	case conset.H_4:
		s = "4H"
	case conset.H_6:
		s = "6H"
	case conset.H_12:
		s = "12H"
	case conset.D_1:
		s = "1D"
	case conset.W_1:
		s = "1W"
	case conset.M_1:
		s = "1M"
	}
	return s
}

func Period(period conset.PERIOD) string {
	var s string
	switch period {
	case conset.SPOT:
		s = "spot"
	case conset.SWAP:
		s = "swap"
	case conset.MARGIN:
		s = "margin"
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

func Symbol(base conset.CCY, quote conset.CCY) string {
	if quote == "" {
		return fmt.Sprintf("%s", strings.ToLower(string(base)))
	}

	return fmt.Sprintf("%s_%s", strings.ToLower(string(base)), strings.ToLower(string(quote)))
}

func (self *Coin) Create() error {
	if _, err := Engine().InsertOne(self); err != nil {
		return err
	}
	return nil
}

func (self *Coin) Last() error {
	if b, err := Engine().Desc("create_time").Get(self); err != nil || !b {
		return errors.New("get")
	}
	return nil
}

func (self *Coin) LastTime() (bool, time.Time) {
	var startTime time.Time

	// coin := &Coin{Plat: self.Plat, Symbol: self.Symbol, Times: self.Times}
	if nil == self.last() {
		startTime = self.CreateTime
	}

	// 最后一条记录是昨天的
	diffHours := time.Now().Sub(startTime).Hours()

	// 是否最新的数据
	switch self.Times {
	case Times(conset.MIN_15):
		if diffHours < 0.5 {
			return false, time.Time{}
		}
	case Times(conset.MIN_30):
		if diffHours < 1 {
			return false, time.Time{}
		}
	case Times(conset.H_1):
		if diffHours < 2 {
			return false, time.Time{}
		}
	case Times(conset.H_2):
		if diffHours < 4 {
			return false, time.Time{}
		}
	case Times(conset.H_4):
		if diffHours < 8 {
			return false, time.Time{}
		}
	case Times(conset.H_6):
		if diffHours < 12 {
			return false, time.Time{}
		}
	case Times(conset.H_12):
		if diffHours < 24 {
			return false, time.Time{}
		}
	case Times(conset.D_1):
		if diffHours < 24*2 {
			return false, time.Time{}
		}
	case Times(conset.W_1):
		if diffHours < 7*24*2 {
			return false, time.Time{}
		}
	case Times(conset.M_1):
		if diffHours < 31*24*2 {
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

func (self *Coin) Lasts(limit int, end time.Time) ([]Coin, error) {
	coins := make([]Coin, 0)

	sql, args, _ := builder.ToSQL(builder.Lt{"create_time": end})
	if err := Engine().Where(sql, args...).Desc("create_time").Limit(limit).Find(&coins, self); err != nil {
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
func (self *Coin) MA(limit int, end time.Time) float32 {
	if coins, err := self.Lasts(limit, end); err != nil {
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
func (self *Coin) EMA(limit int, end time.Time) float32 {
	if coins, err := self.Lasts(limit, end); err != nil {
		return 0.0
	} else {
		log.Print(coins)

		factors := 2.0 / (float32(limit) + 1.0)
		log.Printf("factors : %f", factors)

		var value float32
		for i := 0; i < len(coins); i++ {
			c := coins[i]
			if i == 0 {
				value = c.EMAStart(limit, c.CreateTime)
			} else {
				value = c.Close*factors + value*(1.0-factors)
			}
		}

		return value
	}
}

func (self *Coin) EMAStart(limit int, end time.Time) float32 {
	if coins, err := self.Lasts(limit, end); err != nil {
		return 0.0
	} else {
		log.Print(coins)

		factors := 2.0 / (float32(limit) + 1.0)
		log.Printf("factors : %f", factors)

		var value float32
		for i := 0; i < len(coins); i++ {
			c := coins[i]
			if i == 0 {
				value = c.Close
			} else {
				value = c.Close*factors + value*(1.0-factors)
			}
		}

		return value
	}
}

// 通道
// N日移动平均线=N日收市价之和/N
func (self *Coin) Chanel(limit int, end time.Time) (float32, float32) {
	if coins, err := self.Lasts(limit, end); err != nil {
		return 0.0, 0.0
	} else {
		var low float32
		var high float32

		for k, v := range coins {
			if k == 0 {
				low = v.Low
				high = v.High
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

// 平均波动幅度
// 1、当前交易日的最高价与最低价间的波幅
// 2、前一交易日收盘价与当个交易日最高价间的波幅
// 3、前一交易日收盘价与当个交易日最低价间的波幅
func (self *Coin) ATR(limit int, end time.Time) float32 {
	if coins, err := self.Lasts(limit, end); err != nil {
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
