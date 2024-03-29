package khr_get_physical_device_properties2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"unsafe"
)

// ImageFormatProperties2 specifies image format properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2KHR.html
type ImageFormatProperties2 struct {
	// ImageFormatProperties is a structure in which capabilities are returned
	ImageFormatProperties core1_0.ImageFormatProperties

	common.NextOutData
}

func (o *ImageFormatProperties2) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageFormatProperties2KHR{})))
	}

	data := (*C.VkImageFormatProperties2KHR)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2_KHR
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *ImageFormatProperties2) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkImageFormatProperties2KHR)(cDataPointer)
	o.ImageFormatProperties.MaxExtent = core1_0.Extent3D{
		Width:  int(data.imageFormatProperties.maxExtent.width),
		Height: int(data.imageFormatProperties.maxExtent.height),
		Depth:  int(data.imageFormatProperties.maxExtent.depth),
	}
	o.ImageFormatProperties.MaxMipLevels = int(data.imageFormatProperties.maxMipLevels)
	o.ImageFormatProperties.MaxArrayLayers = int(data.imageFormatProperties.maxArrayLayers)
	o.ImageFormatProperties.SampleCounts = core1_0.SampleCountFlags(data.imageFormatProperties.sampleCounts)
	o.ImageFormatProperties.MaxResourceSize = int(data.imageFormatProperties.maxResourceSize)

	return data.pNext, nil
}
