package operations_test

import (
	"testing"

	"github.com/dfreire/sunny/mocks"
	"github.com/dfreire/sunny/operations"
	"github.com/dfreire/sunny/test"
	"github.com/stretchr/testify/assert"
)

func TestSendContactsToNewsletter(t *testing.T) {
	test.Setup()
	tx := test.CreateDB().Begin()
	mx := &mocks.Mailer{}

	assert.Nil(t, operations.SendContactsToNewsletter(tx, mx))
}
