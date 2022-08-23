package khr_sampler_ycbcr_conversion

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"unsafe"
)

// BindImagePlaneMemoryInfo specifies how to bind an Image plane to DeviceMemory
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImagePlaneMemoryInfo.html
type BindImagePlaneMemoryInfo struct {
	// PlaneAspect specifies the aspect of the disjoint Image plane to bind
	PlaneAspect core1_0.ImageAspectFlags

	common.NextOptions
}

func (o BindImagePlaneMemoryInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBindImagePlaneMemoryInfoKHR{})))
	}

	info := (*C.VkBindImagePlaneMemoryInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO_KHR
	info.pNext = next
	info.planeAspect = C.VkImageAspectFlagBits(o.PlaneAspect)

	return preallocatedPointer, nil
}
