// Code generated by MockGen. DO NOT EDIT.
// Source: extiface.go

// Package mock_external_semaphore_capabilities is a generated GoMock package.
package mock_external_semaphore_capabilities

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	core1_0 "github.com/vkngwrapper/core/v2/core1_0"
	khr_external_semaphore_capabilities "github.com/vkngwrapper/extensions/v2/khr_external_semaphore_capabilities"
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

// PhysicalDeviceExternalSemaphoreProperties mocks base method.
func (m *MockExtension) PhysicalDeviceExternalSemaphoreProperties(physicalDevice core1_0.PhysicalDevice, o khr_external_semaphore_capabilities.PhysicalDeviceExternalSemaphoreInfo, outData *khr_external_semaphore_capabilities.ExternalSemaphoreProperties) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PhysicalDeviceExternalSemaphoreProperties", physicalDevice, o, outData)
	ret0, _ := ret[0].(error)
	return ret0
}

// PhysicalDeviceExternalSemaphoreProperties indicates an expected call of PhysicalDeviceExternalSemaphoreProperties.
func (mr *MockExtensionMockRecorder) PhysicalDeviceExternalSemaphoreProperties(physicalDevice, o, outData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PhysicalDeviceExternalSemaphoreProperties", reflect.TypeOf((*MockExtension)(nil).PhysicalDeviceExternalSemaphoreProperties), physicalDevice, o, outData)
}
