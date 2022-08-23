package khr_sampler_ycbcr_conversion

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

// SamplerYcbcrConversionInfo specifies a Y'CbCr conversion to a Sampler or ImageView
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionInfo.html
type SamplerYcbcrConversionInfo struct {
	// Conversion is a SamplerYcbcrConversion object created from the Device
	Conversion SamplerYcbcrConversion

	common.NextOptions
}

func (o SamplerYcbcrConversionInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSamplerYcbcrConversionInfoKHR{})))
	}

	info := (*C.VkSamplerYcbcrConversionInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO_KHR
	info.pNext = next
	info.conversion = C.VkSamplerYcbcrConversion(unsafe.Pointer(o.Conversion.Handle()))

	return preallocatedPointer, nil
}
