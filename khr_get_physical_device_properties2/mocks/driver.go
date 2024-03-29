// Code generated by MockGen. DO NOT EDIT.
// Source: driver.go

// Package mock_get_physical_device_properties2 is a generated GoMock package.
package mock_get_physical_device_properties2

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/vkngwrapper/core/v2/common"
	driver "github.com/vkngwrapper/core/v2/driver"
	khr_get_physical_device_properties2_driver "github.com/vkngwrapper/extensions/v2/khr_get_physical_device_properties2/driver"
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

// VkGetPhysicalDeviceFeatures2KHR mocks base method.
func (m *MockDriver) VkGetPhysicalDeviceFeatures2KHR(physicalDevice driver.VkPhysicalDevice, pFeatures *khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkGetPhysicalDeviceFeatures2KHR", physicalDevice, pFeatures)
}

// VkGetPhysicalDeviceFeatures2KHR indicates an expected call of VkGetPhysicalDeviceFeatures2KHR.
func (mr *MockDriverMockRecorder) VkGetPhysicalDeviceFeatures2KHR(physicalDevice, pFeatures interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetPhysicalDeviceFeatures2KHR", reflect.TypeOf((*MockDriver)(nil).VkGetPhysicalDeviceFeatures2KHR), physicalDevice, pFeatures)
}

// VkGetPhysicalDeviceFormatProperties2KHR mocks base method.
func (m *MockDriver) VkGetPhysicalDeviceFormatProperties2KHR(physicalDevice driver.VkPhysicalDevice, format driver.VkFormat, pFormatProperties *khr_get_physical_device_properties2_driver.VkFormatProperties2KHR) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkGetPhysicalDeviceFormatProperties2KHR", physicalDevice, format, pFormatProperties)
}

// VkGetPhysicalDeviceFormatProperties2KHR indicates an expected call of VkGetPhysicalDeviceFormatProperties2KHR.
func (mr *MockDriverMockRecorder) VkGetPhysicalDeviceFormatProperties2KHR(physicalDevice, format, pFormatProperties interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetPhysicalDeviceFormatProperties2KHR", reflect.TypeOf((*MockDriver)(nil).VkGetPhysicalDeviceFormatProperties2KHR), physicalDevice, format, pFormatProperties)
}

// VkGetPhysicalDeviceImageFormatProperties2KHR mocks base method.
func (m *MockDriver) VkGetPhysicalDeviceImageFormatProperties2KHR(physicalDevice driver.VkPhysicalDevice, pImageFormatInfo *khr_get_physical_device_properties2_driver.VkPhysicalDeviceImageFormatInfo2KHR, pImageFormatProperties *khr_get_physical_device_properties2_driver.VkImageFormatProperties2KHR) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkGetPhysicalDeviceImageFormatProperties2KHR", physicalDevice, pImageFormatInfo, pImageFormatProperties)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkGetPhysicalDeviceImageFormatProperties2KHR indicates an expected call of VkGetPhysicalDeviceImageFormatProperties2KHR.
func (mr *MockDriverMockRecorder) VkGetPhysicalDeviceImageFormatProperties2KHR(physicalDevice, pImageFormatInfo, pImageFormatProperties interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetPhysicalDeviceImageFormatProperties2KHR", reflect.TypeOf((*MockDriver)(nil).VkGetPhysicalDeviceImageFormatProperties2KHR), physicalDevice, pImageFormatInfo, pImageFormatProperties)
}

// VkGetPhysicalDeviceMemoryProperties2KHR mocks base method.
func (m *MockDriver) VkGetPhysicalDeviceMemoryProperties2KHR(physicalDevice driver.VkPhysicalDevice, pMemoryProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceMemoryProperties2KHR) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkGetPhysicalDeviceMemoryProperties2KHR", physicalDevice, pMemoryProperties)
}

// VkGetPhysicalDeviceMemoryProperties2KHR indicates an expected call of VkGetPhysicalDeviceMemoryProperties2KHR.
func (mr *MockDriverMockRecorder) VkGetPhysicalDeviceMemoryProperties2KHR(physicalDevice, pMemoryProperties interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetPhysicalDeviceMemoryProperties2KHR", reflect.TypeOf((*MockDriver)(nil).VkGetPhysicalDeviceMemoryProperties2KHR), physicalDevice, pMemoryProperties)
}

// VkGetPhysicalDeviceProperties2KHR mocks base method.
func (m *MockDriver) VkGetPhysicalDeviceProperties2KHR(physicalDevice driver.VkPhysicalDevice, pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkGetPhysicalDeviceProperties2KHR", physicalDevice, pProperties)
}

// VkGetPhysicalDeviceProperties2KHR indicates an expected call of VkGetPhysicalDeviceProperties2KHR.
func (mr *MockDriverMockRecorder) VkGetPhysicalDeviceProperties2KHR(physicalDevice, pProperties interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetPhysicalDeviceProperties2KHR", reflect.TypeOf((*MockDriver)(nil).VkGetPhysicalDeviceProperties2KHR), physicalDevice, pProperties)
}

// VkGetPhysicalDeviceQueueFamilyProperties2KHR mocks base method.
func (m *MockDriver) VkGetPhysicalDeviceQueueFamilyProperties2KHR(physicalDevice driver.VkPhysicalDevice, pQueueFamilyPropertyCount *driver.Uint32, pQueueFamilyProperties *khr_get_physical_device_properties2_driver.VkQueueFamilyProperties2KHR) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkGetPhysicalDeviceQueueFamilyProperties2KHR", physicalDevice, pQueueFamilyPropertyCount, pQueueFamilyProperties)
}

// VkGetPhysicalDeviceQueueFamilyProperties2KHR indicates an expected call of VkGetPhysicalDeviceQueueFamilyProperties2KHR.
func (mr *MockDriverMockRecorder) VkGetPhysicalDeviceQueueFamilyProperties2KHR(physicalDevice, pQueueFamilyPropertyCount, pQueueFamilyProperties interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetPhysicalDeviceQueueFamilyProperties2KHR", reflect.TypeOf((*MockDriver)(nil).VkGetPhysicalDeviceQueueFamilyProperties2KHR), physicalDevice, pQueueFamilyPropertyCount, pQueueFamilyProperties)
}

// VkGetPhysicalDeviceSparseImageFormatProperties2KHR mocks base method.
func (m *MockDriver) VkGetPhysicalDeviceSparseImageFormatProperties2KHR(physicalDevice driver.VkPhysicalDevice, pFormatInfo *khr_get_physical_device_properties2_driver.VkPhysicalDeviceSparseImageFormatInfo2KHR, pPropertyCount *driver.Uint32, pProperties *khr_get_physical_device_properties2_driver.VkSparseImageFormatProperties2KHR) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkGetPhysicalDeviceSparseImageFormatProperties2KHR", physicalDevice, pFormatInfo, pPropertyCount, pProperties)
}

// VkGetPhysicalDeviceSparseImageFormatProperties2KHR indicates an expected call of VkGetPhysicalDeviceSparseImageFormatProperties2KHR.
func (mr *MockDriverMockRecorder) VkGetPhysicalDeviceSparseImageFormatProperties2KHR(physicalDevice, pFormatInfo, pPropertyCount, pProperties interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetPhysicalDeviceSparseImageFormatProperties2KHR", reflect.TypeOf((*MockDriver)(nil).VkGetPhysicalDeviceSparseImageFormatProperties2KHR), physicalDevice, pFormatInfo, pPropertyCount, pProperties)
}
