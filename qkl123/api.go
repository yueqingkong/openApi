package qkl123

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	"log"
	"time"
)

var (
	qklApi = "https://gate.qkl123.com"
)

type Api struct {
	ApiKey     string
	SecretKey  string
	Passphrase string
	Auth       string
	Token      string
}

func (self *Api) Candles(pair, level, end_time string) []Candle {
	api := "/w1/ticker/kline"
	api = fmt.Sprintf("%s?pair=%s&exchange=%s", api, pair, "okex")

	if level != "" {
		api = fmt.Sprintf("%s&level=%s", api, level)
	}
	if end_time != "" {
		api = fmt.Sprintf("%s&start_time=%s", api, end_time)
	}
	api = fmt.Sprintf("%s&count=%d&rehabilitation=%v", api, 200, false)

	header := make(map[string]string, 0)
	auth := self.Auth
	token := self.Token

	header["Authorization"] = auth
	header["gate-token"] = token

	var url = qklApi + api
	inst := make([]Candle, 0)

	log.Printf("url : %s", url)
	if err := Get(url, header, &inst); err != nil {
		log.Printf("Candles err: %+v", err)
	}
	return inst
}

// GET
func Get(url string, headers map[string]string, inter interface{}) error {
	if headers == nil {
		headers = make(map[string]string)
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

	resp, err := resty.New().SetTimeout(time.Minute * 1).R().
		SetHeaders(headers).
		Get(url)

	if err != nil {
		return err
	} else {
		log.Printf("Get Body %s", string(resp.Body()))
		if err = json.Unmarshal(resp.Body(), inter); err != nil {
			log.Printf("Get err: %v", err)
		}
	}

	return err
}
