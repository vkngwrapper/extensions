package khr_swapchain

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"math/rand"
	"unsafe"

	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
	khr_swapchain_driver "github.com/vkngwrapper/extensions/v3/khr_swapchain/loader"
)

type Swapchain struct {
	handle khr_swapchain_driver.VkSwapchainKHR
	device loader.VkDevice

	apiVersion common.APIVersion
}

func (s Swapchain) Handle() khr_swapchain_driver.VkSwapchainKHR {
	return s.handle
}

func (s Swapchain) DeviceHandle() loader.VkDevice {
	return s.device
}

func (s Swapchain) APIVersion() common.APIVersion {
	return s.apiVersion
}

func (s Swapchain) Initialized() bool {
	return s.handle != 0
}

func NewDummySwapchain(device core.Device) Swapchain {
	return Swapchain{
		handle:     khr_swapchain_driver.VkSwapchainKHR(unsafe.Pointer(uintptr(rand.Int()))),
		device:     device.Handle(),
		apiVersion: device.APIVersion(),
	}
}
