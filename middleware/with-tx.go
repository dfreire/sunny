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
			log.Println("WithTX Before")

			tx, err := db.Begin()
			if err != nil {
				log.Println("WithTX After")
				return err
			}

			c.Set(TX, tx)
			err = next(c)
			if err != nil {
				log.Printf("WithTX Rollback")
				tx.Rollback()
				log.Println("WithTX After")
				return err
			}

			err = tx.Commit()
			if err != nil {
				log.Printf("WithTX Rollback")
				tx.Rollback()
				log.Println("WithTX After")
				return err
			}
			log.Printf("WithTX Commit")
			log.Println("WithTX After")

			return nil
		}
	}
}
