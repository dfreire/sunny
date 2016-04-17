package crud

import (
	"database/sql"

	"gopkg.in/Masterminds/squirrel.v1"
	"labix.org/v2/mgo/bson"
)

type Record map[string]interface{}

func Create(db *sql.DB, tableName string, record Record) error {
	record["id"] = bson.NewObjectId().Hex()

	_, err := squirrel.
		Insert(tableName).
		SetMap(record).
		RunWith(db).Exec()

	return err
}

func Update(db *sql.DB, tableName string, record Record, ids []string) error {
	_, err := squirrel.
		Update(tableName).
		SetMap(record).
		Where(squirrel.Eq{
			"id": ids,
		}).
		RunWith(db).Exec()

	return err
}

func Delete(db *sql.DB, tableName string, ids []string) error {
	_, err := squirrel.Delete("book").
		Where(squirrel.Eq{
			"id": ids,
		}).
		RunWith(db).Exec()

	return err
}
