// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IController is an autogenerated mock type for the IController type
type IController struct {
	mock.Mock
}

// Check provides a mock function with given fields:
func (_m *IController) Check() {
	_m.Called()
}

// Get provides a mock function with given fields:
func (_m *IController) Get() {
	_m.Called()
}

// Insert provides a mock function with given fields:
func (_m *IController) Insert() {
	_m.Called()
}

// Routes provides a mock function with given fields:
func (_m *IController) Routes() {
	_m.Called()
}

type mockConstructorTestingTNewIController interface {
	mock.TestingT
	Cleanup(func())
}

// NewIController creates a new instance of IController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIController(t mockConstructorTestingTNewIController) *IController {
	mock := &IController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
