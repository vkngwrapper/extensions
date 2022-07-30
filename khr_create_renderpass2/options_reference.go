package khr_create_renderpass2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"unsafe"
)

// AttachmentReference2 specifies an attachment reference
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html
type AttachmentReference2 struct {
	// Attachment identifies an attachment at the corresponding index in
	// RenderPassCreateInfo2.Attachments, or core1_0.AttachmentUnused
	Attachment int
	// Layout specifies the layout the attachment uses during the subpass
	Layout core1_0.ImageLayout
	// AspectMask is a mask of which aspect(s) can be accessed within the specified
	// subpass as an input attachment
	AspectMask core1_0.ImageAspectFlags

	common.NextOptions
}

func (o AttachmentReference2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAttachmentReference2KHR{})))
	}

	info := (*C.VkAttachmentReference2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2_KHR
	info.pNext = next
	info.attachment = C.uint32_t(o.Attachment)
	info.layout = C.VkImageLayout(o.Layout)
	info.aspectMask = C.VkImageAspectFlags(o.AspectMask)

	return preallocatedPointer, nil
}
