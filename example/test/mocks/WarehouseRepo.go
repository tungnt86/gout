// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	model "github.com/tungnt/gout/example/model"
)

// WarehouseRepo is an autogenerated mock type for the WarehouseRepo type
type WarehouseRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, name, cityID
func (_m *WarehouseRepo) Create(ctx context.Context, name string, cityID int64) (*model.Warehouse, error) {
	ret := _m.Called(ctx, name, cityID)

	var r0 *model.Warehouse
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) *model.Warehouse); ok {
		r0 = rf(ctx, name, cityID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Warehouse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, name, cityID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
