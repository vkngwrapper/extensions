package khr_buffer_device_address

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

// MemoryOpaqueCaptureAddressAllocateInfo requests a specific address for a memory allocation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryOpaqueCaptureAddressAllocateInfoKHR.html
type MemoryOpaqueCaptureAddressAllocateInfo struct {
	// OpaqueCaptureAddress is the opaque capture address requested for the memory allocation
	OpaqueCaptureAddress uint64

	common.NextOptions
}

func (o MemoryOpaqueCaptureAddressAllocateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkMemoryOpaqueCaptureAddressAllocateInfoKHR{})))
	}

	info := (*C.VkMemoryOpaqueCaptureAddressAllocateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_MEMORY_OPAQUE_CAPTURE_ADDRESS_ALLOCATE_INFO_KHR
	info.pNext = next
	info.opaqueCaptureAddress = C.uint64_t(o.OpaqueCaptureAddress)

	return preallocatedPointer, nil
}
