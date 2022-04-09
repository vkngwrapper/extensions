// Code generated by MockGen. DO NOT EDIT.
// Source: iface.go

// Package mock_maintenance3 is a generated GoMock package.
package mock_maintenance3

import (
	reflect "reflect"

	core1_0 "github.com/CannibalVox/VKng/core/core1_0"
	khr_maintenance3 "github.com/CannibalVox/VKng/extensions/khr_maintenance3"
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

// DescriptorSetLayoutSupport mocks base method.
func (m *MockExtension) DescriptorSetLayoutSupport(device core1_0.Device, setLayoutOptions core1_0.DescriptorSetLayoutOptions, support *khr_maintenance3.DescriptorSetLayoutSupportOutData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescriptorSetLayoutSupport", device, setLayoutOptions, support)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescriptorSetLayoutSupport indicates an expected call of DescriptorSetLayoutSupport.
func (mr *MockExtensionMockRecorder) DescriptorSetLayoutSupport(device, setLayoutOptions, support interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescriptorSetLayoutSupport", reflect.TypeOf((*MockExtension)(nil).DescriptorSetLayoutSupport), device, setLayoutOptions, support)
}
