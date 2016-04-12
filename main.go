package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/getdiskette/drive-b"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	// log "github.com/mgutz/logxi/v1"
	_ "github.com/mattn/go-sqlite3"
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

const SCHEMA = `
CREATE TABLE IF NOT EXISTS customer_role (
	id  TEXT PRIMARY KEY
);

INSERT INTO customer_role(id) VALUES ('sommelier');
INSERT INTO customer_role(id) VALUES ('restaurant');
INSERT INTO customer_role(id) VALUES ('wine_distribution');
INSERT INTO customer_role(id) VALUES ('wine_shop');
INSERT INTO customer_role(id) VALUES ('wine_lover');
INSERT INTO customer_role(id) VALUES ('other');

CREATE TABLE IF NOT EXISTS customer (
	id          TEXT PRIMARY KEY,
	email       TEXT,
	role_id     TEXT,
	created_at  TEXT,

	FOREIGN KEY(role_id) REFERENCES customer_role(id)
);

CREATE TABLE IF NOT EXISTS customer_wine_comment (
	id           TEXT PRIMARY KEY,
	customer_id  TEXT,
	wine_id      TEXT,
	year         NUMBER,
	created_at   TEXT,
	updated_at   TEXT,
	comment      TEXT,

	UNIQUE(user_id, wine_id, year)
);
`

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	db.Exec(SCHEMA)

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
