// Code generated by mockery v2.14.0. DO NOT EDIT.

package application

import (
	context "context"
	commands "EDA_GO/stores/internal/application/commands"

	mock "github.com/stretchr/testify/mock"
)

// MockCommands is an autogenerated mock type for the Commands type
type MockCommands struct {
	mock.Mock
}

// AddProduct provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) AddProduct(ctx context.Context, cmd commands.AddProduct) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.AddProduct) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateStore provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) CreateStore(ctx context.Context, cmd commands.CreateStore) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.CreateStore) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DecreaseProductPrice provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) DecreaseProductPrice(ctx context.Context, cmd commands.DecreaseProductPrice) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.DecreaseProductPrice) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DisableParticipation provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) DisableParticipation(ctx context.Context, cmd commands.DisableParticipation) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.DisableParticipation) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnableParticipation provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) EnableParticipation(ctx context.Context, cmd commands.EnableParticipation) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.EnableParticipation) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IncreaseProductPrice provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) IncreaseProductPrice(ctx context.Context, cmd commands.IncreaseProductPrice) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.IncreaseProductPrice) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RebrandProduct provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) RebrandProduct(ctx context.Context, cmd commands.RebrandProduct) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.RebrandProduct) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RebrandStore provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) RebrandStore(ctx context.Context, cmd commands.RebrandStore) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.RebrandStore) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveProduct provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) RemoveProduct(ctx context.Context, cmd commands.RemoveProduct) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.RemoveProduct) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockCommands interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockCommands creates a new instance of MockCommands. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockCommands(t mockConstructorTestingTNewMockCommands) *MockCommands {
	mock := &MockCommands{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
