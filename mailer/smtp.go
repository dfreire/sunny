package mailer

import (
	"net/smtp"
	"strconv"
	"strings"

	"github.com/jordan-wright/email"
)

type mailerImpl struct {
	hostAndPort string
	plainAuth   smtp.Auth
}

func NewMailer(host string, port int, login, password string) Mailer {
	hostAndPort := strings.Join([]string{
		host,
		strconv.Itoa(port),
	}, ":")

	plainAuth := smtp.PlainAuth(
		"", // identity
		login,
		password,
		host,
	)

	return &mailerImpl{hostAndPort, plainAuth}
}

func (self *mailerImpl) SendEmail(e *email.Email) error {
	return e.Send(self.hostAndPort, self.plainAuth)
}
