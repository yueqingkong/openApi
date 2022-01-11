package util

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"log"
)

func QQMail(from string, to []string, title string, body string, password string) error {
	m := gomail.NewMessage()
	m.SetHeader(`From`, from) // 发送方
	m.SetHeader(`To`, to...)
	m.SetHeader(`Subject`, title)
	m.SetBody(`text/html`, body)
	dialog := gomail.NewDialer(`smtp.qq.com`, 25, from, password)
	dialog.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := dialog.DialAndSend(m); err != nil {
		log.Printf("QQMail: %v", err)
		return err
	}
	return nil
}
