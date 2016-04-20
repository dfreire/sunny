package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/dfreire/sunny/crud"
	"github.com/dfreire/sunny/middleware"
	"github.com/labstack/echo"
)

// http POST http://localhost:3500/signup-customer-with-newsletter email="dario.freire@gmail.com" roleId="wine_lover"
func SignupCustomerWithNewsletter(c echo.Context) error {
	var reqData struct {
		Email  string `json:"email"`
		RoleId string `json:"roleId"`
	}

	c.Bind(&reqData)

	now := time.Now().Format(time.RFC3339)

	tx := c.Get(middleware.TX).(*sql.Tx)

	_, err := crud.Upsert(
		tx,
		"Customer",
		crud.Record{
			"email": reqData.Email,
		},
		crud.Record{
			"createdAt":      now,
			"signupOriginId": "newsletter",
		},
		crud.Record{
			"roleId": reqData.RoleId,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
		return err
	}

	return c.JSON(http.StatusOK, JsonResponse{Ok: true})
}