package middleware

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

const (
	DB = "DB"
)

func WithDatabase(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			c.Set(DB, db)
			err = next(c)
			return err
		}
	}
}
