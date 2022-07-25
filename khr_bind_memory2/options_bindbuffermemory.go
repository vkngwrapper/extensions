package khr_bind_memory2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"unsafe"
)

// BindBufferMemoryInfo specifies how to bind a Buffer to DeviceMemory
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryInfo.html
type BindBufferMemoryInfo struct {
	// Buffer is the Buffer to be attached to memory
	Buffer core1_0.Buffer
	// Memory describes the DeviceMemory object to attach
	Memory core1_0.DeviceMemory
	// MemoryOffset is the start offset of the region of memory which is to be bound to the Buffer
	MemoryOffset int

	common.NextOptions
}

func (o BindBufferMemoryInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBindBufferMemoryInfoKHR{})))
	}

	createInfo := (*C.VkBindBufferMemoryInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO_KHR
	createInfo.pNext = next
	createInfo.buffer = (C.VkBuffer)(unsafe.Pointer(o.Buffer.Handle()))
	createInfo.memory = (C.VkDeviceMemory)(unsafe.Pointer(o.Memory.Handle()))
	createInfo.memoryOffset = C.VkDeviceSize(o.MemoryOffset)

	return preallocatedPointer, nil
}
