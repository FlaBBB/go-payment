// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	payment "github.com/FlaBBB/go-payment"
	mock "github.com/stretchr/testify/mock"

	subscription "github.com/FlaBBB/go-payment/subscription"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, sub
func (_m *Controller) Create(ctx context.Context, sub *subscription.Subscription) (*subscription.CreateResponse, error) {
	ret := _m.Called(ctx, sub)

	var r0 *subscription.CreateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *subscription.Subscription) *subscription.CreateResponse); ok {
		r0 = rf(ctx, sub)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*subscription.CreateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *subscription.Subscription) error); ok {
		r1 = rf(ctx, sub)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Gateway provides a mock function with given fields:
func (_m *Controller) Gateway() payment.Gateway {
	ret := _m.Called()

	var r0 payment.Gateway
	if rf, ok := ret.Get(0).(func() payment.Gateway); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(payment.Gateway)
	}

	return r0
}

// Pause provides a mock function with given fields: ctx, sub
func (_m *Controller) Pause(ctx context.Context, sub *subscription.Subscription) error {
	ret := _m.Called(ctx, sub)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *subscription.Subscription) error); ok {
		r0 = rf(ctx, sub)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Resume provides a mock function with given fields: ctx, sub
func (_m *Controller) Resume(ctx context.Context, sub *subscription.Subscription) error {
	ret := _m.Called(ctx, sub)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *subscription.Subscription) error); ok {
		r0 = rf(ctx, sub)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields: ctx, stop
func (_m *Controller) Stop(ctx context.Context, stop *subscription.Subscription) error {
	ret := _m.Called(ctx, stop)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *subscription.Subscription) error); ok {
		r0 = rf(ctx, stop)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
