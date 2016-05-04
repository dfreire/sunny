package handlers

import (
	"net/http"

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
		c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
		return err
	}

	err = m.OnSignUpCustomerWithWineComments(reqData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
		return err
	}

	return c.JSON(http.StatusOK, JsonResponse{Ok: true})
}
