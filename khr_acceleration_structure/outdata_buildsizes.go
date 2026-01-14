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

type AccelerationStructureBuildSizesInfo struct {
	AccelerationStructureSize int
	UpdateScratchSize         int
	BuildScratchSize          int

	common.NextOutData
}

func (o *AccelerationStructureBuildSizesInfo) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAccelerationStructureBuildSizesInfoKHR{})))
	}

	info := (*C.VkAccelerationStructureBuildSizesInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_BUILD_SIZES_INFO_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *AccelerationStructureBuildSizesInfo) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkAccelerationStructureBuildSizesInfoKHR)(cDataPointer)

	o.AccelerationStructureSize = int(info.accelerationStructureSize)
	o.UpdateScratchSize = int(info.updateScratchSize)
	o.BuildScratchSize = int(info.buildScratchSize)

	return info.pNext, nil
}
