package khr_device_group

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
)

// DeviceGroupSwapchainCreateInfo specifies parameters of a newly-created khr_swapchain.Swapchain object
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupSwapchainCreateInfoKHR.html
type DeviceGroupSwapchainCreateInfo struct {
	// Modes is a set of modes that the khr_swapchain.Swapchain can be used with
	Modes DeviceGroupPresentModeFlags

	common.NextOptions
}

func (o DeviceGroupSwapchainCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDeviceGroupSwapchainCreateInfoKHR)
	}

	info := (*C.VkDeviceGroupSwapchainCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_SWAPCHAIN_CREATE_INFO_KHR
	info.pNext = next
	info.modes = C.VkDeviceGroupPresentModeFlagsKHR(o.Modes)

	return preallocatedPointer, nil
}
