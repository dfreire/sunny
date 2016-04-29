package middleware

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

const (
	DBX = "DBX"
)

func WithDBX(dbx *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			c.Set(DBX, dbx)
			err = next(c)
			return err
		}
	}
}
