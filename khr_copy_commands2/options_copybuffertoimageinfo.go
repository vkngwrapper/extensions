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

type CopyBufferToImageInfo2 struct {
	SrcBuffer      core1_0.Buffer
	DstImage       core1_0.Image
	DstImageLayout core1_0.ImageLayout

	Regions []BufferImageCopy2

	common.NextOptions
}

func (o CopyBufferToImageInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if !o.SrcBuffer.Initialized() {
		return nil, errors.New("khr_copy_commands2.CopyBufferToImageInfo2.SrcBuffer cannot be left unset")
	}
	if !o.DstImage.Initialized() {
		return nil, errors.New("khr_copy_commands2.CopyBufferToImageInfo2.DstImage cannot be left unset")
	}
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkCopyBufferToImageInfo2KHR{})))
	}

	info := (*C.VkCopyBufferToImageInfo2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_COPY_BUFFER_TO_IMAGE_INFO_2
	info.pNext = next
	info.srcBuffer = C.VkBuffer(unsafe.Pointer(o.SrcBuffer.Handle()))
	info.dstImage = C.VkImage(unsafe.Pointer(o.DstImage.Handle()))
	info.dstImageLayout = C.VkImageLayout(o.DstImageLayout)
	info.regionCount = C.uint32_t(len(o.Regions))
	info.pRegions = nil

	if len(o.Regions) > 0 {
		var err error
		info.pRegions, err = common.AllocOptionSlice[C.VkBufferImageCopy2KHR](allocator, o.Regions)
		if err != nil {
			return nil, err
		}
	}

	return preallocatedPointer, nil
}
