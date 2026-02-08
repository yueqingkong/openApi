package conset

type CCY string

// 推荐使用  dot:= CCY(“DOT”) 方式
// COIN
const (
	BTC CCY = "BTC"
	ETH CCY = "ETH"
	LTC CCY = "LTC"
	DOT CCY = "DOT"
	ADA CCY = "ADA"
	SOL CCY = "SOL"
	BNB CCY = "BNB"
)

// U
const (
	USD  CCY = "USD"
	USDT CCY = "USDT"
	USDC CCY = "USDC"
)

const (
	PURCHASE = "purchase" //申购
	REDEMPT  = "redempt"  //赎回
)

const (
	FUND_ACCOUNT  = "6"  // 6：资金账户
	TRADE_ACCOUNT = "18" // 18：交易账户
)

const (
	TRANSFER_INTERNAL = "0"  // 0：账户内划转
)
