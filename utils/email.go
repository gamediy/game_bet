package utils

import (
	"crypto/tls"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) error {
	mailer := gomail.NewMessage()
	account := viper.GetString("email.account")
	passwd := viper.GetString("email.passwd")
	mailer.SetHeader("From", account)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/plain", body)
	// 配置 Gmail SMTP 服务器信息
	dialer := gomail.NewDialer("smtp.gmail.com", 587, account, passwd)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// 发送邮件
	err := dialer.DialAndSend(mailer)

	return err
}
