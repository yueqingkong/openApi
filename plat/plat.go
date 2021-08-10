package plat

import (
	"github.com/yueqingkong/openApi/conset"
	"github.com/yueqingkong/openApi/db"
	"time"
)

// 平台通用方法  api上的抽象层
type BasePlat interface {
	Inits(...string) // 初始化 apikey...

	Pull(symbol conset.SYMBOL, period conset.PERIOD, times conset.TIMES, start time.Time) []db.Coin                          // 同步Kline
	Price(symbol conset.SYMBOL, period conset.PERIOD) float32                                                      // 当前价格
	DealOrder(symbol conset.SYMBOL, period conset.PERIOD, operation conset.OPERATION, price float32, amount float32) bool // 交易
}
