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

type AccelerationStructureBuildGeometryInfo struct {
	Type  AccelerationStructureType
	Flags BuildAccelerationStructureFlags
	Mode  BuildAccelerationStructureMode

	SrcAccelerationStructure AccelerationStructure
	DstAccelerationStructure AccelerationStructure

	Geometries []AccelerationStructureGeometry

	ScratchData DeviceOrHostAddressConst

	common.NextOptions
}

func (o AccelerationStructureBuildGeometryInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAccelerationStructureBuildGeometryInfoKHR{})))
	}
	if !o.DstAccelerationStructure.Initialized() {
		return nil, errors.New("o.DstAccelerationStructure cannot be uninitialized")
	}

	info := (*C.VkAccelerationStructureBuildGeometryInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_BUILD_GEOMETRY_INFO_KHR
	info.pNext = next

	info._type = C.VkAccelerationStructureTypeKHR(o.Type)
	info.flags = C.VkBuildAccelerationStructureFlagsKHR(o.Flags)
	info.mode = C.VkBuildAccelerationStructureModeKHR(o.Mode)
	if o.SrcAccelerationStructure.Initialized() {
		info.srcAccelerationStructure = C.VkAccelerationStructureKHR(unsafe.Pointer(o.SrcAccelerationStructure.Handle()))
	}
	info.dstAccelerationStructure = C.VkAccelerationStructureKHR(unsafe.Pointer(o.DstAccelerationStructure.Handle()))
	info.geometryCount = C.uint32_t(len(o.Geometries))
	ptr, err := common.AllocOptionSlice[C.VkAccelerationStructureGeometryKHR](allocator, o.Geometries)
	if err != nil {
		return nil, err
	}
	info.pGeometries = ptr
	info.ppGeometries = nil
	if o.ScratchData != nil {
		o.ScratchData.PopulateAddressUnion(unsafe.Pointer(&info.scratchData))
	}

	return preallocatedPointer, nil
}
