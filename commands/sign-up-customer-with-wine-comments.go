package commands

import (
	"github.com/dfreire/sunny/mailer"
	"github.com/dfreire/sunny/model"
	"github.com/jinzhu/gorm"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"labix.org/v2/mgo/bson"
)

type SignupCustomerWithWineCommentsRequestData struct {
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

func SignupCustomerWithWineComments(db *gorm.DB, mx mailer.Mailer, reqData SignupCustomerWithWineCommentsRequestData) error {
	if err := upsertCustomerOnSignupCustomerWithWineComments(db, reqData); err != nil {
		return err
	}

	if err := sendMailOnSignupCustomerWithWineComments(mx, reqData); err != nil {
		return err
	}

	return nil
}

func upsertCustomerOnSignupCustomerWithWineComments(db *gorm.DB, reqData SignupCustomerWithWineCommentsRequestData) error {
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
		Name:       reqData.Name,
		RoleId:     reqData.RoleId,
		LanguageId: reqData.LanguageId,
	}

	if err := db.Model(&customer).Updates(toUpdate).Error; err != nil {
		return err
	}

	for _, comment := range reqData.WineComments {
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

func sendMailOnSignupCustomerWithWineComments(mx mailer.Mailer, reqData SignupCustomerWithWineCommentsRequestData) error {
	e := email.Email{
		To:  []string{reqData.Email},
		Bcc: viper.GetStringSlice("NOTIFICATION_EMAILS"),
	}

	templateId := "on-sign-up-customer-with-wine-comments-email"
	err := mailer.PrepareEmail(&e, reqData.LanguageId, templateId, nil)
	if err != nil {
		return err
	}

	return mx.SendEmail(&e)
}
