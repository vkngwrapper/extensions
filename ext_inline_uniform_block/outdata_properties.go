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

type PhysicalDeviceInlineUniformBlockProperties struct {
	MaxInlineUniformBlockSize                               int
	MaxPerStageDescriptorInlineUniformBlocks                int
	MaxPerStageDescriptorUpdateAfterBindInlineUniformBlocks int
	MaxDescriptorSetInlineUniformBlocks                     int
	MaxDescriptorSetUpdateAfterBindInlineUniformBlocks      int

	common.NextOutData
}

func (o *PhysicalDeviceInlineUniformBlockProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceInlineUniformBlockPropertiesEXT{})))
	}

	properties := (*C.VkPhysicalDeviceInlineUniformBlockPropertiesEXT)(preallocatedPointer)
	properties.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_FEATURES
	properties.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceInlineUniformBlockProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	properties := (*C.VkPhysicalDeviceInlineUniformBlockPropertiesEXT)(cDataPointer)
	o.MaxInlineUniformBlockSize = int(properties.maxInlineUniformBlockSize)
	o.MaxPerStageDescriptorInlineUniformBlocks = int(properties.maxPerStageDescriptorInlineUniformBlocks)
	o.MaxPerStageDescriptorUpdateAfterBindInlineUniformBlocks = int(properties.maxPerStageDescriptorUpdateAfterBindInlineUniformBlocks)
	o.MaxDescriptorSetInlineUniformBlocks = int(properties.maxDescriptorSetInlineUniformBlocks)
	o.MaxDescriptorSetUpdateAfterBindInlineUniformBlocks = int(properties.maxDescriptorSetUpdateAfterBindInlineUniformBlocks)

	return properties.pNext, nil
}
