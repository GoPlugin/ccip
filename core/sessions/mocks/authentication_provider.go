// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	auth "github.com/goplugin/pluginv3.0/v2/core/auth"
	bridges "github.com/goplugin/pluginv3.0/v2/core/bridges"

	context "context"

	mock "github.com/stretchr/testify/mock"

	sessions "github.com/goplugin/pluginv3.0/v2/core/sessions"
)

// AuthenticationProvider is an autogenerated mock type for the AuthenticationProvider type
type AuthenticationProvider struct {
	mock.Mock
}

type AuthenticationProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *AuthenticationProvider) EXPECT() *AuthenticationProvider_Expecter {
	return &AuthenticationProvider_Expecter{mock: &_m.Mock}
}

// AuthorizedUserWithSession provides a mock function with given fields: ctx, sessionID
func (_m *AuthenticationProvider) AuthorizedUserWithSession(ctx context.Context, sessionID string) (sessions.User, error) {
	ret := _m.Called(ctx, sessionID)

	if len(ret) == 0 {
		panic("no return value specified for AuthorizedUserWithSession")
	}

	var r0 sessions.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (sessions.User, error)); ok {
		return rf(ctx, sessionID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) sessions.User); ok {
		r0 = rf(ctx, sessionID)
	} else {
		r0 = ret.Get(0).(sessions.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthenticationProvider_AuthorizedUserWithSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AuthorizedUserWithSession'
type AuthenticationProvider_AuthorizedUserWithSession_Call struct {
	*mock.Call
}

// AuthorizedUserWithSession is a helper method to define mock.On call
//   - ctx context.Context
//   - sessionID string
func (_e *AuthenticationProvider_Expecter) AuthorizedUserWithSession(ctx interface{}, sessionID interface{}) *AuthenticationProvider_AuthorizedUserWithSession_Call {
	return &AuthenticationProvider_AuthorizedUserWithSession_Call{Call: _e.mock.On("AuthorizedUserWithSession", ctx, sessionID)}
}

func (_c *AuthenticationProvider_AuthorizedUserWithSession_Call) Run(run func(ctx context.Context, sessionID string)) *AuthenticationProvider_AuthorizedUserWithSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AuthenticationProvider_AuthorizedUserWithSession_Call) Return(_a0 sessions.User, _a1 error) *AuthenticationProvider_AuthorizedUserWithSession_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthenticationProvider_AuthorizedUserWithSession_Call) RunAndReturn(run func(context.Context, string) (sessions.User, error)) *AuthenticationProvider_AuthorizedUserWithSession_Call {
	_c.Call.Return(run)
	return _c
}

// ClearNonCurrentSessions provides a mock function with given fields: ctx, sessionID
func (_m *AuthenticationProvider) ClearNonCurrentSessions(ctx context.Context, sessionID string) error {
	ret := _m.Called(ctx, sessionID)

	if len(ret) == 0 {
		panic("no return value specified for ClearNonCurrentSessions")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, sessionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthenticationProvider_ClearNonCurrentSessions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClearNonCurrentSessions'
type AuthenticationProvider_ClearNonCurrentSessions_Call struct {
	*mock.Call
}

// ClearNonCurrentSessions is a helper method to define mock.On call
//   - ctx context.Context
//   - sessionID string
func (_e *AuthenticationProvider_Expecter) ClearNonCurrentSessions(ctx interface{}, sessionID interface{}) *AuthenticationProvider_ClearNonCurrentSessions_Call {
	return &AuthenticationProvider_ClearNonCurrentSessions_Call{Call: _e.mock.On("ClearNonCurrentSessions", ctx, sessionID)}
}

func (_c *AuthenticationProvider_ClearNonCurrentSessions_Call) Run(run func(ctx context.Context, sessionID string)) *AuthenticationProvider_ClearNonCurrentSessions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AuthenticationProvider_ClearNonCurrentSessions_Call) Return(_a0 error) *AuthenticationProvider_ClearNonCurrentSessions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthenticationProvider_ClearNonCurrentSessions_Call) RunAndReturn(run func(context.Context, string) error) *AuthenticationProvider_ClearNonCurrentSessions_Call {
	_c.Call.Return(run)
	return _c
}

