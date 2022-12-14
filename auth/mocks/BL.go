// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import spec "github.com/authentication-service/spec"

// BL is an autogenerated mock type for the BL type
type BL struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, loginRequest
func (_m *BL) Login(ctx context.Context, loginRequest spec.LoginRequest) (spec.AuthUserResponse, error) {
	ret := _m.Called(ctx, loginRequest)

	var r0 spec.AuthUserResponse
	if rf, ok := ret.Get(0).(func(context.Context, spec.LoginRequest) spec.AuthUserResponse); ok {
		r0 = rf(ctx, loginRequest)
	} else {
		r0 = ret.Get(0).(spec.AuthUserResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, spec.LoginRequest) error); ok {
		r1 = rf(ctx, loginRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
