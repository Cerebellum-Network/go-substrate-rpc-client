// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	context "context"

	rpc "github.com/Cerebellum-Network/go-substrate-rpc-client/v7/gethrpc"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Call provides a mock function with given fields: result, method, args
func (_m *Client) Call(result interface{}, method string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, result, method)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...interface{}) error); ok {
		r0 = rf(result, method, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: ctx, namespace, subscribeMethodSuffix, unsubscribeMethodSuffix, notificationMethodSuffix, channel, args
func (_m *Client) Subscribe(ctx context.Context, namespace string, subscribeMethodSuffix string, unsubscribeMethodSuffix string, notificationMethodSuffix string, channel interface{}, args ...interface{}) (*rpc.ClientSubscription, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, namespace, subscribeMethodSuffix, unsubscribeMethodSuffix, notificationMethodSuffix, channel)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *rpc.ClientSubscription
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, interface{}, ...interface{}) *rpc.ClientSubscription); ok {
		r0 = rf(ctx, namespace, subscribeMethodSuffix, unsubscribeMethodSuffix, notificationMethodSuffix, channel, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.ClientSubscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, interface{}, ...interface{}) error); ok {
		r1 = rf(ctx, namespace, subscribeMethodSuffix, unsubscribeMethodSuffix, notificationMethodSuffix, channel, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// URL provides a mock function with given fields:
func (_m *Client) URL() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type NewClientT interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t NewClientT) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
