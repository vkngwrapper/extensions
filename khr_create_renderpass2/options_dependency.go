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

// SubpassDependency2 specifies a subpass dependency
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDependency2.html
type SubpassDependency2 struct {
	// SrcSubpass is the subpass index of the first subpass in the dependency, or
	// core1_0.SubpassExternal
	SrcSubpass int
	// DstSubpass is the subpass index of the second subpass in the dependency, or
	// core1_0.SubpassExternal
	DstSubpass int
	// SrcStageMask specifies the source stage mask
	SrcStageMask core1_0.PipelineStageFlags
	// DstStageMask specifies the destination stage mask
	DstStageMask core1_0.PipelineStageFlags
	// SrcAccessMask specifies a source access mask
	SrcAccessMask core1_0.AccessFlags
	// DstAccessMask specifies a source access mask
	DstAccessMask core1_0.AccessFlags
	// DependencyFlags is a set of dependency flags
	DependencyFlags core1_0.DependencyFlags
	// ViewOffset controls which views in the source subpass the views in the destination
	// subpass depend on
	ViewOffset int

	common.NextOptions
}

func (o SubpassDependency2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassDependency2KHR{})))
	}

	info := (*C.VkSubpassDependency2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_DEPENDENCY_2_KHR
	info.pNext = next
	info.srcSubpass = C.uint32_t(o.SrcSubpass)
	info.dstSubpass = C.uint32_t(o.DstSubpass)
	info.srcStageMask = C.VkPipelineStageFlags(o.SrcStageMask)
	info.dstStageMask = C.VkPipelineStageFlags(o.DstStageMask)
	info.srcAccessMask = C.VkAccessFlags(o.SrcAccessMask)
	info.dstAccessMask = C.VkAccessFlags(o.DstAccessMask)
	info.dependencyFlags = C.VkDependencyFlags(o.DependencyFlags)
	info.viewOffset = C.int32_t(o.ViewOffset)

	return preallocatedPointer, nil
}
