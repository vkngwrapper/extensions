package khr_variable_pointers

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

// PhysicalDeviceVariablePointersFeatures describes variable pointer features that can be
// supported by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVariablePointersFeatures.html
type PhysicalDeviceVariablePointersFeatures struct {
	// VariablePointersStorageBuffer specifies whether the implementation supports the SPIR-V
	// VariablePointersStorageBuffer capability
	VariablePointersStorageBuffer bool
	// VariablePointers specifies whether the implementation supports the SPIR-V VariablePointers
	// capability
	VariablePointers bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceVariablePointersFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceVariablePointersFeaturesKHR{})))
	}

	createInfo := (*C.VkPhysicalDeviceVariablePointersFeaturesKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES_KHR
	createInfo.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceVariablePointersFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkPhysicalDeviceVariablePointersFeaturesKHR)(cDataPointer)
	o.VariablePointers = false
	o.VariablePointersStorageBuffer = false

	if createInfo.variablePointersStorageBuffer != C.VkBool32(0) {
		o.VariablePointersStorageBuffer = true
	}
	if createInfo.variablePointers != C.VkBool32(0) {
		o.VariablePointers = true
	}

	return createInfo.pNext, nil
}

func (o PhysicalDeviceVariablePointersFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceVariablePointersFeaturesKHR{})))
	}

	createInfo := (*C.VkPhysicalDeviceVariablePointersFeaturesKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES_KHR
	createInfo.pNext = next
	createInfo.variablePointersStorageBuffer = C.VkBool32(0)
	createInfo.variablePointers = C.VkBool32(0)

	if o.VariablePointersStorageBuffer {
		createInfo.variablePointersStorageBuffer = C.VkBool32(1)
	}
	if o.VariablePointers {
		createInfo.variablePointers = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
