// Code generated by MockGen. DO NOT EDIT.
// Source: driver.go

// Package mock_draw_indirect_count is a generated GoMock package.
package mock_draw_indirect_count

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	driver "github.com/vkngwrapper/core/v2/driver"
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

// VkCmdDrawIndexedIndirectCountKHR mocks base method.
func (m *MockDriver) VkCmdDrawIndexedIndirectCountKHR(commandBuffer driver.VkCommandBuffer, buffer driver.VkBuffer, offset driver.VkDeviceSize, countBuffer driver.VkBuffer, countBufferOffset driver.VkDeviceSize, maxDrawCount, stride driver.Uint32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkCmdDrawIndexedIndirectCountKHR", commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride)
}

// VkCmdDrawIndexedIndirectCountKHR indicates an expected call of VkCmdDrawIndexedIndirectCountKHR.
func (mr *MockDriverMockRecorder) VkCmdDrawIndexedIndirectCountKHR(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkCmdDrawIndexedIndirectCountKHR", reflect.TypeOf((*MockDriver)(nil).VkCmdDrawIndexedIndirectCountKHR), commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride)
}

// VkCmdDrawIndirectCountKHR mocks base method.
func (m *MockDriver) VkCmdDrawIndirectCountKHR(commandBuffer driver.VkCommandBuffer, buffer driver.VkBuffer, offset driver.VkDeviceSize, countBuffer driver.VkBuffer, countBufferOffset driver.VkDeviceSize, maxDrawCount, stride driver.Uint32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkCmdDrawIndirectCountKHR", commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride)
}

// VkCmdDrawIndirectCountKHR indicates an expected call of VkCmdDrawIndirectCountKHR.
func (mr *MockDriverMockRecorder) VkCmdDrawIndirectCountKHR(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkCmdDrawIndirectCountKHR", reflect.TypeOf((*MockDriver)(nil).VkCmdDrawIndirectCountKHR), commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride)
}
