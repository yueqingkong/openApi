package plat

import (
	"github.com/yueqingkong/openApi/conset"
	"github.com/yueqingkong/openApi/db"
	"time"
)

// 平台通用方法  api上的抽象层
type BasePlat interface {
	Inits(...string) // 初始化 apikey...

	Pull(symbol conset.SYMBOL, times conset.TIMES, start time.Time) []db.Coin                    // 同步Kline
	Price(symbol conset.SYMBOL) float32                                                          // 当前价格
	Orders(symbol conset.SYMBOL, operation conset.OPERATION, price float32, amount float32) bool // 交易
}
