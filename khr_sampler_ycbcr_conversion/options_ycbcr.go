package khr_sampler_ycbcr_conversion

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
)

// SamplerYcbcrConversionCreateInfo specifies the parameters of the newly-created conversion
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html
type SamplerYcbcrConversionCreateInfo struct {
	// Format is the format of the Image from which color information will be retrieved
	Format core1_0.Format
	// YcbcrModel describes the color matrix for conversion between color models
	YcbcrModel SamplerYcbcrModelConversion
	// YcbcrRange describes whether the encoded values have headroom and foot room, or whether
	// the encoding uses the full numerical range
	YcbcrRange SamplerYcbcrRange
	// Components applies a swizzle based on core1_0.ComponentSwizzle enums prior to range
	// expansion and color model conversion
	Components core1_0.ComponentMapping
	// XChromaOffset describes the sample location associated with downsampled chroma components
	// in the x dimension
	XChromaOffset ChromaLocation
	// YChromaOffset describes the sample location associated with downsampled chroma components
	// in the y dimension
	YChromaOffset ChromaLocation
	// ChromaFilter is the filter for chroma reconstruction
	ChromaFilter core1_0.Filter
	// ForceExplicitReconstruction can be used to ensure that reconstruction is done explicitly,
	// if supported
	ForceExplicitReconstruction bool

	common.NextOptions
}

func (o SamplerYcbcrConversionCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSamplerYcbcrConversionCreateInfoKHR{})))
	}

	info := (*C.VkSamplerYcbcrConversionCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO_KHR
	info.pNext = next
	info.format = C.VkFormat(o.Format)
	info.ycbcrModel = C.VkSamplerYcbcrModelConversion(o.YcbcrModel)
	info.ycbcrRange = C.VkSamplerYcbcrRange(o.YcbcrRange)
	info.components.r = C.VkComponentSwizzle(o.Components.R)
	info.components.g = C.VkComponentSwizzle(o.Components.G)
	info.components.b = C.VkComponentSwizzle(o.Components.B)
	info.components.a = C.VkComponentSwizzle(o.Components.A)
	info.xChromaOffset = C.VkChromaLocation(o.XChromaOffset)
	info.yChromaOffset = C.VkChromaLocation(o.YChromaOffset)
	info.chromaFilter = C.VkFilter(o.ChromaFilter)
	info.forceExplicitReconstruction = C.VkBool32(0)

	if o.ForceExplicitReconstruction {
		info.forceExplicitReconstruction = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
