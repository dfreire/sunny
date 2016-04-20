package middleware

import (
	"log"

	"github.com/labstack/echo"
)

func LogErr() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			log.Println("LogErr Before")
			err = next(c)
			if err != nil {
				log.Printf("error: %+v", err)
				log.Println("LogErr After")
				return err
			}
			log.Println("LogErr After")
			return nil
		}
	}
}
