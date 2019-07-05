package helper

import (
	"log"
	"net/smtp"
	"os"

	"github.com/amiraliio/goSchedule/model"
)

//SendEmail function
func SendEmail(email model.Email) {

	auth := smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_SERVER"))
	receiver := []string{email.Receiver}
	body := []byte(email.Body)
	err := smtp.SendMail(os.Getenv("MAIL_SERVER")+":"+os.Getenv("MAIL_PORT"), auth, os.Getenv("MAIL_FROM"), receiver, body)
	if err != nil {
		log.Fatal(err.Error())
	}
}
