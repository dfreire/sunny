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

// http POST http://localhost:3500/signup-customer-with-wine-comments email="joe.doe@mailinator.com" roleId="wine_lover" wineComments:='[{"wineId": "wine-1", "wineYear": 2015, "comment": "great"}, {"wineId": "wine-1", "wineYear": 2014, "comment": "fantastic"}]'
func SignupCustomerWithWineComments(c echo.Context) error {
	tx := c.Get(middleware.TX).(*gorm.DB)
	m := c.Get(middleware.MAILER).(mailer.Mailer)

	var reqData commands.SignupCustomerWithWineCommentsRequestData
	c.Bind(&reqData)

	err := commands.SignupCustomerWithWineComments(tx, reqData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{Ok: false})
		return err
	}

	to := []string{reqData.Email}
	templatePath := filepath.Join("templates", "mail", "pt", "on-sign-up-customer-with-wine-comments-email.yaml")
	err = sendEmail(m, to, templatePath, reqData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{Ok: false})
		return err
	}

	return c.JSON(http.StatusOK, jsonResponse{Ok: true})
}
