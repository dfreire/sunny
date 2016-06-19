package operations_test

import (
	"testing"

	"github.com/dfreire/sunny/mocks"
	"github.com/dfreire/sunny/operations"
	"github.com/dfreire/sunny/test"
	"github.com/jordan-wright/email"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSignupCustomerWithWineComments(t *testing.T) {
	test.Setup()
	tx := test.CreateDB().Begin()
	mx := &mocks.Mailer{}

	req := operations.SignupCustomerWithWineCommentsRequest{
		Name:       "Joe Doe",
		Email:      "joe.doe@mailinator.com",
		RoleId:     "wine_lover",
		LanguageId: "pt",
	}

	req.WineComments = append(req.WineComments, operations.WineComment{
		WineId:   "soalheiro-alvarinho-2015",
		WineName: "Soalheiro Alvarinho",
		WineYear: 2015,
		Comment:  "Muito bom!",
	})

	req.WineComments = append(req.WineComments, operations.WineComment{
		WineId:   "soalheiro-primeiras-vinhas-2015",
		WineName: "Soalheiro Primeiras Vinhas",
		WineYear: 2015,
		Comment:  "Fantástico!",
	})

	mx.On("SendEmail", mock.MatchedBy(func(e *email.Email) bool {
		htmlStr := string(e.HTML)
		require.Equal(t, "owner-6f66ed903426@mailinator.com", e.From)
		require.Equal(t, 1, len(e.To))
		require.Equal(t, req.Email, e.To[0])
		require.Equal(t, 0, len(e.Cc))
		require.Equal(t, 2, len(e.Bcc))
		require.Equal(t, "a-6f66ed903426@mailinator.com", e.Bcc[0])
		require.Equal(t, "b-6f66ed903426@mailinator.com", e.Bcc[1])
		require.Equal(t, "Prova de Vinhos Quinta de Soalheiro", e.Subject)
		require.Contains(t, htmlStr, "Aproveitámos esta oportunidade")
		require.Contains(t, htmlStr, "Soalheiro Alvarinho 2015")
		require.Contains(t, htmlStr, "Muito bom!")
		require.Contains(t, htmlStr, "Soalheiro Primeiras Vinhas 2015")
		require.Contains(t, htmlStr, "Fantástico")
		require.Equal(t, 0, len(e.Attachments))
		return true
	})).Return(nil).Once()

	assert.Nil(t, operations.SignupCustomerWithWineComments(tx, mx, req))

	customer, err := operations.GetCustomerByEmail(tx, req.Email)
	assert.Nil(t, err)
	assert.Equal(t, req.Name, customer.Name)
	assert.Equal(t, req.Email, customer.Email)
	assert.Equal(t, req.RoleId, customer.RoleId)
	assert.Equal(t, req.LanguageId, customer.LanguageId)

	comments, err := operations.GetWineCommentsByCustomerId(tx, customer.ID)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(comments))
	for i := 0; i < 2; i++ {
		assert.Equal(t, comments[i].WineId, req.WineComments[i].WineId)
		assert.Equal(t, comments[i].WineYear, req.WineComments[i].WineYear)
		assert.Equal(t, comments[i].Comment, req.WineComments[i].Comment)
	}

	mx.AssertExpectations(t)
}
