package khr_shader_subgroup_extended_types

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

// PhysicalDeviceShaderSubgroupExtendedTypesFeatures describes the extended types subgroups
// support feature for an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures.html
type PhysicalDeviceShaderSubgroupExtendedTypesFeatures struct {
	// ShaderSubgroupExtendedTypes specifies whether subgroup operations can use 8-bit integer,
	// 16-bit integer, 64-bit integer, 16-bit floating-point, and vectors of these types
	// in group operations with subgroup scope, if the implementation supports the types
	ShaderSubgroupExtendedTypes bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceShaderSubgroupExtendedTypesFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderSubgroupExtendedTypesFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceShaderSubgroupExtendedTypesFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_EXTENDED_TYPES_FEATURES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceShaderSubgroupExtendedTypesFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceShaderSubgroupExtendedTypesFeaturesKHR)(cDataPointer)

	o.ShaderSubgroupExtendedTypes = info.shaderSubgroupExtendedTypes != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceShaderSubgroupExtendedTypesFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderSubgroupExtendedTypesFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceShaderSubgroupExtendedTypesFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_EXTENDED_TYPES_FEATURES_KHR
	info.pNext = next
	info.shaderSubgroupExtendedTypes = C.VkBool32(0)

	if o.ShaderSubgroupExtendedTypes {
		info.shaderSubgroupExtendedTypes = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
