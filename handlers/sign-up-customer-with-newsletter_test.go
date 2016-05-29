package handlers

import (
	"testing"

	"github.com/dfreire/sunny/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockedContext struct {
	mock.Mock
}

func (m *mockedContext) Get(key string) interface{} {
	args := m.Called(key)
	return args.Get(0)
}

func TestSignupCustomerWithNewsletter(t *testing.T) {
	c := new(mockedContext)

	db, err := gorm.Open("sqlite3", ":memory:")
	assert.Nil(t, err)
	tx := db.Begin()
	c.On("Get", middleware.TX).Return(tx)
}
