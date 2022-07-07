package khr_separate_depth_stencil_layouts

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"unsafe"
)

type AttachmentReferenceStencilLayout struct {
	StencilLayout core1_0.ImageLayout

	common.NextOptions
}

func (o AttachmentReferenceStencilLayout) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAttachmentReferenceStencilLayoutKHR{})))
	}

	info := (*C.VkAttachmentReferenceStencilLayoutKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_STENCIL_LAYOUT_KHR
	info.pNext = next
	info.stencilLayout = C.VkImageLayout(o.StencilLayout)

	return preallocatedPointer, nil
}
