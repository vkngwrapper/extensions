package khr_multiview

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

// PhysicalDeviceMultiviewFeatures describes multiview features that can be supported by
// an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiviewFeatures.html
type PhysicalDeviceMultiviewFeatures struct {
	// Multiview specifies whether the implementation supports multiview rendering within a
	// RenderPass. If this feature is not enabled, the view mask of each subpass must always
	// be zero
	Multiview bool
	// MultiviewGeometryShader specifies whether the implementation supports multiview rendering
	// within a RenderPass, with geometry shaders
	MultiviewGeometryShader bool
	// MultiviewTessellationShader specifies whether the implementation supports multiview rendering
	// within a RenderPass, with tessellation shaders
	MultiviewTessellationShader bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceMultiviewFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceMultiviewFeaturesKHR{})))
	}
	info := (*C.VkPhysicalDeviceMultiviewFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceMultiviewFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceMultiviewFeaturesKHR)(cDataPointer)

	o.Multiview = info.multiview != C.VkBool32(0)
	o.MultiviewGeometryShader = info.multiviewGeometryShader != C.VkBool32(0)
	o.MultiviewTessellationShader = info.multiviewTessellationShader != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceMultiviewFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceMultiviewFeaturesKHR{})))
	}
	info := (*C.VkPhysicalDeviceMultiviewFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES_KHR
	info.pNext = next
	info.multiview = C.VkBool32(0)
	info.multiviewGeometryShader = C.VkBool32(0)
	info.multiviewTessellationShader = C.VkBool32(0)

	if o.Multiview {
		info.multiview = C.VkBool32(1)
	}

	if o.MultiviewGeometryShader {
		info.multiviewGeometryShader = C.VkBool32(1)
	}

	if o.MultiviewTessellationShader {
		info.multiviewTessellationShader = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
