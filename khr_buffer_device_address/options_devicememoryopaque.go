package khr_buffer_device_address

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"unsafe"
)

// DeviceMemoryOpaqueCaptureAddressInfo specifies the DeviceMemory object to query an address for
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryOpaqueCaptureAddressInfo.html
type DeviceMemoryOpaqueCaptureAddressInfo struct {
	// Memory specifies the DeviceMemory whose address is being queried
	Memory core1_0.DeviceMemory

	common.NextOptions
}

func (o DeviceMemoryOpaqueCaptureAddressInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Memory == nil {
		return nil, errors.New("khr_buffer_device_address.DeviceMemoryOpaqueCaptureAddressInfo.Memory cannot be nil")
	}
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDeviceMemoryOpaqueCaptureAddressInfoKHR{})))
	}

	info := (*C.VkDeviceMemoryOpaqueCaptureAddressInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DEVICE_MEMORY_OPAQUE_CAPTURE_ADDRESS_INFO
	info.pNext = next
	info.memory = C.VkDeviceMemory(unsafe.Pointer(o.Memory.Handle()))

	return preallocatedPointer, nil
}