// CreateAndSetAuthToken provides a mock function with given fields: ctx, user
func (_m *AuthenticationProvider) CreateAndSetAuthToken(ctx context.Context, user *sessions.User) (*auth.Token, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for CreateAndSetAuthToken")
	}

	var r0 *auth.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *sessions.User) (*auth.Token, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *sessions.User) *auth.Token); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *sessions.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthenticationProvider_CreateAndSetAuthToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAndSetAuthToken'
type AuthenticationProvider_CreateAndSetAuthToken_Call struct {
	*mock.Call
}

// CreateAndSetAuthToken is a helper method to define mock.On call
//   - ctx context.Context
//   - user *sessions.User
func (_e *AuthenticationProvider_Expecter) CreateAndSetAuthToken(ctx interface{}, user interface{}) *AuthenticationProvider_CreateAndSetAuthToken_Call {
	return &AuthenticationProvider_CreateAndSetAuthToken_Call{Call: _e.mock.On("CreateAndSetAuthToken", ctx, user)}
}

func (_c *AuthenticationProvider_CreateAndSetAuthToken_Call) Run(run func(ctx context.Context, user *sessions.User)) *AuthenticationProvider_CreateAndSetAuthToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*sessions.User))
	})
	return _c
}

func (_c *AuthenticationProvider_CreateAndSetAuthToken_Call) Return(_a0 *auth.Token, _a1 error) *AuthenticationProvider_CreateAndSetAuthToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthenticationProvider_CreateAndSetAuthToken_Call) RunAndReturn(run func(context.Context, *sessions.User) (*auth.Token, error)) *AuthenticationProvider_CreateAndSetAuthToken_Call {
	_c.Call.Return(run)
	return _c
}

// CreateSession provides a mock function with given fields: ctx, sr
func (_m *AuthenticationProvider) CreateSession(ctx context.Context, sr sessions.SessionRequest) (string, error) {
	ret := _m.Called(ctx, sr)

	if len(ret) == 0 {
		panic("no return value specified for CreateSession")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sessions.SessionRequest) (string, error)); ok {
		return rf(ctx, sr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sessions.SessionRequest) string); ok {
		r0 = rf(ctx, sr)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, sessions.SessionRequest) error); ok {
		r1 = rf(ctx, sr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthenticationProvider_CreateSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSession'
type AuthenticationProvider_CreateSession_Call struct {
	*mock.Call
}

// CreateSession is a helper method to define mock.On call
//   - ctx context.Context
//   - sr sessions.SessionRequest
func (_e *AuthenticationProvider_Expecter) CreateSession(ctx interface{}, sr interface{}) *AuthenticationProvider_CreateSession_Call {
	return &AuthenticationProvider_CreateSession_Call{Call: _e.mock.On("CreateSession", ctx, sr)}
}

func (_c *AuthenticationProvider_CreateSession_Call) Run(run func(ctx context.Context, sr sessions.SessionRequest)) *AuthenticationProvider_CreateSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sessions.SessionRequest))
	})
	return _c
}

func (_c *AuthenticationProvider_CreateSession_Call) Return(_a0 string, _a1 error) *AuthenticationProvider_CreateSession_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthenticationProvider_CreateSession_Call) RunAndReturn(run func(context.Context, sessions.SessionRequest) (string, error)) *AuthenticationProvider_CreateSession_Call {
	_c.Call.Return(run)
	return _c
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *AuthenticationProvider) CreateUser(ctx context.Context, user *sessions.User) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sessions.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthenticationProvider_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type AuthenticationProvider_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - user *sessions.User
func (_e *AuthenticationProvider_Expecter) CreateUser(ctx interface{}, user interface{}) *AuthenticationProvider_CreateUser_Call {
	return &AuthenticationProvider_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, user)}
}

func (_c *AuthenticationProvider_CreateUser_Call) Run(run func(ctx context.Context, user *sessions.User)) *AuthenticationProvider_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*sessions.User))
	})
	return _c
}

