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

type BufferImageCopy2 struct {
	// BufferOffset is the offset in bytes from the start of the Buffer
	BufferOffset int
	// BufferRowLength is the size in texels of the rows of the image stored in the Buffer.
	// 0 indicates that the ImageExtent controls this value
	BufferRowLength int
	// BufferImageHeight is the height in texels of the image stored in the Buffer
	// 0 indicates that the ImageExtent controls this value
	BufferImageHeight int

	// ImageSubresource is used to specify the specific image subresources of the Image
	ImageSubresource core1_0.ImageSubresourceLayers
	// ImageOffset selects the initial x, y, and z offset in texels of the Image subregion
	ImageOffset core1_0.Offset3D
	// ImageExtent is the size in texels of the Image subregion
	ImageExtent core1_0.Extent3D

	common.NextOptions
}

func (o BufferImageCopy2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBufferImageCopy2KHR{})))
	}

	info := (*C.VkBufferImageCopy2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_BUFFER_IMAGE_COPY_2
	info.pNext = next
	info.bufferOffset = C.VkDeviceSize(o.BufferOffset)
	info.bufferRowLength = C.uint32_t(o.BufferRowLength)
	info.bufferImageHeight = C.uint32_t(o.BufferImageHeight)
	info.imageSubresource.aspectMask = C.VkImageAspectFlags(o.ImageSubresource.AspectMask)
	info.imageSubresource.mipLevel = C.uint32_t(o.ImageSubresource.MipLevel)
	info.imageSubresource.baseArrayLayer = C.uint32_t(o.ImageSubresource.BaseArrayLayer)
	info.imageSubresource.layerCount = C.uint32_t(o.ImageSubresource.LayerCount)
	info.imageOffset.x = C.int32_t(o.ImageOffset.X)
	info.imageOffset.y = C.int32_t(o.ImageOffset.Y)
	info.imageOffset.z = C.int32_t(o.ImageOffset.Z)
	info.imageExtent.width = C.uint32_t(o.ImageExtent.Width)
	info.imageExtent.height = C.uint32_t(o.ImageExtent.Height)
	info.imageExtent.depth = C.uint32_t(o.ImageExtent.Depth)

	return preallocatedPointer, nil
}
