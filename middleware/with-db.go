package middleware

import (
	"database/sql"
	"log"

	"github.com/labstack/echo"
)

const (
	DB = "DB"
)

func WithDB(db *sql.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			log.Println("WithDB")
			c.Set(DB, db)
			return next(c)
		}
	}
}
