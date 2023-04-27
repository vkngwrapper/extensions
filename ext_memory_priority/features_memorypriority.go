package ext_memory_priority

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

// PhysicalDeviceMemoryPriorityFeatures describes memory priority feature support
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryPriorityFeaturesEXT.html
type PhysicalDeviceMemoryPriorityFeatures struct {
	MemoryPriority bool

	common.NextOptions
	common.NextOutData
}

func (o PhysicalDeviceMemoryPriorityFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(C.sizeof_struct_VkPhysicalDeviceMemoryPriorityFeaturesEXT))
	}

	info := (*C.VkPhysicalDeviceMemoryPriorityFeaturesEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PRIORITY_FEATURES_EXT
	info.pNext = next
	info.memoryPriority = C.VkBool32(0)

	if o.MemoryPriority {
		info.memoryPriority = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceMemoryPriorityFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(C.sizeof_struct_VkPhysicalDeviceMemoryPriorityFeaturesEXT))
	}
	info := (*C.VkPhysicalDeviceMemoryPriorityFeaturesEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PRIORITY_FEATURES_EXT
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceMemoryPriorityFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDeviceMemoryPriorityFeaturesEXT)(cDataPointer)

	o.MemoryPriority = outData.memoryPriority != C.VkBool32(0)

	return outData.pNext, nil
}
