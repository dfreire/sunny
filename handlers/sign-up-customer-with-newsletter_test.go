package handlers_test

import (
	"net/http"
	"testing"

	"github.com/dfreire/sunny/commands"
	"github.com/dfreire/sunny/handlers"
	"github.com/dfreire/sunny/mailer"
	"github.com/dfreire/sunny/middleware"
	"github.com/dfreire/sunny/mocks"
	"github.com/dfreire/sunny/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func init() {
	viper.SetEnvPrefix("SUNNY")
	viper.AutomaticEnv()
}

func TestSignupCustomerWithNewsletter(t *testing.T) {
	c := createMockContext()

	c.On("Bind", mock.AnythingOfType("*commands.SignupCustomerWithNewsletterRequestData")).
		Run(func(args mock.Arguments) {
			arg := args.Get(0).(*commands.SignupCustomerWithNewsletterRequestData)
			arg.Name = "Joe Doe"
			arg.Email = "joe.doe@mailinator.com"
			arg.RoleId = "wine_lover"
			arg.LanguageId = "pt"
		}).
		Return(nil)

	assert.Nil(t, handlers.SignupCustomerWithNewsletter(c))

	c.AssertExpectations(t)
}

func createMockContext() *mocks.Context {
	c := &mocks.Context{}

	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	model.Initialize(db)

	tx := db.Begin()
	c.On("Get", middleware.TX).Return(tx)

	mx := mailer.NewLogMailer()
	c.On("Get", middleware.MAILER).Return(mx)

	c.On("JSON", http.StatusOK, mock.AnythingOfType("handlers.jsonResponse")).Return(nil)
	// c.On("JSON", http.StatusInternalServerError, mock.AnythingOfType("handlers.jsonResponse")).Return(nil)

	return c
}
