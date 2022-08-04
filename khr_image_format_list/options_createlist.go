package khr_image_format_list

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

// ImageFormatListCreateInfo specifies that an Image can be used with a particular set of formats
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html
type ImageFormatListCreateInfo struct {
	// ViewFormats is a slice of core1_0.Format values specifying all formats which can be used
	// when creating views of this Image
	ViewFormats []core1_0.Format

	common.NextOptions
}

func (o ImageFormatListCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageFormatListCreateInfoKHR{})))
	}

	info := (*C.VkImageFormatListCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO_KHR
	info.pNext = next

	count := len(o.ViewFormats)
	info.viewFormatCount = C.uint32_t(count)
	info.pViewFormats = nil

	if count > 0 {
		info.pViewFormats = (*C.VkFormat)(allocator.Malloc(count * int(unsafe.Sizeof(C.VkFormat(0)))))
		viewFormatSlice := unsafe.Slice(info.pViewFormats, count)

		for i := 0; i < count; i++ {
			viewFormatSlice[i] = C.VkFormat(o.ViewFormats[i])
		}
	}

	return preallocatedPointer, nil
}
