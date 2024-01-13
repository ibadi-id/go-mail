// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	httputil "github.com/ibadi-id/go-mail/internal/httputil"
	mail "github.com/ibadi-id/go-mail/mail"

	mock "github.com/stretchr/testify/mock"
)

// Requester is an autogenerated mock type for the Requester type
type Requester struct {
	mock.Mock
}

// Do provides a mock function with given fields: ctx, r, payload, responder
func (_m *Requester) Do(ctx context.Context, r *httputil.Request, payload httputil.Payload, responder httputil.Responder) (mail.Response, error) {
	ret := _m.Called(ctx, r, payload, responder)

	var r0 mail.Response
	if rf, ok := ret.Get(0).(func(context.Context, *httputil.Request, httputil.Payload, httputil.Responder) mail.Response); ok {
		r0 = rf(ctx, r, payload, responder)
	} else {
		r0 = ret.Get(0).(mail.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *httputil.Request, httputil.Payload, httputil.Responder) error); ok {
		r1 = rf(ctx, r, payload, responder)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
