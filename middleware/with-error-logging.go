package middleware

import (
	"log"

	"github.com/labstack/echo"
)

func ErrorLogging() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			log.Println(c.Request().URI())
			err = next(c)
			if err != nil {
				log.Printf("Error: %+v", err)
				return err
			}
			return nil
		}
	}
}
