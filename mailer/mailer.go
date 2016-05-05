package mailer

import (
	"io/ioutil"
	"net/smtp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/dfreire/sunny/template"
	"github.com/jordan-wright/email"
)

type Mailer interface {
	Send(e email.Email) error
}

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

func (self *mailerImpl) Send(e email.Email) error {
	return e.Send(self.hostAndPort, self.plainAuth)
}

func TemplateToEmail(e *email.Email, templatePath string, templateValues interface{}) error {
	templateData, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return err
	}

	var mt struct {
		From    string
		Cc      []string
		Bcc     []string
		Subject string
		Html    string
	}

	err = yaml.Unmarshal([]byte(templateData), &mt)
	if err != nil {
		return err
	}

	if mt.From != "" {
		e.From = mt.From
	}

	if mt.Cc != nil {
		e.Cc = mt.Cc
	}

	if mt.Bcc != nil {
		e.Bcc = mt.Bcc
	}

	e.Subject = mt.Subject

	html, err := template.Render(mt.Html, templateValues)
	if err != nil {
		return err
	}

	e.HTML = []byte(html)

	return nil
}
