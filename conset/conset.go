package conset

import "errors"

type PLAT int32
type PERIOD int32
type TIMES int32
type OPERATION int32
type CCY string

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
	USD  CCY = "USD"
	USDT CCY = "USDT"
)

const (
	BTC     CCY = "BTC"
	ETH     CCY = "ETH"
	LTC     CCY = "LTC"
	DOT     CCY = "DOT"
	DOGE    CCY = "DOGE"
	ADA     CCY = "ADA"
	FIL     CCY = "FIL"
	ATOM    CCY = "ATOM"
	XRP     CCY = "XRP"
	LINK    CCY = "LINK"
	EOS     CCY = "EOS"
	UNI     CCY = "UNI"
	CRV     CCY = "CRV"
	THETA   CCY = "THETA"
	ALGO    CCY = "ALGO"
	ETC     CCY = "ETC"
	SAND    CCY = "SAND"
	SOL     CCY = "SOL"
	XTZ     CCY = "XTZ"
	DASH    CCY = "DASH"
	TRX     CCY = "TRX"
	XMR     CCY = "XMR"
	MANA    CCY = "MANA"
	SUSHI   CCY = "SUSHI"
	ZEC     CCY = "ZEC"
	LUNA    CCY = "LUNA"
	TONCOIN CCY = "TONCOIN"
	SHIBI   CCY = "SHIBI"
	MATIC   CCY = "MATIC"
	CRO     CCY = "CRO"
	BCH     CCY = "BCH"
	FTM     CCY = "FTM"
	XLM     CCY = "XLM"
	AXS     CCY = "AXS"
	ONE     CCY = "ONE"
	NEAR    CCY = "NEAR"
	ICP     CCY = "ICP"
	LEO     CCY = "LEO"
	IOTA    CCY = "IOTA"
	SNX     CCY = "SNX"
	AVAX    CCY = "AVAX"
	WAVES   CCY = "WAVES"
	AAVE    CCY = "AAVE"
	BSV     CCY = "BSV"
	XCH     CCY = "XCH"
	ENS     CCY = "ENS"
	COMP    CCY = "COMP"
	EGLD    CCY = "EGLD"
	_1INCH  CCY = "1INCH"
	AIDOGE  CCY = "AIDOGE"
	ALPHA   CCY = "ALPHA"
	ANT     CCY = "ANT"
	APE     CCY = "APE"
	API3    CCY = "API3"
	APT     CCY = "APT"
	AR      CCY = "AR"
	ARB     CCY = "ARB"
	BADGER  CCY = "BADGER"
	BAL     CCY = "BAL"
	BAND    CCY = "BAND"
	BAT     CCY = "BAT"
	BICO    CCY = "BICO"
	BLUR    CCY = "BLUR"
	BNB     CCY = "BNB"
	BNT     CCY = "BNT"
	CEL     CCY = "CEL"
	CELO    CCY = "CELO"
	CETUS   CCY = "CETUS"
	CFX     CCY = "CFX"
	CHZ     CCY = "CHZ"
	CORE    CCY = "CORE"
	CSPR    CCY = "CSPR"
	CVC     CCY = "CVC"
	DORA    CCY = "DORA"
	DYDX    CCY = "DYDX"
	ENJ     CCY = "ENJ"
	ETHW    CCY = "ETHW"
	FITFI   CCY = "FITFI"
	FLM     CCY = "FLM"
	FLOKI   CCY = "FLOKI"
	GALA    CCY = "GALA"
	GFT     CCY = "GFT"
	GMT     CCY = "GMT"
	GMX     CCY = "GMX"
	GODS    CCY = "GODS"
	GRT     CCY = "GRT"
	IMX     CCY = "IMX"
	IOST    CCY = "IOST"
	JST     CCY = "JST"
	KISHU   CCY = "KISHU"
	KLAY    CCY = "KLAY"
	KNC     CCY = "KNC"
	KSM     CCY = "KSM"
	LDO     CCY = "LDO"
	LOOKS   CCY = "LOOKS"
	LPT     CCY = "LPT"
	LRC     CCY = "LRC"
	LUNC    CCY = "LUNC"
	MAGIC   CCY = "MAGIC"
	MASK    CCY = "v"
	MINA    CCY = "MINA"
	MKR     CCY = "MKR"
	NEO     CCY = "NEO"
	NFT     CCY = "NFT"
	OMG     CCY = "OMG"
	ONT     CCY = "ONT"
	OP      CCY = "OP"
	ORDI    CCY = "ORDI"
	PEOPLE  CCY = "PEOPLE"
	PERP    CCY = "PERP"
	QTUM    CCY = "QTUM"
	RDNT    CCY = "RDNT"
	REN     CCY = "REN"
	RSR     CCY = "RSR"
	RVN     CCY = "RVN"
	SHIB    CCY = "SHIB"
	SLP     CCY = "SLP"
	STARL   CCY = "STARL"
	STORJ   CCY = "STORJ"
	STX     CCY = "STX"
	SUI     CCY = "SUI"
	SWEAT   CCY = "SWEAT"
	TON     CCY = "TON"
	TRB     CCY = "TRB"
	UMA     CCY = "UMA"
	USTC    CCY = "USTC"
	WLD     CCY = "WLD"
	WOO     CCY = "WOO"
	YFI     CCY = "YFI"
	YFII    CCY = "YFII"
	YGG     CCY = "YGG"
	ZEN     CCY = "ZEN"
	ZIL     CCY = "ZIL"
	ZRX     CCY = "ZRX"
)

// 交易类型
const (
	_            PERIOD = iota
	SPOT                // 现货
	SWAP                // 永续
	MARGIN              // 币币杠杆
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

var (
	NOT_FOUND    = errors.New("not found")
	REQUEST_FAIL = errors.New("requst failt")
)
