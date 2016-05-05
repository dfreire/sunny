package middleware

import (
	"log"

	"github.com/labstack/echo"
)

func Logging() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			var body map[string]interface{}
			c.Bind(&body)
			log.Println(c.Request().URI(), body)

			err = next(c)
			if err != nil {
				log.Printf("Error: %+v", err)
				return err
			}
			return nil
		}
	}
}
