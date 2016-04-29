package commands

import (
	"database/sql"
	"time"

	"github.com/dfreire/sunny/crud"
)

type SignupCustomerWithNewsletterRequestData struct {
	Name   string `json:"name,omitempty"`
	Email  string `json:"email"`
	RoleId string `json:"roleId"`
}

func SignupCustomerWithNewsletter(tx *sql.Tx, reqData SignupCustomerWithNewsletterRequestData) error {
	now := time.Now().Format(time.RFC3339)

	_, err := crud.Upsert(
		tx,
		"Customer",
		crud.Record{
			"email": reqData.Email,
		},
		crud.Record{
			"createdAt":      now,
			"signupOriginId": "newsletter",
			"inMailingList":  false,
		},
		crud.Record{
			"name":   reqData.Name,
			"roleId": reqData.RoleId,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
