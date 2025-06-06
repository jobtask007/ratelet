// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	"ratelet/internal/rate/domain"

	mock "github.com/stretchr/testify/mock"
)

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

type MockService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockService) EXPECT() *MockService_Expecter {
	return &MockService_Expecter{mock: &_m.Mock}
}

// GetRates provides a mock function for the type MockService
func (_mock *MockService) GetRates(currencies []string) (domain.Rates, error) {
	ret := _mock.Called(currencies)

	if len(ret) == 0 {
		panic("no return value specified for GetRates")
	}

	var r0 domain.Rates
	var r1 error
	if returnFunc, ok := ret.Get(0).(func([]string) (domain.Rates, error)); ok {
		return returnFunc(currencies)
	}
	if returnFunc, ok := ret.Get(0).(func([]string) domain.Rates); ok {
		r0 = returnFunc(currencies)
	} else {
		r0 = ret.Get(0).(domain.Rates)
	}
	if returnFunc, ok := ret.Get(1).(func([]string) error); ok {
		r1 = returnFunc(currencies)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockService_GetRates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRates'
type MockService_GetRates_Call struct {
	*mock.Call
}

// GetRates is a helper method to define mock.On call
//   - currencies
func (_e *MockService_Expecter) GetRates(currencies interface{}) *MockService_GetRates_Call {
	return &MockService_GetRates_Call{Call: _e.mock.On("GetRates", currencies)}
}

func (_c *MockService_GetRates_Call) Run(run func(currencies []string)) *MockService_GetRates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *MockService_GetRates_Call) Return(rates domain.Rates, err error) *MockService_GetRates_Call {
	_c.Call.Return(rates, err)
	return _c
}

func (_c *MockService_GetRates_Call) RunAndReturn(run func(currencies []string) (domain.Rates, error)) *MockService_GetRates_Call {
	_c.Call.Return(run)
	return _c
}
