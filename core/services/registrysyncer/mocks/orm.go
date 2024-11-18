// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	registrysyncer "github.com/goplugin/pluginv3.0/v2/core/services/registrysyncer"
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

// AddLocalRegistry provides a mock function with given fields: ctx, localRegistry
func (_m *ORM) AddLocalRegistry(ctx context.Context, localRegistry registrysyncer.LocalRegistry) error {
	ret := _m.Called(ctx, localRegistry)

	if len(ret) == 0 {
		panic("no return value specified for AddLocalRegistry")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, registrysyncer.LocalRegistry) error); ok {
		r0 = rf(ctx, localRegistry)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ORM_AddLocalRegistry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddLocalRegistry'
type ORM_AddLocalRegistry_Call struct {
	*mock.Call
}

// AddLocalRegistry is a helper method to define mock.On call
//   - ctx context.Context
//   - localRegistry registrysyncer.LocalRegistry
func (_e *ORM_Expecter) AddLocalRegistry(ctx interface{}, localRegistry interface{}) *ORM_AddLocalRegistry_Call {
	return &ORM_AddLocalRegistry_Call{Call: _e.mock.On("AddLocalRegistry", ctx, localRegistry)}
}

func (_c *ORM_AddLocalRegistry_Call) Run(run func(ctx context.Context, localRegistry registrysyncer.LocalRegistry)) *ORM_AddLocalRegistry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(registrysyncer.LocalRegistry))
	})
	return _c
}

func (_c *ORM_AddLocalRegistry_Call) Return(_a0 error) *ORM_AddLocalRegistry_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ORM_AddLocalRegistry_Call) RunAndReturn(run func(context.Context, registrysyncer.LocalRegistry) error) *ORM_AddLocalRegistry_Call {
	_c.Call.Return(run)
	return _c
}

// LatestLocalRegistry provides a mock function with given fields: ctx
func (_m *ORM) LatestLocalRegistry(ctx context.Context) (*registrysyncer.LocalRegistry, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for LatestLocalRegistry")
	}

	var r0 *registrysyncer.LocalRegistry
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*registrysyncer.LocalRegistry, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *registrysyncer.LocalRegistry); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*registrysyncer.LocalRegistry)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ORM_LatestLocalRegistry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LatestLocalRegistry'
type ORM_LatestLocalRegistry_Call struct {
	*mock.Call
}

// LatestLocalRegistry is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ORM_Expecter) LatestLocalRegistry(ctx interface{}) *ORM_LatestLocalRegistry_Call {
	return &ORM_LatestLocalRegistry_Call{Call: _e.mock.On("LatestLocalRegistry", ctx)}
}

func (_c *ORM_LatestLocalRegistry_Call) Run(run func(ctx context.Context)) *ORM_LatestLocalRegistry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ORM_LatestLocalRegistry_Call) Return(_a0 *registrysyncer.LocalRegistry, _a1 error) *ORM_LatestLocalRegistry_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ORM_LatestLocalRegistry_Call) RunAndReturn(run func(context.Context) (*registrysyncer.LocalRegistry, error)) *ORM_LatestLocalRegistry_Call {
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
