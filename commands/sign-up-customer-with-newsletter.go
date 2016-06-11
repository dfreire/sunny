package commands

import (
	"github.com/dfreire/sunny/mailer"
	"github.com/dfreire/sunny/model"
	"github.com/jinzhu/gorm"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"labix.org/v2/mgo/bson"
)

type SignupCustomerWithNewsletterRequest struct {
	Name       string `json:"name,omitempty"`
	Email      string `json:"email"`
	RoleId     string `json:"roleId"`
	LanguageId string `json:"language"`
}

func SignupCustomerWithNewsletter(db *gorm.DB, mx mailer.Mailer, req SignupCustomerWithNewsletterRequest) error {
	if err := upsertCustomerOnSignupCustomerWithNewsletter(db, req); err != nil {
		return err
	}

	if err := sendMailOnSignupCustomerWithNewsletter(mx, req); err != nil {
		return err
	}

	return nil
}

func upsertCustomerOnSignupCustomerWithNewsletter(db *gorm.DB, req SignupCustomerWithNewsletterRequest) error {
	toFind := model.Customer{
		Email: req.Email,
	}

	toCreate := model.Customer{
		ID:         bson.NewObjectId().Hex(),
		Email:      req.Email,
		RoleId:     req.RoleId,
		LanguageId: req.LanguageId,
	}

	customer := model.Customer{}
	if err := db.Where(toFind).Attrs(toCreate).FirstOrCreate(&customer).Error; err != nil {
		return err
	}

	toUpdate := model.Customer{
		Name:              req.Name,
		RoleId:            req.RoleId,
		LanguageId:        req.LanguageId,
		OptedInNewsletter: true,
	}

	return db.Model(&customer).Updates(toUpdate).Error
}

func sendMailOnSignupCustomerWithNewsletter(mx mailer.Mailer, req SignupCustomerWithNewsletterRequest) error {
	e := email.Email{
		From: viper.GetString("OWNER_EMAIL"),
		To:   []string{req.Email},
		Bcc:  viper.GetStringSlice("NOTIFICATION_EMAILS"),
	}

	templateId := "on-sign-up-customer-with-newsletter-email"
	if err := mailer.PrepareEmail(&e, req.LanguageId, templateId, nil); err != nil {
		return err
	}

	return mx.SendEmail(&e)
}
