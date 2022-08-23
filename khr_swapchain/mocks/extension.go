// Code generated by MockGen. DO NOT EDIT.
// Source: extension.go

// Package mock_swapchain is a generated GoMock package.
package mock_swapchain

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/vkngwrapper/core/v2/common"
	core1_0 "github.com/vkngwrapper/core/v2/core1_0"
	driver "github.com/vkngwrapper/core/v2/driver"
	khr_swapchain "github.com/vkngwrapper/extensions/v2/khr_swapchain"
	khr_swapchain_driver "github.com/vkngwrapper/extensions/v2/khr_swapchain/driver"
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

// APIVersion mocks base method.
func (m *MockExtension) APIVersion() common.APIVersion {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIVersion")
	ret0, _ := ret[0].(common.APIVersion)
	return ret0
}

// APIVersion indicates an expected call of APIVersion.
func (mr *MockExtensionMockRecorder) APIVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIVersion", reflect.TypeOf((*MockExtension)(nil).APIVersion))
}

// CreateSwapchain mocks base method.
func (m *MockExtension) CreateSwapchain(device core1_0.Device, allocation *driver.AllocationCallbacks, options khr_swapchain.SwapchainCreateInfo) (khr_swapchain.Swapchain, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSwapchain", device, allocation, options)
	ret0, _ := ret[0].(khr_swapchain.Swapchain)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateSwapchain indicates an expected call of CreateSwapchain.
func (mr *MockExtensionMockRecorder) CreateSwapchain(device, allocation, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSwapchain", reflect.TypeOf((*MockExtension)(nil).CreateSwapchain), device, allocation, options)
}

// Driver mocks base method.
func (m *MockExtension) Driver() khr_swapchain_driver.Driver {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Driver")
	ret0, _ := ret[0].(khr_swapchain_driver.Driver)
	return ret0
}

// Driver indicates an expected call of Driver.
func (mr *MockExtensionMockRecorder) Driver() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Driver", reflect.TypeOf((*MockExtension)(nil).Driver))
}

// QueuePresent mocks base method.
func (m *MockExtension) QueuePresent(queue core1_0.Queue, o khr_swapchain.PresentInfo) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueuePresent", queue, o)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueuePresent indicates an expected call of QueuePresent.
func (mr *MockExtensionMockRecorder) QueuePresent(queue, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueuePresent", reflect.TypeOf((*MockExtension)(nil).QueuePresent), queue, o)
}
