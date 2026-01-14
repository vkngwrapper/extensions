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

type AccelerationStructureVersionInfo struct {
	VersionData unsafe.Pointer

	common.NextOptions
}

func (o AccelerationStructureVersionInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAccelerationStructureVersionInfoKHR{})))
	}

	info := (*C.VkAccelerationStructureVersionInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_VERSION_INFO_KHR
	info.pNext = next
	info.pVersionData = (*C.uint8_t)(o.VersionData)

	return preallocatedPointer, nil
}
