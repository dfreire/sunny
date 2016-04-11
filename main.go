package main

import (
	"log"
	"net/http"

	"github.com/getdiskette/drive-b"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	// log "github.com/mgutz/logxi/v1"
	"github.com/spf13/viper"
)

func init() {
	initViper()
}

func initViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Couldn't read the config file")
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// e.SetDebug(true)

	e.Use(driveb.GlobalMiddleware())

	e.Get("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

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

	e.Run(standard.New(":3500"))
}
