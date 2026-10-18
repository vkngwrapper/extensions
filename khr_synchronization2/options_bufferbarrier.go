package khr_synchronization2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
)

type BufferMemoryBarrier2 struct {
	SrcStageMask  PipelineStageFlags2
	SrcAccessMask AccessFlags2
	DstStageMask  PipelineStageFlags2
	DstAccessMask AccessFlags2

	SrcQueueFamilyIndex int
	DstQueueFamilyIndex int

	Buffer core1_0.Buffer
	Offset int
	Size   int

	common.NextOptions
}

func (o BufferMemoryBarrier2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if !o.Buffer.Initialized() {
		return nil, errors.New("khr_synchronization2.BufferMemoryBarrier2.Buffer cannot be uninitialized")
	}
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkBufferMemoryBarrier2{})))
	}
	barrier := (*C.VkBufferMemoryBarrier2)(preallocatedPointer)
	barrier.sType = C.VK_STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER_2
	barrier.pNext = next
	barrier.srcStageMask = C.VkPipelineStageFlags2(o.SrcStageMask)
	barrier.srcAccessMask = C.VkAccessFlags2(o.SrcAccessMask)
	barrier.dstStageMask = C.VkPipelineStageFlags2(o.DstStageMask)
	barrier.dstAccessMask = C.VkAccessFlags2(o.DstAccessMask)
	barrier.srcQueueFamilyIndex = C.uint32_t(o.SrcQueueFamilyIndex)
	barrier.dstQueueFamilyIndex = C.uint32_t(o.DstQueueFamilyIndex)
	barrier.offset = C.VkDeviceSize(o.Offset)
	barrier.size = C.VkDeviceSize(o.Size)
	barrier.buffer = C.VkBuffer(unsafe.Pointer(o.Buffer.Handle()))

	return preallocatedPointer, nil
}
