// Code generated by mockery v2.43.2. DO NOT EDIT.

package prices

import (
	context "context"
	big "math/big"

	ccip "github.com/goplugin/plugin-common/pkg/types/ccip"

	mock "github.com/stretchr/testify/mock"
)

// MockGasPriceEstimator is an autogenerated mock type for the GasPriceEstimator type
type MockGasPriceEstimator struct {
	mock.Mock
}

type MockGasPriceEstimator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGasPriceEstimator) EXPECT() *MockGasPriceEstimator_Expecter {
	return &MockGasPriceEstimator_Expecter{mock: &_m.Mock}
}

// DenoteInUSD provides a mock function with given fields: p, wrappedNativePrice
func (_m *MockGasPriceEstimator) DenoteInUSD(p *big.Int, wrappedNativePrice *big.Int) (*big.Int, error) {
	ret := _m.Called(p, wrappedNativePrice)

	if len(ret) == 0 {
		panic("no return value specified for DenoteInUSD")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int) (*big.Int, error)); ok {
		return rf(p, wrappedNativePrice)
	}
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int) *big.Int); ok {
		r0 = rf(p, wrappedNativePrice)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(*big.Int, *big.Int) error); ok {
		r1 = rf(p, wrappedNativePrice)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGasPriceEstimator_DenoteInUSD_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DenoteInUSD'
type MockGasPriceEstimator_DenoteInUSD_Call struct {
	*mock.Call
}

// DenoteInUSD is a helper method to define mock.On call
//   - p *big.Int
//   - wrappedNativePrice *big.Int
func (_e *MockGasPriceEstimator_Expecter) DenoteInUSD(p interface{}, wrappedNativePrice interface{}) *MockGasPriceEstimator_DenoteInUSD_Call {
	return &MockGasPriceEstimator_DenoteInUSD_Call{Call: _e.mock.On("DenoteInUSD", p, wrappedNativePrice)}
}

func (_c *MockGasPriceEstimator_DenoteInUSD_Call) Run(run func(p *big.Int, wrappedNativePrice *big.Int)) *MockGasPriceEstimator_DenoteInUSD_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*big.Int), args[1].(*big.Int))
	})
	return _c
}

func (_c *MockGasPriceEstimator_DenoteInUSD_Call) Return(_a0 *big.Int, _a1 error) *MockGasPriceEstimator_DenoteInUSD_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGasPriceEstimator_DenoteInUSD_Call) RunAndReturn(run func(*big.Int, *big.Int) (*big.Int, error)) *MockGasPriceEstimator_DenoteInUSD_Call {
	_c.Call.Return(run)
	return _c
}

// Deviates provides a mock function with given fields: p1, p2
func (_m *MockGasPriceEstimator) Deviates(p1 *big.Int, p2 *big.Int) (bool, error) {
	ret := _m.Called(p1, p2)

	if len(ret) == 0 {
		panic("no return value specified for Deviates")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int) (bool, error)); ok {
		return rf(p1, p2)
	}
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int) bool); ok {
		r0 = rf(p1, p2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(*big.Int, *big.Int) error); ok {
		r1 = rf(p1, p2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGasPriceEstimator_Deviates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Deviates'
type MockGasPriceEstimator_Deviates_Call struct {
	*mock.Call
}

// Deviates is a helper method to define mock.On call
//   - p1 *big.Int
//   - p2 *big.Int
func (_e *MockGasPriceEstimator_Expecter) Deviates(p1 interface{}, p2 interface{}) *MockGasPriceEstimator_Deviates_Call {
	return &MockGasPriceEstimator_Deviates_Call{Call: _e.mock.On("Deviates", p1, p2)}
}

func (_c *MockGasPriceEstimator_Deviates_Call) Run(run func(p1 *big.Int, p2 *big.Int)) *MockGasPriceEstimator_Deviates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*big.Int), args[1].(*big.Int))
	})
	return _c
}

func (_c *MockGasPriceEstimator_Deviates_Call) Return(_a0 bool, _a1 error) *MockGasPriceEstimator_Deviates_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGasPriceEstimator_Deviates_Call) RunAndReturn(run func(*big.Int, *big.Int) (bool, error)) *MockGasPriceEstimator_Deviates_Call {
	_c.Call.Return(run)
	return _c
}

