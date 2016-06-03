package mailer

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/dfreire/sunny/template"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
)

type Mailer interface {
	Send(e *email.Email) error
}

func TemplateToEmail(e *email.Email, templateId, languageId string, templateValues interface{}) error {
	switch languageId {
	case "pt", "en":
		break
	default:
		languageId = "en"
	}

	templatePath := filepath.Join(
		viper.GetString("MAILER_TEMPLATES_DIR"),
		"mail",
		languageId,
		strings.Join([]string{templateId, "yaml"}, "."),
	)

	templateData, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return err
	}

	var mt struct {
		To      []string
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

	if mt.To != nil {
		e.To = mt.To
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
