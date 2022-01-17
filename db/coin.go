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
	case conset.LTC_USD:
		s = "ltc_usd"
	case conset.LTC_USDT:
		s = "ltc_usdt"
	case conset.DOT_USD:
		s = "dot_usd"
	case conset.DOT_USDT:
		s = "dot_usdt"
	case conset.DOGE_USD:
		s = "doge_usd"
	case conset.DOGE_USDT:
		s = "doge_usdt"
	case conset.LUNA_USD:
		s = "luna_usd"
	case conset.LUNA_USDT:
		s = "luna_usdt"
	case conset.TONCOIN_USD:
		s = "toncoin_usd"
	case conset.TONCOIN_USDT:
		s = "toncoin_usdt"
	case conset.SHIBI_USD:
		s = "shib_usd"
	case conset.SHIBI_USDT:
		s = "shib_usdt"
	case conset.MATIC_USD:
		s = "matic_usd"
	case conset.MATIC_USDT:
		s = "matic_usdt"
	case conset.CRO_USD:
		s = "cro_usd"
	case conset.CRO_USDT:
		s = "cro_usdt"
	case conset.BCH_USD:
		s = "bch_usd"
	case conset.BCH_USDT:
		s = "bch_usdt"
	case conset.FTM_USD:
		s = "ftm_usd"
	case conset.FTM_USDT:
		s = "ftm_usdt"
	case conset.XLM_USD:
		s = "xlm_usd"
	case conset.XLM_USDT:
		s = "xlm_usdt"
	case conset.AXS_USD:
		s = "axs_usd"
	case conset.AXS_USDT:
		s = "axs_usdt"
	case conset.ONE_USD:
		s = "one_usd"
	case conset.ONE_USDT:
		s = "one_usdt"
	case conset.NEAR_USD:
		s = "near_usd"
	case conset.NEAR_USDT:
		s = "near_usdt"
	case conset.ICP_USD:
		s = "icp_usd"
	case conset.ICP_USDT:
		s = "icp_usdt"
	case conset.LEO_USD:
		s = "leo_usd"
	case conset.LEO_USDT:
		s = "leo_usdt"
	case conset.IOTA_USD:
		s = "iota_usd"
	case conset.IOTA_USDT:
		s = "iota_usdt"
	case conset.ADA_USD:
		s = "ada_usd"
	case conset.ADA_USDT:
		s = "ada_usdt"
	case conset.FIL_USD:
		s = "fil_usd"
	case conset.FIL_USDT:
		s = "fil_usdt"
	case conset.ATOM_USD:
		s = "atom_usd"
	case conset.ATOM_USDT:
		s = "atom_usdt"
	case conset.XRP_USD:
		s = "xrp_usd"
	case conset.XRP_USDT:
		s = "xrp_usdt"
	case conset.LINK_USD:
		s = "link_usd"
	case conset.LINK_USDT:
		s = "link_usdt"
	case conset.EOS_USD:
		s = "eos_usd"
	case conset.EOS_USDT:
		s = "eos_usdt"
	case conset.UNI_USD:
		s = "uni_usd"
	case conset.UNI_USDT:
		s = "uni_usdt"
	case conset.CRV_USD:
		s = "crv_usd"
	case conset.CRV_USDT:
		s = "crv_usdt"
	case conset.THETA_USD:
		s = "theta_usd"
	case conset.THETA_USDT:
		s = "theta_usdt"
	case conset.ALGO_USD:
		s = "algo_usd"
	case conset.ALGO_USDT:
		s = "algo_usdt"
	case conset.ETC_USD:
		s = "etc_usd"
	case conset.ETC_USDT:
		s = "etc_usdt"
	case conset.SAND_USD:
		s = "sand_usd"
	case conset.SAND_USDT:
		s = "sand_usdt"
	case conset.SOL_USD:
		s = "sol_usd"
	case conset.SOL_USDT:
		s = "sol_usdt"
	case conset.XTZ_USD:
		s = "xtz_usd"
	case conset.XTZ_USDT:
		s = "xtz_usdt"
	case conset.DASH_USD:
		s = "dash_usd"
	case conset.DASH_USDT:
		s = "dash_usdt"
	case conset.TRX_USD:
		s = "trx_usd"
	case conset.TRX_USDT:
		s = "trx_usdt"
	case conset.XMR_USD:
		s = "xmr_usd"
	case conset.XMR_USDT:
		s = "xmr_usdt"
	case conset.MANA_USD:
		s = "mana_usd"
	case conset.MANA_USDT:
		s = "mana_usdt"
	case conset.SUSHI_USD:
		s = "sushi_usd"
	case conset.SUSHI_USDT:
		s = "sushi_usdt"
	case conset.ZEC_USD:
		s = "zec_usd"
	case conset.ZEC_USDT:
		s = "zec_usdt"
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
