package khr_create_renderpass2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

// SubpassEndInfo specifies subpass end information
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassEndInfo.html
type SubpassEndInfo struct {
	common.NextOptions
}

func (o SubpassEndInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassEndInfoKHR{})))
	}

	info := (*C.VkSubpassEndInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_END_INFO_KHR
	info.pNext = next

	return preallocatedPointer, nil
}
