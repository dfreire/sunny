package handlers

import (
	"database/sql"
	"net/http"

	"github.com/dfreire/sunny/middleware"
	"github.com/dfreire/sunny/model/commands"
	"github.com/labstack/echo"
)

// http POST http://localhost:3500/signup-customer-with-wine-comments email="dario.freire@gmail.com" roleId="wine_lover" wineComments:='[{"wineId": "wine-1", "wineYear": 2015, "comment": "great"}, {"wineId": "wine-1", "wineYear": 2014, "comment": "fantastic"}]'
func SignupCustomerWithWineComments(c echo.Context) error {
	tx := c.Get(middleware.TX).(*sql.Tx)

	var reqData commands.SignupCustomerWithWineCommentsRequestData
	c.Bind(&reqData)

	err := commands.SignupCustomerWithWineComments(tx, reqData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
		return err
	}

	return c.JSON(http.StatusOK, JsonResponse{Ok: true})
}
