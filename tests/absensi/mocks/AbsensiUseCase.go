// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	absensis "github.com/Risuii/models/absensis"

	mock "github.com/stretchr/testify/mock"

	response "github.com/Risuii/helpers/response"

	token "github.com/Risuii/models/token"
)

// AbsensiUseCase is an autogenerated mock type for the AbsensiUseCase type
type AbsensiUseCase struct {
	mock.Mock
}

// Checkin provides a mock function with given fields: ctx, userID, name
func (_m *AbsensiUseCase) Checkin(ctx context.Context, userID int64, name string) (response.Response, token.Token) {
	ret := _m.Called(ctx, userID, name)

	var r0 response.Response
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) response.Response); ok {
		r0 = rf(ctx, userID, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.Response)
		}
	}

	var r1 token.Token
	if rf, ok := ret.Get(1).(func(context.Context, int64, string) token.Token); ok {
		r1 = rf(ctx, userID, name)
	} else {
		r1 = ret.Get(1).(token.Token)
	}

	return r0, r1
}

// Checkout provides a mock function with given fields: ctx, checkinID
func (_m *AbsensiUseCase) Checkout(ctx context.Context, checkinID int64) response.Response {
	ret := _m.Called(ctx, checkinID)

	var r0 response.Response
	if rf, ok := ret.Get(0).(func(context.Context, int64) response.Response); ok {
		r0 = rf(ctx, checkinID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.Response)
		}
	}

	return r0
}

// Riwayat provides a mock function with given fields: ctx, params
func (_m *AbsensiUseCase) Riwayat(ctx context.Context, params absensis.Riwayat) response.Response {
	ret := _m.Called(ctx, params)

	var r0 response.Response
	if rf, ok := ret.Get(0).(func(context.Context, absensis.Riwayat) response.Response); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.Response)
		}
	}

	return r0
}

type mockConstructorTestingTNewAbsensiUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewAbsensiUseCase creates a new instance of AbsensiUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAbsensiUseCase(t mockConstructorTestingTNewAbsensiUseCase) *AbsensiUseCase {
	mock := &AbsensiUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
