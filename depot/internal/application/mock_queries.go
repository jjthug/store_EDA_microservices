// Code generated by mockery v2.14.0. DO NOT EDIT.

package application

import (
	context "context"
	domain "EDA_GO/depot/internal/domain"

	mock "github.com/stretchr/testify/mock"

	queries "EDA_GO/depot/internal/application/queries"
)

// MockQueries is an autogenerated mock type for the Queries type
type MockQueries struct {
	mock.Mock
}

// GetShoppingList provides a mock function with given fields: ctx, query
func (_m *MockQueries) GetShoppingList(ctx context.Context, query queries.GetShoppingList) (*domain.ShoppingList, error) {
	ret := _m.Called(ctx, query)

	var r0 *domain.ShoppingList
	if rf, ok := ret.Get(0).(func(context.Context, queries.GetShoppingList) *domain.ShoppingList); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ShoppingList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, queries.GetShoppingList) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockQueries interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockQueries creates a new instance of MockQueries. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockQueries(t mockConstructorTestingTNewMockQueries) *MockQueries {
	mock := &MockQueries{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
