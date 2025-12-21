package khr_timeline_semaphore

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

// PhysicalDeviceTimelineSemaphoreProperties describes timeline Semaphore properties that
// can be supported by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTimelineSemaphoreProperties.html
type PhysicalDeviceTimelineSemaphoreProperties struct {
	// MaxTimelineSemaphoreValueDifference indicates the maximum difference allowed by the
	// implementation between the current value of a timeline Semaphore and any pending signal or
	// wait operations
	MaxTimelineSemaphoreValueDifference uint64

	common.NextOutData
}

func (o *PhysicalDeviceTimelineSemaphoreProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceTimelineSemaphorePropertiesKHR{})))
	}

	info := (*C.VkPhysicalDeviceTimelineSemaphorePropertiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_PROPERTIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceTimelineSemaphoreProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceTimelineSemaphorePropertiesKHR)(cDataPointer)

	o.MaxTimelineSemaphoreValueDifference = uint64(info.maxTimelineSemaphoreValueDifference)

	return info.pNext, nil
}
