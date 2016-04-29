package handlers

import (
	"database/sql"
	"net/http"

	"github.com/dfreire/sunny/middleware"
	"github.com/dfreire/sunny/model/queries"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

// http http://localhost:3500/get-customers
func GetCustomers(c echo.Context) error {
	db := c.Get(middleware.DB).(*sql.DB)
	dbx := c.Get(middleware.DBX).(*sqlx.DB)

	data, err := queries.GetCustomers(db, dbx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
		return err
	}
	return c.JSON(http.StatusOK, JsonResponse{Ok: true, Data: data})
}
