package khr_separate_depth_stencil_layouts

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

// PhysicalDeviceSeparateDepthStencilLayoutsFeatures describes whether the implementation
// can do depth and stencil Image barriers separately
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures.html
type PhysicalDeviceSeparateDepthStencilLayoutsFeatures struct {
	// SeparateDepthStencilLayouts indicates whether the implementation supports an
	// ImageMemoryBarrier for a depth/stencil Image with only one of core1_0.ImageAspectDepth or
	// core1_0.ImageAspectStencil, and whether ImageLayoutDepthAttachmentOptimal,
	// ImageLayoutDepthReadOnlyOptimal, ImageLayoutStencilAttachmentOptimal, or
	// ImageLayoutStencilReadOnlyOptimal can be used
	SeparateDepthStencilLayouts bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceSeparateDepthStencilLayoutsFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderAtomicInt64FeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SEPARATE_DEPTH_STENCIL_LAYOUTS_FEATURES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceSeparateDepthStencilLayoutsFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR)(cDataPointer)

	o.SeparateDepthStencilLayouts = info.separateDepthStencilLayouts != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceSeparateDepthStencilLayoutsFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceShaderAtomicInt64FeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SEPARATE_DEPTH_STENCIL_LAYOUTS_FEATURES_KHR
	info.pNext = next
	info.separateDepthStencilLayouts = C.VkBool32(0)

	if o.SeparateDepthStencilLayouts {
		info.separateDepthStencilLayouts = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
