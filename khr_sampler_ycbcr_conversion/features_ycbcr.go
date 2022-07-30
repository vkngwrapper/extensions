package khr_sampler_ycbcr_conversion

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

// PhysicalDeviceSamplerYcbcrConversionFeatures describes Y'CbCr conversion features that can
// be supported by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSamplerYcbcrConversionFeatures.html
type PhysicalDeviceSamplerYcbcrConversionFeatures struct {
	// SamplerYcbcrConversion specifies whether the implementation support sampler Y'CbCr conversion.
	SamplerYcbcrConversion bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceSamplerYcbcrConversionFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceSamplerYcbcrConversionFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR)(cDataPointer)
	o.SamplerYcbcrConversion = info.samplerYcbcrConversion != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceSamplerYcbcrConversionFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES_KHR
	info.pNext = next
	info.samplerYcbcrConversion = C.VkBool32(0)

	if o.SamplerYcbcrConversion {
		info.samplerYcbcrConversion = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
