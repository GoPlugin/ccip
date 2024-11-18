// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	config "github.com/goplugin/pluginv3.0/v2/core/config"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// TelemetryIngress is an autogenerated mock type for the TelemetryIngress type
type TelemetryIngress struct {
	mock.Mock
}

type TelemetryIngress_Expecter struct {
	mock *mock.Mock
}

func (_m *TelemetryIngress) EXPECT() *TelemetryIngress_Expecter {
	return &TelemetryIngress_Expecter{mock: &_m.Mock}
}

// BufferSize provides a mock function with given fields:
func (_m *TelemetryIngress) BufferSize() uint {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for BufferSize")
	}

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// TelemetryIngress_BufferSize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BufferSize'
type TelemetryIngress_BufferSize_Call struct {
	*mock.Call
}

// BufferSize is a helper method to define mock.On call
func (_e *TelemetryIngress_Expecter) BufferSize() *TelemetryIngress_BufferSize_Call {
	return &TelemetryIngress_BufferSize_Call{Call: _e.mock.On("BufferSize")}
}

func (_c *TelemetryIngress_BufferSize_Call) Run(run func()) *TelemetryIngress_BufferSize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TelemetryIngress_BufferSize_Call) Return(_a0 uint) *TelemetryIngress_BufferSize_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TelemetryIngress_BufferSize_Call) RunAndReturn(run func() uint) *TelemetryIngress_BufferSize_Call {
	_c.Call.Return(run)
	return _c
}

// Endpoints provides a mock function with given fields:
func (_m *TelemetryIngress) Endpoints() []config.TelemetryIngressEndpoint {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Endpoints")
	}

	var r0 []config.TelemetryIngressEndpoint
	if rf, ok := ret.Get(0).(func() []config.TelemetryIngressEndpoint); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]config.TelemetryIngressEndpoint)
		}
	}

	return r0
}

// TelemetryIngress_Endpoints_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Endpoints'
type TelemetryIngress_Endpoints_Call struct {
	*mock.Call
}

// Endpoints is a helper method to define mock.On call
func (_e *TelemetryIngress_Expecter) Endpoints() *TelemetryIngress_Endpoints_Call {
	return &TelemetryIngress_Endpoints_Call{Call: _e.mock.On("Endpoints")}
}

func (_c *TelemetryIngress_Endpoints_Call) Run(run func()) *TelemetryIngress_Endpoints_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TelemetryIngress_Endpoints_Call) Return(_a0 []config.TelemetryIngressEndpoint) *TelemetryIngress_Endpoints_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TelemetryIngress_Endpoints_Call) RunAndReturn(run func() []config.TelemetryIngressEndpoint) *TelemetryIngress_Endpoints_Call {
	_c.Call.Return(run)
	return _c
}

// Logging provides a mock function with given fields:
func (_m *TelemetryIngress) Logging() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Logging")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// TelemetryIngress_Logging_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Logging'
type TelemetryIngress_Logging_Call struct {
	*mock.Call
}

// Logging is a helper method to define mock.On call
func (_e *TelemetryIngress_Expecter) Logging() *TelemetryIngress_Logging_Call {
	return &TelemetryIngress_Logging_Call{Call: _e.mock.On("Logging")}
}

func (_c *TelemetryIngress_Logging_Call) Run(run func()) *TelemetryIngress_Logging_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TelemetryIngress_Logging_Call) Return(_a0 bool) *TelemetryIngress_Logging_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TelemetryIngress_Logging_Call) RunAndReturn(run func() bool) *TelemetryIngress_Logging_Call {
	_c.Call.Return(run)
	return _c
}

// MaxBatchSize provides a mock function with given fields:
func (_m *TelemetryIngress) MaxBatchSize() uint {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MaxBatchSize")
	}

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// TelemetryIngress_MaxBatchSize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MaxBatchSize'
type TelemetryIngress_MaxBatchSize_Call struct {
	*mock.Call
}

// MaxBatchSize is a helper method to define mock.On call
func (_e *TelemetryIngress_Expecter) MaxBatchSize() *TelemetryIngress_MaxBatchSize_Call {
	return &TelemetryIngress_MaxBatchSize_Call{Call: _e.mock.On("MaxBatchSize")}
}