func (_c *AuthenticationProvider_CreateUser_Call) Return(_a0 error) *AuthenticationProvider_CreateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthenticationProvider_CreateUser_Call) RunAndReturn(run func(context.Context, *sessions.User) error) *AuthenticationProvider_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAuthToken provides a mock function with given fields: ctx, user
func (_m *AuthenticationProvider) DeleteAuthToken(ctx context.Context, user *sessions.User) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAuthToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sessions.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthenticationProvider_DeleteAuthToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAuthToken'
type AuthenticationProvider_DeleteAuthToken_Call struct {
	*mock.Call
}

// DeleteAuthToken is a helper method to define mock.On call
//   - ctx context.Context
//   - user *sessions.User
func (_e *AuthenticationProvider_Expecter) DeleteAuthToken(ctx interface{}, user interface{}) *AuthenticationProvider_DeleteAuthToken_Call {
	return &AuthenticationProvider_DeleteAuthToken_Call{Call: _e.mock.On("DeleteAuthToken", ctx, user)}
}

func (_c *AuthenticationProvider_DeleteAuthToken_Call) Run(run func(ctx context.Context, user *sessions.User)) *AuthenticationProvider_DeleteAuthToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*sessions.User))
	})
	return _c
}

func (_c *AuthenticationProvider_DeleteAuthToken_Call) Return(_a0 error) *AuthenticationProvider_DeleteAuthToken_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthenticationProvider_DeleteAuthToken_Call) RunAndReturn(run func(context.Context, *sessions.User) error) *AuthenticationProvider_DeleteAuthToken_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUser provides a mock function with given fields: ctx, email
func (_m *AuthenticationProvider) DeleteUser(ctx context.Context, email string) error {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthenticationProvider_DeleteUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUser'
type AuthenticationProvider_DeleteUser_Call struct {
	*mock.Call
}

// DeleteUser is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *AuthenticationProvider_Expecter) DeleteUser(ctx interface{}, email interface{}) *AuthenticationProvider_DeleteUser_Call {
	return &AuthenticationProvider_DeleteUser_Call{Call: _e.mock.On("DeleteUser", ctx, email)}
}

func (_c *AuthenticationProvider_DeleteUser_Call) Run(run func(ctx context.Context, email string)) *AuthenticationProvider_DeleteUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AuthenticationProvider_DeleteUser_Call) Return(_a0 error) *AuthenticationProvider_DeleteUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthenticationProvider_DeleteUser_Call) RunAndReturn(run func(context.Context, string) error) *AuthenticationProvider_DeleteUser_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUserSession provides a mock function with given fields: ctx, sessionID
func (_m *AuthenticationProvider) DeleteUserSession(ctx context.Context, sessionID string) error {
	ret := _m.Called(ctx, sessionID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUserSession")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, sessionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthenticationProvider_DeleteUserSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUserSession'
type AuthenticationProvider_DeleteUserSession_Call struct {
	*mock.Call
}

// DeleteUserSession is a helper method to define mock.On call
//   - ctx context.Context
//   - sessionID string
func (_e *AuthenticationProvider_Expecter) DeleteUserSession(ctx interface{}, sessionID interface{}) *AuthenticationProvider_DeleteUserSession_Call {
	return &AuthenticationProvider_DeleteUserSession_Call{Call: _e.mock.On("DeleteUserSession", ctx, sessionID)}
}

func (_c *AuthenticationProvider_DeleteUserSession_Call) Run(run func(ctx context.Context, sessionID string)) *AuthenticationProvider_DeleteUserSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AuthenticationProvider_DeleteUserSession_Call) Return(_a0 error) *AuthenticationProvider_DeleteUserSession_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthenticationProvider_DeleteUserSession_Call) RunAndReturn(run func(context.Context, string) error) *AuthenticationProvider_DeleteUserSession_Call {
	_c.Call.Return(run)
	return _c
}

