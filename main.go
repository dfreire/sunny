package main

import (
	"database/sql"
	"log"

	"github.com/dfreire/sunny/middleware"
	"github.com/dfreire/sunny/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	echomiddleware "github.com/labstack/echo/middleware"
	// log "github.com/mgutz/logxi/v1"
	"github.com/dfreire/sunny/handlers"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Couldn't read the config file")
	}
}

func main() {
	env := viper.Get("SUNNY_ENV").(string)

	db, err := sql.Open("sqlite3", viper.Get("SUNNY_SQLITE_DB").(string))
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(model.SCHEMA); err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	if env == "development" {
		e.SetDebug(true)
	}
	e.Use(echomiddleware.Gzip())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.Logger())

	logErr := middleware.LogErr()
	withDB := middleware.WithDB(db)
	withTX := middleware.WithTX(db)

	e.Get("/wine-comments/by-customer-id", handlers.GetWineCommentsByCustomerId, logErr, withDB)
	e.Post("/signup-customer-with-wine-comment", handlers.SignupCustomerWithWineComment, logErr, withTX)
	e.Post("/signup-customer-with-newsletter", handlers.SignupCustomerWithNewsletter, logErr, withTX)

	e.Get("/customer/wine-comments", handlers.SignupCustomerWithWineComment, logErr, withTX)
	e.Post("/customer/add-wine-comments", handlers.SignupCustomerWithWineComment, logErr, withTX)
	e.Post("/customer/register-to-newsletter", handlers.SignupCustomerWithNewsletter, logErr, withTX)

	// userService := user.NewService(userCollection, jwtKey)
	// userGroup := e.Group("/user")
	// userGroup.Post("/signup", userService.Signup)
	// userGroup.Post("/confirm", userService.ConfirmSignup)
	// userGroup.Post("/signin", userService.Signin)
	// userGroup.Post("/forgot-password", userService.ForgotPassword)
	// userGroup.Post("/reset-password", userService.ResetPassword)

	// sessionService := session.NewService(userCollection, jwtKey)
	// sessionMiddleware := middleware.CreateSessionMiddleware(userCollection, jwtKey)
	// sessionGroup := e.Group("/session", sessionMiddleware)
	// sessionGroup.Post("/signout", sessionService.Signout)
	// sessionGroup.Post("/change-password", sessionService.ChangePassword)
	// sessionGroup.Post("/change-email", sessionService.ChangeEmail)
	// sessionGroup.Post("/set-profile", sessionService.SetProfile)

	// adminService := admin.NewService(userCollection, jwtKey)
	// adminSessionMiddleware := middleware.CreateAdminSessionMiddleware(userCollection, jwtKey)
	// adminGroup := e.Group("/admin", adminSessionMiddleware)
	// adminGroup.Get("/get-users", adminService.GetUsers)
	// adminGroup.Post("/create-user", adminService.CreateUser)
	// adminGroup.Post("/change-user-password", adminService.ChangeUserPassword)
	// adminGroup.Post("/change-user-email", adminService.ChangeUserEmail)
	// adminGroup.Post("/set-user-roles", adminService.SetUserRoles)
	// adminGroup.Post("/set-user-profile", adminService.SetUserProfile)
	// adminGroup.Delete("/remove-users", adminService.RemoveUsers)
	// adminGroup.Post("/signout-users", adminService.SignoutUsers)
	// adminGroup.Post("/suspend-users", adminService.SuspendUsers)
	// adminGroup.Post("/unsuspend-users", adminService.UnsuspendUsers)
	// adminGroup.Delete("/remove-unconfirmed-users", adminService.RemoveUnconfirmedUsers)
	// adminGroup.Post("/remove-expired-reset-keys", adminService.RemoveExpiredResetKeys)

	port := viper.Get("SUNNY_PORT").(string)
	log.Printf("Running on port %s", port)
	e.Run(standard.New(port))
}
