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
	"github.com/tealeg/xlsx"
)

func TestSendContactsToNewsletter(t *testing.T) {
	test.Setup()
	tx := test.CreateDB().Begin()
	mx := &mocks.Mailer{}

	test.SeedData(tx)

	mx.On("SendEmail", mock.MatchedBy(func(e *email.Email) bool {
		require.Equal(t, "team-6f66ed903426@mailinator.com", e.From)
		require.Equal(t, 1, len(e.To))
		require.Equal(t, "owner-6f66ed903426@mailinator.com", e.To[0])
		require.Equal(t, 0, len(e.Cc))
		require.Equal(t, 2, len(e.Bcc))
		require.Equal(t, "a-6f66ed903426@mailinator.com", e.Bcc[0])
		require.Equal(t, "b-6f66ed903426@mailinator.com", e.Bcc[1])
		require.Equal(t, "Registos recebidos no website", e.Subject)
		require.Contains(t, string(e.HTML), "Este é um mail enviado automaticamente")
		require.Equal(t, 1, len(e.Attachments))

		excelFile, err := xlsx.OpenBinary(e.Attachments[0].Content)
		require.NoError(t, err)
		require.Equal(t, 1, len(excelFile.Sheets))

		sheet := excelFile.Sheets[0]
		require.Equal(t, "Registos", sheet.Name)

		expectedRows := [][]string{
			[]string{"Nome", "Email", "Perfil", "Idioma"},
			[]string{"Joe Doe", "joe.doe@mailinator.com", "wine_lover", "pt"},
		}

		for i, row := range sheet.Rows {
			for j, expectedValue := range expectedRows[i] {
				actualValue, err := row.Cells[j].String()
				require.NoError(t, err)
				require.Equal(t, expectedValue, actualValue)
			}
		}

		return true
	})).Return(nil).Once()

	assert.Nil(t, operations.SendContactsToNewsletter(tx, mx))

	mx.AssertExpectations(t)
}
