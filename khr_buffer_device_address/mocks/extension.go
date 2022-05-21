// Code generated by MockGen. DO NOT EDIT.
// Source: extiface.go

// Package mock_buffer_device_address is a generated GoMock package.
package mock_buffer_device_address

import (
	reflect "reflect"

	core1_0 "github.com/CannibalVox/VKng/core/core1_0"
	khr_buffer_device_address "github.com/CannibalVox/VKng/extensions/khr_buffer_device_address"
	gomock "github.com/golang/mock/gomock"
)

// MockExtension is a mock of Extension interface.
type MockExtension struct {
	ctrl     *gomock.Controller
	recorder *MockExtensionMockRecorder
}

// MockExtensionMockRecorder is the mock recorder for MockExtension.
type MockExtensionMockRecorder struct {
	mock *MockExtension
}

// NewMockExtension creates a new mock instance.
func NewMockExtension(ctrl *gomock.Controller) *MockExtension {
	mock := &MockExtension{ctrl: ctrl}
	mock.recorder = &MockExtensionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExtension) EXPECT() *MockExtensionMockRecorder {
	return m.recorder
}

// GetBufferDeviceAddress mocks base method.
func (m *MockExtension) GetBufferDeviceAddress(device core1_0.Device, o khr_buffer_device_address.BufferDeviceAddressOptions) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBufferDeviceAddress", device, o)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBufferDeviceAddress indicates an expected call of GetBufferDeviceAddress.
func (mr *MockExtensionMockRecorder) GetBufferDeviceAddress(device, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBufferDeviceAddress", reflect.TypeOf((*MockExtension)(nil).GetBufferDeviceAddress), device, o)
}

// GetBufferOpaqueCaptureAddress mocks base method.
func (m *MockExtension) GetBufferOpaqueCaptureAddress(device core1_0.Device, o khr_buffer_device_address.BufferDeviceAddressOptions) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBufferOpaqueCaptureAddress", device, o)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBufferOpaqueCaptureAddress indicates an expected call of GetBufferOpaqueCaptureAddress.
func (mr *MockExtensionMockRecorder) GetBufferOpaqueCaptureAddress(device, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBufferOpaqueCaptureAddress", reflect.TypeOf((*MockExtension)(nil).GetBufferOpaqueCaptureAddress), device, o)
}

// GetDeviceMemoryOpaqueCaptureAddress mocks base method.
func (m *MockExtension) GetDeviceMemoryOpaqueCaptureAddress(device core1_0.Device, o khr_buffer_device_address.DeviceMemoryOpaqueAddressOptions) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceMemoryOpaqueCaptureAddress", device, o)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceMemoryOpaqueCaptureAddress indicates an expected call of GetDeviceMemoryOpaqueCaptureAddress.
func (mr *MockExtensionMockRecorder) GetDeviceMemoryOpaqueCaptureAddress(device, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceMemoryOpaqueCaptureAddress", reflect.TypeOf((*MockExtension)(nil).GetDeviceMemoryOpaqueCaptureAddress), device, o)
}
