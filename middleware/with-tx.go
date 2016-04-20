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
			tx, err := db.Begin()
			if err != nil {
				return err
			}

			c.Set(TX, tx)
			err = next(c)
			if err != nil {
				log.Printf("tx.Rollback")
				tx.Rollback()
				return err
			}

			err = tx.Commit()
			if err != nil {
				log.Printf("tx.Rollback")
				tx.Rollback()
				return err
			}
			log.Printf("tx.Commit")

			return nil
		}
	}
}
