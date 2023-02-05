// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	activitys "github.com/Risuii/models/activitys"

	mock "github.com/stretchr/testify/mock"
)

// ActivityRepository is an autogenerated mock type for the ActivityRepository type
type ActivityRepository struct {
	mock.Mock
}

// AddActivity provides a mock function with given fields: ctx, userID, params
func (_m *ActivityRepository) AddActivity(ctx context.Context, userID int64, params activitys.Activity) (int64, error) {
	ret := _m.Called(ctx, userID, params)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, int64, activitys.Activity) int64); ok {
		r0 = rf(ctx, userID, params)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, activitys.Activity) error); ok {
		r1 = rf(ctx, userID, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ActivityRepository) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByID provides a mock function with given fields: ctx, id
func (_m *ActivityRepository) FindByID(ctx context.Context, id int64) (activitys.Activity, error) {
	ret := _m.Called(ctx, id)

	var r0 activitys.Activity
	if rf, ok := ret.Get(0).(func(context.Context, int64) activitys.Activity); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(activitys.Activity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Riwayat provides a mock function with given fields: ctx, userID, params
func (_m *ActivityRepository) Riwayat(ctx context.Context, userID int64, params activitys.DateReq) ([]activitys.Activity, error) {
	ret := _m.Called(ctx, userID, params)

	var r0 []activitys.Activity
	if rf, ok := ret.Get(0).(func(context.Context, int64, activitys.DateReq) []activitys.Activity); ok {
		r0 = rf(ctx, userID, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]activitys.Activity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, activitys.DateReq) error); ok {
		r1 = rf(ctx, userID, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateActivity provides a mock function with given fields: ctx, id, params
func (_m *ActivityRepository) UpdateActivity(ctx context.Context, id int64, params activitys.Activity) error {
	ret := _m.Called(ctx, id, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, activitys.Activity) error); ok {
		r0 = rf(ctx, id, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewActivityRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewActivityRepository creates a new instance of ActivityRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewActivityRepository(t mockConstructorTestingTNewActivityRepository) *ActivityRepository {
	mock := &ActivityRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
