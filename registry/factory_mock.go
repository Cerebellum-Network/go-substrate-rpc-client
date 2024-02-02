// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package registry

import (
	types "github.com/Cerebellum-Network/go-substrate-rpc-client/v7/types"
	mock "github.com/stretchr/testify/mock"
)

// FactoryMock is an autogenerated mock type for the Factory type
type FactoryMock struct {
	mock.Mock
}

// CreateCallRegistry provides a mock function with given fields: meta
func (_m *FactoryMock) CreateCallRegistry(meta *types.Metadata) (CallRegistry, error) {
	ret := _m.Called(meta)

	var r0 CallRegistry
	if rf, ok := ret.Get(0).(func(*types.Metadata) CallRegistry); ok {
		r0 = rf(meta)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(CallRegistry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.Metadata) error); ok {
		r1 = rf(meta)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateErrorRegistry provides a mock function with given fields: meta
func (_m *FactoryMock) CreateErrorRegistry(meta *types.Metadata) (ErrorRegistry, error) {
	ret := _m.Called(meta)

	var r0 ErrorRegistry
	if rf, ok := ret.Get(0).(func(*types.Metadata) ErrorRegistry); ok {
		r0 = rf(meta)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ErrorRegistry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.Metadata) error); ok {
		r1 = rf(meta)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEventRegistry provides a mock function with given fields: meta
func (_m *FactoryMock) CreateEventRegistry(meta *types.Metadata) (EventRegistry, error) {
	ret := _m.Called(meta)

	var r0 EventRegistry
	if rf, ok := ret.Get(0).(func(*types.Metadata) EventRegistry); ok {
		r0 = rf(meta)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EventRegistry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.Metadata) error); ok {
		r1 = rf(meta)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewFactoryMockT interface {
	mock.TestingT
	Cleanup(func())
}

// NewFactoryMock creates a new instance of FactoryMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFactoryMock(t NewFactoryMockT) *FactoryMock {
	mock := &FactoryMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
