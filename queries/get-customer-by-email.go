package queries

import (
	"github.com/dfreire/sunny/model"
	"github.com/jinzhu/gorm"
)

func GetCustomerByEmail(db *gorm.DB, email string) (*model.Customer, error) {
	customer := model.Customer{}
	err := db.Where("email = ?", email).First(&customer).Error
	return &customer, err
}
