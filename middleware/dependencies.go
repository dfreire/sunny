package middleware

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

const (
	DB  = "DB"
	DBX = "DBX"
)

func Dependencies(db *sql.DB, dbx *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			log.Println("Dependencies")
			c.Set(DB, db)
			c.Set(DBX, dbx)
			return next(c)
		}
	}
}
