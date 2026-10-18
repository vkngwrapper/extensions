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

type ImageMemoryBarrier2 struct {
	SrcStageMask  PipelineStageFlags2
	SrcAccessMask AccessFlags2
	DstStageMask  PipelineStageFlags2
	DstAccessMask AccessFlags2

	OldLayout core1_0.ImageLayout
	NewLayout core1_0.ImageLayout

	SrcQueueFamilyIndex int
	DstQueueFamilyIndex int

	Image            core1_0.Image
	SubresourceRange core1_0.ImageSubresourceRange

	common.NextOptions
}

func (o ImageMemoryBarrier2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if !o.Image.Initialized() {
		return nil, errors.New("khr_synchronization2.ImageMemoryBarrier2.Image cannot be uninitialized")
	}
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkImageMemoryBarrier2{})))
	}

	barrier := (*C.VkImageMemoryBarrier2)(preallocatedPointer)
	barrier.sType = C.VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER_2
	barrier.pNext = next
	barrier.srcStageMask = C.VkPipelineStageFlags2(o.SrcStageMask)
	barrier.srcAccessMask = C.VkAccessFlags2(o.SrcAccessMask)
	barrier.dstStageMask = C.VkPipelineStageFlags2(o.DstStageMask)
	barrier.dstAccessMask = C.VkAccessFlags2(o.DstAccessMask)
	barrier.oldLayout = C.VkImageLayout(o.OldLayout)
	barrier.newLayout = C.VkImageLayout(o.NewLayout)
	barrier.srcQueueFamilyIndex = C.uint32_t(o.SrcQueueFamilyIndex)
	barrier.dstQueueFamilyIndex = C.uint32_t(o.DstQueueFamilyIndex)
	barrier.image = C.VkImage(unsafe.Pointer(o.Image.Handle()))
	barrier.subresourceRange.aspectMask = C.VkImageAspectFlags(o.SubresourceRange.AspectMask)
	barrier.subresourceRange.baseMipLevel = C.uint32_t(o.SubresourceRange.BaseMipLevel)
	barrier.subresourceRange.baseArrayLayer = C.uint32_t(o.SubresourceRange.BaseArrayLayer)
	barrier.subresourceRange.layerCount = C.uint32_t(o.SubresourceRange.LayerCount)
	barrier.subresourceRange.levelCount = C.uint32_t(o.SubresourceRange.LevelCount)

	return preallocatedPointer, nil
}
