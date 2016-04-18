package middleware

import (
	"database/sql"
	"log"

	"github.com/labstack/echo"
)

const (
	TX = "TX"
)

func WithTX(db *sql.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			log.Println("WithTX")

			tx, err := db.Begin()
			if err != nil {
				log.Printf("error: %+v", err)
				return err
			}

			c.Set(TX, tx)
			err = next(c)
			if err != nil {
				log.Printf("Rollback")
				tx.Rollback()
				log.Printf("error: %+v", err)
				return err
			}

			err = tx.Commit()
			if err != nil {
				log.Printf("Rollback")
				tx.Rollback()
				log.Printf("error: %+v", err)
				return err
			}
			log.Printf("Commit")

			return nil
		}
	}
}
