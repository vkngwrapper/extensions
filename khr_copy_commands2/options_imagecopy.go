package khr_copy_commands2

import (
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

type ImageCopy2 struct {
	// SrcSubresource specifies the Image subresources of the Image objects used for the
	// source Image data
	SrcSubresource core1_0.ImageSubresourceLayers
	// SrcOffset selects the initial x, y, and z offsets in texels of the sub-regions of the
	// source Image data
	SrcOffset core1_0.Offset3D
	// DstSubresource specifies the Image subresource of the Image objects used for the
	// destination Image data
	DstSubresource core1_0.ImageSubresourceLayers
	// DstOffset selects the initial x, y, and z offsets in texels of the sub-regions of the
	// destination Image data
	DstOffset core1_0.Offset3D
	// Extent is the size in texels of the Image to copy in width, height, and depth
	Extent core1_0.Extent3D

	common.NextOptions
}

func (o ImageCopy2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageCopy2KHR{})))
	}

	info := (*C.VkImageCopy2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_IMAGE_COPY_2
	info.pNext = next
	info.srcSubresource.aspectMask = C.VkImageAspectFlags(o.SrcSubresource.AspectMask)
	info.srcSubresource.mipLevel = C.uint32_t(o.SrcSubresource.MipLevel)
	info.srcSubresource.baseArrayLayer = C.uint32_t(o.SrcSubresource.BaseArrayLayer)
	info.srcSubresource.layerCount = C.uint32_t(o.SrcSubresource.LayerCount)

	info.dstSubresource.aspectMask = C.VkImageAspectFlags(o.DstSubresource.AspectMask)
	info.dstSubresource.mipLevel = C.uint32_t(o.DstSubresource.MipLevel)
	info.dstSubresource.baseArrayLayer = C.uint32_t(o.DstSubresource.BaseArrayLayer)
	info.dstSubresource.layerCount = C.uint32_t(o.DstSubresource.LayerCount)

	info.srcOffset.x = C.int32_t(o.SrcOffset.X)
	info.srcOffset.y = C.int32_t(o.SrcOffset.Y)
	info.srcOffset.z = C.int32_t(o.SrcOffset.Z)

	info.dstOffset.x = C.int32_t(o.DstOffset.X)
	info.dstOffset.y = C.int32_t(o.DstOffset.Y)
	info.dstOffset.z = C.int32_t(o.DstOffset.Z)

	info.extent.width = C.uint32_t(o.Extent.Width)
	info.extent.height = C.uint32_t(o.Extent.Height)
	info.extent.depth = C.uint32_t(o.Extent.Depth)

	return preallocatedPointer, nil
}
