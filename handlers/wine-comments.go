package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/dfreire/sunny/crud"
	"github.com/dfreire/sunny/middleware"
	"github.com/dfreire/sunny/model"
	"github.com/labstack/echo"
	"gopkg.in/Masterminds/squirrel.v1"
)

// http http://localhost:3500/customer-wine-comments?customerId="customer-1"
func GetCustomerWineComments(c echo.Context) error {
	log.Println("GetCustomerWineComments")

	customerId := c.QueryParam("customerId")

	db := c.Get(middleware.DB).(*sql.DB)

	comments := []model.WineComment{}

	rows, err := squirrel.
		Select("id", "wineId", "wineYear", "comment").
		From("WineComment").
		Where(squirrel.Eq{
			"customerId": customerId,
		}).
		RunWith(db).Query()
	if err != nil {
		log.Printf("error: %+v", err)
		return c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
	}

	for rows.Next() {
		var c model.WineComment
		rows.Scan(&c.Id, &c.WineId, &c.WineYear, &c.Comment)
		comments = append(comments, c)
	}

	return c.JSON(http.StatusOK, JsonResponse{Ok: true, Data: comments})
}

// http POST http://localhost:3500/signup-customer-with-wine-comment email="dario.freire@gmail.com" role="wine_lover" wineComments:='[{"wineId": "wine-1", "wineYear": 2015, "comment": "great"}, {"wineId": "wine-1", "wineYear": 2014, "comment": "fantastic"}]'
func SignupCustomerWithWineComment(c echo.Context) error {
	log.Println("SignupCustomerWithWineComment")

	var reqData struct {
		Email        string `json:"email"`
		Role         string `json:"role"`
		WineComments []struct {
			WineId   string `json:"wineId"`
			WineYear int    `json:"wineYear"`
			Comment  string `json:"comment"`
		} `json:"wineComments"`
	}

	c.Bind(&reqData)

	now := time.Now().Format(time.RFC3339)

	tx := c.Get(middleware.TX).(*sql.Tx)

	customerId, err := crud.UpsertRecord(
		tx,
		"Customer",
		crud.Record{
			"email": reqData.Email,
		},
		crud.Record{
			"createdAt": now,
		},
		crud.Record{
			"role": reqData.Role,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
		return err
	}

	for _, comment := range reqData.WineComments {
		_, err = crud.UpsertRecord(
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