// FindExternalInitiator provides a mock function with given fields: ctx, eia
func (_m *AuthenticationProvider) FindExternalInitiator(ctx context.Context, eia *auth.Token) (*bridges.ExternalInitiator, error) {
	ret := _m.Called(ctx, eia)

	if len(ret) == 0 {
		panic("no return value specified for FindExternalInitiator")
	}

	var r0 *bridges.ExternalInitiator
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *auth.Token) (*bridges.ExternalInitiator, error)); ok {
		return rf(ctx, eia)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *auth.Token) *bridges.ExternalInitiator); ok {
		r0 = rf(ctx, eia)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bridges.ExternalInitiator)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *auth.Token) error); ok {
		r1 = rf(ctx, eia)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthenticationProvider_FindExternalInitiator_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindExternalInitiator'
type AuthenticationProvider_FindExternalInitiator_Call struct {
	*mock.Call
}

// FindExternalInitiator is a helper method to define mock.On call
//   - ctx context.Context
//   - eia *auth.Token
func (_e *AuthenticationProvider_Expecter) FindExternalInitiator(ctx interface{}, eia interface{}) *AuthenticationProvider_FindExternalInitiator_Call {
	return &AuthenticationProvider_FindExternalInitiator_Call{Call: _e.mock.On("FindExternalInitiator", ctx, eia)}
}

func (_c *AuthenticationProvider_FindExternalInitiator_Call) Run(run func(ctx context.Context, eia *auth.Token)) *AuthenticationProvider_FindExternalInitiator_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*auth.Token))
	})
	return _c
}

func (_c *AuthenticationProvider_FindExternalInitiator_Call) Return(initiator *bridges.ExternalInitiator, err error) *AuthenticationProvider_FindExternalInitiator_Call {
	_c.Call.Return(initiator, err)
	return _c
}

func (_c *AuthenticationProvider_FindExternalInitiator_Call) RunAndReturn(run func(context.Context, *auth.Token) (*bridges.ExternalInitiator, error)) *AuthenticationProvider_FindExternalInitiator_Call {
	_c.Call.Return(run)
	return _c
}

// FindUser provides a mock function with given fields: ctx, email
func (_m *AuthenticationProvider) FindUser(ctx context.Context, email string) (sessions.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for FindUser")
	}

	var r0 sessions.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (sessions.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) sessions.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(sessions.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthenticationProvider_FindUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindUser'
type AuthenticationProvider_FindUser_Call struct {
	*mock.Call
}

// FindUser is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *AuthenticationProvider_Expecter) FindUser(ctx interface{}, email interface{}) *AuthenticationProvider_FindUser_Call {
	return &AuthenticationProvider_FindUser_Call{Call: _e.mock.On("FindUser", ctx, email)}
}

func (_c *AuthenticationProvider_FindUser_Call) Run(run func(ctx context.Context, email string)) *AuthenticationProvider_FindUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AuthenticationProvider_FindUser_Call) Return(_a0 sessions.User, _a1 error) *AuthenticationProvider_FindUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthenticationProvider_FindUser_Call) RunAndReturn(run func(context.Context, string) (sessions.User, error)) *AuthenticationProvider_FindUser_Call {
	_c.Call.Return(run)
	return _c
}

// FindUserByAPIToken provides a mock function with given fields: ctx, apiToken
func (_m *AuthenticationProvider) FindUserByAPIToken(ctx context.Context, apiToken string) (sessions.User, error) {
	ret := _m.Called(ctx, apiToken)

	if len(ret) == 0 {
		panic("no return value specified for FindUserByAPIToken")
	}

	var r0 sessions.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (sessions.User, error)); ok {
		return rf(ctx, apiToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) sessions.User); ok {
		r0 = rf(ctx, apiToken)
	} else {
		r0 = ret.Get(0).(sessions.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, apiToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthenticationProvider_FindUserByAPIToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindUserByAPIToken'
type AuthenticationProvider_FindUserByAPIToken_Call struct {
	*mock.Call
}

// FindUserByAPIToken is a helper method to define mock.On call
//   - ctx context.Context
//   - apiToken string
func (_e *AuthenticationProvider_Expecter) FindUserByAPIToken(ctx interface{}, apiToken interface{}) *AuthenticationProvider_FindUserByAPIToken_Call {
	return &AuthenticationProvider_FindUserByAPIToken_Call{Call: _e.mock.On("FindUserByAPIToken", ctx, apiToken)}
}

func (_c *AuthenticationProvider_FindUserByAPIToken_Call) Run(run func(ctx context.Context, apiToken string)) *AuthenticationProvider_FindUserByAPIToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AuthenticationProvider_FindUserByAPIToken_Call) Return(_a0 sessions.User, _a1 error) *AuthenticationProvider_FindUserByAPIToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthenticationProvider_FindUserByAPIToken_Call) RunAndReturn(run func(context.Context, string) (sessions.User, error)) *AuthenticationProvider_FindUserByAPIToken_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserWebAuthn provides a mock function with given fields: ctx, email
func (_m *AuthenticationProvider) GetUserWebAuthn(ctx context.Context, email string) ([]sessions.WebAuthn, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserWebAuthn")
	}

	var r0 []sessions.WebAuthn
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]sessions.WebAuthn, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []sessions.WebAuthn); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sessions.WebAuthn)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthenticationProvider_GetUserWebAuthn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserWebAuthn'
type AuthenticationProvider_GetUserWebAuthn_Call struct {
	*mock.Call
}

