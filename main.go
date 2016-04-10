package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	log "github.com/mgutz/logxi/v1"
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

	// e.SetDebug(true)

	e.Get("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Run(fasthttp.New(":1323"))
}
