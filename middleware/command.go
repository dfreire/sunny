package middleware

import (
	"log"

	"github.com/labstack/echo"
)

func Command() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			log.Println("Command")
			return next(c)
		}
	}
}
