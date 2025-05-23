// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// ClientHandlerMock is an autogenerated mock type for the ClientHandler type
type ClientHandlerMock struct {
	mock.Mock
}

type ClientHandlerMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ClientHandlerMock) EXPECT() *ClientHandlerMock_Expecter {
	return &ClientHandlerMock_Expecter{mock: &_m.Mock}
}

// CreateClient provides a mock function with given fields: ectx
func (_m *ClientHandlerMock) CreateClient(ectx echo.Context) error {
	ret := _m.Called(ectx)

	if len(ret) == 0 {
		panic("no return value specified for CreateClient")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ectx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientHandlerMock_CreateClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateClient'
type ClientHandlerMock_CreateClient_Call struct {
	*mock.Call
}

// CreateClient is a helper method to define mock.On call
//   - ectx echo.Context
func (_e *ClientHandlerMock_Expecter) CreateClient(ectx interface{}) *ClientHandlerMock_CreateClient_Call {
	return &ClientHandlerMock_CreateClient_Call{Call: _e.mock.On("CreateClient", ectx)}
}

func (_c *ClientHandlerMock_CreateClient_Call) Run(run func(ectx echo.Context)) *ClientHandlerMock_CreateClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *ClientHandlerMock_CreateClient_Call) Return(_a0 error) *ClientHandlerMock_CreateClient_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClientHandlerMock_CreateClient_Call) RunAndReturn(run func(echo.Context) error) *ClientHandlerMock_CreateClient_Call {
	_c.Call.Return(run)
	return _c
}

// GetClientContactsByID provides a mock function with given fields: ectx
func (_m *ClientHandlerMock) GetClientContactsByID(ectx echo.Context) error {
	ret := _m.Called(ectx)

	if len(ret) == 0 {
		panic("no return value specified for GetClientContactsByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ectx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientHandlerMock_GetClientContactsByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetClientContactsByID'
type ClientHandlerMock_GetClientContactsByID_Call struct {
	*mock.Call
}

// GetClientContactsByID is a helper method to define mock.On call
//   - ectx echo.Context
func (_e *ClientHandlerMock_Expecter) GetClientContactsByID(ectx interface{}) *ClientHandlerMock_GetClientContactsByID_Call {
	return &ClientHandlerMock_GetClientContactsByID_Call{Call: _e.mock.On("GetClientContactsByID", ectx)}
}

func (_c *ClientHandlerMock_GetClientContactsByID_Call) Run(run func(ectx echo.Context)) *ClientHandlerMock_GetClientContactsByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *ClientHandlerMock_GetClientContactsByID_Call) Return(_a0 error) *ClientHandlerMock_GetClientContactsByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClientHandlerMock_GetClientContactsByID_Call) RunAndReturn(run func(echo.Context) error) *ClientHandlerMock_GetClientContactsByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetClientsWithContact provides a mock function with given fields: ectx
func (_m *ClientHandlerMock) GetClientsWithContact(ectx echo.Context) error {
	ret := _m.Called(ectx)

	if len(ret) == 0 {
		panic("no return value specified for GetClientsWithContact")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ectx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientHandlerMock_GetClientsWithContact_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetClientsWithContact'
type ClientHandlerMock_GetClientsWithContact_Call struct {
	*mock.Call
}

// GetClientsWithContact is a helper method to define mock.On call
//   - ectx echo.Context
func (_e *ClientHandlerMock_Expecter) GetClientsWithContact(ectx interface{}) *ClientHandlerMock_GetClientsWithContact_Call {
	return &ClientHandlerMock_GetClientsWithContact_Call{Call: _e.mock.On("GetClientsWithContact", ectx)}
}

func (_c *ClientHandlerMock_GetClientsWithContact_Call) Run(run func(ectx echo.Context)) *ClientHandlerMock_GetClientsWithContact_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *ClientHandlerMock_GetClientsWithContact_Call) Return(_a0 error) *ClientHandlerMock_GetClientsWithContact_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClientHandlerMock_GetClientsWithContact_Call) RunAndReturn(run func(echo.Context) error) *ClientHandlerMock_GetClientsWithContact_Call {
	_c.Call.Return(run)
	return _c
}

// NewClientHandlerMock creates a new instance of ClientHandlerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientHandlerMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientHandlerMock {
	mock := &ClientHandlerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
