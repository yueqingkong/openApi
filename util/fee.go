package util

import "github.com/yueqingkong/openApi/conset"

// token - > 张数
// 保证金*币的价格*杠杆倍数／合约面值=可开张数
func BuySize(price float32, buyunit float32, value float32) float32 {
	var size float32
	if buyunit == 0.0 {
		size = 0
	} else {
		amout := price * buyunit / value
		if amout < 1.0 {
			size = 1.0
		} else {
			size = Floor(amout)
		}
	}
	return size
}

// 一张 代表的面纸
func ZDollar(symbol conset.SYMBOL) float32 {
	var v float32
	if symbol == conset.BTC_USDT || symbol == conset.BTC_USD {
		v = 100.0
	} else {
		v = 10.0
	}
	return v
}

// 支付手续费
func PayFee(price float32, size, value float32) float32 {
	return value / price * size * 0.0005
}

// 币本位收益 反向合约
// 多仓收益=面值*开仓张数（1／开仓价格-1／平仓价格）
// 空仓收益=面值*开仓张数（1／平仓价格-1／开仓价格）
func Profit(op conset.OPERATION, price float32, lastprice float32, size, value float32) float32 {
	var profit float32
	if op == conset.BUY_HIGH || op == conset.SELL_HIGH {
		profit = (value/lastprice - value/price) * size
	} else if op == conset.BUY_LOW || op == conset.SELL_LOW {
		profit = (value/price - value/lastprice) * size
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
