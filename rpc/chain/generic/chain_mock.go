// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package generic

import (
	types "github.com/Cerebellum-Network/go-substrate-rpc-client/v7/types"
	mock "github.com/stretchr/testify/mock"
)

// ChainMock is an autogenerated mock type for the Chain type
type ChainMock[A interface{}, S interface{}, P interface{}, B GenericSignedBlock[A, S, P]] struct {
	mock.Mock
}

// GetBlock provides a mock function with given fields: blockHash
func (_m *ChainMock[A, S, P, B]) GetBlock(blockHash types.Hash) (B, error) {
	ret := _m.Called(blockHash)

	var r0 B
	if rf, ok := ret.Get(0).(func(types.Hash) B); ok {
		r0 = rf(blockHash)
	} else {
		r0 = ret.Get(0).(B)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Hash) error); ok {
		r1 = rf(blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockLatest provides a mock function with given fields:
func (_m *ChainMock[A, S, P, B]) GetBlockLatest() (B, error) {
	ret := _m.Called()

	var r0 B
	if rf, ok := ret.Get(0).(func() B); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(B)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewChainMockT interface {
	mock.TestingT
	Cleanup(func())
}

// NewChainMock creates a new instance of ChainMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewChainMock[A interface{}, S interface{}, P interface{}, B GenericSignedBlock[A, S, P]](t NewChainMockT) *ChainMock[A, S, P, B] {
	mock := &ChainMock[A, S, P, B]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
