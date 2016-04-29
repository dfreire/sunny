package crud

import (
	"database/sql"
	"errors"

	"github.com/fatih/structs"
	"github.com/guregu/null"

	"gopkg.in/Masterminds/squirrel.v1"
	"labix.org/v2/mgo/bson"
)

type Record map[string]interface{}

func Get2(db *sql.DB, definition interface{}) ([]Record, error) {
	recordDefinition := make(map[string]string)

	s := structs.New(definition)
	for _, field := range s.Fields() {
		recordDefinition[field.Name()] = field.Kind().String()
	}

	return Get(db, s.Name(), recordDefinition)
}

func Get(db *sql.DB, tableName string, recordDefinition map[string]string) ([]Record, error) {
	keys := getKeysFromDefinition(recordDefinition)

	rows, err := squirrel.
		Select(keys...).
		From(tableName).
		RunWith(db).Query()
	if err != nil {
		return nil, err
	}

	var records []Record

	for rows.Next() {
		record, values := getRecordAndValuesFromDefinition(keys, recordDefinition)
		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func GetById(db *sql.DB, tableName string, recordDefinition map[string]string, id string) ([]Record, error) {
	return GetByIds(db, tableName, recordDefinition, []string{id})
}

func GetByIds(db *sql.DB, tableName string, recordDefinition map[string]string, ids []string) ([]Record, error) {
	keys := getKeysFromDefinition(recordDefinition)

	rows, err := squirrel.
		Select(keys...).
		From(tableName).
		Where(squirrel.Eq{
			"id": ids,
		}).
		RunWith(db).Query()
	if err != nil {
		return nil, err
	}

	var records []Record

	for rows.Next() {
		record, values := getRecordAndValuesFromDefinition(keys, recordDefinition)
		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func Create(tx *sql.Tx, tableName string, record Record) error {
	record["id"] = bson.NewObjectId().Hex()

	return sqlResult(squirrel.
		Insert(tableName).
		SetMap(record).
		RunWith(tx).Exec())
}

func UpdateById(tx *sql.Tx, tableName string, record Record, id string) error {
	return UpdateByIds(tx, tableName, record, []string{id})
}

func UpdateByIds(tx *sql.Tx, tableName string, record Record, ids []string) error {
	return sqlResult(squirrel.
		Update(tableName).
		SetMap(record).
		Where(squirrel.Eq{
			"id": ids,
		}).
		RunWith(tx).Exec())
}

func DeleteById(tx *sql.Tx, tableName string, id string) error {
	return DeleteByIds(tx, tableName, []string{id})
}

func DeleteByIds(tx *sql.Tx, tableName string, ids []string) error {
	return sqlResult(squirrel.
		Delete(tableName).
		Where(squirrel.Eq{
			"id": ids,
		}).
		RunWith(tx).Exec())
}

func Upsert(tx *sql.Tx, tableName string, recordToFind Record, recordToInsert Record, recordToUpdate Record) (id string, err error) {
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
		err = UpdateById(tx, tableName, recordToUpdate, id)
	}

	return
}

func RecordId(record Record) (id string) {
	if record["id"] != nil {
		id = record["id"].(string)
	}
	return
}

func getKeysFromDefinition(recordDefinition map[string]string) (keys []string) {
	for k, _ := range recordDefinition {
		keys = append(keys, k)
	}
	return
}

func getRecordAndValuesFromDefinition(keys []string, recordDefinition map[string]string) (record Record, values []interface{}) {
	record = make(Record)
	for _, k := range keys {
		switch recordDefinition[k] {
		case "string":
			record[k] = &null.String{}
		case "bool":
			record[k] = &null.Bool{}
		}
		values = append(values, record[k])
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
