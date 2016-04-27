package queries

import (
	"database/sql"
	"log"

	"github.com/dfreire/sunny/model"

	"gopkg.in/Masterminds/squirrel.v1"
)

func GetCustomers(db *sql.DB) ([]model.Customer, error) {
	rows, err := squirrel.
		Select("id", "email", "roleId", "createdAt", "signupOriginId", "inMailingList").
		From("Customer").
		RunWith(db).Query()
	if err != nil {
		log.Printf("error: %+v", err)
		return nil, err
	}

	customers := []model.Customer{}
	for rows.Next() {
		var customer model.Customer
		err = rows.Scan(
			&customer.Id, &customer.Email, &customer.RoleId,
			&customer.CreatedAt, &customer.SignupOriginId, &customer.InMailingList,
		)
		if err != nil {
			log.Printf("error: %+v", err)
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}
