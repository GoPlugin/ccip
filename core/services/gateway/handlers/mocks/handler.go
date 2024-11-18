// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	api "github.com/goplugin/pluginv3.0/v2/core/services/gateway/api"

	handlers "github.com/goplugin/pluginv3.0/v2/core/services/gateway/handlers"

	mock "github.com/stretchr/testify/mock"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

type Handler_Expecter struct {
	mock *mock.Mock
}

func (_m *Handler) EXPECT() *Handler_Expecter {
	return &Handler_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *Handler) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Handler_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type Handler_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *Handler_Expecter) Close() *Handler_Close_Call {
	return &Handler_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *Handler_Close_Call) Run(run func()) *Handler_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Handler_Close_Call) Return(_a0 error) *Handler_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Handler_Close_Call) RunAndReturn(run func() error) *Handler_Close_Call {
	_c.Call.Return(run)
	return _c
}

// HandleNodeMessage provides a mock function with given fields: ctx, msg, nodeAddr
func (_m *Handler) HandleNodeMessage(ctx context.Context, msg *api.Message, nodeAddr string) error {
	ret := _m.Called(ctx, msg, nodeAddr)

	if len(ret) == 0 {
		panic("no return value specified for HandleNodeMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.Message, string) error); ok {
		r0 = rf(ctx, msg, nodeAddr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Handler_HandleNodeMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleNodeMessage'
type Handler_HandleNodeMessage_Call struct {
	*mock.Call
}

// HandleNodeMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - msg *api.Message
//   - nodeAddr string
func (_e *Handler_Expecter) HandleNodeMessage(ctx interface{}, msg interface{}, nodeAddr interface{}) *Handler_HandleNodeMessage_Call {
	return &Handler_HandleNodeMessage_Call{Call: _e.mock.On("HandleNodeMessage", ctx, msg, nodeAddr)}
}

func (_c *Handler_HandleNodeMessage_Call) Run(run func(ctx context.Context, msg *api.Message, nodeAddr string)) *Handler_HandleNodeMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*api.Message), args[2].(string))
	})
	return _c
}

func (_c *Handler_HandleNodeMessage_Call) Return(_a0 error) *Handler_HandleNodeMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Handler_HandleNodeMessage_Call) RunAndReturn(run func(context.Context, *api.Message, string) error) *Handler_HandleNodeMessage_Call {
	_c.Call.Return(run)
	return _c
}

// HandleUserMessage provides a mock function with given fields: ctx, msg, callbackCh
func (_m *Handler) HandleUserMessage(ctx context.Context, msg *api.Message, callbackCh chan<- handlers.UserCallbackPayload) error {
	ret := _m.Called(ctx, msg, callbackCh)

	if len(ret) == 0 {
		panic("no return value specified for HandleUserMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.Message, chan<- handlers.UserCallbackPayload) error); ok {
		r0 = rf(ctx, msg, callbackCh)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Handler_HandleUserMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleUserMessage'
type Handler_HandleUserMessage_Call struct {
	*mock.Call
}

// HandleUserMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - msg *api.Message
//   - callbackCh chan<- handlers.UserCallbackPayload
func (_e *Handler_Expecter) HandleUserMessage(ctx interface{}, msg interface{}, callbackCh interface{}) *Handler_HandleUserMessage_Call {
	return &Handler_HandleUserMessage_Call{Call: _e.mock.On("HandleUserMessage", ctx, msg, callbackCh)}
}

func (_c *Handler_HandleUserMessage_Call) Run(run func(ctx context.Context, msg *api.Message, callbackCh chan<- handlers.UserCallbackPayload)) *Handler_HandleUserMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*api.Message), args[2].(chan<- handlers.UserCallbackPayload))
	})
	return _c
}

func (_c *Handler_HandleUserMessage_Call) Return(_a0 error) *Handler_HandleUserMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Handler_HandleUserMessage_Call) RunAndReturn(run func(context.Context, *api.Message, chan<- handlers.UserCallbackPayload) error) *Handler_HandleUserMessage_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: _a0
func (_m *Handler) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Handler_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Handler_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Handler_Expecter) Start(_a0 interface{}) *Handler_Start_Call {
	return &Handler_Start_Call{Call: _e.mock.On("Start", _a0)}
}

func (_c *Handler_Start_Call) Run(run func(_a0 context.Context)) *Handler_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Handler_Start_Call) Return(_a0 error) *Handler_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Handler_Start_Call) RunAndReturn(run func(context.Context) error) *Handler_Start_Call {
	_c.Call.Return(run)
	return _c
}

// NewHandler creates a new instance of Handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}