package middleware

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

const (
	DBX = "DBX"
)

func Dependencies(dbx *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			log.Println("SetContext")
			c.Set(DBX, dbx)
			return next(c)
		}
	}
}
