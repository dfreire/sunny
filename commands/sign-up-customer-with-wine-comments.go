package commands

import (
	"github.com/dfreire/sunny/mailer"
	"github.com/dfreire/sunny/model"
	"github.com/jinzhu/gorm"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"labix.org/v2/mgo/bson"
)

type SignupCustomerWithWineCommentsRequest struct {
	Name         string        `json:"name,omitempty"`
	Email        string        `json:"email"`
	RoleId       string        `json:"roleId"`
	LanguageId   string        `json:"language"`
	WineComments []WineComment `json:"wineComments"`
}

type WineComment struct {
	WineName string `json:"wineName"`
	WineId   string `json:"wineId"`
	WineYear int    `json:"wineYear"`
	Comment  string `json:"comment"`
}

func SignupCustomerWithWineComments(db *gorm.DB, mx mailer.Mailer, req SignupCustomerWithWineCommentsRequest) error {
	if err := upsertCustomerOnSignupCustomerWithWineComments(db, req); err != nil {
		return err
	}

	if err := sendMailOnSignupCustomerWithWineComments(mx, req); err != nil {
		return err
	}

	return nil
}

func upsertCustomerOnSignupCustomerWithWineComments(db *gorm.DB, req SignupCustomerWithWineCommentsRequest) error {
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
		Name:       req.Name,
		RoleId:     req.RoleId,
		LanguageId: req.LanguageId,
	}

	if err := db.Model(&customer).Updates(toUpdate).Error; err != nil {
		return err
	}

	for _, comment := range req.WineComments {
		if err := upsertWineCommentOnSignupCustomerWithWineComments(db, customer.ID, comment); err != nil {
			return err
		}
	}

	return nil
}

func upsertWineCommentOnSignupCustomerWithWineComments(db *gorm.DB, customerId string, comment WineComment) error {
	toFind := model.WineComment{
		CustomerId: customerId,
		WineId:     comment.WineId,
		WineYear:   comment.WineYear,
	}

	toCreate := model.WineComment{
		ID: bson.NewObjectId().Hex(),
	}

	wineComment := model.WineComment{}
	err := db.Where(toFind).Attrs(toCreate).FirstOrCreate(&wineComment).Error
	if err != nil {
		return err
	}

	toUpdate := model.WineComment{
		Comment: comment.Comment,
	}

	return db.Model(&wineComment).Updates(toUpdate).Error
}

func sendMailOnSignupCustomerWithWineComments(mx mailer.Mailer, req SignupCustomerWithWineCommentsRequest) error {
	e := email.Email{
		From: viper.GetString("OWNER_EMAIL"),
		To:   []string{req.Email},
		Bcc:  viper.GetStringSlice("NOTIFICATION_EMAILS"),
	}

	templateId := "on-sign-up-customer-with-wine-comments-email"
	err := mailer.PrepareEmail(&e, req.LanguageId, templateId, req)
	if err != nil {
		return err
	}

	return mx.SendEmail(&e)
}
