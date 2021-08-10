package conset

type PLAT int32
type SYMBOL int32
type PERIOD int32
type TIMES int32
type OPERATION int32

const (
	OKEX PLAT = iota + 3
	QKL123
)

const (
	BTC_USD SYMBOL = iota + 10
	BTC_USDT
	ETH_USD
	ETH_USDT
	LTC_USD
	LTC_USDT
)

// 交易类型
const (
	SPOT         PERIOD = iota + 100 // 现货
	SWAP                             // 永续
	WEEK                             // 当周
	WEEK_NEXT                        // 次周
	QUARTER                          // 当季
	QUARTER_NEXT                     // 次季
)

const (
	MIN_5 TIMES = iota + 200 // 5min
	MIN_15
	MIN_30
	H_1
	H_2
	H_4
	H_6
	H_12
	D_1
	D_5
	W_1
	M_1
)

const (
	BUY_HIGH  = 1 // 买入开多/(现货)买入
	BUY_LOW   = 2
	SELL_HIGH = 3 // 平多/(现货)卖出
	SELL_LOW  = 4
)
