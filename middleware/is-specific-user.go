package middleware

import (
	"log"

	"github.com/labstack/echo"
)

func IsSpecificUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			log.Println("IsSpecificUser")
			return next(c)
		}
	}
}
