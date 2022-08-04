package khr_get_physical_device_properties2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"unsafe"
)

// PhysicalDeviceSparseImageFormatInfo2 specifies sparse Image format inputs
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSparseImageFormatInfo2KHR.html
type PhysicalDeviceSparseImageFormatInfo2 struct {
	// Format is the Image format
	Format core1_0.Format
	// Type is the dimensionality of the Image
	Type core1_0.ImageType
	// Samples specifies the number of samples per texel
	Samples core1_0.SampleCountFlags
	// Usage describes the intended usage of the Image
	Usage core1_0.ImageUsageFlags
	// Tiling is the tiling arrangement of the texel blocks in memory
	Tiling core1_0.ImageTiling

	common.NextOptions
}

func (o PhysicalDeviceSparseImageFormatInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceSparseImageFormatInfo2KHR{})))
	}

	createInfo := (*C.VkPhysicalDeviceSparseImageFormatInfo2KHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2_KHR
	createInfo.pNext = next
	createInfo.format = C.VkFormat(o.Format)
	createInfo._type = C.VkImageType(o.Type)
	createInfo.samples = C.VkSampleCountFlagBits(o.Samples)
	createInfo.usage = C.VkImageUsageFlags(o.Usage)
	createInfo.tiling = C.VkImageTiling(o.Tiling)

	return preallocatedPointer, nil
}
