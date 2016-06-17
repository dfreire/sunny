package operations

import (
	"github.com/dfreire/sunny/model"
	"github.com/jinzhu/gorm"
)

func GetCustomers(db *gorm.DB) ([]model.Customer, error) {
	customers := []model.Customer{}
	err := db.Preload("Role").Preload("Language").Find(&customers).Error
	return customers, err
}
