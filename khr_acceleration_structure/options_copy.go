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

type CopyAccelerationStructureInfo struct {
	Src  AccelerationStructure
	Dst  AccelerationStructure
	Mode CopyAccelerationStructureMode

	common.NextOptions
}

func (o CopyAccelerationStructureInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkCopyAccelerationStructureInfoKHR{})))
	}
	if !o.Src.Initialized() {
		return nil, errors.New("o.Src cannot be uninitialized")
	}
	if !o.Dst.Initialized() {
		return nil, errors.New("o.Dst cannot be uninitialized")
	}

	info := (*C.VkCopyAccelerationStructureInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_COPY_ACCELERATION_STRUCTURE_INFO_KHR
	info.pNext = next
	info.src = (C.VkAccelerationStructureKHR)(unsafe.Pointer(o.Src.Handle()))
	info.dst = (C.VkAccelerationStructureKHR)(unsafe.Pointer(o.Dst.Handle()))
	info.mode = C.VkCopyAccelerationStructureModeKHR(o.Mode)

	return preallocatedPointer, nil
}
