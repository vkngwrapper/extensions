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

type PhysicalDeviceAccelerationStructureProperties struct {
	MaxGeometryCount                                           int
	MaxInstanceCount                                           int
	MaxPrimitiveCount                                          int
	MaxPerStageDescriptorAccelerationStructures                int
	MaxPerStageDescriptorUpdateAfterBindAccelerationStructures int
	MaxDescriptorSetAccelerationStructures                     int
	MaxDescriptorSetUpdateAfterBindAccelerationStructures      int
	MinAccelerationStructureScratchOffsetAlignment             int

	common.NextOutData
}

func (o *PhysicalDeviceAccelerationStructureProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceAccelerationStructurePropertiesKHR{})))
	}

	info := (*C.VkPhysicalDeviceAccelerationStructurePropertiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ACCELERATION_STRUCTURE_PROPERTIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceAccelerationStructureProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceAccelerationStructurePropertiesKHR)(cDataPointer)

	o.MaxGeometryCount = int(info.maxGeometryCount)
	o.MaxInstanceCount = int(info.maxInstanceCount)
	o.MaxPrimitiveCount = int(info.maxPrimitiveCount)
	o.MaxPerStageDescriptorAccelerationStructures = int(info.maxPerStageDescriptorAccelerationStructures)
	o.MaxPerStageDescriptorUpdateAfterBindAccelerationStructures = int(info.maxPerStageDescriptorUpdateAfterBindAccelerationStructures)
	o.MaxDescriptorSetAccelerationStructures = int(info.maxDescriptorSetAccelerationStructures)
	o.MaxDescriptorSetUpdateAfterBindAccelerationStructures = int(info.maxDescriptorSetUpdateAfterBindAccelerationStructures)
	o.MinAccelerationStructureScratchOffsetAlignment = int(info.minAccelerationStructureScratchOffsetAlignment)

	return info.pNext, nil
}
