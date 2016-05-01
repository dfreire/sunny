package middleware

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

func HasAppToken(authorizedAppToken string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			appToken := c.QueryParam("appToken")

			if appToken == authorizedAppToken {
				err = next(c)
			} else {
				err = errors.New("Invalid application token")
				c.String(http.StatusUnauthorized, err.Error())
			}

			return err
		}
	}
}
