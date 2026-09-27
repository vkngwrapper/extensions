package khr_copy_commands2

import (
	"errors"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
)

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

type CopyBufferInfo2 struct {
	SrcBuffer core1_0.Buffer
	DstBuffer core1_0.Buffer

	Regions []BufferCopy2

	common.NextOptions
}

func (o CopyBufferInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if !o.SrcBuffer.Initialized() {
		return nil, errors.New("khr_copy_commands2.CopyBufferInfo2.SrcBuffer cannot be left unset")
	}
	if !o.DstBuffer.Initialized() {
		return nil, errors.New("khr_copy_commands2.CopyBufferInfo2.DstBuffer cannot be left unset")
	}

	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkCopyBufferInfo2KHR{})))
	}

	info := (*C.VkCopyBufferInfo2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_COPY_BUFFER_INFO_2
	info.pNext = next
	info.srcBuffer = C.VkBuffer(unsafe.Pointer(o.SrcBuffer.Handle()))
	info.dstBuffer = C.VkBuffer(unsafe.Pointer(o.DstBuffer.Handle()))
	info.regionCount = C.uint32_t(len(o.Regions))
	info.pRegions = nil

	if len(o.Regions) > 0 {
		var err error
		info.pRegions, err = common.AllocOptionSlice[C.VkBufferCopy2KHR](allocator, o.Regions)
		if err != nil {
			return nil, err
		}
	}

	return preallocatedPointer, nil
}
