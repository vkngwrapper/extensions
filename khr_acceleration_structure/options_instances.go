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

type GeometryInstancesData struct {
	ArrayOfPointers bool
	Data            DeviceOrHostAddressConst

	common.NextOptions
}

func (d GeometryInstancesData) IsGeometryType() {}

func (d GeometryInstancesData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAccelerationStructureGeometryInstancesDataKHR{})))
	}

	info := (*C.VkAccelerationStructureGeometryInstancesDataKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_INSTANCES_DATA_KHR
	info.pNext = next
	info.arrayOfPointers = C.VkBool32(0)
	if d.ArrayOfPointers {
		info.arrayOfPointers = C.VkBool32(1)
	}
	if d.Data != nil {
		d.Data.PopulateAddressUnion(unsafe.Pointer(&info.data))
	}

	return preallocatedPointer, nil
}
