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

// BindImageMemorySwapchainInfo specifies swapchain Image memory to bind to
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemorySwapchainInfoKHR.html
type BindImageMemorySwapchainInfo struct {
	// Swapchain is the khr_swapchain.Swapchain whose Image memory will be bound
	Swapchain khr_swapchain.Swapchain
	// ImageIndex is an Image index within Swapchain
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

	if o.Swapchain.Initialized() {
		info.swapchain = C.VkSwapchainKHR(unsafe.Pointer(o.Swapchain.Handle()))
	}

	return preallocatedPointer, nil
}
