package handlers

import (
	"net/http"

	"github.com/dfreire/sunny/commands"
	"github.com/dfreire/sunny/mailer"
	"github.com/dfreire/sunny/middleware"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// http POST http://localhost:3500/signup-customer-with-newsletter email="joe.doe@mailinator.com" roleId="wine_lover" language="pt"
func SignupCustomerWithNewsletter(c echo.Context) error {
	tx := c.Get(middleware.TX).(*gorm.DB)
	mx := c.Get(middleware.MAILER).(mailer.Mailer)

	var req commands.SignupCustomerWithNewsletterRequest
	c.Bind(&req)

	if err := commands.SignupCustomerWithNewsletter(tx, mx, req); err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{Ok: false})
		return err
	}

	return c.JSON(http.StatusOK, jsonResponse{Ok: true})
}
