package khr_external_memory_capabilities

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

// ExternalImageFormatProperties specifies supported external handle properties
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatProperties.html
type ExternalImageFormatProperties struct {
	// ExternalMemoryProperties specifies various capabilities of the external handle type when used
	// with the specified Image creation parameters
	ExternalMemoryProperties ExternalMemoryProperties

	common.NextOutData
}

func (o *ExternalImageFormatProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkExternalImageFormatPropertiesKHR{})))
	}

	info := (*C.VkExternalImageFormatPropertiesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_EXTERNAL_IMAGE_FORMAT_PROPERTIES_KHR
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *ExternalImageFormatProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkExternalImageFormatPropertiesKHR)(cDataPointer)

	err = (&o.ExternalMemoryProperties).PopulateOutData(unsafe.Pointer(&info.externalMemoryProperties))
	return info.pNext, err
}
