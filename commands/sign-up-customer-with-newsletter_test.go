package commands_test

import (
	"testing"

	"github.com/dfreire/sunny/commands"
	"github.com/dfreire/sunny/mailer"
	"github.com/dfreire/sunny/test"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestSignupCustomerWithNewsletter(t *testing.T) {
	test.Setup()
	tx := test.CreateDB().Begin()
	mx := mailer.NewLogMailer()

	reqData := commands.SignupCustomerWithNewsletterRequestData{
		Name:       "Joe Doe",
		Email:      "joe.doe@mailinator.com",
		RoleId:     "wine_lover",
		LanguageId: "pt",
	}

	assert.Nil(t, commands.SignupCustomerWithNewsletter(tx, mx, reqData))
}
