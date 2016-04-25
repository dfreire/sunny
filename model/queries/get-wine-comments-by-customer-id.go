package queries

import (
	"database/sql"
	"log"

	"github.com/dfreire/sunny/model"

	"gopkg.in/Masterminds/squirrel.v1"
)

func GetWineCommentsByCustomerId(db *sql.DB, customerId string) ([]model.WineComment, error) {
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

	comments := []model.WineComment{}
	for rows.Next() {
		var comment model.WineComment
		err = rows.Scan(&comment.Id, &comment.WineId, &comment.WineYear, &comment.Comment)
		if err != nil {
			log.Printf("error: %+v", err)
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