// GetUserWebAuthn is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *AuthenticationProvider_Expecter) GetUserWebAuthn(ctx interface{}, email interface{}) *AuthenticationProvider_GetUserWebAuthn_Call {
	return &AuthenticationProvider_GetUserWebAuthn_Call{Call: _e.mock.On("GetUserWebAuthn", ctx, email)}
}

func (_c *AuthenticationProvider_GetUserWebAuthn_Call) Run(run func(ctx context.Context, email string)) *AuthenticationProvider_GetUserWebAuthn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AuthenticationProvider_GetUserWebAuthn_Call) Return(_a0 []sessions.WebAuthn, _a1 error) *AuthenticationProvider_GetUserWebAuthn_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthenticationProvider_GetUserWebAuthn_Call) RunAndReturn(run func(context.Context, string) ([]sessions.WebAuthn, error)) *AuthenticationProvider_GetUserWebAuthn_Call {
	_c.Call.Return(run)
	return _c
}

// ListUsers provides a mock function with given fields: ctx
func (_m *AuthenticationProvider) ListUsers(ctx context.Context) ([]sessions.User, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListUsers")
	}

	var r0 []sessions.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]sessions.User, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []sessions.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sessions.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthenticationProvider_ListUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListUsers'
type AuthenticationProvider_ListUsers_Call struct {
	*mock.Call
}

// ListUsers is a helper method to define mock.On call
//   - ctx context.Context
func (_e *AuthenticationProvider_Expecter) ListUsers(ctx interface{}) *AuthenticationProvider_ListUsers_Call {
	return &AuthenticationProvider_ListUsers_Call{Call: _e.mock.On("ListUsers", ctx)}
}

func (_c *AuthenticationProvider_ListUsers_Call) Run(run func(ctx context.Context)) *AuthenticationProvider_ListUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *AuthenticationProvider_ListUsers_Call) Return(_a0 []sessions.User, _a1 error) *AuthenticationProvider_ListUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthenticationProvider_ListUsers_Call) RunAndReturn(run func(context.Context) ([]sessions.User, error)) *AuthenticationProvider_ListUsers_Call {
	_c.Call.Return(run)
	return _c
}

// SaveWebAuthn provides a mock function with given fields: ctx, token
func (_m *AuthenticationProvider) SaveWebAuthn(ctx context.Context, token *sessions.WebAuthn) error {
	ret := _m.Called(ctx, token)

	if len(ret) == 0 {
		panic("no return value specified for SaveWebAuthn")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sessions.WebAuthn) error); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthenticationProvider_SaveWebAuthn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveWebAuthn'
type AuthenticationProvider_SaveWebAuthn_Call struct {
	*mock.Call
}

// SaveWebAuthn is a helper method to define mock.On call
//   - ctx context.Context
//   - token *sessions.WebAuthn
func (_e *AuthenticationProvider_Expecter) SaveWebAuthn(ctx interface{}, token interface{}) *AuthenticationProvider_SaveWebAuthn_Call {
	return &AuthenticationProvider_SaveWebAuthn_Call{Call: _e.mock.On("SaveWebAuthn", ctx, token)}
}

