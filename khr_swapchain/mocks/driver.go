// Code generated by MockGen. DO NOT EDIT.
// Source: driver.go

// Package mock_swapchain is a generated GoMock package.
package mock_swapchain

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/vkngwrapper/core/v2/common"
	driver "github.com/vkngwrapper/core/v2/driver"
	khr_surface_driver "github.com/vkngwrapper/extensions/v2/khr_surface/driver"
	khr_swapchain_driver "github.com/vkngwrapper/extensions/v2/khr_swapchain/driver"
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

// VkAcquireNextImage2KHR mocks base method.
func (m *MockDriver) VkAcquireNextImage2KHR(device driver.VkDevice, pAcquireInfo *khr_swapchain_driver.VkAcquireNextImageInfoKHR, pImageIndex *driver.Uint32) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkAcquireNextImage2KHR", device, pAcquireInfo, pImageIndex)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkAcquireNextImage2KHR indicates an expected call of VkAcquireNextImage2KHR.
func (mr *MockDriverMockRecorder) VkAcquireNextImage2KHR(device, pAcquireInfo, pImageIndex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkAcquireNextImage2KHR", reflect.TypeOf((*MockDriver)(nil).VkAcquireNextImage2KHR), device, pAcquireInfo, pImageIndex)
}

// VkAcquireNextImageKHR mocks base method.
func (m *MockDriver) VkAcquireNextImageKHR(device driver.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, timeout driver.Uint64, semaphore driver.VkSemaphore, fence driver.VkFence, pImageIndex *driver.Uint32) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkAcquireNextImageKHR", device, swapchain, timeout, semaphore, fence, pImageIndex)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkAcquireNextImageKHR indicates an expected call of VkAcquireNextImageKHR.
func (mr *MockDriverMockRecorder) VkAcquireNextImageKHR(device, swapchain, timeout, semaphore, fence, pImageIndex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkAcquireNextImageKHR", reflect.TypeOf((*MockDriver)(nil).VkAcquireNextImageKHR), device, swapchain, timeout, semaphore, fence, pImageIndex)
}

// VkCreateSwapchainKHR mocks base method.
func (m *MockDriver) VkCreateSwapchainKHR(device driver.VkDevice, pCreateInfo *khr_swapchain_driver.VkSwapchainCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pSwapchain *khr_swapchain_driver.VkSwapchainKHR) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkCreateSwapchainKHR", device, pCreateInfo, pAllocator, pSwapchain)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkCreateSwapchainKHR indicates an expected call of VkCreateSwapchainKHR.
func (mr *MockDriverMockRecorder) VkCreateSwapchainKHR(device, pCreateInfo, pAllocator, pSwapchain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkCreateSwapchainKHR", reflect.TypeOf((*MockDriver)(nil).VkCreateSwapchainKHR), device, pCreateInfo, pAllocator, pSwapchain)
}

// VkDestroySwapchainKHR mocks base method.
func (m *MockDriver) VkDestroySwapchainKHR(device driver.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, pAllocator *driver.VkAllocationCallbacks) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkDestroySwapchainKHR", device, swapchain, pAllocator)
}

// VkDestroySwapchainKHR indicates an expected call of VkDestroySwapchainKHR.
func (mr *MockDriverMockRecorder) VkDestroySwapchainKHR(device, swapchain, pAllocator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkDestroySwapchainKHR", reflect.TypeOf((*MockDriver)(nil).VkDestroySwapchainKHR), device, swapchain, pAllocator)
}

// VkGetDeviceGroupPresentCapabilitiesKHR mocks base method.
func (m *MockDriver) VkGetDeviceGroupPresentCapabilitiesKHR(device driver.VkDevice, pDeviceGroupPresentCapabilities *khr_swapchain_driver.VkDeviceGroupPresentCapabilitiesKHR) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkGetDeviceGroupPresentCapabilitiesKHR", device, pDeviceGroupPresentCapabilities)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkGetDeviceGroupPresentCapabilitiesKHR indicates an expected call of VkGetDeviceGroupPresentCapabilitiesKHR.
func (mr *MockDriverMockRecorder) VkGetDeviceGroupPresentCapabilitiesKHR(device, pDeviceGroupPresentCapabilities interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetDeviceGroupPresentCapabilitiesKHR", reflect.TypeOf((*MockDriver)(nil).VkGetDeviceGroupPresentCapabilitiesKHR), device, pDeviceGroupPresentCapabilities)
}

// VkGetDeviceGroupSurfacePresentModesKHR mocks base method.
func (m *MockDriver) VkGetDeviceGroupSurfacePresentModesKHR(device driver.VkDevice, surface khr_surface_driver.VkSurfaceKHR, pModes *khr_swapchain_driver.VkDeviceGroupPresentModeFlagsKHR) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkGetDeviceGroupSurfacePresentModesKHR", device, surface, pModes)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkGetDeviceGroupSurfacePresentModesKHR indicates an expected call of VkGetDeviceGroupSurfacePresentModesKHR.
func (mr *MockDriverMockRecorder) VkGetDeviceGroupSurfacePresentModesKHR(device, surface, pModes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetDeviceGroupSurfacePresentModesKHR", reflect.TypeOf((*MockDriver)(nil).VkGetDeviceGroupSurfacePresentModesKHR), device, surface, pModes)
}

// VkGetPhysicalDevicePresentRectanglesKHR mocks base method.
func (m *MockDriver) VkGetPhysicalDevicePresentRectanglesKHR(physicalDevice driver.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pRectCount *driver.Uint32, pRects *driver.VkRect2D) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkGetPhysicalDevicePresentRectanglesKHR", physicalDevice, surface, pRectCount, pRects)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkGetPhysicalDevicePresentRectanglesKHR indicates an expected call of VkGetPhysicalDevicePresentRectanglesKHR.
func (mr *MockDriverMockRecorder) VkGetPhysicalDevicePresentRectanglesKHR(physicalDevice, surface, pRectCount, pRects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetPhysicalDevicePresentRectanglesKHR", reflect.TypeOf((*MockDriver)(nil).VkGetPhysicalDevicePresentRectanglesKHR), physicalDevice, surface, pRectCount, pRects)
}

// VkGetSwapchainImagesKHR mocks base method.
func (m *MockDriver) VkGetSwapchainImagesKHR(device driver.VkDevice, swapchain khr_swapchain_driver.VkSwapchainKHR, pSwapchainImageCount *driver.Uint32, pSwapchainImages *driver.VkImage) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkGetSwapchainImagesKHR", device, swapchain, pSwapchainImageCount, pSwapchainImages)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkGetSwapchainImagesKHR indicates an expected call of VkGetSwapchainImagesKHR.
func (mr *MockDriverMockRecorder) VkGetSwapchainImagesKHR(device, swapchain, pSwapchainImageCount, pSwapchainImages interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetSwapchainImagesKHR", reflect.TypeOf((*MockDriver)(nil).VkGetSwapchainImagesKHR), device, swapchain, pSwapchainImageCount, pSwapchainImages)
}

// VkQueuePresentKHR mocks base method.
func (m *MockDriver) VkQueuePresentKHR(queue driver.VkQueue, pPresentInfo *khr_swapchain_driver.VkPresentInfoKHR) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkQueuePresentKHR", queue, pPresentInfo)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkQueuePresentKHR indicates an expected call of VkQueuePresentKHR.
func (mr *MockDriverMockRecorder) VkQueuePresentKHR(queue, pPresentInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkQueuePresentKHR", reflect.TypeOf((*MockDriver)(nil).VkQueuePresentKHR), queue, pPresentInfo)
}
