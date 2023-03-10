// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	response "github.com/Risuii/helpers/response"
	mock "github.com/stretchr/testify/mock"

	token "github.com/Risuii/models/token"

	users "github.com/Risuii/models/users"
)

// UserUseCase is an autogenerated mock type for the UserUseCase type
type UserUseCase struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, params
func (_m *UserUseCase) Login(ctx context.Context, params users.EmployeeLogin) (response.Response, token.Token) {
	ret := _m.Called(ctx, params)

	var r0 response.Response
	if rf, ok := ret.Get(0).(func(context.Context, users.EmployeeLogin) response.Response); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.Response)
		}
	}

	var r1 token.Token
	if rf, ok := ret.Get(1).(func(context.Context, users.EmployeeLogin) token.Token); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Get(1).(token.Token)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, params
func (_m *UserUseCase) Register(ctx context.Context, params users.Employee) response.Response {
	ret := _m.Called(ctx, params)

	var r0 response.Response
	if rf, ok := ret.Get(0).(func(context.Context, users.Employee) response.Response); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.Response)
		}
	}

	return r0
}

type mockConstructorTestingTNewUserUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUseCase creates a new instance of UserUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUseCase(t mockConstructorTestingTNewUserUseCase) *UserUseCase {
	mock := &UserUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
