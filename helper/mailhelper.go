package helper

import (
	"crypto/tls"
	"log"
	"net/smtp"
	"os"

	"github.com/amiraliio/goSchedule/model"
)

//SendEmail function
func SendEmail(email model.Email) bool {

	auth := smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_SERVER"))
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         os.Getenv("MAIL_SERVER"),
	}
	conn, err := tls.Dial("tcp", os.Getenv("MAIL_SERVER")+":"+os.Getenv("MAIL_PORT"), tlsConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	client, err := smtp.NewClient(conn, os.Getenv("MAIL_SERVER"))
	if err != nil {
		log.Fatal(err.Error())
	}
	if err = client.Auth(auth); err != nil {
		log.Fatal(err.Error())
	}
	if err = client.Mail(os.Getenv("MAIL_USERNAME")); err != nil {
		log.Fatal(err.Error())
	}
	if err = client.Rcpt(email.Receiver); err != nil {
		log.Fatal(err.Error())
	}
	w, err := client.Data()
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = w.Write([]byte(email.Body))
	if err != nil {
		log.Fatal(err.Error())
	}
	if err = w.Close(); err != nil {
		log.Fatal(err.Error())
	}
	if err = client.Quit(); err != nil {
		log.Fatal(err.Error())
	}
	return true
}
