// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package retriever

import (
	parser "github.com/Cerebellum-Network/go-substrate-rpc-client/v9/registry/parser"
	types "github.com/Cerebellum-Network/go-substrate-rpc-client/v9/types"
	mock "github.com/stretchr/testify/mock"
)

// ExtrinsicRetrieverMock is an autogenerated mock type for the ExtrinsicRetriever type
type ExtrinsicRetrieverMock[A interface{}, S interface{}, P interface{}] struct {
	mock.Mock
}

// GetExtrinsics provides a mock function with given fields: blockHash
func (_m *ExtrinsicRetrieverMock[A, S, P]) GetExtrinsics(blockHash types.Hash) ([]*parser.Extrinsic[A, S, P], error) {
	ret := _m.Called(blockHash)

	var r0 []*parser.Extrinsic[A, S, P]
	if rf, ok := ret.Get(0).(func(types.Hash) []*parser.Extrinsic[A, S, P]); ok {
		r0 = rf(blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*parser.Extrinsic[A, S, P])
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Hash) error); ok {
		r1 = rf(blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewExtrinsicRetrieverMockT interface {
	mock.TestingT
	Cleanup(func())
}

// NewExtrinsicRetrieverMock creates a new instance of ExtrinsicRetrieverMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExtrinsicRetrieverMock[A interface{}, S interface{}, P interface{}](t NewExtrinsicRetrieverMockT) *ExtrinsicRetrieverMock[A, S, P] {
	mock := &ExtrinsicRetrieverMock[A, S, P]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
