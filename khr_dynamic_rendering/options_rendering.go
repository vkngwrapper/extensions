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
)

// RenderingInfo specifies the parameters of a render pass instance
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkRenderingInfoKHR.html
type RenderingInfo struct {
	Flags             RenderingFlags
	RenderArea        core1_0.Rect2D
	LayerCount        int
	ViewMask          uint32
	ColorAttachments  []RenderingAttachmentInfo
	DepthAttachment   *RenderingAttachmentInfo
	StencilAttachment *RenderingAttachmentInfo

	common.NextOptions
}

func (o RenderingInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkRenderingInfoKHR{})))
	}

	info := (*C.VkRenderingInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_RENDERING_INFO_KHR
	info.pNext = next
	info.flags = C.VkRenderingFlagsKHR(o.Flags)
	info.renderArea.offset.x = C.int32_t(o.RenderArea.Offset.X)
	info.renderArea.offset.y = C.int32_t(o.RenderArea.Offset.Y)
	info.renderArea.extent.width = C.uint32_t(o.RenderArea.Extent.Width)
	info.renderArea.extent.height = C.uint32_t(o.RenderArea.Extent.Height)
	info.layerCount = C.uint32_t(o.LayerCount)
	info.viewMask = C.uint32_t(o.ViewMask)
	info.colorAttachmentCount = C.uint32_t(len(o.ColorAttachments))
	info.pColorAttachments = nil
	info.pDepthAttachment = nil
	info.pStencilAttachment = nil

	if len(o.ColorAttachments) > 0 {
		attachments, err := common.AllocOptionSlice[C.VkRenderingAttachmentInfoKHR](allocator, o.ColorAttachments)
		if err != nil {
			return nil, err
		}
		info.pColorAttachments = attachments
	}
	if o.DepthAttachment != nil {
		attachment, err := common.AllocOptions(allocator, o.DepthAttachment)
		if err != nil {
			return nil, err
		}
		info.pDepthAttachment = (*C.VkRenderingAttachmentInfoKHR)(attachment)
	}
	if o.StencilAttachment != nil {
		attachment, err := common.AllocOptions(allocator, o.StencilAttachment)
		if err != nil {
			return nil, err
		}
		info.pStencilAttachment = (*C.VkRenderingAttachmentInfoKHR)(attachment)
	}

	return preallocatedPointer, nil
}
