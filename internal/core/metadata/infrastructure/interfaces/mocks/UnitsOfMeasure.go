// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnitsOfMeasure is an autogenerated mock type for the UnitsOfMeasure type
type UnitsOfMeasure struct {
	mock.Mock
}

// Validate provides a mock function with given fields: _a0
func (_m *UnitsOfMeasure) Validate(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
