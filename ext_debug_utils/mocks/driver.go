// Code generated by MockGen. DO NOT EDIT.
// Source: driver.go

// Package mock_debugutils is a generated GoMock package.
package mock_debugutils

import (
	reflect "reflect"

	common "github.com/CannibalVox/VKng/core/common"
	driver "github.com/CannibalVox/VKng/core/driver"
	ext_debug_utils_driver "github.com/CannibalVox/VKng/extensions/ext_debug_utils/driver"
	gomock "github.com/golang/mock/gomock"
)

// MockDriver is a mock of Driver interface.
type MockDriver struct {
	ctrl     *gomock.Controller
	recorder *MockDriverMockRecorder
}

// MockDriverMockRecorder is the mock recorder for MockDriver.
type MockDriverMockRecorder struct {
	mock *MockDriver
}

// NewMockDriver creates a new mock instance.
func NewMockDriver(ctrl *gomock.Controller) *MockDriver {
	mock := &MockDriver{ctrl: ctrl}
	mock.recorder = &MockDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDriver) EXPECT() *MockDriverMockRecorder {
	return m.recorder
}

// VKCmdBeginDebugUtilsLabelEXT mocks base method.
func (m *MockDriver) VKCmdBeginDebugUtilsLabelEXT(commandBuffer driver.VkCommandBuffer, pLabelInfo *ext_debug_utils_driver.VkDebugUtilsLabelEXT) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VKCmdBeginDebugUtilsLabelEXT", commandBuffer, pLabelInfo)
}

// VKCmdBeginDebugUtilsLabelEXT indicates an expected call of VKCmdBeginDebugUtilsLabelEXT.
func (mr *MockDriverMockRecorder) VKCmdBeginDebugUtilsLabelEXT(commandBuffer, pLabelInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VKCmdBeginDebugUtilsLabelEXT", reflect.TypeOf((*MockDriver)(nil).VKCmdBeginDebugUtilsLabelEXT), commandBuffer, pLabelInfo)
}

// VkCmdEndDebugUtilsLabelEXT mocks base method.
func (m *MockDriver) VkCmdEndDebugUtilsLabelEXT(commandBuffer driver.VkCommandBuffer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkCmdEndDebugUtilsLabelEXT", commandBuffer)
}

// VkCmdEndDebugUtilsLabelEXT indicates an expected call of VkCmdEndDebugUtilsLabelEXT.
func (mr *MockDriverMockRecorder) VkCmdEndDebugUtilsLabelEXT(commandBuffer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkCmdEndDebugUtilsLabelEXT", reflect.TypeOf((*MockDriver)(nil).VkCmdEndDebugUtilsLabelEXT), commandBuffer)
}

// VkCmdInsertDebugUtilsLabelEXT mocks base method.
func (m *MockDriver) VkCmdInsertDebugUtilsLabelEXT(commandBuffer driver.VkCommandBuffer, pLabelInfo *ext_debug_utils_driver.VkDebugUtilsLabelEXT) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkCmdInsertDebugUtilsLabelEXT", commandBuffer, pLabelInfo)
}

// VkCmdInsertDebugUtilsLabelEXT indicates an expected call of VkCmdInsertDebugUtilsLabelEXT.
func (mr *MockDriverMockRecorder) VkCmdInsertDebugUtilsLabelEXT(commandBuffer, pLabelInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkCmdInsertDebugUtilsLabelEXT", reflect.TypeOf((*MockDriver)(nil).VkCmdInsertDebugUtilsLabelEXT), commandBuffer, pLabelInfo)
}

// VkCreateDebugUtilsMessengerEXT mocks base method.
func (m *MockDriver) VkCreateDebugUtilsMessengerEXT(instance driver.VkInstance, pCreateInfo *ext_debug_utils_driver.VkDebugUtilsMessengerCreateInfoEXT, pAllocator *driver.VkAllocationCallbacks, pDebugMessenger *ext_debug_utils_driver.VkDebugUtilsMessengerEXT) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkCreateDebugUtilsMessengerEXT", instance, pCreateInfo, pAllocator, pDebugMessenger)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkCreateDebugUtilsMessengerEXT indicates an expected call of VkCreateDebugUtilsMessengerEXT.
func (mr *MockDriverMockRecorder) VkCreateDebugUtilsMessengerEXT(instance, pCreateInfo, pAllocator, pDebugMessenger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkCreateDebugUtilsMessengerEXT", reflect.TypeOf((*MockDriver)(nil).VkCreateDebugUtilsMessengerEXT), instance, pCreateInfo, pAllocator, pDebugMessenger)
}

// VkDestroyDebugUtilsMessengerEXT mocks base method.
func (m *MockDriver) VkDestroyDebugUtilsMessengerEXT(instance driver.VkInstance, debugMessenger ext_debug_utils_driver.VkDebugUtilsMessengerEXT, pAllocator *driver.VkAllocationCallbacks) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkDestroyDebugUtilsMessengerEXT", instance, debugMessenger, pAllocator)
}

