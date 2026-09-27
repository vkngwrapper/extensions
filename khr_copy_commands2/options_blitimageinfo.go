package khr_copy_commands2

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

type BlitImageInfo2 struct {
	SrcImage       core1_0.Image
	SrcImageLayout core1_0.ImageLayout
	DstImage       core1_0.Image
	DstImageLayout core1_0.ImageLayout

	Regions []ImageBlit2
	Filter  core1_0.Filter

	common.NextOptions
}

func (o BlitImageInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if !o.SrcImage.Initialized() {
		return nil, errors.New("khr_copy_commands2.BlitImageInfo2.SrcImage cannot be left unset")
	}
	if !o.DstImage.Initialized() {
		return nil, errors.New("khr_copy_commands2.BlitImageInfo2.DstImage cannot be left unset")
	}
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBlitImageInfo2KHR{})))
	}

	info := (*C.VkBlitImageInfo2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_BLIT_IMAGE_INFO_2
	info.pNext = next
	info.srcImage = C.VkImage(unsafe.Pointer(o.SrcImage.Handle()))
	info.srcImageLayout = C.VkImageLayout(o.SrcImageLayout)
	info.dstImage = C.VkImage(unsafe.Pointer(o.DstImage.Handle()))
	info.dstImageLayout = C.VkImageLayout(o.DstImageLayout)
	info.regionCount = C.uint32_t(len(o.Regions))
	info.pRegions = nil
	info.filter = C.VkFilter(o.Filter)

	if len(o.Regions) > 0 {
		regionPtr, err := common.AllocOptionSlice[C.VkImageBlit2KHR](allocator, o.Regions)
		if err != nil {
			return nil, err
		}
		info.pRegions = regionPtr
	}

	return preallocatedPointer, nil
}
