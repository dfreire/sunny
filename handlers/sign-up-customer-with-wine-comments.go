package handlers

import (
	"net/http"

	"github.com/dfreire/sunny/mailer"
	"github.com/dfreire/sunny/middleware"
	"github.com/dfreire/sunny/operations"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// http POST http://localhost:3500/signup-customer-with-wine-comments email="joe.doe@mailinator.com" roleId="wine_lover" language="en" wineComments:='[{"wineId": "wine-1", "wineName": "Soalheiro Alvarinho", "wineYear": 2015, "comment": "great"}, {"wineId": "wine-1", "wineName": "Soalheiro Alvarinho", "wineYear": 2014, "comment": "fantastic"}]'
func SignupCustomerWithWineComments(c echo.Context) error {
	tx := c.Get(middleware.TX).(*gorm.DB)
	mx := c.Get(middleware.MAILER).(mailer.Mailer)

	var req operations.SignupCustomerWithWineCommentsRequest
	c.Bind(&req)

	err := operations.SignupCustomerWithWineComments(tx, mx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{Ok: false})
		return err
	}

	return c.JSON(http.StatusOK, jsonResponse{Ok: true})
}
