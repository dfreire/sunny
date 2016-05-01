package main

import (
	"log"
	"strings"

	"github.com/dfreire/sunny/handlers"
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
	viper.AutomaticEnv()
	env := viper.Get("ENV").(string)

	viper.AddConfigPath(".")
	viper.SetConfigName(strings.Join([]string{"config", env}, "."))

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config err: %+v", err)
	}
}

func main() {
	env := viper.Get("ENV").(string)
	appToken := viper.Get("appToken").(string)
	database := viper.Get("database").(string)
	port := viper.Get("port").(string)

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
	withDatabase := middleware.WithDatabase(db)
	withTransaction := middleware.WithTransaction(db)
	withErrorLogging := middleware.ErrorLogging()

	e.Post("/signup-customer-with-wine-comments", handlers.SignupCustomerWithWineComments, withErrorLogging, withTransaction)
	e.Post("/signup-customer-with-newsletter", handlers.SignupCustomerWithNewsletter, withErrorLogging, withTransaction)

	e.Get("/get-customers", handlers.GetCustomers, withErrorLogging, hasAppToken, withDatabase)
	e.Get("/get-wine-comments-by-customer-id", handlers.GetWineCommentsByCustomerId, withErrorLogging, hasAppToken, withDatabase)

	if env == "development" {
		db.LogMode(true)
		e.SetDebug(true)
	}

	log.Printf("Running on port %s", port)
	e.Run(standard.New(port))
}
