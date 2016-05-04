package mailer

import (
	"io/ioutil"
	"net/smtp"
	"path/filepath"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/dfreire/sunny/commands"
	"github.com/jordan-wright/email"
)

type Mailer interface {
	OnSignUpCustomerWithNewsletter(reqData commands.SignupCustomerWithNewsletterRequestData) error
	OnSignUpCustomerWithWineComments(reqData commands.SignupCustomerWithWineCommentsRequestData) error
}

type mailerImpl struct {
	from        string
	hostAndPort string
	plainAuth   smtp.Auth
}

func NewMailer(host string, port int, from, login, password string) Mailer {
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

	return &mailerImpl{from, hostAndPort, plainAuth}
}

func (self *mailerImpl) send(e *email.Email) error {
	e.From = self.from
	return e.Send(self.hostAndPort, self.plainAuth)
}

type mailTemplate struct {
	Subject string
	Body    string
}

func (self *mailerImpl) OnSignUpCustomerWithNewsletter(reqData commands.SignupCustomerWithNewsletterRequestData) error {
	data, err := ioutil.ReadFile(filepath.Join("templates", "on-sign-up-customer-with-newsletter-email.pt.yaml"))
	if err != nil {
		return err
	}

	var mt mailTemplate
	err = yaml.Unmarshal([]byte(data), &mt)
	if err != nil {
		return err
	}

	e := email.NewEmail()
	e.To = []string{reqData.Email}
	// e.Bcc = mail.Bcc
	e.Subject = mt.Subject
	e.HTML = []byte(mt.Body)
	return self.send(e)
}

func (self *mailerImpl) OnSignUpCustomerWithWineComments(reqData commands.SignupCustomerWithWineCommentsRequestData) error {
	data, err := ioutil.ReadFile(filepath.Join("templates", "on-sign-up-customer-with-wine-comments-email.pt.yaml"))
	if err != nil {
		return err
	}

	var mt mailTemplate
	err = yaml.Unmarshal([]byte(data), &mt)
	if err != nil {
		return err
	}

	e := email.NewEmail()
	e.To = []string{reqData.Email}
	// e.Bcc = mail.Bcc
	e.Subject = mt.Subject
	e.HTML = []byte(mt.Body)
	return self.send(e)
}
