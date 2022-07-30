package khr_external_memory_capabilities

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

// PhysicalDeviceExternalImageFormatInfo specifies external Image creation parameters
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalImageFormatInfo.html
type PhysicalDeviceExternalImageFormatInfo struct {
	// HandleType specifies the memory handle type that will be used with the memory associated
	// with the Image
	HandleType ExternalMemoryHandleTypeFlags

	common.NextOptions
}

func (o PhysicalDeviceExternalImageFormatInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceExternalImageFormatInfoKHR{})))
	}

	info := (*C.VkPhysicalDeviceExternalImageFormatInfoKHR)(preallocatedPointer)

	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_IMAGE_FORMAT_INFO_KHR
	info.pNext = next
	info.handleType = C.VkExternalMemoryHandleTypeFlagBits(o.HandleType)

	return preallocatedPointer, nil
}
