package mock_swapchain

import (
	"math/rand"
	"unsafe"

	khr_swapchain_driver "github.com/vkngwrapper/extensions/v3/khr_swapchain/driver"
	gomock "go.uber.org/mock/gomock"
)

func NewFakeSwapchain() khr_swapchain_driver.VkSwapchainKHR {
	return khr_swapchain_driver.VkSwapchainKHR(unsafe.Pointer(uintptr(rand.Int())))
}

func EasyMockSwapchain(ctrl *gomock.Controller) *MockSwapchain {
	swapchain := NewMockSwapchain(ctrl)
	swapchain.EXPECT().Handle().Return(NewFakeSwapchain()).AnyTimes()

	return swapchain
}
