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
	"github.com/tealeg/xlsx"
)

func TestSendContactsToNewsletter(t *testing.T) {
	test.Setup()
	tx := test.CreateDB().Begin()
	mx := &mocks.Mailer{}

	test.SeedData(tx)

	mx.On("SendEmail", mock.MatchedBy(func(e *email.Email) bool {
		mailOk := e.From == "team-6f66ed903426@mailinator.com" &&
			len(e.To) == 1 &&
			e.To[0] == "owner-6f66ed903426@mailinator.com" &&
			len(e.Cc) == 0 &&
			len(e.Bcc) == 2 &&
			e.Bcc[0] == "a-6f66ed903426@mailinator.com" &&
			e.Bcc[1] == "b-6f66ed903426@mailinator.com" &&
			e.Subject == "Registos recebidos no website" &&
			strings.Contains(string(e.HTML), "Este é um mail enviado automaticamente") &&
			len(e.Attachments) == 1

		if !mailOk {
			return false
		}

		excelFile, err := xlsx.OpenBinary(e.Attachments[0].Content)
		if err != nil {
			return false
		}

		if len(excelFile.Sheets) != 1 {
			return false
		}

		sheet := excelFile.Sheets[0]
		if sheet.Name != "Registos" {
			return false
		}

		return true
	})).Return(nil).Once()

	assert.Nil(t, operations.SendContactsToNewsletter(tx, mx))

	mx.AssertExpectations(t)
}
