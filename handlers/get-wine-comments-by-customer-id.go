package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dfreire/sunny/middleware"
	"github.com/dfreire/sunny/model"
	"github.com/labstack/echo"
	"gopkg.in/Masterminds/squirrel.v1"
)

// http http://localhost:3500/get-wine-comments-by-customer-id?customerId=customer-1
func GetWineCommentsByCustomerId(c echo.Context) error {
	customerId := c.QueryParam("customerId")

	db := c.Get(middleware.DB).(*sql.DB)

	rows, err := squirrel.
		Select("id", "wineId", "wineYear", "comment").
		From("WineComment").
		Where(squirrel.Eq{
			"customerId": customerId,
		}).
		RunWith(db).Query()
	if err != nil {
		log.Printf("error: %+v", err)
		c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
		return err
	}

	comments := []model.WineComment{}
	for rows.Next() {
		var comment model.WineComment
		err = rows.Scan(&comment.Id, &comment.WineId, &comment.WineYear, &comment.Comment)
		if err != nil {
			log.Printf("error: %+v", err)
			c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
			return err
		}
		comments = append(comments, comment)
	}

	return c.JSON(http.StatusOK, JsonResponse{Ok: true, Data: comments})
}
