package khr_descriptor_update_template

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/core1_0"
	"unsafe"
)

// DescriptorUpdateTemplateEntry describes a single descriptor update of the DescriptorUpdateTemplate
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateEntry.html
type DescriptorUpdateTemplateEntry struct {
	// DstBinding is the descriptor binding to update when using this DescriptorUpdateTemplate
	DstBinding int
	// DstArrayElement is the starting element in the array belonging to DstBinding
	DstArrayElement int
	// DescriptorCount is the number of descriptors to update
	DescriptorCount int

	// DescriptorType specifies the type of the descriptor
	DescriptorType core1_0.DescriptorType

	// Offset is the offset in bytes of the first binding in the raw data structure
	Offset int
	// Stride is the stride in bytes between two consecutive array elements of the
	// descriptor update informations in the raw data structure
	Stride int
}

func (e DescriptorUpdateTemplateEntry) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDescriptorUpdateTemplateEntryKHR{})))
	}

	entry := (*C.VkDescriptorUpdateTemplateEntryKHR)(preallocatedPointer)
	entry.dstBinding = C.uint32_t(e.DstBinding)
	entry.dstArrayElement = C.uint32_t(e.DstArrayElement)
	entry.descriptorCount = C.uint32_t(e.DescriptorCount)
	entry.descriptorType = C.VkDescriptorType(e.DescriptorType)
	entry.offset = C.size_t(e.Offset)
	entry.stride = C.size_t(e.Stride)

	return preallocatedPointer, nil
}
