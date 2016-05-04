package handlers

import (
	"net/http"

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
		c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
		return err
	}

	c.JSON(http.StatusOK, JsonResponse{Ok: true})

	return m.OnSignUpCustomerWithNewsletter(reqData)
}
