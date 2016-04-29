package queries

import (
	"database/sql"

	"github.com/dfreire/sunny/crud"
	"github.com/jmoiron/sqlx"
)

func GetCustomers(db *sql.DB, dbx *sqlx.DB) ([]crud.Record, error) {

	recordDefinition := make(map[string]string)
	recordDefinition["id"] = "string"
	recordDefinition["name"] = "string"
	recordDefinition["email"] = "string"
	recordDefinition["roleId"] = "string"
	recordDefinition["createdAt"] = "string"
	recordDefinition["signupOriginId"] = "string"
	recordDefinition["inMailingList"] = "bool"

	return crud.Get(db, "Customer", recordDefinition)
}
