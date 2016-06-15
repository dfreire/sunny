package operations

import (
	"github.com/dfreire/sunny/mailer"
	"github.com/dfreire/sunny/model"
	"github.com/jinzhu/gorm"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"github.com/tealeg/xlsx"
)

func SendToNewsletter(db *gorm.DB, mx mailer.Mailer) error {
	customers := []model.Customer{}
	err := db.Where(map[string]interface{}{
		"opted_out_newsletter": false,
		"sent_to_newsletter":   false,
	}).Find(&customers).Error
	if err != nil {
		return err
	}

	if len(customers) > 0 {
		ids := []string{}
		for _, customer := range customers {
			ids = append(ids, customer.ID)
		}

		err = db.Model(&model.Customer{}).
			Where("id IN (?)", ids).
			Update("sent_to_newsletter", true).
			Error
		if err != nil {
			return err
		}

		fileName := "emails.xlsx"

		if err = exportEmailsToFile(customers, fileName); err != nil {
			return err
		}

		if err = sendMailToNewsletter(mx, fileName); err != nil {
			return err
		}
	}

	return nil
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
	e := email.Email{
		From: viper.GetString("TEAM_EMAIL"),
		To:   []string{viper.GetString("OWNER_EMAIL")},
		Bcc:  viper.GetStringSlice("NOTIFICATION_EMAILS"),
	}

	e.AttachFile(fileName)

	languageId := "pt"
	templateId := "send-to-newsletter-email"
	err := mailer.PrepareEmail(&e, languageId, templateId, nil)
	if err != nil {
		return err
	}

	return m.SendEmail(&e)
}
