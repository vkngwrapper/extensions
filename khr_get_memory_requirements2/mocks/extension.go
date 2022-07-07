// Code generated by MockGen. DO NOT EDIT.
// Source: extiface.go

// Package mock_get_memory_requirements2 is a generated GoMock package.
package mock_get_memory_requirements2

import (
	reflect "reflect"

	core1_0 "github.com/vkngwrapper/core/core1_0"
	khr_get_memory_requirements2 "github.com/vkngwrapper/extensions/khr_get_memory_requirements2"
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

// BufferMemoryRequirements mocks base method.
func (m *MockExtension) BufferMemoryRequirements2(device core1_0.Device, o khr_get_memory_requirements2.BufferMemoryRequirementsInfo2, out *khr_get_memory_requirements2.MemoryRequirements2) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BufferMemoryRequirements2", device, o, out)
	ret0, _ := ret[0].(error)
	return ret0
}

// BufferMemoryRequirements indicates an expected call of BufferMemoryRequirements.
func (mr *MockExtensionMockRecorder) BufferMemoryRequirements(device, o, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BufferMemoryRequirements2", reflect.TypeOf((*MockExtension)(nil).BufferMemoryRequirements2), device, o, out)
}

// ImageMemoryRequirements mocks base method.
func (m *MockExtension) ImageMemoryRequirements2(device core1_0.Device, o khr_get_memory_requirements2.ImageMemoryRequirementsInfo2, out *khr_get_memory_requirements2.MemoryRequirements2) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageMemoryRequirements2", device, o, out)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImageMemoryRequirements indicates an expected call of ImageMemoryRequirements.
func (mr *MockExtensionMockRecorder) ImageMemoryRequirements(device, o, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageMemoryRequirements2", reflect.TypeOf((*MockExtension)(nil).ImageMemoryRequirements2), device, o, out)
}

// SparseImageMemoryRequirements mocks base method.
func (m *MockExtension) ImageSparseMemoryRequirements2(device core1_0.Device, o khr_get_memory_requirements2.ImageSparseMemoryRequirementsInfo2, outDataFactory func() *khr_get_memory_requirements2.SparseImageMemoryRequirements2) ([]*khr_get_memory_requirements2.SparseImageMemoryRequirements2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageSparseMemoryRequirements2", device, o, outDataFactory)
	ret0, _ := ret[0].([]*khr_get_memory_requirements2.SparseImageMemoryRequirements2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SparseImageMemoryRequirements indicates an expected call of SparseImageMemoryRequirements.
func (mr *MockExtensionMockRecorder) SparseImageMemoryRequirements(device, o, outDataFactory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageSparseMemoryRequirements2", reflect.TypeOf((*MockExtension)(nil).ImageSparseMemoryRequirements2), device, o, outDataFactory)
}
