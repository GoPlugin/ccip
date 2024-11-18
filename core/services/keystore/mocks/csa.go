// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	csakey "github.com/goplugin/pluginv3.0/v2/core/services/keystore/keys/csakey"

	mock "github.com/stretchr/testify/mock"
)

// CSA is an autogenerated mock type for the CSA type
type CSA struct {
	mock.Mock
}

type CSA_Expecter struct {
	mock *mock.Mock
}

func (_m *CSA) EXPECT() *CSA_Expecter {
	return &CSA_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: ctx, key
func (_m *CSA) Add(ctx context.Context, key csakey.KeyV2) error {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, csakey.KeyV2) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CSA_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type CSA_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - ctx context.Context
//   - key csakey.KeyV2
func (_e *CSA_Expecter) Add(ctx interface{}, key interface{}) *CSA_Add_Call {
	return &CSA_Add_Call{Call: _e.mock.On("Add", ctx, key)}
}

func (_c *CSA_Add_Call) Run(run func(ctx context.Context, key csakey.KeyV2)) *CSA_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(csakey.KeyV2))
	})
	return _c
}

func (_c *CSA_Add_Call) Return(_a0 error) *CSA_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CSA_Add_Call) RunAndReturn(run func(context.Context, csakey.KeyV2) error) *CSA_Add_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx
func (_m *CSA) Create(ctx context.Context) (csakey.KeyV2, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 csakey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (csakey.KeyV2, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) csakey.KeyV2); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(csakey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CSA_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type CSA_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
func (_e *CSA_Expecter) Create(ctx interface{}) *CSA_Create_Call {
	return &CSA_Create_Call{Call: _e.mock.On("Create", ctx)}
}

func (_c *CSA_Create_Call) Run(run func(ctx context.Context)) *CSA_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *CSA_Create_Call) Return(_a0 csakey.KeyV2, _a1 error) *CSA_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CSA_Create_Call) RunAndReturn(run func(context.Context) (csakey.KeyV2, error)) *CSA_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *CSA) Delete(ctx context.Context, id string) (csakey.KeyV2, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 csakey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (csakey.KeyV2, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) csakey.KeyV2); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(csakey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CSA_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type CSA_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *CSA_Expecter) Delete(ctx interface{}, id interface{}) *CSA_Delete_Call {
	return &CSA_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *CSA_Delete_Call) Run(run func(ctx context.Context, id string)) *CSA_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *CSA_Delete_Call) Return(_a0 csakey.KeyV2, _a1 error) *CSA_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CSA_Delete_Call) RunAndReturn(run func(context.Context, string) (csakey.KeyV2, error)) *CSA_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// EnsureKey provides a mock function with given fields: ctx
func (_m *CSA) EnsureKey(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for EnsureKey")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CSA_EnsureKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnsureKey'
type CSA_EnsureKey_Call struct {
	*mock.Call
}

// EnsureKey is a helper method to define mock.On call
//   - ctx context.Context
func (_e *CSA_Expecter) EnsureKey(ctx interface{}) *CSA_EnsureKey_Call {
	return &CSA_EnsureKey_Call{Call: _e.mock.On("EnsureKey", ctx)}
}

func (_c *CSA_EnsureKey_Call) Run(run func(ctx context.Context)) *CSA_EnsureKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *CSA_EnsureKey_Call) Return(_a0 error) *CSA_EnsureKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CSA_EnsureKey_Call) RunAndReturn(run func(context.Context) error) *CSA_EnsureKey_Call {
	_c.Call.Return(run)
	return _c
}

// Export provides a mock function with given fields: id, password
func (_m *CSA) Export(id string, password string) ([]byte, error) {
	ret := _m.Called(id, password)

	if len(ret) == 0 {
		panic("no return value specified for Export")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]byte, error)); ok {
		return rf(id, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(id, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CSA_Export_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Export'
type CSA_Export_Call struct {
	*mock.Call
}

// Export is a helper method to define mock.On call
//   - id string
//   - password string
func (_e *CSA_Expecter) Export(id interface{}, password interface{}) *CSA_Export_Call {
	return &CSA_Export_Call{Call: _e.mock.On("Export", id, password)}
}

func (_c *CSA_Export_Call) Run(run func(id string, password string)) *CSA_Export_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *CSA_Export_Call) Return(_a0 []byte, _a1 error) *CSA_Export_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CSA_Export_Call) RunAndReturn(run func(string, string) ([]byte, error)) *CSA_Export_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: id
func (_m *CSA) Get(id string) (csakey.KeyV2, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 csakey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (csakey.KeyV2, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) csakey.KeyV2); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(csakey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CSA_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type CSA_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - id string
func (_e *CSA_Expecter) Get(id interface{}) *CSA_Get_Call {
	return &CSA_Get_Call{Call: _e.mock.On("Get", id)}
}

func (_c *CSA_Get_Call) Run(run func(id string)) *CSA_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *CSA_Get_Call) Return(_a0 csakey.KeyV2, _a1 error) *CSA_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CSA_Get_Call) RunAndReturn(run func(string) (csakey.KeyV2, error)) *CSA_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields:
func (_m *CSA) GetAll() ([]csakey.KeyV2, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []csakey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]csakey.KeyV2, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []csakey.KeyV2); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]csakey.KeyV2)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CSA_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type CSA_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
func (_e *CSA_Expecter) GetAll() *CSA_GetAll_Call {
	return &CSA_GetAll_Call{Call: _e.mock.On("GetAll")}
}

func (_c *CSA_GetAll_Call) Run(run func()) *CSA_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CSA_GetAll_Call) Return(_a0 []csakey.KeyV2, _a1 error) *CSA_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CSA_GetAll_Call) RunAndReturn(run func() ([]csakey.KeyV2, error)) *CSA_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// Import provides a mock function with given fields: ctx, keyJSON, password
func (_m *CSA) Import(ctx context.Context, keyJSON []byte, password string) (csakey.KeyV2, error) {
	ret := _m.Called(ctx, keyJSON, password)

	if len(ret) == 0 {
		panic("no return value specified for Import")
	}

	var r0 csakey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string) (csakey.KeyV2, error)); ok {
		return rf(ctx, keyJSON, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string) csakey.KeyV2); ok {
		r0 = rf(ctx, keyJSON, password)
	} else {
		r0 = ret.Get(0).(csakey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte, string) error); ok {
		r1 = rf(ctx, keyJSON, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CSA_Import_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Import'
type CSA_Import_Call struct {
	*mock.Call
}

// Import is a helper method to define mock.On call
//   - ctx context.Context
//   - keyJSON []byte
//   - password string
func (_e *CSA_Expecter) Import(ctx interface{}, keyJSON interface{}, password interface{}) *CSA_Import_Call {
	return &CSA_Import_Call{Call: _e.mock.On("Import", ctx, keyJSON, password)}
}

func (_c *CSA_Import_Call) Run(run func(ctx context.Context, keyJSON []byte, password string)) *CSA_Import_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]byte), args[2].(string))
	})
	return _c
}

func (_c *CSA_Import_Call) Return(_a0 csakey.KeyV2, _a1 error) *CSA_Import_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CSA_Import_Call) RunAndReturn(run func(context.Context, []byte, string) (csakey.KeyV2, error)) *CSA_Import_Call {
	_c.Call.Return(run)
	return _c
}

// NewCSA creates a new instance of CSA. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCSA(t interface {
	mock.TestingT
	Cleanup(func())
}) *CSA {
	mock := &CSA{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
