package plat

import (
	"encoding/json"
	"github.com/go-resty/resty"
	"log"
	"time"
)

type Response struct {
	Code string        `json:"code"`
	Msg  string        `json:"msg"`
	Data []interface{} `json:"data"`
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
		data := &Response{}

		// log.Printf("Get Body %s", string(resp.Body()))
		if err = json.Unmarshal(resp.Body(), data); err != nil {
			log.Printf("Get err: %v", err)
		} else {
			dataBytes, _ := json.Marshal(data.Data)
			if err = json.Unmarshal(dataBytes, inter); err != nil {
				log.Printf("Get err: %v", err)
			}
		}
	}

	return err
}

// POST
func Post(url string, headers map[string]string, params interface{}, inter interface{}) error {
	if headers == nil {
		headers = make(map[string]string)
		headers["Content-Type"] = "application/x-www-form-urlencoded"
		headers["User-Agent"] = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36"
	}

	resp, err := resty.New().SetTimeout(time.Minute * 1).R().
		SetHeaders(headers).
		SetBody(params).
		Post(url)

	if err != nil {
		log.Printf("Post err: %+v", err)
		return err
	} else {
		data := &Response{}

		// log.Printf("Post Bodyody %s", string(resp.Body()))
		if err = json.Unmarshal(resp.Body(), data); err != nil {
			log.Printf("Post err: %v", err)
		} else {
			dataBytes, _ := json.Marshal(data.Data)
			if err = json.Unmarshal(dataBytes, inter); err != nil {
				log.Printf("Post err: %v", err)
			}
		}
	}
	return err
}