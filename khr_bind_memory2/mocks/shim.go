// Code generated by MockGen. DO NOT EDIT.
// Source: shim.go

// Package mock_bind_memory2 is a generated GoMock package.
package mock_bind_memory2

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/vkngwrapper/core/v2/common"
	core1_1 "github.com/vkngwrapper/core/v2/core1_1"
)

// MockShim is a mock of Shim interface.
type MockShim struct {
	ctrl     *gomock.Controller
	recorder *MockShimMockRecorder
}

// MockShimMockRecorder is the mock recorder for MockShim.
type MockShimMockRecorder struct {
	mock *MockShim
}

// NewMockShim creates a new mock instance.
func NewMockShim(ctrl *gomock.Controller) *MockShim {
	mock := &MockShim{ctrl: ctrl}
	mock.recorder = &MockShimMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShim) EXPECT() *MockShimMockRecorder {
	return m.recorder
}

// BindBufferMemory2 mocks base method.
func (m *MockShim) BindBufferMemory2(options []core1_1.BindBufferMemoryInfo) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindBufferMemory2", options)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BindBufferMemory2 indicates an expected call of BindBufferMemory2.
func (mr *MockShimMockRecorder) BindBufferMemory2(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindBufferMemory2", reflect.TypeOf((*MockShim)(nil).BindBufferMemory2), options)
}

// BindImageMemory2 mocks base method.
func (m *MockShim) BindImageMemory2(options []core1_1.BindImageMemoryInfo) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindImageMemory2", options)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BindImageMemory2 indicates an expected call of BindImageMemory2.
func (mr *MockShimMockRecorder) BindImageMemory2(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindImageMemory2", reflect.TypeOf((*MockShim)(nil).BindImageMemory2), options)
}
