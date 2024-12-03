// Code generated by mockery v2.43.2. DO NOT EDIT.

package session

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockManager is an autogenerated mock type for the Manager type
type MockManager struct {
	mock.Mock
}

// CreateSession provides a mock function with given fields: ctx, userID
func (_m *MockManager) CreateSession(ctx context.Context, userID string) (*Session, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for CreateSession")
	}

	var r0 *Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*Session, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *Session); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Session)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateSession provides a mock function with given fields: ctx, sessionID
func (_m *MockManager) ValidateSession(ctx context.Context, sessionID string) (*Session, error) {
	ret := _m.Called(ctx, sessionID)

	if len(ret) == 0 {
		panic("no return value specified for ValidateSession")
	}

	var r0 *Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*Session, error)); ok {
		return rf(ctx, sessionID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *Session); ok {
		r0 = rf(ctx, sessionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Session)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockManager creates a new instance of MockManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockManager {
	mock := &MockManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}