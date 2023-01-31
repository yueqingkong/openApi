package conset

type PLAT int32
type PERIOD int32
type TIMES int32
type OPERATION int32
type CCY int32

const (
	_ = iota
	ENABLE
	DISABLE
)

const (
	OKEX PLAT = iota
	QKL123
)

const (
	_ CCY = iota
	USD
	USDT
	BTC
	ETH
	LTC
	DOT
	DOGE
	ADA
	FIL
	ATOM
	XRP
	LINK
	EOS
	UNI
	CRV
	THETA
	ALGO
	ETC
	SAND
	SOL
	XTZ
	DASH
	TRX
	XMR
	MANA
	SUSHI
	ZEC
	LUNA
	TONCOIN
	SHIBI
	MATIC
	CRO
	BCH
	FTM
	XLM
	AXS
	ONE
	NEAR
	ICP
	LEO
	IOTA
	SNX
	AVAX
	WAVES
	AAVE
	BSV
	XCH
	ENS
	COMP
	EGLD
)

// 交易类型
const (
	_            PERIOD = iota
	SPOT                // 现货
	SWAP                // 永续
	WEEK                // 当周
	WEEK_NEXT           // 次周
	QUARTER             // 当季
	QUARTER_NEXT        // 次季
)

const (
	_     TIMES = iota
	MIN_5       // 5min
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
