package operations_test

import (
	"strings"
	"testing"

	"github.com/dfreire/sunny/mocks"
	"github.com/dfreire/sunny/operations"
	"github.com/dfreire/sunny/test"
	"github.com/jordan-wright/email"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSendContactsToNewsletter(t *testing.T) {
	test.Setup()
	tx := test.CreateDB().Begin()
	mx := &mocks.Mailer{}

	test.SeedData(tx)

	// customers, err := operations.GetCustomers(tx)
	// assert.Nil(t, err)
	// log.Printf("CUSTOMERS2: %+v", customers)

	mx.On("SendEmail", mock.MatchedBy(func(e *email.Email) bool {
		return e.From == "team-6f66ed903426@mailinator.com" &&
			len(e.To) == 1 &&
			e.To[0] == "owner-6f66ed903426@mailinator.com" &&
			len(e.Cc) == 0 &&
			len(e.Bcc) == 2 &&
			e.Bcc[0] == "a-6f66ed903426@mailinator.com" &&
			e.Bcc[1] == "b-6f66ed903426@mailinator.com" &&
			e.Subject == "Registos recebidos no website" &&
			strings.Contains(string(e.HTML), "Este é um mail enviado automaticamente") &&
			len(e.Attachments) == 1
	})).Return(nil).Once()

	assert.Nil(t, operations.SendContactsToNewsletter(tx, mx))

	mx.AssertExpectations(t)
}
