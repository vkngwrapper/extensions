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

type AccelerationStructureDeviceAddressInfo struct {
	AccelerationStructure AccelerationStructure

	common.NextOptions
}

func (o AccelerationStructureDeviceAddressInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAccelerationStructureDeviceAddressInfoKHR{})))
	}
	if !o.AccelerationStructure.Initialized() {
		return nil, errors.New("o.AccelerationStructure cannot be uninitialized")
	}

	info := (*C.VkAccelerationStructureDeviceAddressInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_DEVICE_ADDRESS_INFO_KHR
	info.pNext = next
	info.accelerationStructure = C.VkAccelerationStructureKHR(unsafe.Pointer(o.AccelerationStructure.handle))

	return preallocatedPointer, nil
}
