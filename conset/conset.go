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
	DOT_USD
	DOT_USDT
	DOGE_USD
	DOGE_USDT
	ADA_USD
	ADA_USDT
	FIL_USD
	FIL_USDT
	ATOM_USD
	ATOM_USDT
	XRP_USD
	XRP_USDT
	LINK_USD
	LINK_USDT
	EOS_USD
	EOS_USDT
	UNI_USD
	UNI_USDT
	CRV_USD
	CRV_USDT
	THETA_USD
	THETA_USDT
	ALGO_USD
	ALGO_USDT
	ETC_USD
	ETC_USDT
	SAND_USD
	SAND_USDT
	SOL_USD
	SOL_USDT
	XTZ_USD
	XTZ_USDT
	DASH_USD
	DASH_USDT
	TRX_USD
	TRX_USDT
	XMR_USD
	XMR_USDT
	MANA_USD
	MANA_USDT
	SUSHI_USD
	SUSHI_USDT
	ZEC_USD
	ZEC_USDT
	LUNA_USD
	LUNA_USDT
	TONCOIN_USD
	TONCOIN_USDT
	SHIBI_USD
	SHIBI_USDT
	MATIC_USD
	MATIC_USDT
	CRO_USD
	CRO_USDT
	BCH_USD
	BCH_USDT
	FTM_USD
	FTM_USDT
	XLM_USD
	XLM_USDT
	AXS_USD
	AXS_USDT
	ONE_USD
	ONE_USDT
	NEAR_USD
	NEAR_USDT
	ICP_USD
	ICP_USDT
	LEO_USD
	LEO_USDT
	IOTA_USD
	IOTA_USDT
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

const (
	BUY   = "buy"
	SELL  = "sell"
	LONG  = "long"
	SHORT = "short"
)
