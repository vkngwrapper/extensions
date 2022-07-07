// Code generated by MockGen. DO NOT EDIT.
// Source: driver.go

// Package mock_maintenance3 is a generated GoMock package.
package mock_maintenance3

import (
	reflect "reflect"

	driver "github.com/vkngwrapper/core/driver"
	khr_maintenance3_driver "github.com/vkngwrapper/extensions/khr_maintenance3/driver"
	gomock "github.com/golang/mock/gomock"
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

// VkGetDescriptorSetLayoutSupportKHR dummies base method.
func (m *MockDriver) VkGetDescriptorSetLayoutSupportKHR(device driver.VkDevice, pCreateInfo *driver.VkDescriptorSetLayoutCreateInfo, pSupport *khr_maintenance3_driver.VkDescriptorSetLayoutSupportKHR) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkGetDescriptorSetLayoutSupportKHR", device, pCreateInfo, pSupport)
}

// VkGetDescriptorSetLayoutSupportKHR indicates an expected call of VkGetDescriptorSetLayoutSupportKHR.
func (mr *MockDriverMockRecorder) VkGetDescriptorSetLayoutSupportKHR(device, pCreateInfo, pSupport any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetDescriptorSetLayoutSupportKHR", reflect.TypeOf((*MockDriver)(nil).VkGetDescriptorSetLayoutSupportKHR), device, pCreateInfo, pSupport)
}
