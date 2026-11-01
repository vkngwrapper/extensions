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

type DescriptorPoolInlineUniformBlockCreateInfo struct {
	MaxInlineUniformBlockBindings int

	common.NextOptions
}

func (o DescriptorPoolInlineUniformBlockCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkDescriptorPoolInlineUniformBlockCreateInfoEXT{})))
	}

	features := (*C.VkDescriptorPoolInlineUniformBlockCreateInfoEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_INLINE_UNIFORM_BLOCK_CREATE_INFO
	features.pNext = next
	features.maxInlineUniformBlockBindings = C.uint32_t(o.MaxInlineUniformBlockBindings)

	return preallocatedPointer, nil
}
