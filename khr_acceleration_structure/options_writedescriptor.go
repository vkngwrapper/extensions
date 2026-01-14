package khr_acceleration_structure

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

type WriteDescriptorSetAccelerationStructure struct {
	AccelerationStructures []AccelerationStructure

	common.NextOptions
}

func (o WriteDescriptorSetAccelerationStructure) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkWriteDescriptorSetAccelerationStructureKHR{})))
	}

	info := (*C.VkWriteDescriptorSetAccelerationStructureKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_ACCELERATION_STRUCTURE_KHR
	info.pNext = next

	count := len(o.AccelerationStructures)
	info.accelerationStructureCount = C.uint32_t(count)

	ptr := (*C.VkAccelerationStructureKHR)(allocator.Malloc(int(unsafe.Sizeof([1]C.VkAccelerationStructureKHR{})) * count))
	structureSlice := unsafe.Slice(ptr, count)

	for index, structure := range o.AccelerationStructures {
		structureSlice[index] = C.VkAccelerationStructureKHR(unsafe.Pointer(structure.Handle()))
	}
	info.pAccelerationStructures = ptr

	return preallocatedPointer, nil
}
