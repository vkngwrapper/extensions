package khr_buffer_device_address

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"unsafe"
)

// BufferDeviceAddressInfo specifies the Buffer to query an address for
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferDeviceAddressInfo.html
type BufferDeviceAddressInfo struct {
	// Buffer specifies the Buffer whose address is being queried
	Buffer core1_0.Buffer

	common.NextOptions
}

func (o BufferDeviceAddressInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Buffer == nil {
		return nil, errors.New("khr_buffer_device_address.DeviceMemoryAddressOptions.Buffer cannot be nil")
	}

	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBufferDeviceAddressInfoKHR{})))
	}

	info := (*C.VkBufferDeviceAddressInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO_KHR
	info.pNext = next
	info.buffer = C.VkBuffer(unsafe.Pointer(o.Buffer.Handle()))

	return preallocatedPointer, nil
}
