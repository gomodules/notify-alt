// Code generated by mockery 2.10.0. DO NOT EDIT.

package syslog

import mock "github.com/stretchr/testify/mock"

// mockSyslogWriter is an autogenerated mock type for the mockSyslogWriter type
type mockSyslogWriter struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *mockSyslogWriter) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Write provides a mock function with given fields: p
func (_m *mockSyslogWriter) Write(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
