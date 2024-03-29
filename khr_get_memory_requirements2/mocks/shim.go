// Code generated by MockGen. DO NOT EDIT.
// Source: shim.go

// Package mock_get_memory_requirements2 is a generated GoMock package.
package mock_get_memory_requirements2

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
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

// BufferMemoryRequirements2 mocks base method.
func (m *MockShim) BufferMemoryRequirements2(o core1_1.BufferMemoryRequirementsInfo2, out *core1_1.MemoryRequirements2) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BufferMemoryRequirements2", o, out)
	ret0, _ := ret[0].(error)
	return ret0
}

// BufferMemoryRequirements2 indicates an expected call of BufferMemoryRequirements2.
func (mr *MockShimMockRecorder) BufferMemoryRequirements2(o, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BufferMemoryRequirements2", reflect.TypeOf((*MockShim)(nil).BufferMemoryRequirements2), o, out)
}

// ImageMemoryRequirements2 mocks base method.
func (m *MockShim) ImageMemoryRequirements2(o core1_1.ImageMemoryRequirementsInfo2, out *core1_1.MemoryRequirements2) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageMemoryRequirements2", o, out)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImageMemoryRequirements2 indicates an expected call of ImageMemoryRequirements2.
func (mr *MockShimMockRecorder) ImageMemoryRequirements2(o, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageMemoryRequirements2", reflect.TypeOf((*MockShim)(nil).ImageMemoryRequirements2), o, out)
}

// ImageSparseMemoryRequirements2 mocks base method.
func (m *MockShim) ImageSparseMemoryRequirements2(o core1_1.ImageSparseMemoryRequirementsInfo2, outDataFactory func() *core1_1.SparseImageMemoryRequirements2) ([]*core1_1.SparseImageMemoryRequirements2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageSparseMemoryRequirements2", o, outDataFactory)
	ret0, _ := ret[0].([]*core1_1.SparseImageMemoryRequirements2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageSparseMemoryRequirements2 indicates an expected call of ImageSparseMemoryRequirements2.
func (mr *MockShimMockRecorder) ImageSparseMemoryRequirements2(o, outDataFactory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageSparseMemoryRequirements2", reflect.TypeOf((*MockShim)(nil).ImageSparseMemoryRequirements2), o, outDataFactory)
}
