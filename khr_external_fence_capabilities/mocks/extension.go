// Code generated by MockGen. DO NOT EDIT.
// Source: extiface.go

// Package mock_external_fence_capabilities is a generated GoMock package.
package mock_external_fence_capabilities

import (
	reflect "reflect"

	core1_0 "github.com/CannibalVox/VKng/core/core1_0"
	khr_external_fence_capabilities "github.com/CannibalVox/VKng/extensions/khr_external_fence_capabilities"
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

// PhysicalDeviceExternalFenceProperties mocks base method.
func (m *MockExtension) ExternalFenceProperties(physicalDevice core1_0.PhysicalDevice, o khr_external_fence_capabilities.ExternalFencePropertiesOptions, outData *khr_external_fence_capabilities.ExternalFencePropertiesOutData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExternalFenceProperties", physicalDevice, o, outData)
	ret0, _ := ret[0].(error)
	return ret0
}

// PhysicalDeviceExternalFenceProperties indicates an expected call of PhysicalDeviceExternalFenceProperties.
func (mr *MockExtensionMockRecorder) PhysicalDeviceExternalFenceProperties(physicalDevice, o, outData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExternalFenceProperties", reflect.TypeOf((*MockExtension)(nil).ExternalFenceProperties), physicalDevice, o, outData)
}