package ext_descriptor_indexing

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

// DescriptorSetVariableDescriptorCountLayoutSupport returns information about whether a
// DescriptorSetLayout can be supported
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetVariableDescriptorCountLayoutSupport.html
type DescriptorSetVariableDescriptorCountLayoutSupport struct {
	// MaxVariableDescriptorCount indicates the maximum number of descriptors supported in the
	// highest numbered binding of the layout, if that binding is variable-sized
	MaxVariableDescriptorCount int

	common.NextOutData
}

func (o *DescriptorSetVariableDescriptorCountLayoutSupport) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDescriptorSetVariableDescriptorCountLayoutSupportEXT{})))
	}

	info := (*C.VkDescriptorSetVariableDescriptorCountLayoutSupportEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT_EXT
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *DescriptorSetVariableDescriptorCountLayoutSupport) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkDescriptorSetVariableDescriptorCountLayoutSupportEXT)(cDataPointer)

	o.MaxVariableDescriptorCount = int(info.maxVariableDescriptorCount)

	return info.pNext, nil
}
