package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/dfreire/sunny/commands"
	"github.com/dfreire/sunny/mailer"
	"github.com/dfreire/sunny/middleware"
	"github.com/jinzhu/gorm"
	"github.com/jordan-wright/email"
	"github.com/labstack/echo"
)

// http POST http://localhost:3500/signup-customer-with-newsletter email="joe.doe@mailinator.com" roleId="wine_lover"
func SignupCustomerWithNewsletter(c echo.Context) error {
	tx := c.Get(middleware.TX).(*gorm.DB)

	var reqData commands.SignupCustomerWithNewsletterRequestData
	c.Bind(&reqData)

	err := commands.SignupCustomerWithNewsletter(tx, reqData)

	type mailTemplate struct {
		Subject string
		Body    string
	}
	var mt mailTemplate

	data, err := ioutil.ReadFile(filepath.Join("templates", "on-sign-up-customer-with-newsletter-email.pt.yaml"))
	err = yaml.Unmarshal([]byte(data), &mt)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(mt)

	e := email.NewEmail()
	e.To = []string{reqData.Email}
	// e.Bcc = mail.Bcc
	// e.Subject = mail.Subject
	// e.HTML = []byte(mail.Body)
	c.Get(middleware.MAILER).(mailer.Mailer).Send(e)

	if err != nil {
		c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
		return err
	}

	return c.JSON(http.StatusOK, JsonResponse{Ok: true})
}