func (_c *TelemetryIngress_MaxBatchSize_Call) Run(run func()) *TelemetryIngress_MaxBatchSize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TelemetryIngress_MaxBatchSize_Call) Return(_a0 uint) *TelemetryIngress_MaxBatchSize_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TelemetryIngress_MaxBatchSize_Call) RunAndReturn(run func() uint) *TelemetryIngress_MaxBatchSize_Call {
	_c.Call.Return(run)
	return _c
}

// SendInterval provides a mock function with given fields:
func (_m *TelemetryIngress) SendInterval() time.Duration {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SendInterval")
	}

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// TelemetryIngress_SendInterval_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendInterval'
type TelemetryIngress_SendInterval_Call struct {
	*mock.Call
}

// SendInterval is a helper method to define mock.On call
func (_e *TelemetryIngress_Expecter) SendInterval() *TelemetryIngress_SendInterval_Call {
	return &TelemetryIngress_SendInterval_Call{Call: _e.mock.On("SendInterval")}
}

func (_c *TelemetryIngress_SendInterval_Call) Run(run func()) *TelemetryIngress_SendInterval_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TelemetryIngress_SendInterval_Call) Return(_a0 time.Duration) *TelemetryIngress_SendInterval_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TelemetryIngress_SendInterval_Call) RunAndReturn(run func() time.Duration) *TelemetryIngress_SendInterval_Call {
	_c.Call.Return(run)
	return _c
}

// SendTimeout provides a mock function with given fields:
func (_m *TelemetryIngress) SendTimeout() time.Duration {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SendTimeout")
	}

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// TelemetryIngress_SendTimeout_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendTimeout'
type TelemetryIngress_SendTimeout_Call struct {
	*mock.Call
}

// SendTimeout is a helper method to define mock.On call
func (_e *TelemetryIngress_Expecter) SendTimeout() *TelemetryIngress_SendTimeout_Call {
	return &TelemetryIngress_SendTimeout_Call{Call: _e.mock.On("SendTimeout")}
}

func (_c *TelemetryIngress_SendTimeout_Call) Run(run func()) *TelemetryIngress_SendTimeout_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TelemetryIngress_SendTimeout_Call) Return(_a0 time.Duration) *TelemetryIngress_SendTimeout_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TelemetryIngress_SendTimeout_Call) RunAndReturn(run func() time.Duration) *TelemetryIngress_SendTimeout_Call {
	_c.Call.Return(run)
	return _c
}

// UniConn provides a mock function with given fields:
func (_m *TelemetryIngress) UniConn() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for UniConn")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// TelemetryIngress_UniConn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UniConn'
type TelemetryIngress_UniConn_Call struct {
	*mock.Call
}

// UniConn is a helper method to define mock.On call
func (_e *TelemetryIngress_Expecter) UniConn() *TelemetryIngress_UniConn_Call {
	return &TelemetryIngress_UniConn_Call{Call: _e.mock.On("UniConn")}
}

func (_c *TelemetryIngress_UniConn_Call) Run(run func()) *TelemetryIngress_UniConn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TelemetryIngress_UniConn_Call) Return(_a0 bool) *TelemetryIngress_UniConn_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TelemetryIngress_UniConn_Call) RunAndReturn(run func() bool) *TelemetryIngress_UniConn_Call {
	_c.Call.Return(run)
	return _c
}

// UseBatchSend provides a mock function with given fields:
func (_m *TelemetryIngress) UseBatchSend() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for UseBatchSend")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// TelemetryIngress_UseBatchSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UseBatchSend'
type TelemetryIngress_UseBatchSend_Call struct {
	*mock.Call
}

// UseBatchSend is a helper method to define mock.On call
func (_e *TelemetryIngress_Expecter) UseBatchSend() *TelemetryIngress_UseBatchSend_Call {
	return &TelemetryIngress_UseBatchSend_Call{Call: _e.mock.On("UseBatchSend")}
}

func (_c *TelemetryIngress_UseBatchSend_Call) Run(run func()) *TelemetryIngress_UseBatchSend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TelemetryIngress_UseBatchSend_Call) Return(_a0 bool) *TelemetryIngress_UseBatchSend_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TelemetryIngress_UseBatchSend_Call) RunAndReturn(run func() bool) *TelemetryIngress_UseBatchSend_Call {
	_c.Call.Return(run)
	return _c
}

// NewTelemetryIngress creates a new instance of TelemetryIngress. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTelemetryIngress(t interface {
	mock.TestingT
	Cleanup(func())
}) *TelemetryIngress {
	mock := &TelemetryIngress{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}