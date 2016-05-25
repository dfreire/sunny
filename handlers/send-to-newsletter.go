package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/dfreire/sunny/mailer"
	"github.com/dfreire/sunny/middleware"
	"github.com/dfreire/sunny/model"
	"github.com/jinzhu/gorm"
	"github.com/jordan-wright/email"
	"github.com/labstack/echo"
	"github.com/tealeg/xlsx"
)

// http POST http://localhost:3500/send-to-newsletter
func SendToNewsletter(c echo.Context) error {
	tx := c.Get(middleware.TX).(*gorm.DB)
	m := c.Get(middleware.MAILER).(mailer.Mailer)

	customers := []model.Customer{}
	err := tx.Where(map[string]interface{}{
		"opted_out_newsletter": false,
		"sent_to_newsletter":   false,
	}).Find(&customers).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{Ok: false})
		return err
	}

	ids := []string{}
	for _, customer := range customers {
		ids = append(ids, customer.ID)
	}

	err = tx.Model(&model.Customer{}).
		Where("id IN (?)", ids).
		Update("sent_to_newsletter", true).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{Ok: false})
		return err
	}

	fileName := "emails.xlsx"

	err = exportEmailsToFile(customers, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{Ok: false})
		return err
	}

	err = sendMailToNewsletter(m, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{Ok: false})
		return err
	}

	return c.JSON(http.StatusOK, jsonResponse{Ok: true, Result: customers})
}

func exportEmailsToFile(customers []model.Customer, fileName string) error {
	file := xlsx.NewFile()

	sheet, err := file.AddSheet("Registos")
	if err != nil {
		return err
	}

	row := sheet.AddRow()
	row.AddCell().SetString("Nome")
	row.AddCell().SetString("Email")
	row.AddCell().SetString("Perfil")
	row.AddCell().SetString("Idioma")

	for _, customer := range customers {
		row := sheet.AddRow()
		row.AddCell().SetString(customer.Name)
		row.AddCell().SetString(customer.Email)
		row.AddCell().SetString(customer.RoleId)
		row.AddCell().SetString(customer.LanguageId)
	}

	sheet.SetColWidth(0, 5, 25)

	return file.Save(fileName)
}

func sendMailToNewsletter(m mailer.Mailer, fileName string) error {
	e := email.Email{}
	e.AttachFile(fileName)
	templatePath := filepath.Join("templates", "mail", "pt", "send-to-newsletter-email.yaml")

	err := mailer.TemplateToEmail(&e, templatePath, nil)
	if err != nil {
		return err
	}

	return m.Send(&e)
}
