package okex

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/yueqingkong/openApi/plat"
	"github.com/yueqingkong/openApi/util"
	"log"
	"strings"
	"time"
)

var (
	okApi = "https://www.okx.com"
)

type Api struct {
	ApiKey     string
	SecretKey  string
	Passphrase string
}

// 初始化 Key
func (self *Api) InitApiKeys(apikey, secretkey, passphrase string) {
	self.ApiKey = apikey
	self.SecretKey = secretkey
	self.Passphrase = passphrase
}

func (self *Api) header(request string, path string, body interface{}) map[string]string {
	var paramString string
	if body != nil && body != "" {
		bodyBytes, _ := json.Marshal(body)
		paramString = string(bodyBytes)
	}

	timestamp := util.IsoTime(time.Now())
	comnination := timestamp + strings.ToUpper(request) + path + paramString
	sign, err := HmacSha256Base64Signer(comnination, self.SecretKey)

	if err != nil {
		log.Print("签名失败", err)
	}

	var headerMap = make(map[string]string, 0)
	headerMap["Accept"] = "application/json"
	headerMap["Content-Type"] = "application/json; charset=UTF-8"
	headerMap["Cookie"] = "locale=" + "en_US"

	headerMap["OK-ACCESS-KEY"] = self.ApiKey
	headerMap["OK-ACCESS-SIGN"] = sign
	headerMap["OK-ACCESS-TIMESTAMP"] = timestamp
	headerMap["OK-ACCESS-PASSPHRASE"] = self.Passphrase
	return headerMap
}

func HmacSha256Base64Signer(message string, secretKey string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(message))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}

func (self *Api) SysTime() []*SysTime {
	api := "/api/v5/public/time"

	var url = okApi + api
	inst := make([]*SysTime, 0)
	plat.Get(url, nil, &inst)
	return inst
}

// 获取所有产品行情信息
func (self *Api) Tickers(instType, uly string) []*Ticker {
	api := "/api/v5/market/tickers"

	params := make(map[string]string)
	params["instType"] = instType
	if uly != "" {
		params["uly"] = uly
	}
	api = api + parseParams(params)

	var url = okApi + api
	inst := make([]*Ticker, 0)
	plat.Get(url, nil, &inst)
	return inst
}

// 获取单个产品行情信息
func (self *Api) Ticker(instId string) []*Ticker {
	api := "/api/v5/market/ticker"

	params := make(map[string]string)
	params["instId"] = instId
	api = api + parseParams(params)

	var url = okApi + api
	inst := make([]*Ticker, 0)
	plat.Get(url, nil, &inst)
	return inst
}

// 获取所有可交易产品的信息列表
func (self *Api) Instruments(instType, uly, instId string) []*Instrument {
	api := "/api/v5/public/instruments"

	params := make(map[string]string)
	params["instType"] = instType
	if uly != "" {
		params["uly"] = uly
	}
	if instId != "" {
		params["instId"] = instId
	}
	api = api + parseParams(params)

	var url = okApi + api
	inst := make([]*Instrument, 0)
	plat.Get(url, nil, &inst)
	return inst
}

// 查询单个交易产品的最高买价和最低卖价
// 仅适用于交割/永续/期权
func (self *Api) PriceLimit(instId string) []*PriceLimit {
	api := "/api/v5/public/price-limit"

	params := make(map[string]string)
	params["instId"] = instId
	api = api + parseParams(params)

	var url = okApi + api
	inst := make([]*PriceLimit, 0)
	plat.Get(url, nil, &inst)
	return inst
}

// 获取预估交割/行权价格
// 获取交割合约和期权预估交割/行权价。交割/行权预估价只有交割/行权前一小时才有返回值
func (self *Api) EstimatedPrice(instId string) []*EstimatedPrice {
	api := "/api/v5/public/estimated-price"

	params := make(map[string]string)
	params["instId"] = instId
	api = api + parseParams(params)

	var url = okApi + api
	inst := make([]*EstimatedPrice, 0)
	plat.Get(url, nil, &inst)
	return inst
}

