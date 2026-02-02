package util

import (
	"github.com/yueqingkong/openApi/conset"
)

// coin/u -> 张
func Size(bs, quote conset.CCY, period conset.PERIOD, price float32, buyunit float32) float32 {
	if price == 0 || buyunit == 0 {
		return 0
	}

	var amount float32
	switch period {
	case conset.SPOT: // 现货
		// 仅u本位
		amount = buyunit / price
	case conset.SWAP:
		switch quote {
		case conset.USD: // 币本位
			amount = UsdSize(price, buyunit, CcyValue(bs, quote))
		case conset.USDT: // u本位
			amount = UsdtSize(price, buyunit, CcyValue(bs, quote))
		}
	}
	return amount
}

// token - > 张数
// 保证金*币的价格*杠杆倍数／合约面值=可开张数
func UsdSize(price float32, buyunit float32, ccyValue float32) float32 {
	var size float32
	amout := price * buyunit / ccyValue
	if amout < 1.0 {
		size = 1.0
	} else {
		size = Floor(amout)
	}
	return size
}

func UsdtSize(price float32, buyunit float32, ccyValue float32) float32 {
	var size float32
	amout := buyunit / price / ccyValue
	if amout < 1.0 {
		size = 1.0
	} else {
		size = Floor(amout)
	}
	return size
}

// 面值
func CcyValue(bs, quote conset.CCY) float32 {
	var v float32
	switch quote {
	case conset.USD:
		switch bs {
		case conset.BTC:
			v = 100.0
		default:
			v = 10.0
		}
	case conset.USDT:
		switch bs {
		case conset.BTC:
			v = 0.01
		default:
			v = 0.1
		}
	}
	return v
}

// 支付手续费
// 用户等级 lv1 吃单费率 0.05%，挂单费率 0.02%
// 币本位合约开仓手续费=面值*开仓张数/开仓价格*手续费费率
// 币本位合约平仓手续费=面值*平仓张数/平仓价格*手续费费率
// USDT合约开仓手续费=面值*开仓张数*开仓价格*手续费费率
// USDT合约平仓手续费=面值*平仓张数*平仓价格*手续费费率
func PayFee(price, size, value float32, fee ...float32) float32 {
	realFee := float32(0.0005)
	if len(fee) > 0 {
		realFee = fee[0]
	}

	return value / price * size * realFee
}

// 合约收益
func Profit(bs, quote conset.CCY, op conset.OPERATION, price float32, lastprice float32, size float32) float32 {
	var amount float32
	if quote == conset.USD {
		amount = UsdProfit(op, price, lastprice, size, CcyValue(bs, quote))
	} else if quote == conset.USDT {
		amount = UsdtProfit(op, price, lastprice, size, CcyValue(bs, quote))
	}
	return amount
}

// 多仓收益=面值*开仓张数（1／开仓价格-1／平仓价格）
// 空仓收益=面值*开仓张数（1／平仓价格-1／开仓价格）
func UsdProfit(op conset.OPERATION, price float32, lastprice float32, size, value float32) float32 {
	var profit float32
	if op == conset.BUY_HIGH || op == conset.SELL_HIGH {
		profit = value * size * (1.0/lastprice - 1.0/price)
	} else if op == conset.BUY_LOW || op == conset.SELL_LOW {
		profit = value * size * (1.0/price - 1.0/lastprice)
	}
	return profit
}

// USDT收益 正向合约
// 做多：收益=（平仓价-开仓价）*面值*张数=（平仓价-开仓价）*数量
// 做空：收益=（开仓价-平仓价）*面值*张数=（开仓价-平仓价）*数量
func UsdtProfit(op conset.OPERATION, closePrice, openPrice float32, size, value float32) float32 {
	var profit float32
	if op == conset.BUY_HIGH || op == conset.SELL_HIGH {
		profit = (closePrice - openPrice) * value * size
	} else if op == conset.BUY_LOW || op == conset.SELL_LOW {
		profit = (openPrice - closePrice) * value * size
	}
	return profit
}

// 收益率
// op 3 平多 4 平空
func ProfitRate(profit float32, total float32) float32 {
	if profit == 0 || total == 0 {
		return 0.0
	}
	rate := profit / total * 100
	return rate
}
