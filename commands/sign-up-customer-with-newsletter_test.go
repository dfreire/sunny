package commands_test

import (
	"testing"

	"github.com/dfreire/sunny/commands"
	"github.com/dfreire/sunny/mailer"
	"github.com/dfreire/sunny/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
	viper.SetEnvPrefix("SUNNY")
	viper.AutomaticEnv()
}

func TestSignupCustomerWithNewsletter(t *testing.T) {
	tx := newDB().Begin()
	mx := mailer.NewLogMailer()

	reqData := commands.SignupCustomerWithNewsletterRequestData{
		Name:       "Joe Doe",
		Email:      "joe.doe@mailinator.com",
		RoleId:     "wine_lover",
		LanguageId: "pt",
	}

	assert.Nil(t, commands.SignupCustomerWithNewsletter(tx, mx, reqData))
}

func newDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	model.Initialize(db)

	return db
}
