package mailer

import (
	"log"

	"github.com/jordan-wright/email"
)

type fakeMailer struct {
}

func NewFakeMailer() Mailer {
	return &fakeMailer{}
}

func (self *fakeMailer) SendEmail(e *email.Email) error {
	m := make(map[string]interface{})
	m["from"] = e.From
	m["to"] = e.To
	m["cc"] = e.Cc
	m["bcc"] = e.Bcc
	m["subject"] = e.Subject
	m["html"] = string(e.HTML)
	log.Printf("Fake sent email: %+v", m)
	return nil
}
