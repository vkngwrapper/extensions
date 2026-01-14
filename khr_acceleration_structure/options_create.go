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
	"github.com/vkngwrapper/core/v3/core1_0"
)

type AccelerationStructureCreateInfo struct {
	CreateFlags   AccelerationStructureCreateFlags
	Buffer        core1_0.Buffer
	Offset        int
	Size          int
	Type          AccelerationStructureType
	DeviceAddress uint64

	common.NextOptions
}

func (o AccelerationStructureCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAccelerationStructureCreateInfoKHR{})))
	}

	if !o.Buffer.Initialized() {
		return nil, errors.New("o.Buffer cannot be uninitialized")
	}

	info := (*C.VkAccelerationStructureCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_CREATE_INFO_KHR
	info.pNext = next
	info.createFlags = C.VkAccelerationStructureCreateFlagsKHR(o.CreateFlags)
	info.buffer = C.VkBuffer(unsafe.Pointer(o.Buffer.Handle()))
	info.offset = C.VkDeviceSize(o.Offset)
	info.size = C.VkDeviceSize(o.Size)
	info._type = C.VkAccelerationStructureTypeKHR(o.Type)
	info.deviceAddress = C.VkDeviceAddress(o.DeviceAddress)

	return preallocatedPointer, nil
}
