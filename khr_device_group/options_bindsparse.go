package khr_device_group

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

// DeviceGroupBindSparseInfo indicates which instances are bound
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupBindSparseInfo.html
type DeviceGroupBindSparseInfo struct {
	// ResourceDeviceIndex is a Device index indicating which instance of the resource is bound
	ResourceDeviceIndex int
	// MemoryDeviceIndex is a Device index indicating which instance of the memory the resource instance is bound to
	MemoryDeviceIndex int

	common.NextOptions
}

func (o DeviceGroupBindSparseInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceGroupBindSparseInfoKHR{})))
	}

	createInfo := (*C.VkDeviceGroupBindSparseInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO_KHR
	createInfo.pNext = next
	createInfo.resourceDeviceIndex = C.uint32_t(o.ResourceDeviceIndex)
	createInfo.memoryDeviceIndex = C.uint32_t(o.MemoryDeviceIndex)

	return preallocatedPointer, nil
}
