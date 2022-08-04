package khr_maintenance3

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

// DescriptorSetLayoutSupport returns information about whether a DescriptorSetLayout can be
// supported
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutSupport.html
type DescriptorSetLayoutSupport struct {
	// Supported specifies whether the DescriptorSetLayout can be created
	Supported bool

	common.NextOutData
}

func (o *DescriptorSetLayoutSupport) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDescriptorSetLayoutSupportKHR{})))
	}

	outData := (*C.VkDescriptorSetLayoutSupportKHR)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT_KHR
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *DescriptorSetLayoutSupport) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkDescriptorSetLayoutSupportKHR)(cDataPointer)
	o.Supported = outData.supported != C.VkBool32(0)

	return outData.pNext, nil
}
