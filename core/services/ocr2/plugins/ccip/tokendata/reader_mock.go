// Code generated by mockery v2.43.2. DO NOT EDIT.

package tokendata

import (
	context "context"

	ccip "github.com/goplugin/plugin-common/pkg/types/ccip"

	mock "github.com/stretchr/testify/mock"
)

// MockReader is an autogenerated mock type for the Reader type
type MockReader struct {
	mock.Mock
}

type MockReader_Expecter struct {
	mock *mock.Mock
}

func (_m *MockReader) EXPECT() *MockReader_Expecter {
	return &MockReader_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockReader) Close() error {
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

// MockReader_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockReader_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockReader_Expecter) Close() *MockReader_Close_Call {
	return &MockReader_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockReader_Close_Call) Run(run func()) *MockReader_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockReader_Close_Call) Return(_a0 error) *MockReader_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockReader_Close_Call) RunAndReturn(run func() error) *MockReader_Close_Call {
	_c.Call.Return(run)
	return _c
}

// ReadTokenData provides a mock function with given fields: ctx, msg, tokenIndex
func (_m *MockReader) ReadTokenData(ctx context.Context, msg ccip.EVM2EVMOnRampCCIPSendRequestedWithMeta, tokenIndex int) ([]byte, error) {
	ret := _m.Called(ctx, msg, tokenIndex)

	if len(ret) == 0 {
		panic("no return value specified for ReadTokenData")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ccip.EVM2EVMOnRampCCIPSendRequestedWithMeta, int) ([]byte, error)); ok {
		return rf(ctx, msg, tokenIndex)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ccip.EVM2EVMOnRampCCIPSendRequestedWithMeta, int) []byte); ok {
		r0 = rf(ctx, msg, tokenIndex)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ccip.EVM2EVMOnRampCCIPSendRequestedWithMeta, int) error); ok {
		r1 = rf(ctx, msg, tokenIndex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReader_ReadTokenData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadTokenData'
type MockReader_ReadTokenData_Call struct {
	*mock.Call
}

// ReadTokenData is a helper method to define mock.On call
//   - ctx context.Context
//   - msg ccip.EVM2EVMOnRampCCIPSendRequestedWithMeta
//   - tokenIndex int
func (_e *MockReader_Expecter) ReadTokenData(ctx interface{}, msg interface{}, tokenIndex interface{}) *MockReader_ReadTokenData_Call {
	return &MockReader_ReadTokenData_Call{Call: _e.mock.On("ReadTokenData", ctx, msg, tokenIndex)}
}

func (_c *MockReader_ReadTokenData_Call) Run(run func(ctx context.Context, msg ccip.EVM2EVMOnRampCCIPSendRequestedWithMeta, tokenIndex int)) *MockReader_ReadTokenData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ccip.EVM2EVMOnRampCCIPSendRequestedWithMeta), args[2].(int))
	})
	return _c
}

func (_c *MockReader_ReadTokenData_Call) Return(tokenData []byte, err error) *MockReader_ReadTokenData_Call {
	_c.Call.Return(tokenData, err)
	return _c
}

func (_c *MockReader_ReadTokenData_Call) RunAndReturn(run func(context.Context, ccip.EVM2EVMOnRampCCIPSendRequestedWithMeta, int) ([]byte, error)) *MockReader_ReadTokenData_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockReader creates a new instance of MockReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockReader {
	mock := &MockReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}