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

type CopyAccelerationStructureToMemoryInfo struct {
	Src  AccelerationStructure
	Dst  DeviceOrHostAddressConst
	Mode CopyAccelerationStructureMode

	common.NextOptions
}

func (o CopyAccelerationStructureToMemoryInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkCopyAccelerationStructureToMemoryInfoKHR{})))
	}
	if !o.Src.Initialized() {
		return nil, errors.New("o.Src cannot be uninitialized")
	}
	if o.Dst == nil {
		return nil, errors.New("o.Dst cannot be nil")
	}

	info := (*C.VkCopyAccelerationStructureToMemoryInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_COPY_ACCELERATION_STRUCTURE_TO_MEMORY_INFO_KHR
	info.pNext = next
	info.src = C.VkAccelerationStructureKHR(unsafe.Pointer(o.Src.Handle()))
	o.Dst.PopulateAddressUnion(unsafe.Pointer(&info.dst))
	info.mode = C.VkCopyAccelerationStructureModeKHR(o.Mode)

	return preallocatedPointer, nil
}
