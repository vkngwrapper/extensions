package khr_copy_commands2

import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
)

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type BufferCopy2 struct {
	SrcOffset int
	DstOffset int
	Size      int

	common.NextOptions
}

func (o BufferCopy2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBufferCopy2KHR{})))
	}

	info := (*C.VkBufferCopy2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_BUFFER_COPY_2
	info.pNext = next
	info.srcOffset = C.VkDeviceSize(o.SrcOffset)
	info.dstOffset = C.VkDeviceSize(o.DstOffset)
	info.size = C.VkDeviceSize(o.Size)

	return preallocatedPointer, nil
}
