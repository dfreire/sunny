package commands_test

import (
	"strings"
	"testing"

	"github.com/dfreire/sunny/commands"
	"github.com/dfreire/sunny/mocks"
	"github.com/dfreire/sunny/test"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jordan-wright/email"
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

	mx.On("SendEmail", mock.MatchedBy(func(e *email.Email) bool {
		return e.From == "owner-6f66ed903426@mailinator.com" &&
			len(e.To) == 1 &&
			e.To[0] == reqData.Email &&
			len(e.Cc) == 0 &&
			len(e.Bcc) == 2 &&
			e.Bcc[0] == "a-6f66ed903426@mailinator.com" &&
			e.Bcc[1] == "b-6f66ed903426@mailinator.com" &&
			e.Subject == "Newsletter Quinta de Soalheiro" &&
			strings.Contains(string(e.HTML), "A sua subscrição") &&
			len(e.Attachments) == 0
	})).Return(nil).Once()

	assert.Nil(t, commands.SignupCustomerWithNewsletter(tx, mx, reqData))
	mx.AssertExpectations(t)
}