func (_c *AuthenticationProvider_SaveWebAuthn_Call) Run(run func(ctx context.Context, token *sessions.WebAuthn)) *AuthenticationProvider_SaveWebAuthn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*sessions.WebAuthn))
	})
	return _c
}

func (_c *AuthenticationProvider_SaveWebAuthn_Call) Return(_a0 error) *AuthenticationProvider_SaveWebAuthn_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthenticationProvider_SaveWebAuthn_Call) RunAndReturn(run func(context.Context, *sessions.WebAuthn) error) *AuthenticationProvider_SaveWebAuthn_Call {
	_c.Call.Return(run)
	return _c
}

// Sessions provides a mock function with given fields: ctx, offset, limit
func (_m *AuthenticationProvider) Sessions(ctx context.Context, offset int, limit int) ([]sessions.Session, error) {
	ret := _m.Called(ctx, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for Sessions")
	}

	var r0 []sessions.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]sessions.Session, error)); ok {
		return rf(ctx, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []sessions.Session); ok {
		r0 = rf(ctx, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sessions.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthenticationProvider_Sessions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sessions'
type AuthenticationProvider_Sessions_Call struct {
	*mock.Call
}

// Sessions is a helper method to define mock.On call
//   - ctx context.Context
//   - offset int
//   - limit int
func (_e *AuthenticationProvider_Expecter) Sessions(ctx interface{}, offset interface{}, limit interface{}) *AuthenticationProvider_Sessions_Call {
	return &AuthenticationProvider_Sessions_Call{Call: _e.mock.On("Sessions", ctx, offset, limit)}
}

func (_c *AuthenticationProvider_Sessions_Call) Run(run func(ctx context.Context, offset int, limit int)) *AuthenticationProvider_Sessions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int))
	})
	return _c
}

func (_c *AuthenticationProvider_Sessions_Call) Return(_a0 []sessions.Session, _a1 error) *AuthenticationProvider_Sessions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthenticationProvider_Sessions_Call) RunAndReturn(run func(context.Context, int, int) ([]sessions.Session, error)) *AuthenticationProvider_Sessions_Call {
	_c.Call.Return(run)
	return _c
}

// SetAuthToken provides a mock function with given fields: ctx, user, token
func (_m *AuthenticationProvider) SetAuthToken(ctx context.Context, user *sessions.User, token *auth.Token) error {
	ret := _m.Called(ctx, user, token)

	if len(ret) == 0 {
		panic("no return value specified for SetAuthToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sessions.User, *auth.Token) error); ok {
		r0 = rf(ctx, user, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthenticationProvider_SetAuthToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetAuthToken'
type AuthenticationProvider_SetAuthToken_Call struct {
	*mock.Call
}

// SetAuthToken is a helper method to define mock.On call
//   - ctx context.Context
//   - user *sessions.User
//   - token *auth.Token
func (_e *AuthenticationProvider_Expecter) SetAuthToken(ctx interface{}, user interface{}, token interface{}) *AuthenticationProvider_SetAuthToken_Call {
	return &AuthenticationProvider_SetAuthToken_Call{Call: _e.mock.On("SetAuthToken", ctx, user, token)}
}

func (_c *AuthenticationProvider_SetAuthToken_Call) Run(run func(ctx context.Context, user *sessions.User, token *auth.Token)) *AuthenticationProvider_SetAuthToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*sessions.User), args[2].(*auth.Token))
	})
	return _c
}

func (_c *AuthenticationProvider_SetAuthToken_Call) Return(_a0 error) *AuthenticationProvider_SetAuthToken_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthenticationProvider_SetAuthToken_Call) RunAndReturn(run func(context.Context, *sessions.User, *auth.Token) error) *AuthenticationProvider_SetAuthToken_Call {
	_c.Call.Return(run)
	return _c
}

