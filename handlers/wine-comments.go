package handlers

import (
	"database/sql"
	"errors"
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
		log.Printf("error: %+v", err)
		return c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
	}

	for rows.Next() {
		var c model.WineComment
		rows.Scan(&c.Id, &c.CustomerId, &c.WineId, &c.WineYear, &c.Comment)
		comments = append(comments, c)
	}

	return c.JSON(http.StatusOK, JsonResponse{Ok: true, Data: comments})
}

func UpsertWineComment(c echo.Context) error {
	log.Println("UpsertWineComment")

	db := c.Get(middleware.DB).(*sql.DB)

	comment := make(map[string]interface{})
	c.Bind(&comment)

	if comment["id"] == nil || comment["id"] == "" {
		result, err := insertWineComment(db, comment)
		return insertOrUpdateResponse(c, result, err)
	} else {
		result, err := updateWineComment(db, comment)
		return insertOrUpdateResponse(c, result, err)
	}
}

func insertWineComment(db *sql.DB, comment map[string]interface{}) (sql.Result, error) {
	log.Println("insertWineComment")

	comment["id"] = bson.NewObjectId().Hex()

	createdAt := time.Now().Format(time.RFC3339)
	comment["createdAt"] = createdAt
	comment["updatedAt"] = createdAt

	return squirrel.
		Insert("WineComment").
		SetMap(comment).
		RunWith(db).Exec()
}

func updateWineComment(db *sql.DB, comment map[string]interface{}) (sql.Result, error) {
	log.Println("updateWineComment")

	updatedAt := time.Now().Format(time.RFC3339)
	comment["updatedAt"] = updatedAt

	return squirrel.
		Update("WineComment").
		SetMap(comment).
		Where(squirrel.Eq{
			"id": comment["id"],
		}).
		RunWith(db).Exec()
}

func insertOrUpdateResponse(c echo.Context, result sql.Result, err error) error {
	if err != nil {
		log.Printf("error: %+v", err)
		return c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
	} else {
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Printf("error: %+v", err)
			return c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
		} else if rowsAffected == 0 {
			err = errors.New("No rows affected")
			log.Printf("error: %+v", err)
			return c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
		}
	}

	return c.JSON(http.StatusOK, JsonResponse{Ok: true})
}

type JsonResponse struct {
	Ok    bool        `json:"ok"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}
