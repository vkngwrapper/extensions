package khr_bind_memory2

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

// BindImageMemoryInfo specifies how to bind an Image to DeviceMemory
type BindImageMemoryInfo struct {
	// Image is the image to be attached to DeviceMemory
	Image core1_0.Image
	// Memory describes the DeviceMemory to attach
	Memory core1_0.DeviceMemory
	// MemoryOffset is the start offset of the region of DeviceMemory to be bound to the Image
	MemoryOffset uint64

	common.NextOptions
}

func (o BindImageMemoryInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Image == nil {
		return nil, errors.New("khr_bind_memory2.BindImageMemoryInfo.Image cannot be nil")
	}
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBindImageMemoryInfoKHR{})))
	}

	createInfo := (*C.VkBindImageMemoryInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO_KHR
	createInfo.pNext = next
	createInfo.image = (C.VkImage)(unsafe.Pointer(o.Image.Handle()))
	createInfo.memory = nil
	if o.Memory != nil {
		createInfo.memory = (C.VkDeviceMemory)(unsafe.Pointer(o.Memory.Handle()))
	}
	createInfo.memoryOffset = C.VkDeviceSize(o.MemoryOffset)

	return preallocatedPointer, nil
}
