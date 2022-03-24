package khr_get_physical_device_properties2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type QueueFamilyOutData struct {
	QueueFamily core1_0.QueueFamily

	common.HaveNext
}

func (o *QueueFamilyOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkQueueFamilyProperties2KHR{})))
	}

	data := (*C.VkQueueFamilyProperties2KHR)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2_KHR
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *QueueFamilyOutData) PopulateOutData(cDataPointer unsafe.Pointer) (next unsafe.Pointer, err error) {
	data := (*C.VkQueueFamilyProperties2KHR)(cDataPointer)

	o.QueueFamily.Flags = common.QueueFlags(data.queueFamilyProperties.queueFlags)
	o.QueueFamily.QueueCount = int(data.queueFamilyProperties.queueCount)
	o.QueueFamily.TimestampValidBits = uint32(data.queueFamilyProperties.timestampValidBits)
	o.QueueFamily.MinImageTransferGranularity = common.Extent3D{
		Width:  int(data.queueFamilyProperties.minImageTransferGranularity.width),
		Height: int(data.queueFamilyProperties.minImageTransferGranularity.height),
		Depth:  int(data.queueFamilyProperties.minImageTransferGranularity.depth),
	}

	return data.pNext, nil
}
