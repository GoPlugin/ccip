// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	proto "github.com/goplugin/pluginv3.0/v2/core/services/feeds/proto"
	mock "github.com/stretchr/testify/mock"
)

// FeedsManagerClient is an autogenerated mock type for the FeedsManagerClient type
type FeedsManagerClient struct {
	mock.Mock
}

type FeedsManagerClient_Expecter struct {
	mock *mock.Mock
}

func (_m *FeedsManagerClient) EXPECT() *FeedsManagerClient_Expecter {
	return &FeedsManagerClient_Expecter{mock: &_m.Mock}
}

// ApprovedJob provides a mock function with given fields: ctx, in
func (_m *FeedsManagerClient) ApprovedJob(ctx context.Context, in *proto.ApprovedJobRequest) (*proto.ApprovedJobResponse, error) {
	ret := _m.Called(ctx, in)

	if len(ret) == 0 {
		panic("no return value specified for ApprovedJob")
	}

	var r0 *proto.ApprovedJobResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ApprovedJobRequest) (*proto.ApprovedJobResponse, error)); ok {
		return rf(ctx, in)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ApprovedJobRequest) *proto.ApprovedJobResponse); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.ApprovedJobResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.ApprovedJobRequest) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FeedsManagerClient_ApprovedJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApprovedJob'
type FeedsManagerClient_ApprovedJob_Call struct {
	*mock.Call
}

// ApprovedJob is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.ApprovedJobRequest
func (_e *FeedsManagerClient_Expecter) ApprovedJob(ctx interface{}, in interface{}) *FeedsManagerClient_ApprovedJob_Call {
	return &FeedsManagerClient_ApprovedJob_Call{Call: _e.mock.On("ApprovedJob", ctx, in)}
}

func (_c *FeedsManagerClient_ApprovedJob_Call) Run(run func(ctx context.Context, in *proto.ApprovedJobRequest)) *FeedsManagerClient_ApprovedJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.ApprovedJobRequest))
	})
	return _c
}

func (_c *FeedsManagerClient_ApprovedJob_Call) Return(_a0 *proto.ApprovedJobResponse, _a1 error) *FeedsManagerClient_ApprovedJob_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FeedsManagerClient_ApprovedJob_Call) RunAndReturn(run func(context.Context, *proto.ApprovedJobRequest) (*proto.ApprovedJobResponse, error)) *FeedsManagerClient_ApprovedJob_Call {
	_c.Call.Return(run)
	return _c
}

// CancelledJob provides a mock function with given fields: ctx, in
func (_m *FeedsManagerClient) CancelledJob(ctx context.Context, in *proto.CancelledJobRequest) (*proto.CancelledJobResponse, error) {
	ret := _m.Called(ctx, in)

	if len(ret) == 0 {
		panic("no return value specified for CancelledJob")
	}

	var r0 *proto.CancelledJobResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.CancelledJobRequest) (*proto.CancelledJobResponse, error)); ok {
		return rf(ctx, in)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.CancelledJobRequest) *proto.CancelledJobResponse); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.CancelledJobResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.CancelledJobRequest) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FeedsManagerClient_CancelledJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CancelledJob'
type FeedsManagerClient_CancelledJob_Call struct {
	*mock.Call
}

// CancelledJob is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.CancelledJobRequest
func (_e *FeedsManagerClient_Expecter) CancelledJob(ctx interface{}, in interface{}) *FeedsManagerClient_CancelledJob_Call {
	return &FeedsManagerClient_CancelledJob_Call{Call: _e.mock.On("CancelledJob", ctx, in)}
}

func (_c *FeedsManagerClient_CancelledJob_Call) Run(run func(ctx context.Context, in *proto.CancelledJobRequest)) *FeedsManagerClient_CancelledJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.CancelledJobRequest))
	})
	return _c
}

func (_c *FeedsManagerClient_CancelledJob_Call) Return(_a0 *proto.CancelledJobResponse, _a1 error) *FeedsManagerClient_CancelledJob_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FeedsManagerClient_CancelledJob_Call) RunAndReturn(run func(context.Context, *proto.CancelledJobRequest) (*proto.CancelledJobResponse, error)) *FeedsManagerClient_CancelledJob_Call {
	_c.Call.Return(run)
	return _c
}

// Healthcheck provides a mock function with given fields: ctx, in
func (_m *FeedsManagerClient) Healthcheck(ctx context.Context, in *proto.HealthcheckRequest) (*proto.HealthcheckResponse, error) {
	ret := _m.Called(ctx, in)

	if len(ret) == 0 {
		panic("no return value specified for Healthcheck")
	}

	var r0 *proto.HealthcheckResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.HealthcheckRequest) (*proto.HealthcheckResponse, error)); ok {
		return rf(ctx, in)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.HealthcheckRequest) *proto.HealthcheckResponse); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.HealthcheckResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.HealthcheckRequest) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FeedsManagerClient_Healthcheck_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Healthcheck'
