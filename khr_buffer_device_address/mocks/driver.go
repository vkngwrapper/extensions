// Code generated by MockGen. DO NOT EDIT.
// Source: driver.go

// Package mock_buffer_device_address is a generated GoMock package.
package mock_buffer_device_address

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	driver "github.com/vkngwrapper/core/v2/driver"
	khr_buffer_device_address_driver "github.com/vkngwrapper/extensions/v2/khr_buffer_device_address/driver"
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

// VkGetBufferDeviceAddressKHR mocks base method.
func (m *MockDriver) VkGetBufferDeviceAddressKHR(device driver.VkDevice, pInfo *khr_buffer_device_address_driver.VkBufferDeviceAddressInfoKHR) driver.VkDeviceAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkGetBufferDeviceAddressKHR", device, pInfo)
	ret0, _ := ret[0].(driver.VkDeviceAddress)
	return ret0
}

// VkGetBufferDeviceAddressKHR indicates an expected call of VkGetBufferDeviceAddressKHR.
func (mr *MockDriverMockRecorder) VkGetBufferDeviceAddressKHR(device, pInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetBufferDeviceAddressKHR", reflect.TypeOf((*MockDriver)(nil).VkGetBufferDeviceAddressKHR), device, pInfo)
}

// VkGetBufferOpaqueCaptureAddressKHR mocks base method.
func (m *MockDriver) VkGetBufferOpaqueCaptureAddressKHR(device driver.VkDevice, pInfo *khr_buffer_device_address_driver.VkBufferDeviceAddressInfoKHR) driver.Uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkGetBufferOpaqueCaptureAddressKHR", device, pInfo)
	ret0, _ := ret[0].(driver.Uint64)
	return ret0
}

// VkGetBufferOpaqueCaptureAddressKHR indicates an expected call of VkGetBufferOpaqueCaptureAddressKHR.
func (mr *MockDriverMockRecorder) VkGetBufferOpaqueCaptureAddressKHR(device, pInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetBufferOpaqueCaptureAddressKHR", reflect.TypeOf((*MockDriver)(nil).VkGetBufferOpaqueCaptureAddressKHR), device, pInfo)
}

// VkGetDeviceMemoryOpaqueCaptureAddressKHR mocks base method.
func (m *MockDriver) VkGetDeviceMemoryOpaqueCaptureAddressKHR(device driver.VkDevice, pInfo *khr_buffer_device_address_driver.VkDeviceMemoryOpaqueCaptureAddressInfoKHR) driver.Uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkGetDeviceMemoryOpaqueCaptureAddressKHR", device, pInfo)
	ret0, _ := ret[0].(driver.Uint64)
	return ret0
}

// VkGetDeviceMemoryOpaqueCaptureAddressKHR indicates an expected call of VkGetDeviceMemoryOpaqueCaptureAddressKHR.
func (mr *MockDriverMockRecorder) VkGetDeviceMemoryOpaqueCaptureAddressKHR(device, pInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetDeviceMemoryOpaqueCaptureAddressKHR", reflect.TypeOf((*MockDriver)(nil).VkGetDeviceMemoryOpaqueCaptureAddressKHR), device, pInfo)
}
