package khr_zero_initialize_workgroup_memory

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
)

type PhysicalDeviceZeroInitializeWorkgroupMemoryFeatures struct {
	ShaderZeroInitializeWorkgroupMemory bool

	common.NextOutData
	common.NextOptions
}

func (o PhysicalDeviceZeroInitializeWorkgroupMemoryFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeaturesKHR{})))
	}

	features := (*C.VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeaturesKHR)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ZERO_INITIALIZE_WORKGROUP_MEMORY_FEATURES
	features.pNext = next
	features.shaderZeroInitializeWorkgroupMemory = C.VkBool32(0)
	if o.ShaderZeroInitializeWorkgroupMemory {
		features.shaderZeroInitializeWorkgroupMemory = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceZeroInitializeWorkgroupMemoryFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeaturesKHR{})))
	}

	features := (*C.VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeaturesKHR)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ZERO_INITIALIZE_WORKGROUP_MEMORY_FEATURES
	features.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceZeroInitializeWorkgroupMemoryFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	features := (*C.VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeaturesKHR)(cDataPointer)
	o.ShaderZeroInitializeWorkgroupMemory = features.shaderZeroInitializeWorkgroupMemory != C.VkBool32(0)

	return features.pNext, nil
}
