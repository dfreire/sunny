package handlers

import (
	"database/sql"
	"net/http"

	"github.com/dfreire/sunny/middleware"
	"github.com/dfreire/sunny/model/queries"
	"github.com/labstack/echo"
)

// http http://localhost:3500/get-customers
func GetCustomers(c echo.Context) error {
	db := c.Get(middleware.DB).(*sql.DB)

	data, err := queries.GetCustomers(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
		return err
	}
	return c.JSON(http.StatusOK, JsonResponse{Ok: true, Data: data})
}
