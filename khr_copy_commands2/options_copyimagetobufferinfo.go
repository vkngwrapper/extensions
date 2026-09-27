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

type CopyImageToBufferInfo2 struct {
	SrcImage       core1_0.Image
	SrcImageLayout core1_0.ImageLayout
	DstBuffer      core1_0.Buffer
	Regions        []BufferImageCopy2

	common.NextOptions
}

func (o CopyImageToBufferInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if !o.SrcImage.Initialized() {
		return nil, errors.New("khr_copy_commands2.CopyImageToBufferInfo2.SrcImage cannot be left unset")
	}
	if !o.DstBuffer.Initialized() {
		return nil, errors.New("khr_copy_commands2.CopyImageToBufferInfo2.DstBuffer cannot be left unset")
	}
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkCopyImageToBufferInfo2KHR{})))
	}

	info := (*C.VkCopyImageToBufferInfo2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_COPY_IMAGE_TO_BUFFER_INFO_2
	info.pNext = next
	info.srcImage = C.VkImage(unsafe.Pointer(o.SrcImage.Handle()))
	info.srcImageLayout = C.VkImageLayout(o.SrcImageLayout)
	info.dstBuffer = C.VkBuffer(unsafe.Pointer(o.DstBuffer.Handle()))
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
