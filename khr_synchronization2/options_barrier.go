package khr_synchronization2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
)

type MemoryBarrier2 struct {
	SrcStageMask  PipelineStageFlags2
	SrcAccessMask AccessFlags2
	DstStageMask  PipelineStageFlags2
	DstAccessMask AccessFlags2

	common.NextOptions
}

func (o MemoryBarrier2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkMemoryBarrier2{})))
	}

	barrier := (*C.VkMemoryBarrier2)(preallocatedPointer)
	barrier.sType = C.VK_STRUCTURE_TYPE_MEMORY_BARRIER_2
	barrier.pNext = next
	barrier.srcStageMask = (C.VkPipelineStageFlags2)(o.SrcStageMask)
	barrier.srcAccessMask = (C.VkAccessFlags2)(o.SrcAccessMask)
	barrier.dstStageMask = (C.VkPipelineStageFlags2)(o.DstStageMask)
	barrier.dstAccessMask = (C.VkAccessFlags2)(o.DstAccessMask)

	return preallocatedPointer, nil
}
