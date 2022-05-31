package khr_create_renderpass2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type SubpassDependencyOptions struct {
	SrcSubpassIndex int
	DstSubpassIndex int
	SrcStageMask    common.PipelineStages
	DstStageMask    common.PipelineStages
	SrcAccessMask   common.AccessFlags
	DstAccessMask   common.AccessFlags
	DependencyFlags common.DependencyFlags
	ViewOffset      int

	common.HaveNext
}

func (o SubpassDependencyOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassDependency2KHR{})))
	}

	info := (*C.VkSubpassDependency2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2_KHR
	info.pNext = next
	info.srcSubpass = C.uint32_t(o.SrcSubpassIndex)
	info.dstSubpass = C.uint32_t(o.DstSubpassIndex)
	info.srcStageMask = C.VkPipelineStageFlags(o.SrcStageMask)
	info.dstStageMask = C.VkPipelineStageFlags(o.DstStageMask)
	info.srcAccessMask = C.VkAccessFlags(o.SrcAccessMask)
	info.dstAccessMask = C.VkAccessFlags(o.DstAccessMask)
	info.dependencyFlags = C.VkDependencyFlags(o.DependencyFlags)
	info.viewOffset = C.int32_t(o.ViewOffset)

	return preallocatedPointer, nil
}

func (o SubpassDependencyOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkSubpassDependency2KHR)(cDataPointer)
	return info.pNext, nil
}