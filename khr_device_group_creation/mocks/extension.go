// Code generated by MockGen. DO NOT EDIT.
// Source: extiface.go

// Package mock_device_group_creation is a generated GoMock package.
package mock_device_group_creation

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/vkngwrapper/core/v2/common"
	core1_0 "github.com/vkngwrapper/core/v2/core1_0"
	khr_device_group_creation "github.com/vkngwrapper/extensions/v2/khr_device_group_creation"
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

// EnumeratePhysicalDeviceGroups mocks base method.
func (m *MockExtension) EnumeratePhysicalDeviceGroups(instance core1_0.Instance, outDataFactory func() *khr_device_group_creation.PhysicalDeviceGroupProperties) ([]*khr_device_group_creation.PhysicalDeviceGroupProperties, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnumeratePhysicalDeviceGroups", instance, outDataFactory)
	ret0, _ := ret[0].([]*khr_device_group_creation.PhysicalDeviceGroupProperties)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EnumeratePhysicalDeviceGroups indicates an expected call of EnumeratePhysicalDeviceGroups.
func (mr *MockExtensionMockRecorder) EnumeratePhysicalDeviceGroups(instance, outDataFactory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnumeratePhysicalDeviceGroups", reflect.TypeOf((*MockExtension)(nil).EnumeratePhysicalDeviceGroups), instance, outDataFactory)
}
