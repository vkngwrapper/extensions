package khr_maintenance2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
)

// ImageViewUsageCreateInfo specifies the intended usage of an ImageView
type ImageViewUsageCreateInfo struct {
	// Usage specifies allowed usages of the ImageView
	Usage core1_0.ImageUsageFlags

	common.NextOptions
}

func (o ImageViewUsageCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageViewUsageCreateInfoKHR{})))
	}

	createInfo := (*C.VkImageViewUsageCreateInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO_KHR
	createInfo.pNext = next
	createInfo.usage = C.VkImageUsageFlags(o.Usage)

	return preallocatedPointer, nil
}
