package khr_uniform_buffer_standard_layout

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

// PhysicalDeviceUniformBufferStandardLayoutFeatures indicates support for std430-like
// packing in uniform Buffer objects
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceUniformBufferStandardLayoutFeatures.html
type PhysicalDeviceUniformBufferStandardLayoutFeatures struct {
	// UniformBufferStandardLayout indicates that the implementation supports the same layouts
	// for uniform Buffer objects as for storage and other kinds of Buffer objects
	UniformBufferStandardLayout bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceUniformBufferStandardLayoutFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceUniformBufferStandardLayoutFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceUniformBufferStandardLayoutFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFORM_BUFFER_STANDARD_LAYOUT_FEATURES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceUniformBufferStandardLayoutFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceUniformBufferStandardLayoutFeaturesKHR)(cDataPointer)

	o.UniformBufferStandardLayout = info.uniformBufferStandardLayout != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceUniformBufferStandardLayoutFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceUniformBufferStandardLayoutFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceUniformBufferStandardLayoutFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_UNIFORM_BUFFER_STANDARD_LAYOUT_FEATURES_KHR
	info.pNext = next
	info.uniformBufferStandardLayout = C.VkBool32(0)

	if o.UniformBufferStandardLayout {
		info.uniformBufferStandardLayout = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