// SetPassword provides a mock function with given fields: ctx, user, newPassword
func (_m *AuthenticationProvider) SetPassword(ctx context.Context, user *sessions.User, newPassword string) error {
	ret := _m.Called(ctx, user, newPassword)

	if len(ret) == 0 {
		panic("no return value specified for SetPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sessions.User, string) error); ok {
		r0 = rf(ctx, user, newPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthenticationProvider_SetPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPassword'
type AuthenticationProvider_SetPassword_Call struct {
	*mock.Call
}

// SetPassword is a helper method to define mock.On call
//   - ctx context.Context
//   - user *sessions.User
//   - newPassword string
func (_e *AuthenticationProvider_Expecter) SetPassword(ctx interface{}, user interface{}, newPassword interface{}) *AuthenticationProvider_SetPassword_Call {
	return &AuthenticationProvider_SetPassword_Call{Call: _e.mock.On("SetPassword", ctx, user, newPassword)}
}

func (_c *AuthenticationProvider_SetPassword_Call) Run(run func(ctx context.Context, user *sessions.User, newPassword string)) *AuthenticationProvider_SetPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*sessions.User), args[2].(string))
	})
	return _c
}

func (_c *AuthenticationProvider_SetPassword_Call) Return(_a0 error) *AuthenticationProvider_SetPassword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthenticationProvider_SetPassword_Call) RunAndReturn(run func(context.Context, *sessions.User, string) error) *AuthenticationProvider_SetPassword_Call {
	_c.Call.Return(run)
	return _c
}

// TestPassword provides a mock function with given fields: ctx, email, password
func (_m *AuthenticationProvider) TestPassword(ctx context.Context, email string, password string) error {
	ret := _m.Called(ctx, email, password)

	if len(ret) == 0 {
		panic("no return value specified for TestPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, email, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthenticationProvider_TestPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TestPassword'
type AuthenticationProvider_TestPassword_Call struct {
	*mock.Call
}

// TestPassword is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
//   - password string
func (_e *AuthenticationProvider_Expecter) TestPassword(ctx interface{}, email interface{}, password interface{}) *AuthenticationProvider_TestPassword_Call {
	return &AuthenticationProvider_TestPassword_Call{Call: _e.mock.On("TestPassword", ctx, email, password)}
}

func (_c *AuthenticationProvider_TestPassword_Call) Run(run func(ctx context.Context, email string, password string)) *AuthenticationProvider_TestPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *AuthenticationProvider_TestPassword_Call) Return(_a0 error) *AuthenticationProvider_TestPassword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthenticationProvider_TestPassword_Call) RunAndReturn(run func(context.Context, string, string) error) *AuthenticationProvider_TestPassword_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateRole provides a mock function with given fields: ctx, email, newRole
func (_m *AuthenticationProvider) UpdateRole(ctx context.Context, email string, newRole string) (sessions.User, error) {
	ret := _m.Called(ctx, email, newRole)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRole")
	}

	var r0 sessions.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (sessions.User, error)); ok {
		return rf(ctx, email, newRole)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) sessions.User); ok {
		r0 = rf(ctx, email, newRole)
	} else {
		r0 = ret.Get(0).(sessions.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, newRole)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthenticationProvider_UpdateRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateRole'
type AuthenticationProvider_UpdateRole_Call struct {
	*mock.Call
}

// UpdateRole is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
//   - newRole string
func (_e *AuthenticationProvider_Expecter) UpdateRole(ctx interface{}, email interface{}, newRole interface{}) *AuthenticationProvider_UpdateRole_Call {
	return &AuthenticationProvider_UpdateRole_Call{Call: _e.mock.On("UpdateRole", ctx, email, newRole)}
}

func (_c *AuthenticationProvider_UpdateRole_Call) Run(run func(ctx context.Context, email string, newRole string)) *AuthenticationProvider_UpdateRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *AuthenticationProvider_UpdateRole_Call) Return(_a0 sessions.User, _a1 error) *AuthenticationProvider_UpdateRole_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthenticationProvider_UpdateRole_Call) RunAndReturn(run func(context.Context, string, string) (sessions.User, error)) *AuthenticationProvider_UpdateRole_Call {
	_c.Call.Return(run)
	return _c
}

// NewAuthenticationProvider creates a new instance of AuthenticationProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthenticationProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthenticationProvider {
	mock := &AuthenticationProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
