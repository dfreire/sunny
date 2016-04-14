package handlers

import (
	"log"
	"net/http"

	"github.com/dfreire/sunny/middleware"
	"github.com/dfreire/sunny/model"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"gopkg.in/Masterminds/squirrel.v1"
)

func GetWineComments(c echo.Context) error {
	log.Println("GetWineComments")

	dbx := c.Get(middleware.DBX).(*sqlx.DB)

	comments := []model.WineComment{}

	sql, args, _ := squirrel.Select("*").From("WineComment").ToSql()
	log.Println(sql, args)
	rows, _ := dbx.Queryx(sql, args...)
	for rows.Next() {
		var c model.WineComment
		rows.StructScan(&c)
		comments = append(comments, c)
	}

	return c.JSON(http.StatusOK, comments)
}
