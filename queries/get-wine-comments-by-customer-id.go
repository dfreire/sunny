package queries

import (
	"github.com/dfreire/sunny/model"
	"github.com/jinzhu/gorm"
)

func GetWineCommentsByCustomerId(db *gorm.DB, customerId string) ([]model.WineComment, error) {
	var comments []model.WineComment
	err := db.Where("customer_id = ?", customerId).Find(&comments).Error
	return comments, err
}
