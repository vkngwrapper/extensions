package khr_acceleration_structure

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
)

type GeometryAABBData struct {
	Data   DeviceOrHostAddressConst
	Stride int

	common.NextOptions
}

func (a GeometryAABBData) IsGeometryType() {}

func (a GeometryAABBData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAccelerationStructureGeometryAabbsDataKHR{})))
	}

	if a.Stride%8 != 0 {
		return nil, errors.New("stride must be a multiple of 8")
	}

	info := (*C.VkAccelerationStructureGeometryAabbsDataKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_AABBS_DATA_KHR
	info.pNext = next
	info.stride = C.VkDeviceSize(a.Stride)
	if a.Data != nil {
		a.Data.PopulateAddressUnion(unsafe.Pointer(&info.data))
	}

	return preallocatedPointer, nil
}
