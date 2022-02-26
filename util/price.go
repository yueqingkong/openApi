package util

// 价格相等 合理偏差 0.001
// dest 目标价格
func PriceEqual(dest float32, price float32) bool {
	return Abs(dest-price)/dest < 0.005
}

// 价格趋势
// 1 上升 2 下降 3 波动
func PriceTend(prices []float32) int32 {
	var lastTend int32 = 3

	var lastValue float32
	for k, value := range prices {
		if k == 0 {
			lastValue = value
		} else if k == 1 {
			if value >= lastValue*(1+0.001) {
				lastTend = 1
			} else if value*(1+0.001) <= lastValue {
				lastTend = 2
			} else {
				break
			}

			lastValue = value
		} else {
			if lastTend == 1 && value >= lastValue*(1+0.001) {
				lastTend = 1
			} else if lastTend == 2 && value*(1+0.001) <= lastValue {
				lastTend = 2
			} else {
				lastTend = 3
				break
			}

			lastValue = value
		}
	}

	return lastTend
}

/**
 * 买入单位,回撤 ATR 亏损比例
 * rate*Total = atr/price * Unit
 * rate: 总账户最大亏损(总账户 满足连续最大亏损40次)
 * atr/price 回撤(高位最大回撤 0.04)
 */
func Unit(total float32, price float32) float32 {
	MaxLoss := float32(4000.0) // 最高点最大回撤

	unit := (total * price) / (20.0 * MaxLoss)
	return unit
}
