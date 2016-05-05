package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type CustomerRole struct {
	ID string `gorm:"primary_key" json:"id"`
}

type Customer struct {
	ID                 string       `gorm:"primary_key" json:"id"`
	CreatedAt          time.Time    `json:"createdAt"`
	UpdatedAt          time.Time    `json:"updatedAt"`
	Name               string       `json:"name"`
	Email              string       `json:"email"`
	Role               CustomerRole `json:"-"`
	RoleId             string       `json:"roleId"`
	OptedInNewsletter  bool         `json:"optedInNewsletter"`
	OptedOutNewsletter bool         `json:"optedOutNewsletter"`
	SentToNewsletter   bool         `json:"sentToNewsletter"`
}

type WineComment struct {
	ID         string    `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Customer   Customer  `json:"-"`
	CustomerId string    `json:"customerId"`
	WineId     string    `json:"wineId"`
	WineYear   int       `json:"wineYear"`
	Comment    string    `json:"comment"`
}

func Initialize(db *gorm.DB) {
	db.SingularTable(true)
	db.Exec(SCHEMA)
}
