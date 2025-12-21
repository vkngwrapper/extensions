package khr_create_renderpass2

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

// SubpassBeginInfo specifies subpass begin information
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassBeginInfoKHR.html
type SubpassBeginInfo struct {
	// Contents specifies how the commands in the next subpass will be provided
	Contents core1_0.SubpassContents

	common.NextOptions
}

func (o SubpassBeginInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassBeginInfoKHR{})))
	}

	info := (*C.VkSubpassBeginInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_BEGIN_INFO_KHR
	info.pNext = next
	info.contents = C.VkSubpassContents(o.Contents)

	return preallocatedPointer, nil
}
