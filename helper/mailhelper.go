package helper

import (
	"log"
	"net/smtp"
	"os"

	"github.com/amiraliio/goSchedule/model"
)

//SendEmail function
func SendEmail(email model.Email) bool {

	msg := "From: " + os.Getenv("MAIL_USERNAME") + "\n" +
		"To: " + email.Receiver + "\n" +
		"Subject:" + email.Subject + "\n\n" + email.Body

	err := smtp.SendMail(os.Getenv("MAIL_SERVER")+":"+os.Getenv("MAIL_PORT"),
		smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_SERVER")),
		os.Getenv("MAIL_USERNAME"), []string{email.Receiver}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return false
	}

	return true
}
