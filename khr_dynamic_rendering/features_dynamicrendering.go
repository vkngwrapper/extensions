package khr_dynamic_rendering

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

// PhysicalDeviceDynamicRenderingFeatures describes support for dynamic rendering
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkPhysicalDeviceDynamicRenderingFeaturesKHR.html
type PhysicalDeviceDynamicRenderingFeatures struct {
	DynamicRendering bool

	common.NextOptions
	common.NextOutData
}

func (o PhysicalDeviceDynamicRenderingFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceDynamicRenderingFeaturesKHR{})))
	}
	info := (*C.VkPhysicalDeviceDynamicRenderingFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DYNAMIC_RENDERING_FEATURES_KHR
	info.pNext = next
	info.dynamicRendering = C.VkBool32(0)

	if o.DynamicRendering {
		info.dynamicRendering = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceDynamicRenderingFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceDynamicRenderingFeaturesKHR{})))
	}
	info := (*C.VkPhysicalDeviceDynamicRenderingFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DYNAMIC_RENDERING_FEATURES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceDynamicRenderingFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDeviceDynamicRenderingFeaturesKHR)(cDataPointer)
	o.DynamicRendering = outData.dynamicRendering != C.VkBool32(0)

	return outData.pNext, nil
}
