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
	"github.com/vkngwrapper/extensions/v3/khr_swapchain"
)

// ImageSwapchainCreateInfo specifies that an Image will be bound to swapchain memory
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSwapchainCreateInfoKHR.html
type ImageSwapchainCreateInfo struct {
	// Swapchain is a khr_swapchain.Swapchain object that the Image will be bound to
	Swapchain khr_swapchain.Swapchain

	common.NextOptions
}

func (o ImageSwapchainCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageSwapchainCreateInfoKHR{})))
	}

	info := (*C.VkImageSwapchainCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_IMAGE_SWAPCHAIN_CREATE_INFO_KHR
	info.pNext = next
	info.swapchain = nil

	if o.Swapchain.Handle() != 0 {
		info.swapchain = C.VkSwapchainKHR(unsafe.Pointer(o.Swapchain.Handle()))
	}

	return preallocatedPointer, nil
}
