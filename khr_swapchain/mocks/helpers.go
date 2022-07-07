package mock_swapchain

import (
	"github.com/golang/mock/gomock"
	khr_swapchain_driver "github.com/vkngwrapper/extensions/khr_swapchain/driver"
	"math/rand"
	"unsafe"
)

func NewFakeSwapchain() khr_swapchain_driver.VkSwapchainKHR {
	return khr_swapchain_driver.VkSwapchainKHR(unsafe.Pointer(uintptr(rand.Int())))
}

func EasyMockSwapchain(ctrl *gomock.Controller) *MockSwapchain {
	swapchain := NewMockSwapchain(ctrl)
	swapchain.EXPECT().Handle().Return(NewFakeSwapchain()).AnyTimes()

	return swapchain
}
