package okex

import (
	"bytes"
	"compress/flate"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/yueqingkong/openApi/conset"
	"github.com/yueqingkong/openApi/db"
	"github.com/yueqingkong/openApi/util"
	"github.com/yueqingkong/openApi/ws"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"
)

var (
	wsPublicURL = "wss://ws.okx.com:8443/ws/v5/public"
)

type wsResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Arg  struct {
		Channel string `json:"channel"`
		InstID  string `json:"instId"`
	} `json:"arg"`
	Data []struct {
		InstType  string `json:"instType"`
		InstID    string `json:"instId"`
		Last      string `json:"last"`
		LastSz    string `json:"lastSz"`
		AskPx     string `json:"askPx"`
		AskSz     string `json:"askSz"`
		BidPx     string `json:"bidPx"`
		BidSz     string `json:"bidSz"`
		Open24H   string `json:"open24h"`
		High24H   string `json:"high24h"`
		Low24H    string `json:"low24h"`
		SodUtc0   string `json:"sodUtc0"`
		SodUtc8   string `json:"sodUtc8"`
		VolCcy24H string `json:"volCcy24h"`
		Vol24H    string `json:"vol24h"`
		Ts        string `json:"ts"`
	} `json:"data"`
}

type Ws struct {
	*ws.WsBuilder
	once       *sync.Once
	WsConn     *ws.WsConn
	respHandle func(symbol conset.SYMBOL, price float32)

	ApiKey     string
	SecretKey  string
	Passphrase string
}

// 初始化 Key
func (self *Ws) InitWsKeys(apikey, secretkey, passphrase string) {
	self.ApiKey = apikey
	self.SecretKey = secretkey
	self.Passphrase = passphrase

	self.once = new(sync.Once)

	self.WsBuilder = ws.NewWsBuilder().
		WsUrl(wsPublicURL).
		ReconnectInterval(time.Second).
		AutoReconnect().
		Heartbeat(func() []byte { return []byte("ping") }, 28*time.Second).
		DecompressFunc(FlateDecompress).ProtoHandleFunc(self.handle)
}

func (self *Ws) UUID() string {
	return strings.Replace(uuid.New().String(), "-", "", 32)
}

func FlateDecompress(data []byte) ([]byte, error) {
	return ioutil.ReadAll(flate.NewReader(bytes.NewReader(data)))
}

func (self *Ws) ConnectWs() {
	self.once.Do(func() {
		self.WsConn = self.WsBuilder.Build()
	})
}

func (self *Ws) handle(msg []byte) error {
	log.Printf("[ws] [response] %v", string(msg))
	if string(msg) == "pong" {
		return nil
	}

	var wsResp wsResponse
	err := json.Unmarshal(msg, &wsResp)
	if err != nil {
		log.Print(err)
		return err
	}

	if wsResp.Code != "" {
		log.Print(string(msg))
		return fmt.Errorf("%s", string(msg))
	}

	if len(wsResp.Data) != 0 {
		for _, data := range wsResp.Data {
			lastPos := strings.LastIndex(data.InstID, "-")
			s := data.InstID[0:lastPos]
			s = strings.ReplaceAll(s, "-", "_")
			s = strings.ToLower(s)

			symbol := db.StringToSymbol(s)
			price := util.Float32(data.Last)
			log.Printf("handle ws data symbol: %v, price: %v", db.SymbolToString(symbol), price)

			if self.respHandle != nil {
				self.respHandle(symbol, price)
			}
		}
		return nil
	}

	return fmt.Errorf("unknown websocket message: %v", wsResp)
}

func (self *Ws) Subscribe(sub map[string]interface{}) error {
	self.ConnectWs()
	return self.WsConn.Subscribe(sub)
}

func (self *Ws) WsTickers(instIds []string, f func(conset.SYMBOL, float32)) error {
	self.ConnectWs()
	self.respHandle = f

	params := make(map[string]interface{})
	params["op"] = "subscribe"

	args := make([]interface{}, 0)
	for _, instId := range instIds {
		arg := make(map[string]interface{})
		arg["channel"] = "tickers"
		arg["instId"] = instId
		args = append(args, arg)
	}
	params["args"] = args

	return self.WsConn.Subscribe(params)
}
