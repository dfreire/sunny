package handlers

import (
	"net/http"

	"github.com/dfreire/sunny/commands"
	"github.com/dfreire/sunny/mailer"
	"github.com/dfreire/sunny/middleware"
	"github.com/jinzhu/gorm"
	"github.com/jordan-wright/email"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

// http POST http://localhost:3500/signup-customer-with-wine-comments email="joe.doe@mailinator.com" roleId="wine_lover" language="en" wineComments:='[{"wineId": "wine-1", "wineName": "Soalheiro Alvarinho", "wineYear": 2015, "comment": "great"}, {"wineId": "wine-1", "wineName": "Soalheiro Alvarinho", "wineYear": 2014, "comment": "fantastic"}]'
func SignupCustomerWithWineComments(c echo.Context) error {
	tx := c.Get(middleware.TX).(*gorm.DB)
	m := c.Get(middleware.MAILER).(mailer.Mailer)

	var reqData commands.SignupCustomerWithWineCommentsRequestData
	c.Bind(&reqData)

	err := commands.SignupCustomerWithWineComments(tx, reqData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{Ok: false})
		return err
	}

	err = sendMailAfterSignupCustomerWithWineComments(m, reqData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{Ok: false})
		return err
	}

	return c.JSON(http.StatusOK, jsonResponse{Ok: true})
}

func sendMailAfterSignupCustomerWithWineComments(m mailer.Mailer, reqData commands.SignupCustomerWithWineCommentsRequestData) error {
	e := email.Email{
		To:  []string{reqData.Email},
		Bcc: viper.GetStringSlice("NOTIFICATION_EMAILS"),
	}

	templateId := "on-sign-up-customer-with-wine-comments-email"
	languageId := reqData.LanguageId
	err := mailer.TemplateToEmail(&e, templateId, languageId, nil)
	if err != nil {
		return err
	}

	return m.Send(&e)
}
