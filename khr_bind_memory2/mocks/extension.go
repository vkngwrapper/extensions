// Code generated by MockGen. DO NOT EDIT.
// Source: extiface.go

// Package mock_bind_memory2 is a generated GoMock package.
package mock_bind_memory2

import (
	reflect "reflect"

	common "github.com/vkngwrapper/core/common"
	core1_0 "github.com/vkngwrapper/core/core1_0"
	khr_bind_memory2 "github.com/vkngwrapper/extensions/khr_bind_memory2"
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

// BindBufferMemory mocks base method.
func (m *MockExtension) BindBufferMemory2(device core1_0.Device, options []khr_bind_memory2.BindBufferMemoryInfo) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindBufferMemory2", device, options)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BindBufferMemory indicates an expected call of BindBufferMemory.
func (mr *MockExtensionMockRecorder) BindBufferMemory(device, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindBufferMemory2", reflect.TypeOf((*MockExtension)(nil).BindBufferMemory2), device, options)
}

// BindImageMemory mocks base method.
func (m *MockExtension) BindImageMemory2(device core1_0.Device, options []khr_bind_memory2.BindImageMemoryInfo) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindImageMemory2", device, options)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BindImageMemory indicates an expected call of BindImageMemory.
func (mr *MockExtensionMockRecorder) BindImageMemory(device, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindImageMemory2", reflect.TypeOf((*MockExtension)(nil).BindImageMemory2), device, options)
}
