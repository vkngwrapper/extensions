// Code generated by MockGen. DO NOT EDIT.
// Source: surface.go

// Package mock_surface is a generated GoMock package.
package mock_surface

import (
	reflect "reflect"

	loader "github.com/CannibalVox/VKng/core/loader"
	resource "github.com/CannibalVox/VKng/core/resource"
	ext_surface "github.com/CannibalVox/VKng/extensions/surface"
	cgoalloc "github.com/CannibalVox/cgoalloc"
	gomock "github.com/golang/mock/gomock"
)

// MockSurface is a mock of Surface interface.
type MockSurface struct {
	ctrl     *gomock.Controller
	recorder *MockSurfaceMockRecorder
}

// MockSurfaceMockRecorder is the mock recorder for MockSurface.
type MockSurfaceMockRecorder struct {
	mock *MockSurface
}

// NewMockSurface creates a new mock instance.
func NewMockSurface(ctrl *gomock.Controller) *MockSurface {
	mock := &MockSurface{ctrl: ctrl}
	mock.recorder = &MockSurfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSurface) EXPECT() *MockSurfaceMockRecorder {
	return m.recorder
}

// Capabilities mocks base method.
func (m *MockSurface) Capabilities(allocator cgoalloc.Allocator, device resource.PhysicalDevice) (*ext_surface.Capabilities, loader.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Capabilities", allocator, device)
	ret0, _ := ret[0].(*ext_surface.Capabilities)
	ret1, _ := ret[1].(loader.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Capabilities indicates an expected call of Capabilities.
func (mr *MockSurfaceMockRecorder) Capabilities(allocator, device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capabilities", reflect.TypeOf((*MockSurface)(nil).Capabilities), allocator, device)
}

// Destroy mocks base method.
func (m *MockSurface) Destroy() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Destroy")
}

// Destroy indicates an expected call of Destroy.
func (mr *MockSurfaceMockRecorder) Destroy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockSurface)(nil).Destroy))
}

// Formats mocks base method.
func (m *MockSurface) Formats(allocator cgoalloc.Allocator, device resource.PhysicalDevice) ([]ext_surface.Format, loader.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Formats", allocator, device)
	ret0, _ := ret[0].([]ext_surface.Format)
	ret1, _ := ret[1].(loader.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Formats indicates an expected call of Formats.
func (mr *MockSurfaceMockRecorder) Formats(allocator, device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Formats", reflect.TypeOf((*MockSurface)(nil).Formats), allocator, device)
}

// Handle mocks base method.
func (m *MockSurface) Handle() ext_surface.Handle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle")
	ret0, _ := ret[0].(ext_surface.Handle)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockSurfaceMockRecorder) Handle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockSurface)(nil).Handle))
}

// PresentModes mocks base method.
func (m *MockSurface) PresentModes(allocator cgoalloc.Allocator, device resource.PhysicalDevice) ([]ext_surface.PresentMode, loader.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PresentModes", allocator, device)
	ret0, _ := ret[0].([]ext_surface.PresentMode)
	ret1, _ := ret[1].(loader.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PresentModes indicates an expected call of PresentModes.
func (mr *MockSurfaceMockRecorder) PresentModes(allocator, device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PresentModes", reflect.TypeOf((*MockSurface)(nil).PresentModes), allocator, device)
}

// SupportsDevice mocks base method.
func (m *MockSurface) SupportsDevice(physicalDevice resource.PhysicalDevice, queueFamilyIndex int) (bool, loader.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupportsDevice", physicalDevice, queueFamilyIndex)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(loader.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SupportsDevice indicates an expected call of SupportsDevice.
func (mr *MockSurfaceMockRecorder) SupportsDevice(physicalDevice, queueFamilyIndex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportsDevice", reflect.TypeOf((*MockSurface)(nil).SupportsDevice), physicalDevice, queueFamilyIndex)
}
