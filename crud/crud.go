package crud

import (
	"database/sql"
	"errors"

	"gopkg.in/Masterminds/squirrel.v1"
	"labix.org/v2/mgo/bson"
)

type Record map[string]interface{}

func Create(tx *sql.Tx, tableName string, record Record) error {
	record["id"] = bson.NewObjectId().Hex()

	return sqlResult(squirrel.
		Insert(tableName).
		SetMap(record).
		RunWith(tx).Exec())
}

func Update(tx *sql.Tx, tableName string, record Record, ids []string) error {
	return sqlResult(squirrel.
		Update(tableName).
		SetMap(record).
		Where(squirrel.Eq{
			"id": ids,
		}).
		RunWith(tx).Exec())
}

func Delete(tx *sql.Tx, tableName string, ids []string) error {
	return sqlResult(squirrel.Delete("book").
		Where(squirrel.Eq{
			"id": ids,
		}).
		RunWith(tx).Exec())
}

func UpsertRecord(tx *sql.Tx, tableName string, recordToFind Record, recordToInsert Record, recordToUpdate Record) (id string, err error) {
	rows, err := squirrel.
		Select("id").
		From(tableName).
		Where(recordToMap(recordToFind)).
		RunWith(tx).Query()
	if err != nil {
		return
	}

	for rows.Next() {
		rows.Scan(&id)
	}

	if id == "" {
		recordToInsert = mergeRecords(recordToFind, recordToUpdate, recordToInsert)
		err = Create(tx, tableName, recordToInsert)
		id = RecordId(recordToInsert)
	} else {
		recordToUpdate = mergeRecords(recordToFind, recordToUpdate)
		err = Update(tx, tableName, recordToUpdate, []string{id})
	}

	return
}

func RecordId(record Record) (id string) {
	if record["id"] != nil {
		id = record["id"].(string)
	}
	return
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

func recordToMap(record Record) map[string]interface{} {
	return record
}

func mergeRecords(records ...Record) Record {
	dest := Record{}
	for _, record := range records {
		for k, v := range record {
			dest[k] = v
		}
	}
	return dest
}
