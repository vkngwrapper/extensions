package khr_device_group

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

// MemoryAllocateFlagsInfo controls how many instances of memory will be allocated
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagsInfoKHR.html
type MemoryAllocateFlagsInfo struct {
	// Flags controls the allocation
	Flags MemoryAllocateFlags
	// DeviceMask is a mask of PhysicalDevice objects in the logical Device, indicating that
	// memory must be allocated on each Device in the mask, if MemoryAllocateDeviceMask is set
	// in flags
	DeviceMask uint32

	common.NextOptions
}

func (o MemoryAllocateFlagsInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkMemoryAllocateFlagsInfoKHR{})))
	}

	createInfo := (*C.VkMemoryAllocateFlagsInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO_KHR
	createInfo.pNext = next
	createInfo.flags = C.VkMemoryAllocateFlags(o.Flags)
	createInfo.deviceMask = C.uint32_t(o.DeviceMask)

	return preallocatedPointer, nil
}