// VkDestroyDebugUtilsMessengerEXT indicates an expected call of VkDestroyDebugUtilsMessengerEXT.
func (mr *MockDriverMockRecorder) VkDestroyDebugUtilsMessengerEXT(instance, debugMessenger, pAllocator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkDestroyDebugUtilsMessengerEXT", reflect.TypeOf((*MockDriver)(nil).VkDestroyDebugUtilsMessengerEXT), instance, debugMessenger, pAllocator)
}

// VkQueueBeginDebugUtilsLabelEXT mocks base method.
func (m *MockDriver) VkQueueBeginDebugUtilsLabelEXT(queue driver.VkQueue, pLabelInfo *ext_debug_utils_driver.VkDebugUtilsLabelEXT) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkQueueBeginDebugUtilsLabelEXT", queue, pLabelInfo)
}

// VkQueueBeginDebugUtilsLabelEXT indicates an expected call of VkQueueBeginDebugUtilsLabelEXT.
func (mr *MockDriverMockRecorder) VkQueueBeginDebugUtilsLabelEXT(queue, pLabelInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkQueueBeginDebugUtilsLabelEXT", reflect.TypeOf((*MockDriver)(nil).VkQueueBeginDebugUtilsLabelEXT), queue, pLabelInfo)
}

// VkQueueEndDebugUtilsLabelEXT mocks base method.
func (m *MockDriver) VkQueueEndDebugUtilsLabelEXT(queue driver.VkQueue) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkQueueEndDebugUtilsLabelEXT", queue)
}

// VkQueueEndDebugUtilsLabelEXT indicates an expected call of VkQueueEndDebugUtilsLabelEXT.
func (mr *MockDriverMockRecorder) VkQueueEndDebugUtilsLabelEXT(queue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkQueueEndDebugUtilsLabelEXT", reflect.TypeOf((*MockDriver)(nil).VkQueueEndDebugUtilsLabelEXT), queue)
}

// VkQueueInsertDebugUtilsLabelEXT mocks base method.
func (m *MockDriver) VkQueueInsertDebugUtilsLabelEXT(queue driver.VkQueue, pLabelInfo *ext_debug_utils_driver.VkDebugUtilsLabelEXT) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkQueueInsertDebugUtilsLabelEXT", queue, pLabelInfo)
}

// VkQueueInsertDebugUtilsLabelEXT indicates an expected call of VkQueueInsertDebugUtilsLabelEXT.
func (mr *MockDriverMockRecorder) VkQueueInsertDebugUtilsLabelEXT(queue, pLabelInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkQueueInsertDebugUtilsLabelEXT", reflect.TypeOf((*MockDriver)(nil).VkQueueInsertDebugUtilsLabelEXT), queue, pLabelInfo)
}

// VkSetDebugUtilsObjectNameEXT mocks base method.
func (m *MockDriver) VkSetDebugUtilsObjectNameEXT(device driver.VkDevice, pNameInfo *ext_debug_utils_driver.VkDebugUtilsObjectNameInfoEXT) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkSetDebugUtilsObjectNameEXT", device, pNameInfo)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkSetDebugUtilsObjectNameEXT indicates an expected call of VkSetDebugUtilsObjectNameEXT.
func (mr *MockDriverMockRecorder) VkSetDebugUtilsObjectNameEXT(device, pNameInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkSetDebugUtilsObjectNameEXT", reflect.TypeOf((*MockDriver)(nil).VkSetDebugUtilsObjectNameEXT), device, pNameInfo)
}

// VkSetDebugUtilsObjectTagEXT mocks base method.
func (m *MockDriver) VkSetDebugUtilsObjectTagEXT(device driver.VkDevice, pTagInfo *ext_debug_utils_driver.VkDebugUtilsObjectTagInfoEXT) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkSetDebugUtilsObjectTagEXT", device, pTagInfo)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkSetDebugUtilsObjectTagEXT indicates an expected call of VkSetDebugUtilsObjectTagEXT.
func (mr *MockDriverMockRecorder) VkSetDebugUtilsObjectTagEXT(device, pTagInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkSetDebugUtilsObjectTagEXT", reflect.TypeOf((*MockDriver)(nil).VkSetDebugUtilsObjectTagEXT), device, pTagInfo)
}

// VkSubmitDebugUtilsMessageEXT mocks base method.
func (m *MockDriver) VkSubmitDebugUtilsMessageEXT(instance driver.VkInstance, messageSeverity ext_debug_utils_driver.VkDebugUtilsMessageSeverityFlagBitsEXT, messageTypes ext_debug_utils_driver.VkDebugUtilsMessageTypeFlagsEXT, pCallbackData *ext_debug_utils_driver.VkDebugUtilsMessengerCallbackDataEXT) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkSubmitDebugUtilsMessageEXT", instance, messageSeverity, messageTypes, pCallbackData)
}

// VkSubmitDebugUtilsMessageEXT indicates an expected call of VkSubmitDebugUtilsMessageEXT.
func (mr *MockDriverMockRecorder) VkSubmitDebugUtilsMessageEXT(instance, messageSeverity, messageTypes, pCallbackData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkSubmitDebugUtilsMessageEXT", reflect.TypeOf((*MockDriver)(nil).VkSubmitDebugUtilsMessageEXT), instance, messageSeverity, messageTypes, pCallbackData)
}
