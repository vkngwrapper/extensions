// Code generated by MockGen. DO NOT EDIT.
// Source: driver.go

// Package mock_host_query_reset is a generated GoMock package.
package mock_host_query_reset

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	driver "github.com/vkngwrapper/core/v2/driver"
)

// MockDriver is a mock of Driver interface.
type MockDriver struct {
	ctrl     *gomock.Controller
	recorder *MockDriverMockRecorder
}

// MockDriverMockRecorder is the mock recorder for MockDriver.
type MockDriverMockRecorder struct {
	mock *MockDriver
}

// NewMockDriver creates a new mock instance.
func NewMockDriver(ctrl *gomock.Controller) *MockDriver {
	mock := &MockDriver{ctrl: ctrl}
	mock.recorder = &MockDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDriver) EXPECT() *MockDriverMockRecorder {
	return m.recorder
}

// VkResetQueryPoolEXT mocks base method.
func (m *MockDriver) VkResetQueryPoolEXT(device driver.VkDevice, queryPool driver.VkQueryPool, firstQuery, queryCount driver.Uint32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkResetQueryPoolEXT", device, queryPool, firstQuery, queryCount)
}

// VkResetQueryPoolEXT indicates an expected call of VkResetQueryPoolEXT.
func (mr *MockDriverMockRecorder) VkResetQueryPoolEXT(device, queryPool, firstQuery, queryCount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkResetQueryPoolEXT", reflect.TypeOf((*MockDriver)(nil).VkResetQueryPoolEXT), device, queryPool, firstQuery, queryCount)
}
