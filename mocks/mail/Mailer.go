// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mail "github.com/ainsleyclark/go-mail/mail"
	mock "github.com/stretchr/testify/mock"
)

// Mailer is an autogenerated mock type for the Mailer type
type Mailer struct {
	mock.Mock
}

// Send provides a mock function with given fields: t
func (_m *Mailer) Send(t *mail.Transmission) (mail.Response, error) {
	ret := _m.Called(t)

	var r0 mail.Response
	if rf, ok := ret.Get(0).(func(*mail.Transmission) mail.Response); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Get(0).(mail.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*mail.Transmission) error); ok {
		r1 = rf(t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}