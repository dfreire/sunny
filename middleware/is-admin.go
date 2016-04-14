package middleware

import (
	"log"

	"github.com/labstack/echo"
)

func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			log.Println("IsAdmin")
			return next(c)
		}
	}
}
