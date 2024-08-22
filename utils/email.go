package utils

import "gopkg.in/gomail.v2"

func SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "pawarganesh9510@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text.plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "pawarganesh9510@gmail.com", "nvxr drxz npaw gisg")
	return d.DialAndSend(m)
}
