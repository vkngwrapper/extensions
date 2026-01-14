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

type CopyMemoryToAccelerationStructureInfo struct {
	Src  DeviceOrHostAddressConst
	Dst  AccelerationStructure
	Mode CopyAccelerationStructureMode

	common.NextOptions
}

func (o CopyMemoryToAccelerationStructureInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkCopyMemoryToAccelerationStructureInfoKHR{})))
	}
	if o.Src == nil {
		return nil, errors.New("o.Src cannot be nil")
	}
	if !o.Dst.Initialized() {
		return nil, errors.New("o.Dst cannot be uninitialized")
	}

	info := (*C.VkCopyMemoryToAccelerationStructureInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_COPY_MEMORY_TO_ACCELERATION_STRUCTURE_INFO_KHR
	info.pNext = next
	o.Src.PopulateAddressUnion(unsafe.Pointer(&info.src))
	info.dst = C.VkAccelerationStructureKHR(unsafe.Pointer(o.Dst.Handle()))
	info.mode = C.VkCopyAccelerationStructureModeKHR(o.Mode)

	return preallocatedPointer, nil
}