type FeedsManagerClient_Healthcheck_Call struct {
	*mock.Call
}

// Healthcheck is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.HealthcheckRequest
func (_e *FeedsManagerClient_Expecter) Healthcheck(ctx interface{}, in interface{}) *FeedsManagerClient_Healthcheck_Call {
	return &FeedsManagerClient_Healthcheck_Call{Call: _e.mock.On("Healthcheck", ctx, in)}
}

func (_c *FeedsManagerClient_Healthcheck_Call) Run(run func(ctx context.Context, in *proto.HealthcheckRequest)) *FeedsManagerClient_Healthcheck_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.HealthcheckRequest))
	})
	return _c
}

func (_c *FeedsManagerClient_Healthcheck_Call) Return(_a0 *proto.HealthcheckResponse, _a1 error) *FeedsManagerClient_Healthcheck_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FeedsManagerClient_Healthcheck_Call) RunAndReturn(run func(context.Context, *proto.HealthcheckRequest) (*proto.HealthcheckResponse, error)) *FeedsManagerClient_Healthcheck_Call {
	_c.Call.Return(run)
	return _c
}

// RejectedJob provides a mock function with given fields: ctx, in
func (_m *FeedsManagerClient) RejectedJob(ctx context.Context, in *proto.RejectedJobRequest) (*proto.RejectedJobResponse, error) {
	ret := _m.Called(ctx, in)

	if len(ret) == 0 {
		panic("no return value specified for RejectedJob")
	}

	var r0 *proto.RejectedJobResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.RejectedJobRequest) (*proto.RejectedJobResponse, error)); ok {
		return rf(ctx, in)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.RejectedJobRequest) *proto.RejectedJobResponse); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.RejectedJobResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.RejectedJobRequest) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FeedsManagerClient_RejectedJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RejectedJob'
type FeedsManagerClient_RejectedJob_Call struct {
	*mock.Call
}

// RejectedJob is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.RejectedJobRequest
func (_e *FeedsManagerClient_Expecter) RejectedJob(ctx interface{}, in interface{}) *FeedsManagerClient_RejectedJob_Call {
	return &FeedsManagerClient_RejectedJob_Call{Call: _e.mock.On("RejectedJob", ctx, in)}
}

func (_c *FeedsManagerClient_RejectedJob_Call) Run(run func(ctx context.Context, in *proto.RejectedJobRequest)) *FeedsManagerClient_RejectedJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.RejectedJobRequest))
	})
	return _c
}

func (_c *FeedsManagerClient_RejectedJob_Call) Return(_a0 *proto.RejectedJobResponse, _a1 error) *FeedsManagerClient_RejectedJob_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FeedsManagerClient_RejectedJob_Call) RunAndReturn(run func(context.Context, *proto.RejectedJobRequest) (*proto.RejectedJobResponse, error)) *FeedsManagerClient_RejectedJob_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateNode provides a mock function with given fields: ctx, in
func (_m *FeedsManagerClient) UpdateNode(ctx context.Context, in *proto.UpdateNodeRequest) (*proto.UpdateNodeResponse, error) {
	ret := _m.Called(ctx, in)

	if len(ret) == 0 {
		panic("no return value specified for UpdateNode")
	}

	var r0 *proto.UpdateNodeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.UpdateNodeRequest) (*proto.UpdateNodeResponse, error)); ok {
		return rf(ctx, in)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.UpdateNodeRequest) *proto.UpdateNodeResponse); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.UpdateNodeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.UpdateNodeRequest) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FeedsManagerClient_UpdateNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateNode'
type FeedsManagerClient_UpdateNode_Call struct {
	*mock.Call
}

// UpdateNode is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.UpdateNodeRequest
func (_e *FeedsManagerClient_Expecter) UpdateNode(ctx interface{}, in interface{}) *FeedsManagerClient_UpdateNode_Call {
	return &FeedsManagerClient_UpdateNode_Call{Call: _e.mock.On("UpdateNode", ctx, in)}
}

func (_c *FeedsManagerClient_UpdateNode_Call) Run(run func(ctx context.Context, in *proto.UpdateNodeRequest)) *FeedsManagerClient_UpdateNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.UpdateNodeRequest))
	})
	return _c
}

func (_c *FeedsManagerClient_UpdateNode_Call) Return(_a0 *proto.UpdateNodeResponse, _a1 error) *FeedsManagerClient_UpdateNode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FeedsManagerClient_UpdateNode_Call) RunAndReturn(run func(context.Context, *proto.UpdateNodeRequest) (*proto.UpdateNodeResponse, error)) *FeedsManagerClient_UpdateNode_Call {
	_c.Call.Return(run)
	return _c
}

// NewFeedsManagerClient creates a new instance of FeedsManagerClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFeedsManagerClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *FeedsManagerClient {
	mock := &FeedsManagerClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}