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

// PipelineRenderingCreateInfo specifies the attachment formats used by a graphics pipeline
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineRenderingCreateInfoKHR.html
type PipelineRenderingCreateInfo struct {
	ViewMask                uint32
	ColorAttachmentFormats  []core1_0.Format
	DepthAttachmentFormat   core1_0.Format
	StencilAttachmentFormat core1_0.Format

	common.NextOptions
}

func (o PipelineRenderingCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPipelineRenderingCreateInfoKHR{})))
	}

	info := (*C.VkPipelineRenderingCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PIPELINE_RENDERING_CREATE_INFO_KHR
	info.pNext = next
	info.viewMask = C.uint32_t(o.ViewMask)
	info.depthAttachmentFormat = C.VkFormat(o.DepthAttachmentFormat)
	info.stencilAttachmentFormat = C.VkFormat(o.StencilAttachmentFormat)

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
