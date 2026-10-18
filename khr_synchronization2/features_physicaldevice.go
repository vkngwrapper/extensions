package khr_synchronization2

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

type PhysicalDeviceSynchronization2Features struct {
	Synchronization2 bool

	common.NextOptions
	common.NextOutData
}

func (o PhysicalDeviceSynchronization2Features) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceSynchronization2Features{})))
	}

	features := (*C.VkPhysicalDeviceSynchronization2Features)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SYNCHRONIZATION_2_FEATURES
	features.pNext = next
	features.synchronization2 = C.VkBool32(0)
	if o.Synchronization2 {
		features.synchronization2 = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceSynchronization2Features) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceSynchronization2Features{})))
	}

	features := (*C.VkPhysicalDeviceSynchronization2Features)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SYNCHRONIZATION_2_FEATURES
	features.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceSynchronization2Features) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	features := (*C.VkPhysicalDeviceSynchronization2Features)(cDataPointer)
	o.Synchronization2 = features.synchronization2 != C.VkBool32(0)

	return features.pNext, nil
}
