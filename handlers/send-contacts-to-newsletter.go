package handlers

import (
	"net/http"

	"github.com/dfreire/sunny/mailer"
	"github.com/dfreire/sunny/middleware"
	"github.com/dfreire/sunny/operations"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// http POST http://localhost:3500/send-contacts-to-newsletter?appToken=2fe9a70a-46f2-4d00-88f2-6f66ed903426
func SendContactsToNewsletter(c echo.Context) error {
	tx := c.Get(middleware.TX).(*gorm.DB)
	mx := c.Get(middleware.MAILER).(mailer.Mailer)

	if err := operations.SendContactsToNewsletter(tx, mx); err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{Ok: false})
		return err
	}

	return c.JSON(http.StatusOK, jsonResponse{Ok: true})
}
