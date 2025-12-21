package khr_external_memory_capabilities

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
)

// ExternalBufferProperties specifies supported external handle capabilities
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalBufferProperties.html
type ExternalBufferProperties struct {
	// ExternalMemoryProperties specifies various capabilities of the external handle type when
	// used with the specified Buffer creation parameters
	ExternalMemoryProperties ExternalMemoryProperties

	common.NextOutData
}

func (o *ExternalBufferProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalBufferPropertiesKHR{})))
	}
	info := (*C.VkExternalBufferPropertiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_BUFFER_PROPERTIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *ExternalBufferProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalBufferPropertiesKHR)(cDataPointer)

	err = (&o.ExternalMemoryProperties).PopulateOutData(unsafe.Pointer(&info.externalMemoryProperties))
	return info.pNext, nil
}
