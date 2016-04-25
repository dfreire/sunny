package model

import (
	"database/sql"
	"log"

	"gopkg.in/Masterminds/squirrel.v1"
)

func GetWineCommentsByCustomerId(db *sql.DB, customerId string) ([]WineComment, error) {
	rows, err := squirrel.
		Select("id", "wineId", "wineYear", "comment").
		From("WineComment").
		Where(squirrel.Eq{
			"customerId": customerId,
		}).
		RunWith(db).Query()
	if err != nil {
		log.Printf("error: %+v", err)
		return nil, err
	}

	comments := []WineComment{}
	for rows.Next() {
		var comment WineComment
		err = rows.Scan(&comment.Id, &comment.WineId, &comment.WineYear, &comment.Comment)
		if err != nil {
			log.Printf("error: %+v", err)
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
