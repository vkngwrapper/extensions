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

type Geometry interface {
	common.Options
	IsGeometryType()
}

type AccelerationStructureGeometry struct {
	Type     GeometryType
	Flags    GeometryFlags
	Geometry Geometry

	common.NextOptions
}

func (g AccelerationStructureGeometry) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAccelerationStructureGeometryKHR{})))
	}

	info := (*C.VkAccelerationStructureGeometryKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_KHR
	info.pNext = next
	info.geometryType = C.VkGeometryTypeKHR(g.Type)
	info.flags = C.VkGeometryFlagsKHR(g.Flags)

	_, err := common.AllocOptions(allocator, g.Geometry, unsafe.Pointer(&info.geometry))
	if err != nil {
		return nil, err
	}

	return preallocatedPointer, nil
}
