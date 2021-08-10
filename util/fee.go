package util

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
func ZDollar(symbol string) float32 {
	var v float32
	if symbol == "btc" {
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

// 收益
// op 3 平多 4 平空
func Profit(op int32, price float32, lastprice float32, size, value float32) float32 {
	var profit float32
	if op == 3 {
		profit = (value/lastprice - value/price) * size
	} else if op == 4 {
		profit = (value/price - value/lastprice) * size
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
