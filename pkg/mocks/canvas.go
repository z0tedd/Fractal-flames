// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// Canvas is an autogenerated mock type for the Canvas type
type Canvas struct {
	mock.Mock
}

type Canvas_Expecter struct {
	mock *mock.Mock
}

func (_m *Canvas) EXPECT() *Canvas_Expecter {
	return &Canvas_Expecter{mock: &_m.Mock}
}

// Canvas provides a mock function with no fields
func (_m *Canvas) Canvas() [][]domain.Pixel {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Canvas")
	}

	var r0 [][]domain.Pixel
	if rf, ok := ret.Get(0).(func() [][]domain.Pixel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]domain.Pixel)
		}
	}

	return r0
}

// Canvas_Canvas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Canvas'
type Canvas_Canvas_Call struct {
	*mock.Call
}

// Canvas is a helper method to define mock.On call
func (_e *Canvas_Expecter) Canvas() *Canvas_Canvas_Call {
	return &Canvas_Canvas_Call{Call: _e.mock.On("Canvas")}
}

func (_c *Canvas_Canvas_Call) Run(run func()) *Canvas_Canvas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Canvas_Canvas_Call) Return(_a0 [][]domain.Pixel) *Canvas_Canvas_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Canvas_Canvas_Call) RunAndReturn(run func() [][]domain.Pixel) *Canvas_Canvas_Call {
	_c.Call.Return(run)
	return _c
}

// NewCanvas creates a new instance of Canvas. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCanvas(t interface {
	mock.TestingT
	Cleanup(func())
}) *Canvas {
	mock := &Canvas{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
