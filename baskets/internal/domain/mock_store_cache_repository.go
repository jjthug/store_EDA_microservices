// Code generated by mockery v2.14.0. DO NOT EDIT.

package domain

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockStoreCacheRepository is an autogenerated mock type for the StoreCacheRepository type
type MockStoreCacheRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, storeID, name
func (_m *MockStoreCacheRepository) Add(ctx context.Context, storeID string, name string) error {
	ret := _m.Called(ctx, storeID, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, storeID, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctx, storeID
func (_m *MockStoreCacheRepository) Find(ctx context.Context, storeID string) (*Store, error) {
	ret := _m.Called(ctx, storeID)

	var r0 *Store
	if rf, ok := ret.Get(0).(func(context.Context, string) *Store); ok {
		r0 = rf(ctx, storeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Store)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, storeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Rename provides a mock function with given fields: ctx, storeID, name
func (_m *MockStoreCacheRepository) Rename(ctx context.Context, storeID string, name string) error {
	ret := _m.Called(ctx, storeID, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, storeID, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockStoreCacheRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockStoreCacheRepository creates a new instance of MockStoreCacheRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockStoreCacheRepository(t mockConstructorTestingTNewMockStoreCacheRepository) *MockStoreCacheRepository {
	mock := &MockStoreCacheRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
