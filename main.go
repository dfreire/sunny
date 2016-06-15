package main

import (
	"log"

	"github.com/dfreire/sunny/handlers"
	"github.com/dfreire/sunny/mailer"
	"github.com/dfreire/sunny/middleware"
	"github.com/dfreire/sunny/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	echomiddleware "github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("SUNNY")
	viper.AutomaticEnv()
}

func main() {
	debug := viper.GetBool("DEBUG") == true
	appToken := viper.GetString("APP_TOKEN")
	database := viper.GetString("DATABASE")
	port := viper.GetString("PORT")

	db, err := gorm.Open("sqlite3", database)
	if err != nil {
		log.Fatalf("open connection to the database err: %+v", err)
	}

	model.Initialize(db)

	e := echo.New()

	e.Use(echomiddleware.Gzip())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.Logger())

	hasAppToken := middleware.HasAppToken(appToken)
	withMailer := createMiddlewareWithMailer()
	withDatabase := middleware.WithDatabase(db)
	withTransaction := middleware.WithTransaction(db)
	withLogging := middleware.Logging()

	e.Post("/signup-customer-with-wine-comments", handlers.SignupCustomerWithWineComments, withLogging, withTransaction, withMailer)
	e.Post("/signup-customer-with-newsletter", handlers.SignupCustomerWithNewsletter, withLogging, withTransaction, withMailer)

	e.Get("/get-customers", handlers.GetCustomers, withLogging, hasAppToken, withDatabase)
	e.Get("/get-wine-comments-by-customer-id", handlers.GetWineCommentsByCustomerId, withLogging, hasAppToken, withDatabase)

	e.Post("/send-contacts-to-newsletter", handlers.SendContactsToNewsletter, withLogging, hasAppToken, withTransaction, withMailer)

	if debug {
		db.LogMode(true)
		e.SetDebug(true)
	}

	log.Printf("Running on port %s", port)
	e.Run(standard.New(port))
}

func createMiddlewareWithMailer() echo.MiddlewareFunc {
	if viper.GetString("MAILER") == "log" {
		return middleware.WithMailer(mailer.NewLogMailer())
	} else {
		smtpHost := viper.GetString("smtp.host")
		smtpPort := viper.GetInt("smtp.port")
		smtpLogin := viper.GetString("smtp.login")
		smtpPassword := viper.GetString("smtp.password")
		return middleware.WithMailer(mailer.NewMailer(smtpHost, smtpPort, smtpLogin, smtpPassword))
	}
}
