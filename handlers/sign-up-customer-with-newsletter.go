package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/dfreire/sunny/commands"
	"github.com/dfreire/sunny/mailer"
	"github.com/dfreire/sunny/middleware"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// http POST http://localhost:3500/signup-customer-with-newsletter email="joe.doe@mailinator.com" roleId="wine_lover"
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

	to := []string{reqData.Email}
	templatePath := filepath.Join("templates", "mail", "pt", "on-sign-up-customer-with-newsletter-email.yaml")
	err = sendEmail(m, to, templatePath, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{Ok: false})
		return err
	}

	return c.JSON(http.StatusOK, jsonResponse{Ok: true})
}
