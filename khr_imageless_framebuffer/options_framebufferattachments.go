package khr_imageless_framebuffer

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

// FramebufferAttachmentsCreateInfo specifies parameters of Image objects that will be used with
// a Framebuffer
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html
type FramebufferAttachmentsCreateInfo struct {
	// AttachmentImageInfos is a slice of FramebufferAttachmentInfo structures, each structure
	// describing a number of parameters of the corresponding attachment in a RenderPass instance
	AttachmentImageInfos []FramebufferAttachmentImageInfo

	common.NextOptions
}

func (o FramebufferAttachmentsCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkFramebufferAttachmentsCreateInfoKHR{})))
	}

	info := (*C.VkFramebufferAttachmentsCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENTS_CREATE_INFO_KHR
	info.pNext = next

	count := len(o.AttachmentImageInfos)
	info.attachmentImageInfoCount = C.uint32_t(count)
	info.pAttachmentImageInfos = nil

	if count > 0 {
		infosPtr, err := common.AllocOptionSlice[C.VkFramebufferAttachmentImageInfoKHR, FramebufferAttachmentImageInfo](allocator, o.AttachmentImageInfos)
		if err != nil {
			return nil, err
		}

		info.pAttachmentImageInfos = infosPtr
	}

	return preallocatedPointer, nil
}
