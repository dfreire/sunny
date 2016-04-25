package handlers

import (
	"database/sql"
	"net/http"

	"github.com/dfreire/sunny/middleware"
	"github.com/dfreire/sunny/model/commands"
	"github.com/labstack/echo"
)

// http POST http://localhost:3500/signup-customer-with-newsletter email="dario.freire@gmail.com" roleId="wine_lover"
func SignupCustomerWithNewsletter(c echo.Context) error {
	tx := c.Get(middleware.TX).(*sql.Tx)

	var reqData commands.SignupCustomerWithNewsletterRequestData
	c.Bind(&reqData)

	err := commands.SignupCustomerWithNewsletter(tx, reqData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
		return err
	}

	return c.JSON(http.StatusOK, JsonResponse{Ok: true})
}
