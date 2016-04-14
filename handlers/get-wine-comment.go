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

func GetWineComments(c echo.Context) error {
	log.Println("GetWineComments")

	db := c.Get(middleware.DB).(*sql.DB)

	comments := []model.WineComment{}

	rows, err := squirrel.
		Select("id", "wineId", "wineYear", "comment").
		From("WineComment").
		RunWith(db).Query()
	if err != nil {
		return err
	}

	for rows.Next() {
		var c model.WineComment
		rows.Scan(&c.Id, &c.WineId, &c.WineYear, &c.Comment)
		comments = append(comments, c)
	}

	return c.JSON(http.StatusOK, comments)
}
