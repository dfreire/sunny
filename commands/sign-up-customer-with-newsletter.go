package commands

import (
	"github.com/dfreire/sunny/mailer"
	"github.com/dfreire/sunny/model"
	"github.com/jinzhu/gorm"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"labix.org/v2/mgo/bson"
)

type SignupCustomerWithNewsletterRequestData struct {
	Name       string `json:"name,omitempty"`
	Email      string `json:"email"`
	RoleId     string `json:"roleId"`
	LanguageId string `json:"language"`
}

func SignupCustomerWithNewsletter(db *gorm.DB, mx mailer.Mailer, reqData SignupCustomerWithNewsletterRequestData) error {
	if err := upsertCustomerWithNewsletter(db, reqData); err != nil {
		return err
	}

	if err := sendMailOnSignupCustomerWithNewsletter(mx, reqData); err != nil {
		return err
	}

	return nil
}

func upsertCustomerWithNewsletter(db *gorm.DB, reqData SignupCustomerWithNewsletterRequestData) error {
	toFind := model.Customer{
		Email: reqData.Email,
	}

	toCreate := model.Customer{
		ID:         bson.NewObjectId().Hex(),
		Email:      reqData.Email,
		RoleId:     reqData.RoleId,
		LanguageId: reqData.LanguageId,
	}

	customer := model.Customer{}
	if err := db.Where(toFind).Attrs(toCreate).FirstOrCreate(&customer).Error; err != nil {
		return err
	}

	toUpdate := model.Customer{
		Name:              reqData.Name,
		RoleId:            reqData.RoleId,
		LanguageId:        reqData.LanguageId,
		OptedInNewsletter: true,
	}

	return db.Model(&customer).Updates(toUpdate).Error
}

func sendMailOnSignupCustomerWithNewsletter(mx mailer.Mailer, reqData SignupCustomerWithNewsletterRequestData) error {
	e := email.Email{
		From: viper.GetString("OWNER_EMAIL"),
		To:   []string{reqData.Email},
		Bcc:  viper.GetStringSlice("NOTIFICATION_EMAILS"),
	}

	templateId := "on-sign-up-customer-with-newsletter-email"
	if err := mailer.PrepareEmail(&e, reqData.LanguageId, templateId, nil); err != nil {
		return err
	}

	return mx.SendEmail(&e)
}
