package ext_descriptor_indexing

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

// DescriptorSetVariableDescriptorCountAllocateInfo specifies additional allocation parameters
// for DescriptorSet objects
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetVariableDescriptorCountAllocateInfo.html
type DescriptorSetVariableDescriptorCountAllocateInfo struct {
	// DescriptorCounts is a slice of descriptor counts, with each member specifying the number
	// of descriptors in a variable-sized descriptor binding in the corresponding DescriptorSet
	// being allocated
	DescriptorCounts []int

	common.NextOptions
}

func (o DescriptorSetVariableDescriptorCountAllocateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDescriptorSetVariableDescriptorCountAllocateInfoEXT{})))
	}

	info := (*C.VkDescriptorSetVariableDescriptorCountAllocateInfoEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO_EXT
	info.pNext = next

	count := len(o.DescriptorCounts)
	info.descriptorSetCount = C.uint32_t(count)
	info.pDescriptorCounts = nil

	if count > 0 {
		info.pDescriptorCounts = (*C.uint32_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint32_t(0)))))
		descriptorCountSlice := unsafe.Slice(info.pDescriptorCounts, count)
		for i := 0; i < count; i++ {
			descriptorCountSlice[i] = C.uint32_t(o.DescriptorCounts[i])
		}
	}

	return preallocatedPointer, nil
}
