package operations

import (
	"github.com/dfreire/sunny/model"
	"github.com/jinzhu/gorm"
)

func GetWineCommentsByCustomerId(db *gorm.DB, customerId string) ([]model.WineComment, error) {
	comments := []model.WineComment{}
	err := db.Where("customer_id = ?", customerId).Preload("Customer").Preload("Customer.Role").Find(&comments).Error
	return comments, err
}
