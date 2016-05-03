package mailer

import (
	"net/smtp"
	"strconv"
	"strings"

	"github.com/jordan-wright/email"
)

type Mailer interface {
	Send(mail email.Email) error
}

type mailerImpl struct {
	hostAndPort string
	plainAuth   smtp.Auth
}

func NewMailer(host string, port int, email, password string) Mailer {
	hostAndPort := strings.Join([]string{
		host,
		strconv.Itoa(port),
	}, ":")

	plainAuth := smtp.PlainAuth(
		"", // identity
		email,
		password,
		host,
	)

	return &mailerImpl{hostAndPort, plainAuth}
}

func (self *mailerImpl) Send(e email.Email) error {
	return e.Send(self.hostAndPort, self.plainAuth)
}
