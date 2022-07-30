package khr_external_memory

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/extensions/khr_external_memory_capabilities"
	"unsafe"
)

// ExternalMemoryImageCreateInfo specifies that an Image may be backed by external memory
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html
type ExternalMemoryImageCreateInfo struct {
	// HandleTypes specifies one or more external memory handle types
	HandleTypes khr_external_memory_capabilities.ExternalMemoryHandleTypeFlags

	common.NextOptions
}

func (o ExternalMemoryImageCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalMemoryImageCreateInfoKHR{})))
	}

	info := (*C.VkExternalMemoryImageCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_KHR
	info.pNext = next
	info.handleTypes = C.VkExternalMemoryHandleTypeFlagsKHR(o.HandleTypes)

	return preallocatedPointer, nil
}

func (o ExternalMemoryImageCreateInfo) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalMemoryImageCreateInfoKHR)(cDataPointer)
	return info.pNext, nil
}
