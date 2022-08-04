package khr_external_fence_capabilities

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

// PhysicalDeviceExternalFenceInfo specifies Fence creation parameters
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalFenceInfo.html
type PhysicalDeviceExternalFenceInfo struct {
	// HandleType specifies an external Fence handle type for which capabilities will be
	// returned
	HandleType ExternalFenceHandleTypeFlags

	common.NextOptions
}

func (o PhysicalDeviceExternalFenceInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceExternalFenceInfoKHR{})))
	}
	info := (*C.VkPhysicalDeviceExternalFenceInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_FENCE_INFO_KHR
	info.pNext = next
	info.handleType = C.VkExternalFenceHandleTypeFlagBits(o.HandleType)

	return preallocatedPointer, nil
}
