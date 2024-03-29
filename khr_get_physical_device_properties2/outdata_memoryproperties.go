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

// PhysicalDeviceMemoryProperties2 specifies PhysicalDevice memory properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties2.html
type PhysicalDeviceMemoryProperties2 struct {
	// MemoryProperties is a structure which is populated with the same values as in
	// PhysicalDevice.MemoryProperties
	MemoryProperties core1_0.PhysicalDeviceMemoryProperties

	common.NextOutData
}

func (o *PhysicalDeviceMemoryProperties2) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceMemoryProperties2KHR{})))
	}
	data := (*C.VkPhysicalDeviceMemoryProperties2KHR)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2_KHR
	data.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceMemoryProperties2) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	data := (*C.VkPhysicalDeviceMemoryProperties2KHR)(cDataPointer)

	memoryTypeCount := int(data.memoryProperties.memoryTypeCount)
	o.MemoryProperties.MemoryTypes = make([]core1_0.MemoryType, memoryTypeCount)

	for i := 0; i < memoryTypeCount; i++ {
		o.MemoryProperties.MemoryTypes[i].PropertyFlags = core1_0.MemoryPropertyFlags(data.memoryProperties.memoryTypes[i].propertyFlags)
		o.MemoryProperties.MemoryTypes[i].HeapIndex = int(data.memoryProperties.memoryTypes[i].heapIndex)
	}

	memoryHeapCount := int(data.memoryProperties.memoryHeapCount)
	o.MemoryProperties.MemoryHeaps = make([]core1_0.MemoryHeap, memoryHeapCount)

	for i := 0; i < memoryHeapCount; i++ {
		o.MemoryProperties.MemoryHeaps[i].Size = int(data.memoryProperties.memoryHeaps[i].size)
		o.MemoryProperties.MemoryHeaps[i].Flags = core1_0.MemoryHeapFlags(data.memoryProperties.memoryHeaps[i].flags)
	}

	return data.pNext, nil
}
