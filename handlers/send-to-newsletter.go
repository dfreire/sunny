package handlers

import (
	"fmt"
	"net/http"

	"github.com/dfreire/sunny/middleware"
	"github.com/dfreire/sunny/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/tealeg/xlsx"
)

// http POST http://localhost:3500/send-to-newsletter
func SendToNewsletter(c echo.Context) error {
	tx := c.Get(middleware.TX).(*gorm.DB)
	// m := c.Get(middleware.MAILER).(mailer.Mailer)

	customers := []model.Customer{}
	err := tx.Where(map[string]interface{}{
		"opted_out_newsletter": false,
		"sent_to_newsletter":   false,
	}).Find(&customers).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsonResponse{Ok: false})
		return err
	}

	exportToExcel(customers)

	return c.JSON(http.StatusOK, jsonResponse{Ok: true, Result: customers})
}

func exportToExcel(customers []model.Customer) {

	type export struct {
		Name     string
		Email    string
		Role     string
		Language string
	}

	file := xlsx.NewFile()

	sheet, err := file.AddSheet("Registos")
	if err != nil {
		fmt.Printf(err.Error())
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

	sheet.SetColWidth(0, 3, 30)

	err = file.Save("file.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
