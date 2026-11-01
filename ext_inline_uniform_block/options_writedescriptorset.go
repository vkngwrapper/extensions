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

type WriteDescriptorSetInlineUniformBlock struct {
	Data []byte

	common.NextOptions
}

func (o WriteDescriptorSetInlineUniformBlock) WriteDescriptorSetCount() int {
	return len(o.Data)
}

func (o WriteDescriptorSetInlineUniformBlock) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkWriteDescriptorSetInlineUniformBlockEXT{})))
	}

	features := (*C.VkWriteDescriptorSetInlineUniformBlockEXT)(preallocatedPointer)
	features.sType = C.VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_INLINE_UNIFORM_BLOCK
	features.pNext = next

	dataLen := len(o.Data)
	features.dataSize = C.uint32_t(dataLen)
	features.pData = nil

	if dataLen > 0 {
		dataUnsafe := allocator.Malloc(dataLen)
		dataPtr := (*byte)(dataUnsafe)
		dataSlice := unsafe.Slice(dataPtr, dataLen)
		copy(dataSlice, o.Data)
		features.pData = dataUnsafe
	}

	return preallocatedPointer, nil
}
