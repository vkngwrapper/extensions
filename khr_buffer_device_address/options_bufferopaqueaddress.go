package khr_buffer_device_address

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

// BufferOpaqueCaptureAddressCreateInfo requests a specific address for a Buffer
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferOpaqueCaptureAddressCreateInfo.html
type BufferOpaqueCaptureAddressCreateInfo struct {
	// OpaqueCaptureAddress is the opaque capture address requested for the Buffer
	OpaqueCaptureAddress uint64

	common.NextOptions
}

func (o BufferOpaqueCaptureAddressCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBufferOpaqueCaptureAddressCreateInfoKHR{})))
	}

	info := (*C.VkBufferOpaqueCaptureAddressCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_BUFFER_OPAQUE_CAPTURE_ADDRESS_CREATE_INFO
	info.pNext = next
	info.opaqueCaptureAddress = C.uint64_t(o.OpaqueCaptureAddress)

	return preallocatedPointer, nil
}