// EstimateMsgCostUSD provides a mock function with given fields: p, wrappedNativePrice, msg
func (_m *MockGasPriceEstimator) EstimateMsgCostUSD(p *big.Int, wrappedNativePrice *big.Int, msg ccip.EVM2EVMOnRampCCIPSendRequestedWithMeta) (*big.Int, error) {
	ret := _m.Called(p, wrappedNativePrice, msg)

	if len(ret) == 0 {
		panic("no return value specified for EstimateMsgCostUSD")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int, ccip.EVM2EVMOnRampCCIPSendRequestedWithMeta) (*big.Int, error)); ok {
		return rf(p, wrappedNativePrice, msg)
	}
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int, ccip.EVM2EVMOnRampCCIPSendRequestedWithMeta) *big.Int); ok {
		r0 = rf(p, wrappedNativePrice, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(*big.Int, *big.Int, ccip.EVM2EVMOnRampCCIPSendRequestedWithMeta) error); ok {
		r1 = rf(p, wrappedNativePrice, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGasPriceEstimator_EstimateMsgCostUSD_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EstimateMsgCostUSD'
type MockGasPriceEstimator_EstimateMsgCostUSD_Call struct {
	*mock.Call
}

// EstimateMsgCostUSD is a helper method to define mock.On call
//   - p *big.Int
//   - wrappedNativePrice *big.Int
//   - msg ccip.EVM2EVMOnRampCCIPSendRequestedWithMeta
func (_e *MockGasPriceEstimator_Expecter) EstimateMsgCostUSD(p interface{}, wrappedNativePrice interface{}, msg interface{}) *MockGasPriceEstimator_EstimateMsgCostUSD_Call {
	return &MockGasPriceEstimator_EstimateMsgCostUSD_Call{Call: _e.mock.On("EstimateMsgCostUSD", p, wrappedNativePrice, msg)}
}

func (_c *MockGasPriceEstimator_EstimateMsgCostUSD_Call) Run(run func(p *big.Int, wrappedNativePrice *big.Int, msg ccip.EVM2EVMOnRampCCIPSendRequestedWithMeta)) *MockGasPriceEstimator_EstimateMsgCostUSD_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*big.Int), args[1].(*big.Int), args[2].(ccip.EVM2EVMOnRampCCIPSendRequestedWithMeta))
	})
	return _c
}

func (_c *MockGasPriceEstimator_EstimateMsgCostUSD_Call) Return(_a0 *big.Int, _a1 error) *MockGasPriceEstimator_EstimateMsgCostUSD_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGasPriceEstimator_EstimateMsgCostUSD_Call) RunAndReturn(run func(*big.Int, *big.Int, ccip.EVM2EVMOnRampCCIPSendRequestedWithMeta) (*big.Int, error)) *MockGasPriceEstimator_EstimateMsgCostUSD_Call {
	_c.Call.Return(run)
	return _c
}

// GetGasPrice provides a mock function with given fields: ctx
func (_m *MockGasPriceEstimator) GetGasPrice(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetGasPrice")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*big.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGasPriceEstimator_GetGasPrice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGasPrice'
type MockGasPriceEstimator_GetGasPrice_Call struct {
	*mock.Call
}

// GetGasPrice is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockGasPriceEstimator_Expecter) GetGasPrice(ctx interface{}) *MockGasPriceEstimator_GetGasPrice_Call {
	return &MockGasPriceEstimator_GetGasPrice_Call{Call: _e.mock.On("GetGasPrice", ctx)}
}

func (_c *MockGasPriceEstimator_GetGasPrice_Call) Run(run func(ctx context.Context)) *MockGasPriceEstimator_GetGasPrice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockGasPriceEstimator_GetGasPrice_Call) Return(_a0 *big.Int, _a1 error) *MockGasPriceEstimator_GetGasPrice_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGasPriceEstimator_GetGasPrice_Call) RunAndReturn(run func(context.Context) (*big.Int, error)) *MockGasPriceEstimator_GetGasPrice_Call {
	_c.Call.Return(run)
	return _c
}

// Median provides a mock function with given fields: gasPrices
func (_m *MockGasPriceEstimator) Median(gasPrices []*big.Int) (*big.Int, error) {
	ret := _m.Called(gasPrices)

	if len(ret) == 0 {
		panic("no return value specified for Median")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func([]*big.Int) (*big.Int, error)); ok {
		return rf(gasPrices)
	}
	if rf, ok := ret.Get(0).(func([]*big.Int) *big.Int); ok {
		r0 = rf(gasPrices)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func([]*big.Int) error); ok {
		r1 = rf(gasPrices)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGasPriceEstimator_Median_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Median'
type MockGasPriceEstimator_Median_Call struct {
	*mock.Call
}

// Median is a helper method to define mock.On call
//   - gasPrices []*big.Int
func (_e *MockGasPriceEstimator_Expecter) Median(gasPrices interface{}) *MockGasPriceEstimator_Median_Call {
	return &MockGasPriceEstimator_Median_Call{Call: _e.mock.On("Median", gasPrices)}
}

func (_c *MockGasPriceEstimator_Median_Call) Run(run func(gasPrices []*big.Int)) *MockGasPriceEstimator_Median_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]*big.Int))
	})
	return _c
}

func (_c *MockGasPriceEstimator_Median_Call) Return(_a0 *big.Int, _a1 error) *MockGasPriceEstimator_Median_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGasPriceEstimator_Median_Call) RunAndReturn(run func([]*big.Int) (*big.Int, error)) *MockGasPriceEstimator_Median_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGasPriceEstimator creates a new instance of MockGasPriceEstimator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGasPriceEstimator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGasPriceEstimator {
	mock := &MockGasPriceEstimator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
