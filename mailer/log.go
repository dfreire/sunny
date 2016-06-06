package mailer

import (
	"log"

	"github.com/jordan-wright/email"
)

type logMailer struct {
}

func NewLogMailer() Mailer {
	return &logMailer{}
}

func (self *logMailer) SendEmail(e *email.Email) error {
	log.Printf("From: %+v", e.From)
	log.Printf("To: %+v", e.To)
	log.Printf("Cc: %+v", e.Cc)
	log.Printf("Bcc: %+v", e.Bcc)
	log.Printf("Subject: %+v", e.Subject)
	log.Printf("HTML: %+v", string(e.HTML))
	for _, attachment := range e.Attachments {
		log.Printf("Attachment: %+v", attachment.Filename)
	}
	return nil
}
