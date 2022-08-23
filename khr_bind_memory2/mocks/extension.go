// Code generated by MockGen. DO NOT EDIT.
// Source: extiface.go

// Package mock_bind_memory2 is a generated GoMock package.
package mock_bind_memory2

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/vkngwrapper/core/v2/common"
	core1_0 "github.com/vkngwrapper/core/v2/core1_0"
	khr_bind_memory2 "github.com/vkngwrapper/extensions/v2/khr_bind_memory2"
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

// BindBufferMemory2 mocks base method.
func (m *MockExtension) BindBufferMemory2(device core1_0.Device, options []khr_bind_memory2.BindBufferMemoryInfo) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindBufferMemory2", device, options)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BindBufferMemory2 indicates an expected call of BindBufferMemory2.
func (mr *MockExtensionMockRecorder) BindBufferMemory2(device, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindBufferMemory2", reflect.TypeOf((*MockExtension)(nil).BindBufferMemory2), device, options)
}

// BindImageMemory2 mocks base method.
func (m *MockExtension) BindImageMemory2(device core1_0.Device, options []khr_bind_memory2.BindImageMemoryInfo) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindImageMemory2", device, options)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BindImageMemory2 indicates an expected call of BindImageMemory2.
func (mr *MockExtensionMockRecorder) BindImageMemory2(device, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindImageMemory2", reflect.TypeOf((*MockExtension)(nil).BindImageMemory2), device, options)
}
