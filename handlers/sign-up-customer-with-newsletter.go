package handlers

import (
	"net/http"

	"github.com/dfreire/sunny/commands"
	"github.com/dfreire/sunny/mailer"
	"github.com/dfreire/sunny/middleware"
	"github.com/jinzhu/gorm"
	"github.com/jordan-wright/email"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

// http POST http://localhost:3500/signup-customer-with-newsletter email="joe.doe@mailinator.com" roleId="wine_lover" language="pt"
func SignupCustomerWithNewsletter(c echo.Context) error {
	tx := c.Get(middleware.TX).(*gorm.DB)
	m := c.Get(middleware.MAILER).(mailer.Mailer)

	var reqData commands.SignupCustomerWithNewsletterRequestData
	c.Bind(&reqData)

	err := commands.SignupCustomerWithNewsletter(tx, reqData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{Ok: false})
		return err
	}

	err = sendMailAfterSignupCustomerWithNewsletter(m, reqData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{Ok: false})
		return err
	}

	return c.JSON(http.StatusOK, jsonResponse{Ok: true})
}

func sendMailAfterSignupCustomerWithNewsletter(m mailer.Mailer, reqData commands.SignupCustomerWithNewsletterRequestData) error {
	e := email.Email{
		To:  []string{reqData.Email},
		Bcc: viper.GetStringSlice("NOTIFICATION_EMAILS"),
	}

	err := mailer.TemplateToEmail(&e, "on-sign-up-customer-with-newsletter-email", "pt", nil)
	if err != nil {
		return err
	}

	return m.Send(&e)
}
