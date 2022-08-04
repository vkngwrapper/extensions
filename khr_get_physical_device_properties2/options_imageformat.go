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

// PhysicalDeviceImageFormatInfo2 specifies Image creation parameters
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2KHR.html
type PhysicalDeviceImageFormatInfo2 struct {
	// Format indicates the Image format, corresponding to ImageCreateInfo.Format
	Format core1_0.Format
	// Type indicates the ImageType, corresponding to ImageCreateInfo.ImageType
	Type core1_0.ImageType
	// Tiling indicates the Image tiling, corresponding to ImageCreateInfo.Tiling
	Tiling core1_0.ImageTiling
	// Usage indicates the intended usage of the Image, corresponding to ImageCreateInfo.Usage
	Usage core1_0.ImageUsageFlags
	// Flags indicates additional parameters of the Image, corresponding to ImageCreateInfo.Flags
	Flags core1_0.ImageCreateFlags

	common.NextOptions
}

func (o PhysicalDeviceImageFormatInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceImageFormatInfo2KHR{})))
	}
	info := (*C.VkPhysicalDeviceImageFormatInfo2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2_KHR
	info.pNext = next
	info.format = C.VkFormat(o.Format)
	info._type = C.VkImageType(o.Type)
	info.tiling = C.VkImageTiling(o.Tiling)
	info.usage = C.VkImageUsageFlags(o.Usage)
	info.flags = C.VkImageCreateFlags(o.Flags)

	return preallocatedPointer, nil
}
