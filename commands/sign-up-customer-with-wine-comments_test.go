package commands_test

import (
	"strings"
	"testing"

	"github.com/dfreire/sunny/commands"
	"github.com/dfreire/sunny/mocks"
	"github.com/dfreire/sunny/queries"
	"github.com/dfreire/sunny/test"
	"github.com/jordan-wright/email"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignupCustomerWithWineComments(t *testing.T) {
	test.Setup()
	tx := test.CreateDB().Begin()
	mx := &mocks.Mailer{}

	req := commands.SignupCustomerWithWineCommentsRequest{
		Name:       "Joe Doe",
		Email:      "joe.doe@mailinator.com",
		RoleId:     "wine_lover",
		LanguageId: "pt",
	}

	req.WineComments = append(req.WineComments, commands.WineComment{
		WineId:   "soalheiro-alvarinho-2015",
		WineName: "Soalheiro Alvarinho",
		WineYear: 2015,
		Comment:  "Muito bom!",
	})

	req.WineComments = append(req.WineComments, commands.WineComment{
		WineId:   "soalheiro-primeiras-vinhas-2015",
		WineName: "Soalheiro Primeiras Vinhas",
		WineYear: 2015,
		Comment:  "Fantástico!",
	})

	mx.On("SendEmail", mock.MatchedBy(func(e *email.Email) bool {
		htmlStr := string(e.HTML)
		return e.From == "owner-6f66ed903426@mailinator.com" &&
			len(e.To) == 1 &&
			e.To[0] == req.Email &&
			len(e.Cc) == 0 &&
			len(e.Bcc) == 2 &&
			e.Bcc[0] == "a-6f66ed903426@mailinator.com" &&
			e.Bcc[1] == "b-6f66ed903426@mailinator.com" &&
			e.Subject == "Prova de Vinhos Quinta de Soalheiro" &&
			strings.Contains(htmlStr, "Aproveitámos esta oportunidade") &&
			strings.Contains(htmlStr, "Soalheiro Alvarinho 2015") &&
			strings.Contains(htmlStr, "Muito bom!") &&
			strings.Contains(htmlStr, "Soalheiro Primeiras Vinhas 2015") &&
			strings.Contains(htmlStr, "Fantástico!") &&
			len(e.Attachments) == 0
	})).Return(nil).Once()

	assert.Nil(t, commands.SignupCustomerWithWineComments(tx, mx, req))

	customer, err := queries.GetCustomerByEmail(tx, req.Email)
	assert.Nil(t, err)
	assert.Equal(t, req.Name, customer.Name)
	assert.Equal(t, req.Email, customer.Email)
	assert.Equal(t, req.RoleId, customer.RoleId)
	assert.Equal(t, req.LanguageId, customer.LanguageId)

	comments, err := queries.GetWineCommentsByCustomerId(tx, customer.ID)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(comments))
	for i := 0; i < 2; i++ {
		assert.Equal(t, comments[i].WineId, req.WineComments[i].WineId)
		assert.Equal(t, comments[i].WineYear, req.WineComments[i].WineYear)
		assert.Equal(t, comments[i].Comment, req.WineComments[i].Comment)
	}

	mx.AssertExpectations(t)
}
