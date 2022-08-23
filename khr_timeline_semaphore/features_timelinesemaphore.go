package khr_timeline_semaphore

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

// PhysicalDeviceTimelineSemaphoreFeatures describes timeline Semaphore features that can be
// supported by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTimelineSemaphoreFeatures.html
type PhysicalDeviceTimelineSemaphoreFeatures struct {
	// TimelineSemaphore indicates whether Semaphore objects created with a SemaphoreType
	// of SemaphoreTypeTimeline are supported
	TimelineSemaphore bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceTimelineSemaphoreFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceTimelineSemaphoreFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceTimelineSemaphoreFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_FEATURES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceTimelineSemaphoreFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceTimelineSemaphoreFeaturesKHR)(cDataPointer)

	o.TimelineSemaphore = info.timelineSemaphore != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceTimelineSemaphoreFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceTimelineSemaphoreFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceTimelineSemaphoreFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_FEATURES_KHR
	info.pNext = next
	info.timelineSemaphore = C.VkBool32(0)

	if o.TimelineSemaphore {
		info.timelineSemaphore = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
