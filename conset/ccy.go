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
