package crud

import (
	"database/sql"
	"errors"

	"gopkg.in/Masterminds/squirrel.v1"
	"labix.org/v2/mgo/bson"
)

type Record map[string]interface{}

func Create(db *sql.DB, tableName string, record Record) error {
	record["id"] = bson.NewObjectId().Hex()

	return sqlResult(squirrel.
		Insert(tableName).
		SetMap(record).
		RunWith(db).Exec())
}

func Update(db *sql.DB, tableName string, record Record, ids []string) error {
	return sqlResult(squirrel.
		Update(tableName).
		SetMap(record).
		Where(squirrel.Eq{
			"id": ids,
		}).
		RunWith(db).Exec())
}

func Delete(db *sql.DB, tableName string, ids []string) error {
	return sqlResult(squirrel.Delete("book").
		Where(squirrel.Eq{
			"id": ids,
		}).
		RunWith(db).Exec())
}

func sqlResult(result sql.Result, err error) error {
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return errors.New("No rows changed.")
	}

	return nil
}

func EnsureRecord(db *sql.DB, tableName string, recordToFind squirrel.Eq, recordToCreate Record) (id string, err error) {
	rows, err := squirrel.
		Select("id").
		From(tableName).
		Where(recordToFind).
		RunWith(db).Query()
	if err != nil {
		return
	}

	for rows.Next() {
		rows.Scan(&id)
	}

	if id == "" {
		err = Create(db, tableName, recordToCreate)
		if err != nil {
			return
		}
		id = recordToCreate["id"].(string)
	}

	return
}
