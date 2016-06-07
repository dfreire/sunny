package commands_test

import (
	"testing"

	"github.com/dfreire/sunny/commands"
	"github.com/dfreire/sunny/mocks"
	"github.com/dfreire/sunny/test"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignupCustomerWithNewsletter(t *testing.T) {
	test.Setup()
	tx := test.CreateDB().Begin()
	mx := &mocks.Mailer{}

	reqData := commands.SignupCustomerWithNewsletterRequestData{
		Name:       "Joe Doe",
		Email:      "joe.doe@mailinator.com",
		RoleId:     "wine_lover",
		LanguageId: "pt",
	}

	mx.On("SendEmail", mock.AnythingOfType("*email.Email")).Return(nil)

	assert.Nil(t, commands.SignupCustomerWithNewsletter(tx, mx, reqData))
}
