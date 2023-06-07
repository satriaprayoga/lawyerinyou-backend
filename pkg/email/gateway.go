package email

import (
	"lawyerinyou-backend/pkg/settings"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, htmlBody string) error {
	//fmt.Printf("%s %s %s", to, subject, htmlBody)
	smtp := settings.AppConfigSetting.SMTP
	//from := mail.Address{
	//	Name:    smtp.Identity,
	//	Address: smtp.Sender,
	//}
	m := gomail.NewMessage()
	m.Reset()
	m.SetHeader("From", smtp.Sender)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)

	m.SetBody("text/html", htmlBody)

	d := gomail.NewDialer(smtp.Server, smtp.Port, smtp.User, smtp.Passwd)
	return d.DialAndSend(m)
}
