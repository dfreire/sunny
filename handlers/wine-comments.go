package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/dfreire/sunny/crud"
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

	var comment crud.Record
	c.Bind(&comment)

	var err error

	if comment["id"] == nil || comment["id"] == "" {
		createdAt := time.Now().Format(time.RFC3339)
		comment["createdAt"] = createdAt
		comment["updatedAt"] = createdAt
		err = crud.Create(db, "WineComment", comment)

	} else {
		id := comment["id"].(string)
		updatedAt := time.Now().Format(time.RFC3339)
		comment["updatedAt"] = updatedAt
		delete(comment, "customerId")
		delete(comment, "wineId")
		delete(comment, "wineYear")
		err = crud.Update(db, "WineComment", comment, []string{id})
	}

	if err != nil {
		log.Printf("error: %+v", err)
		return c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
	}

	return c.JSON(http.StatusOK, JsonResponse{Ok: true})
}
