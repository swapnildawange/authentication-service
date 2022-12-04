// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"

import spec "github.com/authentication-service/spec"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetUser provides a mock function with given fields: ctx, userId
func (_m *Repository) GetUser(ctx context.Context, userId int) (spec.User, error) {
	ret := _m.Called(ctx, userId)

	var r0 spec.User
	if rf, ok := ret.Get(0).(func(context.Context, int) spec.User); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Get(0).(spec.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserFromAuth provides a mock function with given fields: ctx, email
func (_m *Repository) GetUserFromAuth(ctx context.Context, email string) (spec.AuthUserResponse, error) {
	ret := _m.Called(ctx, email)

	var r0 spec.AuthUserResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) spec.AuthUserResponse); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(spec.AuthUserResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
