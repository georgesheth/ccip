// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ServiceCtx is an autogenerated mock type for the ServiceCtx type
type ServiceCtx struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *ServiceCtx) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *ServiceCtx) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewServiceCtx interface {
	mock.TestingT
	Cleanup(func())
}

// NewServiceCtx creates a new instance of ServiceCtx. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServiceCtx(t mockConstructorTestingTNewServiceCtx) *ServiceCtx {
	mock := &ServiceCtx{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
