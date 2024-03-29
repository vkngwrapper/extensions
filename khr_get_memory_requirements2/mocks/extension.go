// Code generated by MockGen. DO NOT EDIT.
// Source: extiface.go

// Package mock_get_memory_requirements2 is a generated GoMock package.
package mock_get_memory_requirements2

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	core1_0 "github.com/vkngwrapper/core/v2/core1_0"
	khr_get_memory_requirements2 "github.com/vkngwrapper/extensions/v2/khr_get_memory_requirements2"
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

// BufferMemoryRequirements2 mocks base method.
func (m *MockExtension) BufferMemoryRequirements2(device core1_0.Device, o khr_get_memory_requirements2.BufferMemoryRequirementsInfo2, out *khr_get_memory_requirements2.MemoryRequirements2) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BufferMemoryRequirements2", device, o, out)
	ret0, _ := ret[0].(error)
	return ret0
}

// BufferMemoryRequirements2 indicates an expected call of BufferMemoryRequirements2.
func (mr *MockExtensionMockRecorder) BufferMemoryRequirements2(device, o, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BufferMemoryRequirements2", reflect.TypeOf((*MockExtension)(nil).BufferMemoryRequirements2), device, o, out)
}

// ImageMemoryRequirements2 mocks base method.
func (m *MockExtension) ImageMemoryRequirements2(device core1_0.Device, o khr_get_memory_requirements2.ImageMemoryRequirementsInfo2, out *khr_get_memory_requirements2.MemoryRequirements2) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageMemoryRequirements2", device, o, out)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImageMemoryRequirements2 indicates an expected call of ImageMemoryRequirements2.
func (mr *MockExtensionMockRecorder) ImageMemoryRequirements2(device, o, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageMemoryRequirements2", reflect.TypeOf((*MockExtension)(nil).ImageMemoryRequirements2), device, o, out)
}

// ImageSparseMemoryRequirements2 mocks base method.
func (m *MockExtension) ImageSparseMemoryRequirements2(device core1_0.Device, o khr_get_memory_requirements2.ImageSparseMemoryRequirementsInfo2, outDataFactory func() *khr_get_memory_requirements2.SparseImageMemoryRequirements2) ([]*khr_get_memory_requirements2.SparseImageMemoryRequirements2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageSparseMemoryRequirements2", device, o, outDataFactory)
	ret0, _ := ret[0].([]*khr_get_memory_requirements2.SparseImageMemoryRequirements2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageSparseMemoryRequirements2 indicates an expected call of ImageSparseMemoryRequirements2.
func (mr *MockExtensionMockRecorder) ImageSparseMemoryRequirements2(device, o, outDataFactory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageSparseMemoryRequirements2", reflect.TypeOf((*MockExtension)(nil).ImageSparseMemoryRequirements2), device, o, outDataFactory)
}
