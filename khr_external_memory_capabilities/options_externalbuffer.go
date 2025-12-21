package khr_external_memory_capabilities

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

// PhysicalDeviceExternalBufferInfo specifies Buffer creation parameters
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalBufferInfo.html
type PhysicalDeviceExternalBufferInfo struct {
	// Flags describes additional parameters of the Buffer, corresponding to BufferCreateInfo.Flags
	Flags core1_0.BufferCreateFlags
	// Usage describes the intended usage of the Buffer, corresponding to BufferCreateInfo.Usage
	Usage core1_0.BufferUsageFlags
	// HandleType specifies the memory handle type that will be used with the memory
	// associated with the Buffer
	HandleType ExternalMemoryHandleTypeFlags

	common.NextOptions
}

func (o PhysicalDeviceExternalBufferInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceExternalBufferInfoKHR{})))
	}

	info := (*C.VkPhysicalDeviceExternalBufferInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_BUFFER_INFO_KHR
	info.pNext = next
	info.flags = (C.VkBufferCreateFlags)(o.Flags)
	info.usage = (C.VkBufferUsageFlags)(o.Usage)
	info.handleType = (C.VkExternalMemoryHandleTypeFlagBits)(o.HandleType)

	return preallocatedPointer, nil
}
