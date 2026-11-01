package ext_inline_uniform_block

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

type PhysicalDeviceInlineUniformBlockFeatures struct {
	InlineUniformBlock                                 bool
	DescriptorBindingInlineUniformBlockUpdateAfterBind bool

	common.NextOptions
	common.NextOutData
}

func (o PhysicalDeviceInlineUniformBlockFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceInlineUniformBlockFeaturesEXT{})))
	}

	features := (*C.VkPhysicalDeviceInlineUniformBlockFeaturesEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_FEATURES
	features.pNext = next
	features.inlineUniformBlock = C.VkBool32(0)
	if o.InlineUniformBlock {
		features.inlineUniformBlock = C.VkBool32(1)
	}
	features.descriptorBindingInlineUniformBlockUpdateAfterBind = C.VkBool32(0)
	if o.DescriptorBindingInlineUniformBlockUpdateAfterBind {
		features.descriptorBindingInlineUniformBlockUpdateAfterBind = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceInlineUniformBlockFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceInlineUniformBlockFeaturesEXT{})))
	}

	features := (*C.VkPhysicalDeviceInlineUniformBlockFeaturesEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_FEATURES
	features.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceInlineUniformBlockFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	features := (*C.VkPhysicalDeviceInlineUniformBlockFeaturesEXT)(cDataPointer)
	o.InlineUniformBlock = features.inlineUniformBlock != C.VkBool32(0)
	o.DescriptorBindingInlineUniformBlockUpdateAfterBind = features.descriptorBindingInlineUniformBlockUpdateAfterBind != C.VkBool32(0)

	return features.pNext, nil
}
