// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	model "github.com/tungnt/gout/example/model"
)

// CityRepo is an autogenerated mock type for the CityRepo type
type CityRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, name
func (_m *CityRepo) Create(ctx context.Context, name string) (*model.City, error) {
	ret := _m.Called(ctx, name)

	var r0 *model.City
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.City); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.City)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
