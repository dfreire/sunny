package mocks

import "github.com/stretchr/testify/mock"

import "github.com/jordan-wright/email"

// Mailer is an autogenerated mock type for the Mailer type
type Mailer struct {
	mock.Mock
}

// SendEmail provides a mock function with given fields: e
func (_m *Mailer) SendEmail(e *email.Email) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(*email.Email) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}