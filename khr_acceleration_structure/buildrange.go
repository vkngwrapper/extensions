package khr_acceleration_structure

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
)

type AccelerationStructureBuildRangeInfo struct {
	PrimitiveCount  int
	PrimitiveOffset int
	FirstVertex     int
	TransformOffset int
}

func (i AccelerationStructureBuildRangeInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]*C.VkAccelerationStructureBuildRangeInfoKHR{})))
	}

	info := (*C.VkAccelerationStructureBuildRangeInfoKHR)(preallocatedPointer)
	info.primitiveCount = C.uint32_t(i.PrimitiveCount)
	info.primitiveOffset = C.uint32_t(i.PrimitiveOffset)
	info.firstVertex = C.uint32_t(i.FirstVertex)
	info.transformOffset = C.uint32_t(i.TransformOffset)

	return preallocatedPointer, nil
}
