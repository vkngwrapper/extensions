package khr_external_memory

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/extensions/v3/khr_external_memory_capabilities"
)

// ExternalMemoryBufferCreateInfo specifies that a Buffer may be backed by external memory
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryBufferCreateInfo.html
type ExternalMemoryBufferCreateInfo struct {
	// HandleTypes specifies one or more external memory handle types
	HandleTypes khr_external_memory_capabilities.ExternalMemoryHandleTypeFlags

	common.NextOptions
}

func (o ExternalMemoryBufferCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalMemoryBufferCreateInfoKHR{})))
	}

	info := (*C.VkExternalMemoryBufferCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_BUFFER_CREATE_INFO_KHR
	info.pNext = next
	info.handleTypes = C.VkExternalMemoryHandleTypeFlagsKHR(o.HandleTypes)

	return preallocatedPointer, nil
}

func (o ExternalMemoryBufferCreateInfo) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalMemoryBufferCreateInfoKHR)(cDataPointer)
	return info.pNext, nil
}
