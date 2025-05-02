package binance

// https://api.binance.com
// https://api-gcp.binance.com
// https://api1.binance.com
// https://api2.binance.com
// https://api3.binance.com
// https://api4.binance.com
var (
	bianceApi = "https://api.binance.com"
)

type Api struct {
	ApiKey     string
	SecretKey  string
	Passphrase string
}
