package khr_imageless_framebuffer

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
)

// RenderPassAttachmentBeginInfo specifies Image objects to be used as Framebuffer attachments
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassAttachmentBeginInfo.html
type RenderPassAttachmentBeginInfo struct {
	// Attachments is a slice of ImageView objects, each of which will be used as the corresponding
	// attachment in the RenderPass instance
	Attachments []core.ImageView

	common.NextOptions
}

func (o RenderPassAttachmentBeginInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkRenderPassAttachmentBeginInfoKHR{})))
	}

	info := (*C.VkRenderPassAttachmentBeginInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_RENDER_PASS_ATTACHMENT_BEGIN_INFO_KHR
	info.pNext = next

	count := len(o.Attachments)
	info.attachmentCount = C.uint32_t(count)
	info.pAttachments = nil

	if count > 0 {
		info.pAttachments = (*C.VkImageView)(allocator.Malloc(count * int(unsafe.Sizeof([1]C.VkImageView{}))))
		attachmentSlice := unsafe.Slice(info.pAttachments, count)
		for i := 0; i < count; i++ {
			if o.Attachments[i].Handle() == 0 {
				return nil, errors.Errorf("khr_imageless_framebuffer.RenderPassAttachmentBeginInfo.Attachments "+
					"cannot contain uninitialized elements, but element %d is uninitialized", i)
			}
			attachmentSlice[i] = C.VkImageView(unsafe.Pointer(o.Attachments[i].Handle()))
		}
	}

	return preallocatedPointer, nil
}
