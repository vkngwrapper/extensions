package khr_device_group

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/extensions/khr_swapchain"
	"unsafe"
)

type BindImageMemorySwapchainInfo struct {
	Swapchain  khr_swapchain.Swapchain
	ImageIndex int

	common.NextOptions
}

func (o BindImageMemorySwapchainInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBindImageMemorySwapchainInfoKHR{})))
	}

	info := (*C.VkBindImageMemorySwapchainInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_SWAPCHAIN_INFO_KHR
	info.pNext = next
	info.swapchain = nil
	info.imageIndex = C.uint32_t(o.ImageIndex)

	if o.Swapchain != nil {
		info.swapchain = C.VkSwapchainKHR(unsafe.Pointer(o.Swapchain.Handle()))
	}

	return preallocatedPointer, nil
}
