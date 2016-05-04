package commands

import (
	"github.com/dfreire/sunny/model"
	"github.com/jinzhu/gorm"
	"labix.org/v2/mgo/bson"
)

type SignupCustomerWithWineCommentsRequestData struct {
	Name         string        `json:"name,omitempty"`
	Email        string        `json:"email"`
	RoleId       string        `json:"roleId"`
	WineComments []WineComment `json:"wineComments"`
}

type WineComment struct {
	WineName string `json:"wineName"`
	WineId   string `json:"wineId"`
	WineYear int    `json:"wineYear"`
	Comment  string `json:"comment"`
}

func SignupCustomerWithWineComments(db *gorm.DB, reqData SignupCustomerWithWineCommentsRequestData) error {
	toFind := model.Customer{
		Email: reqData.Email,
	}

	toCreate := model.Customer{
		ID:     bson.NewObjectId().Hex(),
		Email:  reqData.Email,
		RoleId: reqData.RoleId,
	}

	customer := model.Customer{}
	err := db.Where(toFind).Attrs(toCreate).FirstOrCreate(&customer).Error
	if err != nil {
		return err
	}

	toUpdate := model.Customer{
		Name:   reqData.Name,
		RoleId: reqData.RoleId,
	}

	err = db.Model(&customer).Updates(toUpdate).Error
	if err != nil {
		return err
	}

	for _, comment := range reqData.WineComments {
		err = upsertWineComment(db, customer.ID, comment)
		if err != nil {
			return err
		}
	}

	return nil
}

func upsertWineComment(db *gorm.DB, customerId string, comment WineComment) error {
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