// 获取法币汇率
// 该接口提供的是2周的平均汇率数据
func (self *Api) ExchangeRate() []*UsdCny {
	api := "/api/v5/market/exchange-rate"

	var url = okApi + api
	inst := make([]*UsdCny, 0)
	plat.Get(url, nil, &inst)
	return inst
}

// 获取所有交易产品K线数据
// 获取K线数据。K线数据按请求的粒度分组返回，K线数据每个粒度最多可获取最近1440条。
// bar [1m/3m/5m/15m/30m/1H/2H/4H/6H/12H/1D/1W/1M/3M/6M/1Y]
func (self *Api) Candles(instId, bar, before string) [][]string {
	api := "/api/v5/market/candles"

	params := make(map[string]string)
	params["instId"] = instId
	if bar != "" {
		params["bar"] = bar
	}
	if before != "" {
		params["before"] = before
	}
	api = api + parseParams(params)

	var url = okApi + api
	inst := make([][]string, 0)
	plat.Get(url, nil, &inst)
	return inst
}

// 获取账户中资金余额信息
func (self *Api) balance(ccy string) []*Balance {
	var api = "/api/v5/account/balance"

	params := make(map[string]string)
	if ccy != "" {
		params["ccy"] = ccy
	}
	api = api + parseParams(params)

	var url = okApi + api
	result := make([]*Balance, 0)
	plat.Get(url, self.header("get", api, nil), &result)
	return result
}

// 获取订单信息
func (self *Api) OrderInfo(instId, orderId string) *OrderInfo {
	var api = "/api/v5/trade/order"

	params := make(map[string]string)
	if instId != "" {
		params["instId"] = instId
	}
	if orderId != "" {
		params["ordId"] = orderId
	}
	api = api + parseParams(params)

	var url = okApi + api
	result := &OrderInfo{}
	plat.Get(url, self.header("get", api, nil), result)
	return result
}

// 设置杠杆倍数
func (self *Api) setLeverage(instId, lever, mgnMode, posSide string) []*SetLeverage {
	var api = "/api/v5/account/set-leverage"
	var url = okApi + api

	params := make(map[string]interface{})
	params["instId"] = instId
	params["lever"] = lever
	params["mgnMode"] = mgnMode
	params["posSide"] = posSide

	results := make([]*SetLeverage, 0)
	plat.Post(url, self.header("post", api, params), params, &results)
	return results
}

// 下单
// 持仓方向，单向持仓模式下此参数非必填，如果填写仅可以选择net；在双向持仓模式下必填，且仅可选择 long 或 short。
// 双向持仓模式下，side和posSide需要进行组合
// 开多：买入开多（side 填写 buy； posSide 填写 long ）
// 开空：卖出开空（side 填写 sell； posSide 填写 short ）
// 平多：卖出平多（side 填写 sell；posSide 填写 long ）
// 平空：买入平空（side 填写 buy； posSide 填写 short ）

// 交易数量，表示要购买或者出售的数量。
// 当币币/币币杠杆以限价买入和卖出时，指交易货币数量。
// 当币币/币币杠杆以市价买入时，指计价货币的数量。
// 当币币/币币杠杆以市价卖出时，指交易货币的数量。
// 当交割、永续、期权买入和卖出时，指合约张数。
func (self *Api) Order(instId, tdMode, side, posSide string, px, sz float32) []*OrderRes {
	var api = "/api/v5/trade/order"
	var url = okApi + api

	params := make(map[string]interface{}, 0)
	params["instId"] = instId
	params["tdMode"] = tdMode
	params["posSide"] = posSide

	params["side"] = side
	params["ordType"] = "limit"
	params["sz"] = sz
	params["px"] = px

	results := make([]*OrderRes, 0)
	plat.Post(url, self.header("post", api, params), params, &results)
	return results
}

func parseParams(params map[string]string) string {
	url := "?"
	for k, v := range params {
		url = url + k + "=" + v + "&"
	}
	return url[:len(url)-1]
}
