// Code generated by MockGen. DO NOT EDIT.
// Source: messenger.go

// Package mock_debugutils is a generated GoMock package.
package mock_debugutils

import (
	reflect "reflect"

	ext_debug_utils "github.com/CannibalVox/VKng/extensions/ext_debug_utils"
	gomock "github.com/golang/mock/gomock"
)

// MockMessenger is a mock of Messenger interface.
type MockMessenger struct {
	ctrl     *gomock.Controller
	recorder *MockMessengerMockRecorder
}

// MockMessengerMockRecorder is the mock recorder for MockMessenger.
type MockMessengerMockRecorder struct {
	mock *MockMessenger
}

// NewMockMessenger creates a new mock instance.
func NewMockMessenger(ctrl *gomock.Controller) *MockMessenger {
	mock := &MockMessenger{ctrl: ctrl}
	mock.recorder = &MockMessengerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessenger) EXPECT() *MockMessengerMockRecorder {
	return m.recorder
}

// Destroy mocks base method.
func (m *MockMessenger) Destroy() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Destroy")
}

// Destroy indicates an expected call of Destroy.
func (mr *MockMessengerMockRecorder) Destroy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockMessenger)(nil).Destroy))
}

// Handle mocks base method.
func (m *MockMessenger) Handle() ext_debug_utils.VkDebugUtilsMessengerEXT {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle")
	ret0, _ := ret[0].(ext_debug_utils.VkDebugUtilsMessengerEXT)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockMessengerMockRecorder) Handle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockMessenger)(nil).Handle))
}
