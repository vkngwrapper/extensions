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

// ImagePlaneMemoryRequirementsInfo specifies Image plane for memory requirements
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePlaneMemoryRequirementsInfo.html
type ImagePlaneMemoryRequirementsInfo struct {
	// PlaneAspect specifies the aspect corresponding to the Image plane
	// to query
	PlaneAspect core1_0.ImageAspectFlags

	common.NextOptions
}

func (o ImagePlaneMemoryRequirementsInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImagePlaneMemoryRequirementsInfoKHR{})))
	}

	info := (*C.VkImagePlaneMemoryRequirementsInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO_KHR
	info.pNext = next
	info.planeAspect = C.VkImageAspectFlagBits(o.PlaneAspect)

	return preallocatedPointer, nil
}
