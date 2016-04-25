package handlers

import (
	"database/sql"
	"net/http"

	"github.com/dfreire/sunny/middleware"
	"github.com/dfreire/sunny/model/queries"
	"github.com/labstack/echo"
)

// http http://localhost:3500/get-wine-comments-by-customer-id?customerId=customer-1
func GetWineCommentsByCustomerId(c echo.Context) error {
	db := c.Get(middleware.DB).(*sql.DB)
	customerId := c.QueryParam("customerId")
	comments, err := queries.GetWineCommentsByCustomerId(db, customerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
		return err
	}
	return c.JSON(http.StatusOK, JsonResponse{Ok: true, Data: comments})
}
