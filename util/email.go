package util

import (
	"gopkg.in/gomail.v2"
	"log"
)

func QQMail(from string, to []string, title string, body string, password string) {
	m := gomail.NewMessage()
	m.SetHeader(`From`, from) // 发送方
	m.SetHeader(`To`, to...)
	m.SetHeader(`Subject`, title)
	m.SetBody(`text/html`, body)
	err := gomail.NewDialer(`smtp.qq.com`, 25, from, password).DialAndSend(m)
	if err != nil {
		log.Printf("QQMail: %v", err)
	}
}
