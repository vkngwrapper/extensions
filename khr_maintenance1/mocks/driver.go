// Code generated by MockGen. DO NOT EDIT.
// Source: driver.go

// Package mock_maintenance1 is a generated GoMock package.
package mock_maintenance1

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	driver "github.com/vkngwrapper/core/v2/driver"
	khr_maintenance1_driver "github.com/vkngwrapper/extensions/v2/khr_maintenance1/driver"
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

// VkTrimCommandPoolKHR mocks base method.
func (m *MockDriver) VkTrimCommandPoolKHR(device driver.VkDevice, commandPool driver.VkCommandPool, flags khr_maintenance1_driver.VkCommandPoolTrimFlagsKHR) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkTrimCommandPoolKHR", device, commandPool, flags)
}

// VkTrimCommandPoolKHR indicates an expected call of VkTrimCommandPoolKHR.
func (mr *MockDriverMockRecorder) VkTrimCommandPoolKHR(device, commandPool, flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkTrimCommandPoolKHR", reflect.TypeOf((*MockDriver)(nil).VkTrimCommandPoolKHR), device, commandPool, flags)
}
