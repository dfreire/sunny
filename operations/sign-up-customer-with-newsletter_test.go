package operations_test

import (
	"testing"

	"github.com/dfreire/sunny/mocks"
	"github.com/dfreire/sunny/operations"
	"github.com/dfreire/sunny/test"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jordan-wright/email"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSignupCustomerWithNewsletter(t *testing.T) {
	test.Setup()
	tx := test.CreateDB().Begin()
	mx := &mocks.Mailer{}

	req := operations.SignupCustomerWithNewsletterRequest{
		Name:       "Joe Doe",
		Email:      "joe.doe@mailinator.com",
		RoleId:     "wine_lover",
		LanguageId: "pt",
	}

	mx.On("SendEmail", mock.MatchedBy(func(e *email.Email) bool {
		require.Equal(t, "owner-6f66ed903426@mailinator.com", e.From)
		require.Equal(t, 1, len(e.To))
		require.Equal(t, req.Email, e.To[0])
		require.Equal(t, 0, len(e.Cc))
		require.Equal(t, 2, len(e.Bcc))
		require.Equal(t, "a-6f66ed903426@mailinator.com", e.Bcc[0])
		require.Equal(t, "b-6f66ed903426@mailinator.com", e.Bcc[1])
		require.Equal(t, "Newsletter Quinta de Soalheiro", e.Subject)
		require.Contains(t, string(e.HTML), "A sua subscrição")
		require.Equal(t, 0, len(e.Attachments))
		return true
	})).Return(nil).Once()

	assert.Nil(t, operations.SignupCustomerWithNewsletter(tx, mx, req))

	customer, err := operations.GetCustomerByEmail(tx, req.Email)
	assert.Nil(t, err)
	assert.Equal(t, req.Name, customer.Name)
	assert.Equal(t, req.Email, customer.Email)
	assert.Equal(t, req.RoleId, customer.RoleId)
	assert.Equal(t, req.LanguageId, customer.LanguageId)

	mx.AssertExpectations(t)
}
