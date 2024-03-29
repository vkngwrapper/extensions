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

// QueueFamilyProperties2 provides information about a Queue family
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties2KHR.html
type QueueFamilyProperties2 struct {
	// QueueFamilyProperties is populated with the same values as in
	// PhysicalDevice.QueueFamilyProperties
	QueueFamilyProperties core1_0.QueueFamilyProperties

	common.NextOutData
}

func (o *QueueFamilyProperties2) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkQueueFamilyProperties2KHR{})))
	}

	data := (*C.VkQueueFamilyProperties2KHR)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2_KHR
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *QueueFamilyProperties2) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkQueueFamilyProperties2KHR)(cDataPointer)

	o.QueueFamilyProperties.QueueFlags = core1_0.QueueFlags(data.queueFamilyProperties.queueFlags)
	o.QueueFamilyProperties.QueueCount = int(data.queueFamilyProperties.queueCount)
	o.QueueFamilyProperties.TimestampValidBits = uint32(data.queueFamilyProperties.timestampValidBits)
	o.QueueFamilyProperties.MinImageTransferGranularity = core1_0.Extent3D{
		Width:  int(data.queueFamilyProperties.minImageTransferGranularity.width),
		Height: int(data.queueFamilyProperties.minImageTransferGranularity.height),
		Depth:  int(data.queueFamilyProperties.minImageTransferGranularity.depth),
	}

	return data.pNext, nil
}
