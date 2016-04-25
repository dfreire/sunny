package commands

import (
	"database/sql"
	"time"

	"github.com/dfreire/sunny/crud"
)

type SignupCustomerWithNewsletterRequestData struct {
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
		},
		crud.Record{
			"roleId": reqData.RoleId,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
