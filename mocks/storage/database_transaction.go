// Code generated by mockery v1.0.0. DO NOT EDIT.

package storage

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// DatabaseTransaction is an autogenerated mock type for the DatabaseTransaction type
type DatabaseTransaction struct {
	mock.Mock
}

// Commit provides a mock function with given fields: _a0
func (_m *DatabaseTransaction) Commit(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *DatabaseTransaction) Delete(_a0 context.Context, _a1 []byte) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Discard provides a mock function with given fields: _a0
func (_m *DatabaseTransaction) Discard(_a0 context.Context) {
	_m.Called(_a0)
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *DatabaseTransaction) Get(_a0 context.Context, _a1 []byte) (bool, []byte, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, []byte) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(context.Context, []byte) []byte); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []byte) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Scan provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5
func (_m *DatabaseTransaction) Scan(_a0 context.Context, _a1 []byte, _a2 []byte, _a3 func([]byte, []byte) error, _a4 bool, _a5 bool) (int, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, []byte, []byte, func([]byte, []byte) error, bool, bool) int); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte, []byte, func([]byte, []byte) error, bool, bool) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *DatabaseTransaction) Set(_a0 context.Context, _a1 []byte, _a2 []byte, _a3 bool) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, []byte, bool) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}