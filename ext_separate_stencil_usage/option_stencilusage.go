package ext_separate_stencil_usage

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

// ImageStencilUsageCreateInfo specifies separate usage flags for the stencil aspect of a
// depth-stencil Image
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html
type ImageStencilUsageCreateInfo struct {
	// StencilUsage describes the intended usage of the stencil aspect of the Image
	StencilUsage core1_0.ImageUsageFlags

	common.NextOptions
}

func (o ImageStencilUsageCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageStencilUsageCreateInfoEXT{})))
	}

	info := (*C.VkImageStencilUsageCreateInfoEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_IMAGE_STENCIL_USAGE_CREATE_INFO_EXT
	info.pNext = next
	info.stencilUsage = C.VkImageUsageFlags(o.StencilUsage)

	return preallocatedPointer, nil
}
