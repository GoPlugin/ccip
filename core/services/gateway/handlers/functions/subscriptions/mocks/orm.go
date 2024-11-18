// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	subscriptions "github.com/goplugin/pluginv3.0/v2/core/services/gateway/handlers/functions/subscriptions"
	mock "github.com/stretchr/testify/mock"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

type ORM_Expecter struct {
	mock *mock.Mock
}

func (_m *ORM) EXPECT() *ORM_Expecter {
	return &ORM_Expecter{mock: &_m.Mock}
}

// GetSubscriptions provides a mock function with given fields: ctx, offset, limit
func (_m *ORM) GetSubscriptions(ctx context.Context, offset uint, limit uint) ([]subscriptions.StoredSubscription, error) {
	ret := _m.Called(ctx, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetSubscriptions")
	}

	var r0 []subscriptions.StoredSubscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, uint) ([]subscriptions.StoredSubscription, error)); ok {
		return rf(ctx, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint, uint) []subscriptions.StoredSubscription); ok {
		r0 = rf(ctx, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]subscriptions.StoredSubscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint, uint) error); ok {
		r1 = rf(ctx, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_GetSubscriptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubscriptions'
type ORM_GetSubscriptions_Call struct {
	*mock.Call
}

// GetSubscriptions is a helper method to define mock.On call
//   - ctx context.Context
//   - offset uint
//   - limit uint
func (_e *ORM_Expecter) GetSubscriptions(ctx interface{}, offset interface{}, limit interface{}) *ORM_GetSubscriptions_Call {
	return &ORM_GetSubscriptions_Call{Call: _e.mock.On("GetSubscriptions", ctx, offset, limit)}
}

func (_c *ORM_GetSubscriptions_Call) Run(run func(ctx context.Context, offset uint, limit uint)) *ORM_GetSubscriptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint), args[2].(uint))
	})
	return _c
}

func (_c *ORM_GetSubscriptions_Call) Return(_a0 []subscriptions.StoredSubscription, _a1 error) *ORM_GetSubscriptions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_GetSubscriptions_Call) RunAndReturn(run func(context.Context, uint, uint) ([]subscriptions.StoredSubscription, error)) *ORM_GetSubscriptions_Call {
	_c.Call.Return(run)
	return _c
}

// UpsertSubscription provides a mock function with given fields: ctx, subscription
func (_m *ORM) UpsertSubscription(ctx context.Context, subscription subscriptions.StoredSubscription) error {
	ret := _m.Called(ctx, subscription)

	if len(ret) == 0 {
		panic("no return value specified for UpsertSubscription")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, subscriptions.StoredSubscription) error); ok {
		r0 = rf(ctx, subscription)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_UpsertSubscription_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpsertSubscription'
type ORM_UpsertSubscription_Call struct {
	*mock.Call
}

// UpsertSubscription is a helper method to define mock.On call
//   - ctx context.Context
//   - subscription subscriptions.StoredSubscription
func (_e *ORM_Expecter) UpsertSubscription(ctx interface{}, subscription interface{}) *ORM_UpsertSubscription_Call {
	return &ORM_UpsertSubscription_Call{Call: _e.mock.On("UpsertSubscription", ctx, subscription)}
}

func (_c *ORM_UpsertSubscription_Call) Run(run func(ctx context.Context, subscription subscriptions.StoredSubscription)) *ORM_UpsertSubscription_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(subscriptions.StoredSubscription))
	})
	return _c
}

func (_c *ORM_UpsertSubscription_Call) Return(_a0 error) *ORM_UpsertSubscription_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_UpsertSubscription_Call) RunAndReturn(run func(context.Context, subscriptions.StoredSubscription) error) *ORM_UpsertSubscription_Call {
	_c.Call.Return(run)
	return _c
}

// NewORM creates a new instance of ORM. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewORM(t interface {
	mock.TestingT
	Cleanup(func())
}) *ORM {
	mock := &ORM{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}