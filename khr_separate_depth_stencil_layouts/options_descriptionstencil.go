package khr_separate_depth_stencil_layouts

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

// AttachmentDescriptionStencilLayout specifies an attachment description
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionStencilLayout.html
type AttachmentDescriptionStencilLayout struct {
	// StencilInitialLayout is the layout of the stencil aspect of the attachment Image
	// subresource will be in when a RenderPass instance begins
	StencilInitialLayout core1_0.ImageLayout
	// StencilFinalLayout is the layout the stencil aspect of the attachment Image subresource
	// will be transitioned to when a RenderPass instance ends
	StencilFinalLayout core1_0.ImageLayout

	common.NextOptions
}

func (o AttachmentDescriptionStencilLayout) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAttachmentDescriptionStencilLayoutKHR{})))
	}

	info := (*C.VkAttachmentDescriptionStencilLayoutKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_STENCIL_LAYOUT_KHR
	info.pNext = next
	info.stencilInitialLayout = C.VkImageLayout(o.StencilInitialLayout)
	info.stencilFinalLayout = C.VkImageLayout(o.StencilFinalLayout)

	return preallocatedPointer, nil
}
