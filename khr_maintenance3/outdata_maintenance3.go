package khr_maintenance3

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

// PhysicalDeviceMaintenance3Properties describes DescriptorSet properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance3Properties.html
type PhysicalDeviceMaintenance3Properties struct {
	// MaxPerSetDescriptors is a maximum number of descriptors in a single DescriptorSet that is
	// guaranteed to satisfy any implementation-dependent constraints on the size of a
	// DescriptorSet itself
	MaxPerSetDescriptors int
	// MaxMemoryAllocationSize is the maximum size of a memory allocation that can be created,
	// even if the is more space available in the heap
	MaxMemoryAllocationSize int

	common.NextOutData
}

func (o *PhysicalDeviceMaintenance3Properties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceMaintenance3PropertiesKHR{})))
	}

	outData := (*C.VkPhysicalDeviceMaintenance3PropertiesKHR)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES_KHR
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceMaintenance3Properties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDeviceMaintenance3PropertiesKHR)(cDataPointer)

	o.MaxMemoryAllocationSize = int(outData.maxMemoryAllocationSize)
	o.MaxPerSetDescriptors = int(outData.maxPerSetDescriptors)

	return outData.pNext, nil
}
