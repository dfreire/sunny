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

	type export struct {
		Name     string
		Email    string
		Role     string
		Language string
	}

	return c.JSON(http.StatusOK, jsonResponse{Ok: true, Result: customers})
}

func exportToExcel(customers []model.Customer) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	for _, customer := range customers {
		row := sheet.AddRow()

		cell := row.AddCell()
		cell.Value = customer.Email
	}

	err = file.Save("file.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
