// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	config "github.com/goplugin/plugin-common/pkg/config"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Config is an autogenerated mock type for the Config type
type Config struct {
	mock.Mock
}

type Config_Expecter struct {
	mock *mock.Mock
}

func (_m *Config) EXPECT() *Config_Expecter {
	return &Config_Expecter{mock: &_m.Mock}
}

// DefaultHTTPLimit provides a mock function with given fields:
func (_m *Config) DefaultHTTPLimit() int64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DefaultHTTPLimit")
	}

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Config_DefaultHTTPLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DefaultHTTPLimit'
type Config_DefaultHTTPLimit_Call struct {
	*mock.Call
}

// DefaultHTTPLimit is a helper method to define mock.On call
func (_e *Config_Expecter) DefaultHTTPLimit() *Config_DefaultHTTPLimit_Call {
	return &Config_DefaultHTTPLimit_Call{Call: _e.mock.On("DefaultHTTPLimit")}
}

func (_c *Config_DefaultHTTPLimit_Call) Run(run func()) *Config_DefaultHTTPLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Config_DefaultHTTPLimit_Call) Return(_a0 int64) *Config_DefaultHTTPLimit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Config_DefaultHTTPLimit_Call) RunAndReturn(run func() int64) *Config_DefaultHTTPLimit_Call {
	_c.Call.Return(run)
	return _c
}

// DefaultHTTPTimeout provides a mock function with given fields:
func (_m *Config) DefaultHTTPTimeout() config.Duration {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DefaultHTTPTimeout")
	}

	var r0 config.Duration
	if rf, ok := ret.Get(0).(func() config.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(config.Duration)
	}

	return r0
}

// Config_DefaultHTTPTimeout_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DefaultHTTPTimeout'
type Config_DefaultHTTPTimeout_Call struct {
	*mock.Call
}

// DefaultHTTPTimeout is a helper method to define mock.On call
func (_e *Config_Expecter) DefaultHTTPTimeout() *Config_DefaultHTTPTimeout_Call {
	return &Config_DefaultHTTPTimeout_Call{Call: _e.mock.On("DefaultHTTPTimeout")}
}

func (_c *Config_DefaultHTTPTimeout_Call) Run(run func()) *Config_DefaultHTTPTimeout_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Config_DefaultHTTPTimeout_Call) Return(_a0 config.Duration) *Config_DefaultHTTPTimeout_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Config_DefaultHTTPTimeout_Call) RunAndReturn(run func() config.Duration) *Config_DefaultHTTPTimeout_Call {
	_c.Call.Return(run)
	return _c
}

// MaxRunDuration provides a mock function with given fields:
func (_m *Config) MaxRunDuration() time.Duration {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MaxRunDuration")
	}

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// Config_MaxRunDuration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MaxRunDuration'
type Config_MaxRunDuration_Call struct {
	*mock.Call
}

// MaxRunDuration is a helper method to define mock.On call
func (_e *Config_Expecter) MaxRunDuration() *Config_MaxRunDuration_Call {
	return &Config_MaxRunDuration_Call{Call: _e.mock.On("MaxRunDuration")}
}

func (_c *Config_MaxRunDuration_Call) Run(run func()) *Config_MaxRunDuration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Config_MaxRunDuration_Call) Return(_a0 time.Duration) *Config_MaxRunDuration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Config_MaxRunDuration_Call) RunAndReturn(run func() time.Duration) *Config_MaxRunDuration_Call {
	_c.Call.Return(run)
	return _c
}

// ReaperInterval provides a mock function with given fields:
func (_m *Config) ReaperInterval() time.Duration {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReaperInterval")
	}

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// Config_ReaperInterval_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReaperInterval'
type Config_ReaperInterval_Call struct {
	*mock.Call
}

// ReaperInterval is a helper method to define mock.On call
func (_e *Config_Expecter) ReaperInterval() *Config_ReaperInterval_Call {
	return &Config_ReaperInterval_Call{Call: _e.mock.On("ReaperInterval")}
}

func (_c *Config_ReaperInterval_Call) Run(run func()) *Config_ReaperInterval_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Config_ReaperInterval_Call) Return(_a0 time.Duration) *Config_ReaperInterval_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Config_ReaperInterval_Call) RunAndReturn(run func() time.Duration) *Config_ReaperInterval_Call {
	_c.Call.Return(run)
	return _c
}

// ReaperThreshold provides a mock function with given fields:
func (_m *Config) ReaperThreshold() time.Duration {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReaperThreshold")
	}

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// Config_ReaperThreshold_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReaperThreshold'
type Config_ReaperThreshold_Call struct {
	*mock.Call
}

// ReaperThreshold is a helper method to define mock.On call
func (_e *Config_Expecter) ReaperThreshold() *Config_ReaperThreshold_Call {
	return &Config_ReaperThreshold_Call{Call: _e.mock.On("ReaperThreshold")}
}

func (_c *Config_ReaperThreshold_Call) Run(run func()) *Config_ReaperThreshold_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Config_ReaperThreshold_Call) Return(_a0 time.Duration) *Config_ReaperThreshold_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Config_ReaperThreshold_Call) RunAndReturn(run func() time.Duration) *Config_ReaperThreshold_Call {
	_c.Call.Return(run)
	return _c
}

// VerboseLogging provides a mock function with given fields:
func (_m *Config) VerboseLogging() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for VerboseLogging")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Config_VerboseLogging_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerboseLogging'
type Config_VerboseLogging_Call struct {
	*mock.Call
}

// VerboseLogging is a helper method to define mock.On call
func (_e *Config_Expecter) VerboseLogging() *Config_VerboseLogging_Call {
	return &Config_VerboseLogging_Call{Call: _e.mock.On("VerboseLogging")}
}

func (_c *Config_VerboseLogging_Call) Run(run func()) *Config_VerboseLogging_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Config_VerboseLogging_Call) Return(_a0 bool) *Config_VerboseLogging_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Config_VerboseLogging_Call) RunAndReturn(run func() bool) *Config_VerboseLogging_Call {
	_c.Call.Return(run)
	return _c
}

// NewConfig creates a new instance of Config. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConfig(t interface {
	mock.TestingT
	Cleanup(func())
}) *Config {
	mock := &Config{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
