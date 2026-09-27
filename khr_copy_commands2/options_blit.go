package khr_copy_commands2

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

type ImageBlit2 struct {
	SrcSubresource core1_0.ImageSubresourceLayers
	SrcOffsets     [2]core1_0.Offset3D

	DstSubresource core1_0.ImageSubresourceLayers
	DstOffsets     [2]core1_0.Offset3D

	common.NextOptions
}

func (o ImageBlit2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageBlit2KHR{})))
	}

	info := (*C.VkImageBlit2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_IMAGE_BLIT_2
	info.pNext = next

	info.srcSubresource.aspectMask = C.VkImageAspectFlags(o.SrcSubresource.AspectMask)
	info.srcSubresource.mipLevel = C.uint32_t(o.SrcSubresource.MipLevel)
	info.srcSubresource.baseArrayLayer = C.uint32_t(o.SrcSubresource.BaseArrayLayer)
	info.srcSubresource.layerCount = C.uint32_t(o.SrcSubresource.LayerCount)

	info.dstSubresource.aspectMask = C.VkImageAspectFlags(o.DstSubresource.AspectMask)
	info.dstSubresource.mipLevel = C.uint32_t(o.DstSubresource.MipLevel)
	info.dstSubresource.baseArrayLayer = C.uint32_t(o.DstSubresource.BaseArrayLayer)
	info.dstSubresource.layerCount = C.uint32_t(o.DstSubresource.LayerCount)

	info.srcOffsets[0].x = C.int32_t(o.SrcOffsets[0].X)
	info.srcOffsets[0].y = C.int32_t(o.SrcOffsets[0].Y)
	info.srcOffsets[0].z = C.int32_t(o.SrcOffsets[0].Z)
	info.srcOffsets[1].x = C.int32_t(o.SrcOffsets[1].X)
	info.srcOffsets[1].y = C.int32_t(o.SrcOffsets[1].Y)
	info.srcOffsets[1].z = C.int32_t(o.SrcOffsets[1].Z)

	info.dstOffsets[0].x = C.int32_t(o.DstOffsets[0].X)
	info.dstOffsets[0].y = C.int32_t(o.DstOffsets[0].Y)
	info.dstOffsets[0].z = C.int32_t(o.DstOffsets[0].Z)
	info.dstOffsets[1].x = C.int32_t(o.DstOffsets[1].X)
	info.dstOffsets[1].y = C.int32_t(o.DstOffsets[1].Y)
	info.dstOffsets[1].z = C.int32_t(o.DstOffsets[1].Z)

	return preallocatedPointer, nil
}
