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

// CommandBufferInheritanceRenderingInfo specifies rendering parameters inherited by a secondary CommandBuffer
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkCommandBufferInheritanceRenderingInfoKHR.html
type CommandBufferInheritanceRenderingInfo struct {
	Flags                   RenderingFlags
	ViewMask                uint32
	ColorAttachmentFormats  []core1_0.Format
	DepthAttachmentFormat   core1_0.Format
	StencilAttachmentFormat core1_0.Format
	RasterizationSamples    core1_0.SampleCountFlags

	common.NextOptions
}

func (o CommandBufferInheritanceRenderingInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkCommandBufferInheritanceRenderingInfoKHR{})))
	}

	info := (*C.VkCommandBufferInheritanceRenderingInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_RENDERING_INFO_KHR
	info.pNext = next
	info.flags = C.VkRenderingFlagsKHR(o.Flags)
	info.viewMask = C.uint32_t(o.ViewMask)
	info.depthAttachmentFormat = C.VkFormat(o.DepthAttachmentFormat)
	info.stencilAttachmentFormat = C.VkFormat(o.StencilAttachmentFormat)
	info.rasterizationSamples = C.VkSampleCountFlagBits(o.RasterizationSamples)

	count := len(o.ColorAttachmentFormats)
	info.colorAttachmentCount = C.uint32_t(count)
	info.pColorAttachmentFormats = nil
	if count > 0 {
		formats := (*C.VkFormat)(allocator.Malloc(count * int(unsafe.Sizeof(C.VkFormat(0)))))
		formatSlice := unsafe.Slice(formats, count)
		for index, format := range o.ColorAttachmentFormats {
			formatSlice[index] = C.VkFormat(format)
		}
		info.pColorAttachmentFormats = formats
	}

	return preallocatedPointer, nil
}
