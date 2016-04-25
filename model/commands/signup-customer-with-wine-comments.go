package commands

import (
	"database/sql"
	"time"

	"github.com/dfreire/sunny/crud"
)

type SignupCustomerWithWineCommentsRequestData struct {
	Email        string `json:"email"`
	RoleId       string `json:"roleId"`
	WineComments []struct {
		WineId   string `json:"wineId"`
		WineYear int    `json:"wineYear"`
		Comment  string `json:"comment"`
	} `json:"wineComments"`
}

func SignupCustomerWithWineComments(tx *sql.Tx, reqData SignupCustomerWithWineCommentsRequestData) error {
	now := time.Now().Format(time.RFC3339)

	customerId, err := crud.Upsert(
		tx,
		"Customer",
		crud.Record{
			"email": reqData.Email,
		},
		crud.Record{
			"createdAt":      now,
			"signupOriginId": "wine_comment",
		},
		crud.Record{
			"roleId": reqData.RoleId,
		},
	)
	if err != nil {
		return err
	}

	for _, comment := range reqData.WineComments {
		_, err = crud.Upsert(
			tx,
			"WineComment",
			crud.Record{
				"customerId": customerId,
				"wineId":     comment.WineId,
				"wineYear":   comment.WineYear,
			},
			crud.Record{
				"createdAt": now,
			},
			crud.Record{
				"updatedAt": now,
				"comment":   comment.Comment,
			},
		)
		if err != nil {
			return err
		}
	}

	return nil
}
