package khr_dynamic_rendering

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
	"github.com/vkngwrapper/extensions/v3/khr_depth_stencil_resolve"
)

// RenderingAttachmentInfo specifies an attachment used in a render pass instance
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkRenderingAttachmentInfoKHR.html
type RenderingAttachmentInfo struct {
	ImageView          core1_0.ImageView
	ImageLayout        core1_0.ImageLayout
	ResolveMode        khr_depth_stencil_resolve.ResolveModeFlags
	ResolveImageView   core1_0.ImageView
	ResolveImageLayout core1_0.ImageLayout
	LoadOp             core1_0.AttachmentLoadOp
	StoreOp            core1_0.AttachmentStoreOp
	ClearValue         core1_0.ClearValue

	common.NextOptions
}

func (o RenderingAttachmentInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkRenderingAttachmentInfoKHR{})))
	}

	info := (*C.VkRenderingAttachmentInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_RENDERING_ATTACHMENT_INFO_KHR
	info.pNext = next
	info.imageView = nil
	info.resolveImageView = nil
	if o.ImageView.Initialized() {
		info.imageView = C.VkImageView(unsafe.Pointer(o.ImageView.Handle()))
	}
	if o.ResolveImageView.Initialized() {
		info.resolveImageView = C.VkImageView(unsafe.Pointer(o.ResolveImageView.Handle()))
	}
	info.imageLayout = C.VkImageLayout(o.ImageLayout)
	info.resolveMode = C.VkResolveModeFlagBitsKHR(o.ResolveMode)
	info.resolveImageLayout = C.VkImageLayout(o.ResolveImageLayout)
	info.loadOp = C.VkAttachmentLoadOp(o.LoadOp)
	info.storeOp = C.VkAttachmentStoreOp(o.StoreOp)
	info.clearValue = C.VkClearValue{}
	if o.ClearValue != nil {
		o.ClearValue.PopulateValueUnion(unsafe.Pointer(&info.clearValue))
	}

	return preallocatedPointer, nil
}
