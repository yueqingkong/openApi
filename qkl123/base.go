package qkl123

import (
	"fmt"
	"github.com/yueqingkong/openApi/conset"
	"github.com/yueqingkong/openApi/db"
	"log"
	"time"
)

type Base struct {
	*Api
}

func (self *Base) Plat() conset.PLAT {
	return conset.QKL123
}

func (self *Base) instId(symbol conset.SYMBOL, period conset.PERIOD) string {
	var s string
	switch period {
	case conset.SPOT:
		switch symbol {
		case conset.BTC_USD, conset.BTC_USDT:
			s = "btc_usd"
		case conset.ETH_USD, conset.ETH_USDT:
			s = "eth_usd"
		}
	case conset.SWAP:
		switch symbol {
		case conset.BTC_USD, conset.BTC_USDT:
			s = "btc_usd_swaps"
		case conset.ETH_USD, conset.ETH_USDT:
			s = "eth_usd_swaps"
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

func (self *Base) Times(times conset.TIMES) string {
	var s string
	switch times {
	case conset.MIN_15:
		s = "15min"
	case conset.MIN_30:
		s = "30min"
	case conset.H_1:
		s = "1hour"
	case conset.H_6:
		s = "6hour"
	case conset.H_12:
		s = "12hour"
	case conset.D_1:
		s = "1day"
	}
	return s
}

func (self *Api) InitKeys(apikey, secretkey, passphrase, auth string) {
	self.ApiKey = apikey
	self.SecretKey = secretkey
	self.Passphrase = passphrase
	self.Auth = auth
}

func (self *Base) Pull(symbol conset.SYMBOL, period conset.PERIOD, times conset.TIMES, start time.Time) bool {
	var end string
	if start.IsZero() {
		end = fmt.Sprintf("%d", 1577808000)
	} else {
		end = fmt.Sprintf("%d", start.Unix())
	}
	candles := self.Candles(self.instId(symbol, period), self.Times(times), end)

	if len(candles) == 0 {
		log.Printf("Pull : 同步完成")
		return true
	}

	for k, value := range candles {
		if k != len(candles)-1 { // 最近时间一条有效的K线不保存
			coin := &db.Coin{}
			if err := coin.Create(self.Plat(), symbol, period, times, value.Open, value.Close, value.High, value.Low, value.FundsOut, value.Date*1000); err != nil {
				log.Printf("Create err: %+v", err)
			}
		}
	}
	return false
}
