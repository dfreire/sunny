package middleware

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

const (
	TX = "TX"
)

func WithTransaction(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			tx := db.Begin()
			c.Set(TX, tx)

			err = next(c)
			if err != nil {
				tx.Rollback()
				log.Printf("tx.Rollback")
				return err
			}

			tx.Commit()
			log.Printf("tx.Commit")

			return nil
		}
	}
}
