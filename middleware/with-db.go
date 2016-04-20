package middleware

import (
	"database/sql"

	"github.com/labstack/echo"
)

const (
	DB = "DB"
)

func WithDB(db *sql.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			c.Set(DB, db)
			err = next(c)
			return err
		}
	}
}
