// Code generated by MockGen. DO NOT EDIT.
// Source: iface.go

// Package mock_descriptor_update_template is a generated GoMock package.
package mock_descriptor_update_template

import (
	reflect "reflect"

	core1_0 "github.com/vkngwrapper/core/core1_0"
	driver "github.com/vkngwrapper/core/driver"
	khr_descriptor_update_template "github.com/vkngwrapper/extensions/khr_descriptor_update_template"
	khr_descriptor_update_template_driver "github.com/vkngwrapper/extensions/khr_descriptor_update_template/driver"
	gomock "github.com/golang/mock/gomock"
)

// MockDescriptorUpdateTemplate is a mock of DescriptorUpdateTemplate interface.
type MockDescriptorUpdateTemplate struct {
	ctrl     *gomock.Controller
	recorder *MockDescriptorUpdateTemplateMockRecorder
}

// MockDescriptorUpdateTemplateMockRecorder is the mock recorder for MockDescriptorUpdateTemplate.
type MockDescriptorUpdateTemplateMockRecorder struct {
	mock *MockDescriptorUpdateTemplate
}

// NewMockDescriptorUpdateTemplate creates a new mock instance.
func NewMockDescriptorUpdateTemplate(ctrl *gomock.Controller) *MockDescriptorUpdateTemplate {
	mock := &MockDescriptorUpdateTemplate{ctrl: ctrl}
	mock.recorder = &MockDescriptorUpdateTemplateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDescriptorUpdateTemplate) EXPECT() *MockDescriptorUpdateTemplateMockRecorder {
	return m.recorder
}

// Destroy mocks base method.
func (m *MockDescriptorUpdateTemplate) Destroy(allocator *driver.AllocationCallbacks) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Destroy", allocator)
}

// Destroy indicates an expected call of Destroy.
func (mr *MockDescriptorUpdateTemplateMockRecorder) Destroy(allocator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockDescriptorUpdateTemplate)(nil).Destroy), allocator)
}

// Handle mocks base method.
func (m *MockDescriptorUpdateTemplate) Handle() khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle")
	ret0, _ := ret[0].(khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockDescriptorUpdateTemplateMockRecorder) Handle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockDescriptorUpdateTemplate)(nil).Handle))
}

// UpdateDescriptorSetFromBuffer mocks base method.
func (m *MockDescriptorUpdateTemplate) UpdateDescriptorSetFromBuffer(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorBufferInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateDescriptorSetFromBuffer", descriptorSet, data)
}

// UpdateDescriptorSetFromBuffer indicates an expected call of UpdateDescriptorSetFromBuffer.
func (mr *MockDescriptorUpdateTemplateMockRecorder) UpdateDescriptorSetFromBuffer(descriptorSet, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDescriptorSetFromBuffer", reflect.TypeOf((*MockDescriptorUpdateTemplate)(nil).UpdateDescriptorSetFromBuffer), descriptorSet, data)
}

// UpdateDescriptorSetFromImage mocks base method.
func (m *MockDescriptorUpdateTemplate) UpdateDescriptorSetFromImage(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorImageInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateDescriptorSetFromImage", descriptorSet, data)
}

// UpdateDescriptorSetFromImage indicates an expected call of UpdateDescriptorSetFromImage.
func (mr *MockDescriptorUpdateTemplateMockRecorder) UpdateDescriptorSetFromImage(descriptorSet, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDescriptorSetFromImage", reflect.TypeOf((*MockDescriptorUpdateTemplate)(nil).UpdateDescriptorSetFromImage), descriptorSet, data)
}

// UpdateDescriptorSetFromObjectHandle mocks base method.
func (m *MockDescriptorUpdateTemplate) UpdateDescriptorSetFromObjectHandle(descriptorSet core1_0.DescriptorSet, data driver.VulkanHandle) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateDescriptorSetFromObjectHandle", descriptorSet, data)
}

// UpdateDescriptorSetFromObjectHandle indicates an expected call of UpdateDescriptorSetFromObjectHandle.
func (mr *MockDescriptorUpdateTemplateMockRecorder) UpdateDescriptorSetFromObjectHandle(descriptorSet, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDescriptorSetFromObjectHandle", reflect.TypeOf((*MockDescriptorUpdateTemplate)(nil).UpdateDescriptorSetFromObjectHandle), descriptorSet, data)
}

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

// CreateDescriptorUpdateTemplate mocks base method.
func (m *MockExtension) CreateDescriptorUpdateTemplate(device core1_0.Device, o khr_descriptor_update_template.DescriptorUpdateTemplateCreateInfo, allocator *driver.AllocationCallbacks) (khr_descriptor_update_template.DescriptorUpdateTemplate, driver.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDescriptorUpdateTemplate", device, o, allocator)
	ret0, _ := ret[0].(khr_descriptor_update_template.DescriptorUpdateTemplate)
	ret1, _ := ret[1].(driver.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateDescriptorUpdateTemplate indicates an expected call of CreateDescriptorUpdateTemplate.
func (mr *MockExtensionMockRecorder) CreateDescriptorUpdateTemplate(device, o, allocator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDescriptorUpdateTemplate", reflect.TypeOf((*MockExtension)(nil).CreateDescriptorUpdateTemplate), device, o, allocator)
}
