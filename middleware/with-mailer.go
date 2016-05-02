package middleware

import (
	"github.com/dfreire/sunny/mailer"
	"github.com/labstack/echo"
)

const (
	MAILER = "MAILER"
)

func WithMailer(mailer *mailer.Mailer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			c.Set(MAILER, mailer)
			err = next(c)
			return err
		}
	}
}
