package sys

import (
	"context"
	"crypto/tls"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/genv"
	gomail "gopkg.in/mail.v2"
)

func SendEmail(ctx context.Context, mailto []string, title string, body string) error {
	user := genv.Get("freekey_uname").String()
	pass := genv.Get("freekey_pass").String()
	host := "smtp.gmail.com"
	port := 587
	m := gomail.NewMessage(gomail.SetEncoding(gomail.Base64))
	m.SetHeader("From", m.FormatAddress(user, "freekey"))
	m.SetHeader("To", mailto...)
	m.SetHeader("Subject", title)
	m.SetBody("text/plain", body)
	d := gomail.NewDialer(host, port, user, pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
