// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	types "github.com/Cerebellum-Network/go-substrate-rpc-client/v7/types"
	mock "github.com/stretchr/testify/mock"
)

// MMR is an autogenerated mock type for the MMR type
type MMR struct {
	mock.Mock
}

// GenerateProof provides a mock function with given fields: leafIndex, blockHash
func (_m *MMR) GenerateProof(leafIndex uint64, blockHash types.Hash) (types.GenerateMMRProofResponse, error) {
	ret := _m.Called(leafIndex, blockHash)

	var r0 types.GenerateMMRProofResponse
	if rf, ok := ret.Get(0).(func(uint64, types.Hash) types.GenerateMMRProofResponse); ok {
		r0 = rf(leafIndex, blockHash)
	} else {
		r0 = ret.Get(0).(types.GenerateMMRProofResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, types.Hash) error); ok {
		r1 = rf(leafIndex, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateProofLatest provides a mock function with given fields: leafIndex
func (_m *MMR) GenerateProofLatest(leafIndex uint64) (types.GenerateMMRProofResponse, error) {
	ret := _m.Called(leafIndex)

	var r0 types.GenerateMMRProofResponse
	if rf, ok := ret.Get(0).(func(uint64) types.GenerateMMRProofResponse); ok {
		r0 = rf(leafIndex)
	} else {
		r0 = ret.Get(0).(types.GenerateMMRProofResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(leafIndex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewMMRT interface {
	mock.TestingT
	Cleanup(func())
}

// NewMMR creates a new instance of MMR. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMMR(t NewMMRT) *MMR {
	mock := &MMR{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
