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

// DescriptorSetLayoutBindingFlagsCreateInfo specifies parameters of a newly-created
// DescriptorSetLayout
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutBindingFlagsCreateInfo.html
type DescriptorSetLayoutBindingFlagsCreateInfo struct {
	// BindingFlags is a slice of DescriptorBindingFlags, one for each DescriptorSetLayout binding
	BindingFlags []DescriptorBindingFlags

	common.NextOptions
}

func (o DescriptorSetLayoutBindingFlagsCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDescriptorSetLayoutBindingFlagsCreateInfoEXT{})))
	}

	info := (*C.VkDescriptorSetLayoutBindingFlagsCreateInfoEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO_EXT
	info.pNext = next

	count := len(o.BindingFlags)
	info.bindingCount = C.uint32_t(count)
	info.pBindingFlags = nil

	if count > 0 {
		info.pBindingFlags = (*C.VkDescriptorBindingFlags)(allocator.Malloc(count * int(unsafe.Sizeof(C.VkDescriptorBindingFlags(0)))))
		flagSlice := unsafe.Slice(info.pBindingFlags, count)

		for i := 0; i < count; i++ {
			flagSlice[i] = C.VkDescriptorBindingFlags(o.BindingFlags[i])
		}
	}

	return preallocatedPointer, nil
}
