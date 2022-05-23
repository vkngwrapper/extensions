// Code generated by MockGen. DO NOT EDIT.
// Source: extiface.go

// Package mock_get_physical_device_properties2 is a generated GoMock package.
package mock_get_physical_device_properties2

import (
	reflect "reflect"

	common "github.com/CannibalVox/VKng/core/common"
	core1_0 "github.com/CannibalVox/VKng/core/core1_0"
	khr_get_physical_device_properties2 "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2"
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

// PhysicalDeviceFeatures2 mocks base method.
func (m *MockExtension) PhysicalDeviceFeatures2(physicalDevice core1_0.PhysicalDevice, out *khr_get_physical_device_properties2.DeviceFeaturesOutData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PhysicalDeviceFeatures2", physicalDevice, out)
	ret0, _ := ret[0].(error)
	return ret0
}

// PhysicalDeviceFeatures2 indicates an expected call of PhysicalDeviceFeatures2.
func (mr *MockExtensionMockRecorder) PhysicalDeviceFeatures2(physicalDevice, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PhysicalDeviceFeatures2", reflect.TypeOf((*MockExtension)(nil).PhysicalDeviceFeatures2), physicalDevice, out)
}

// PhysicalDeviceFormatProperties2 mocks base method.
func (m *MockExtension) PhysicalDeviceFormatProperties2(physicalDevice core1_0.PhysicalDevice, format common.DataFormat, out *khr_get_physical_device_properties2.FormatPropertiesOutData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PhysicalDeviceFormatProperties2", physicalDevice, format, out)
	ret0, _ := ret[0].(error)
	return ret0
}

// PhysicalDeviceFormatProperties2 indicates an expected call of PhysicalDeviceFormatProperties2.
func (mr *MockExtensionMockRecorder) PhysicalDeviceFormatProperties2(physicalDevice, format, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PhysicalDeviceFormatProperties2", reflect.TypeOf((*MockExtension)(nil).PhysicalDeviceFormatProperties2), physicalDevice, format, out)
}

// PhysicalDeviceImageFormatProperties2 mocks base method.
func (m *MockExtension) PhysicalDeviceImageFormatProperties2(physicalDevice core1_0.PhysicalDevice, options khr_get_physical_device_properties2.ImageFormatOptions, out *khr_get_physical_device_properties2.ImageFormatPropertiesOutData) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PhysicalDeviceImageFormatProperties2", physicalDevice, options, out)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PhysicalDeviceImageFormatProperties2 indicates an expected call of PhysicalDeviceImageFormatProperties2.
func (mr *MockExtensionMockRecorder) PhysicalDeviceImageFormatProperties2(physicalDevice, options, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PhysicalDeviceImageFormatProperties2", reflect.TypeOf((*MockExtension)(nil).PhysicalDeviceImageFormatProperties2), physicalDevice, options, out)
}

// PhysicalDeviceMemoryProperties2 mocks base method.
func (m *MockExtension) PhysicalDeviceMemoryProperties2(physicalDevice core1_0.PhysicalDevice, out *khr_get_physical_device_properties2.MemoryPropertiesOutData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PhysicalDeviceMemoryProperties2", physicalDevice, out)
	ret0, _ := ret[0].(error)
	return ret0
}

// PhysicalDeviceMemoryProperties2 indicates an expected call of PhysicalDeviceMemoryProperties2.
func (mr *MockExtensionMockRecorder) PhysicalDeviceMemoryProperties2(physicalDevice, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PhysicalDeviceMemoryProperties2", reflect.TypeOf((*MockExtension)(nil).PhysicalDeviceMemoryProperties2), physicalDevice, out)
}

// PhysicalDeviceProperties2 mocks base method.
func (m *MockExtension) PhysicalDeviceProperties2(physicalDevice core1_0.PhysicalDevice, out *khr_get_physical_device_properties2.DevicePropertiesOutData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PhysicalDeviceProperties2", physicalDevice, out)
	ret0, _ := ret[0].(error)
	return ret0
}

// PhysicalDeviceProperties2 indicates an expected call of PhysicalDeviceProperties2.
func (mr *MockExtensionMockRecorder) PhysicalDeviceProperties2(physicalDevice, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PhysicalDeviceProperties2", reflect.TypeOf((*MockExtension)(nil).PhysicalDeviceProperties2), physicalDevice, out)
}

// PhysicalDeviceQueueFamilyProperties2 mocks base method.
func (m *MockExtension) PhysicalDeviceQueueFamilyProperties2(physicalDevice core1_0.PhysicalDevice, outDataFactory func() *khr_get_physical_device_properties2.QueueFamilyOutData) ([]*khr_get_physical_device_properties2.QueueFamilyOutData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PhysicalDeviceQueueFamilyProperties2", physicalDevice, outDataFactory)
	ret0, _ := ret[0].([]*khr_get_physical_device_properties2.QueueFamilyOutData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PhysicalDeviceQueueFamilyProperties2 indicates an expected call of PhysicalDeviceQueueFamilyProperties2.
func (mr *MockExtensionMockRecorder) PhysicalDeviceQueueFamilyProperties2(physicalDevice, outDataFactory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PhysicalDeviceQueueFamilyProperties2", reflect.TypeOf((*MockExtension)(nil).PhysicalDeviceQueueFamilyProperties2), physicalDevice, outDataFactory)
}

// PhysicalDeviceSparseImageFormatProperties2 mocks base method.
func (m *MockExtension) PhysicalDeviceSparseImageFormatProperties2(physicalDevice core1_0.PhysicalDevice, options khr_get_physical_device_properties2.SparseImageFormatOptions, outDataFactory func() *khr_get_physical_device_properties2.SparseImageFormatPropertiesOutData) ([]*khr_get_physical_device_properties2.SparseImageFormatPropertiesOutData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PhysicalDeviceSparseImageFormatProperties2", physicalDevice, options, outDataFactory)
	ret0, _ := ret[0].([]*khr_get_physical_device_properties2.SparseImageFormatPropertiesOutData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PhysicalDeviceSparseImageFormatProperties2 indicates an expected call of PhysicalDeviceSparseImageFormatProperties2.
func (mr *MockExtensionMockRecorder) PhysicalDeviceSparseImageFormatProperties2(physicalDevice, options, outDataFactory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PhysicalDeviceSparseImageFormatProperties2", reflect.TypeOf((*MockExtension)(nil).PhysicalDeviceSparseImageFormatProperties2), physicalDevice, options, outDataFactory)
}
