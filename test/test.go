package test

import (
	"github.com/dfreire/sunny/model"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("SUNNY")
	viper.AutomaticEnv()
}

func Setup() {
	// nothing to do yet, but present to force init()
}

func CreateDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	model.Initialize(db)

	return db
}
