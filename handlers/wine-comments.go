package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"labix.org/v2/mgo/bson"

	"github.com/dfreire/sunny/middleware"
	"github.com/dfreire/sunny/model"
	"github.com/labstack/echo"
	"gopkg.in/Masterminds/squirrel.v1"
)

func GetWineComments(c echo.Context) error {
	log.Println("GetWineComments")

	db := c.Get(middleware.DB).(*sql.DB)

	comments := []model.WineComment{}

	rows, err := squirrel.
		Select("id", "customerId", "wineId", "wineYear", "comment").
		From("WineComment").
		RunWith(db).Query()
	if err != nil {
		return err
	}

	for rows.Next() {
		var c model.WineComment
		rows.Scan(&c.Id, &c.CustomerId, &c.WineId, &c.WineYear, &c.Comment)
		comments = append(comments, c)
	}

	return c.JSON(http.StatusOK, comments)
}

func UpsertWineComment(c echo.Context) error {
	log.Println("UpsertWineComment")

	db := c.Get(middleware.DB).(*sql.DB)

	comment := make(map[string]interface{})
	c.Bind(&comment)

	var rowsAffected int64
	var err error
	if comment["id"] == nil || comment["id"] == "" {
		rowsAffected, err = insertWineComment(db, comment)
	} else {
		rowsAffected, err = updateWineComment(db, comment)
	}

	if rowsAffected == 0 || err != nil {
		log.Printf("error: %+v", err)
		return err
	}

	return c.JSON(http.StatusOK, comment)
}

func insertWineComment(db *sql.DB, comment map[string]interface{}) (int64, error) {
	log.Println("insertWineComment")

	comment["id"] = bson.NewObjectId().Hex()

	createdAt := time.Now().Format(time.RFC3339)
	comment["createdAt"] = createdAt
	comment["updatedAt"] = createdAt

	result, err := squirrel.
		Insert("WineComment").
		SetMap(comment).
		RunWith(db).Exec()

	return result.RowsAffected()
}

func updateWineComment(db *sql.DB, comment map[string]interface{}) (int64, error) {
	log.Println("updateWineComment")

	updatedAt := time.Now().Format(time.RFC3339)
	comment["updatedAt"] = updatedAt

	result, err := squirrel.
		Update("WineComment").
		SetMap(comment).
		Where(squirrel.Eq{
			"id": comment["id"],
		}).
		RunWith(db).Exec()

	return result.RowsAffected()
}
