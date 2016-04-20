package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/dfreire/sunny/crud"
	"github.com/dfreire/sunny/middleware"
	"github.com/labstack/echo"
)

// http POST http://localhost:3500/signup-customer-with-wine-comment email="dario.freire@gmail.com" roleId="wine_lover" wineComments:='[{"wineId": "wine-1", "wineYear": 2015, "comment": "great"}, {"wineId": "wine-1", "wineYear": 2014, "comment": "fantastic"}]'
func SignupCustomerWithWineComment(c echo.Context) error {
	log.Println("SignupCustomerWithWineComment")

	var reqData struct {
		Email        string `json:"email"`
		RoleId       string `json:"roleId"`
		WineComments []struct {
			WineId   string `json:"wineId"`
			WineYear int    `json:"wineYear"`
			Comment  string `json:"comment"`
		} `json:"wineComments"`
	}

	c.Bind(&reqData)

	now := time.Now().Format(time.RFC3339)

	tx := c.Get(middleware.TX).(*sql.Tx)

	customerId, err := crud.Upsert(
		tx,
		"Customer",
		crud.Record{
			"email": reqData.Email,
		},
		crud.Record{
			"createdAt":      now,
			"signupOriginId": "wine_comment",
		},
		crud.Record{
			"roleId": reqData.RoleId,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
		return err
	}

	for _, comment := range reqData.WineComments {
		_, err = crud.Upsert(
			tx,
			"WineComment",
			crud.Record{
				"customerId": customerId,
				"wineId":     comment.WineId,
				"wineYear":   comment.WineYear,
			},
			crud.Record{
				"createdAt": now,
			},
			crud.Record{
				"updatedAt": now,
				"comment":   comment.Comment,
			},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
			return err
		}
	}

	log.Printf("customerId: %+v", customerId)

	return c.JSON(http.StatusOK, JsonResponse{Ok: true})
}
