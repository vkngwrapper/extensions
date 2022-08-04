package khr_imageless_framebuffer

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"unsafe"
)

// FramebufferAttachmentImageInfo specifies parameters of an Image that will be used with a
// Framebuffer
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentImageInfo.html
type FramebufferAttachmentImageInfo struct {
	// Flags matches the value of ImageCreateInfo.Flags used to create an Image that will be used
	// with this Framebuffer
	Flags core1_0.ImageCreateFlags
	// Usage matches the value of ImageCreateInfo.Usage used to create an Image used with this
	// Framebuffer
	Usage core1_0.ImageUsageFlags
	// Width is the width of the ImageView used for rendering
	Width int
	// Height is the height of ImageView used for rendering
	Height int
	// LayerCount is the number of array layers of the ImageView used for rendering
	LayerCount int

	// ViewFormats is a slice of core1_0.Format values specifying all of the formats which
	// can be used when creating views of the Image, matching the value of
	// ImageFormatListCreateInfo.ViewFormats used to create an Image used with this
	// Framebuffer
	ViewFormats []core1_0.Format

	common.NextOptions
}

func (o FramebufferAttachmentImageInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkFramebufferAttachmentImageInfoKHR{})))
	}

	info := (*C.VkFramebufferAttachmentImageInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO_KHR
	info.pNext = next
	info.flags = C.VkImageCreateFlags(o.Flags)
	info.usage = C.VkImageUsageFlags(o.Usage)
	info.width = C.uint32_t(o.Width)
	info.height = C.uint32_t(o.Height)
	info.layerCount = C.uint32_t(o.LayerCount)

	count := len(o.ViewFormats)
	info.viewFormatCount = C.uint32_t(count)
	info.pViewFormats = nil

	if count > 0 {
		info.pViewFormats = (*C.VkFormat)(allocator.Malloc(count * int(unsafe.Sizeof(C.VkFormat(0)))))
		viewFormatSlice := ([]C.VkFormat)(unsafe.Slice(info.pViewFormats, count))
		for i := 0; i < count; i++ {
			viewFormatSlice[i] = C.VkFormat(o.ViewFormats[i])
		}
	}

	return preallocatedPointer, nil
}
