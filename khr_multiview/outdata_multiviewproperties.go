package khr_multiview

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

// PhysicalDeviceMultiviewProperties describes multiview limits that can be supported by an
// implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiviewProperties.html
type PhysicalDeviceMultiviewProperties struct {
	// MaxMultiviewViewCount is one greater than the maximum view index that can be used in
	// a subpass
	MaxMultiviewViewCount int
	// MaxMultiviewInstanceIndex is the maximum
	MaxMultiviewInstanceIndex int

	common.NextOutData
}

func (o *PhysicalDeviceMultiviewProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceMultiviewPropertiesKHR{})))
	}

	info := (*C.VkPhysicalDeviceMultiviewPropertiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceMultiviewProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceMultiviewPropertiesKHR)(cDataPointer)
	o.MaxMultiviewViewCount = int(info.maxMultiviewViewCount)
	o.MaxMultiviewInstanceIndex = int(info.maxMultiviewInstanceIndex)

	return info.pNext, nil
}
