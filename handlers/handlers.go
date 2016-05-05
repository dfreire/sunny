package handlers

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/dfreire/sunny/mailer"
	"github.com/dfreire/sunny/template"
	"github.com/jordan-wright/email"
)

type jsonResponse struct {
	Ok    bool        `json:"ok"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func sendEmail(m mailer.Mailer, to []string, templatePath string, templateValues interface{}) error {
	templateData, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return err
	}

	e := email.Email{}
	err = yaml.Unmarshal([]byte(templateData), &e)
	if err != nil {
		return err
	}

	e.HTML, err = template.RenderBytes(e.HTML, templateValues)
	if err != nil {
		return err
	}

	return m.Send(e)
}
