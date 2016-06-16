package handlers

import (
	"net/http"

	"github.com/dfreire/sunny/middleware"
	"github.com/dfreire/sunny/operations"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// http http://localhost:3500/get-customers?appToken=2fe9a70a-46f2-4d00-88f2-6f66ed903426
func GetCustomers(c echo.Context) error {
	db := c.Get(middleware.DB).(*gorm.DB)

	customers, err := operations.GetCustomers(db)

	if err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{Ok: false})
		return err
	}

	return c.JSON(http.StatusOK, jsonResponse{Ok: true, Result: customers})
}
