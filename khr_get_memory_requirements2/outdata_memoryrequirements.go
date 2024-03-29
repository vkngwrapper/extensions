package khr_get_memory_requirements2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"unsafe"
)

// MemoryRequirements2 specifies memory requirements
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html
type MemoryRequirements2 struct {
	// MemoryRequirements describes the memory requirements of the resource
	MemoryRequirements core1_0.MemoryRequirements

	common.NextOutData
}

func (o *MemoryRequirements2) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkMemoryRequirements2KHR{})))
	}

	outData := (*C.VkMemoryRequirements2KHR)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2_KHR
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *MemoryRequirements2) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkMemoryRequirements2KHR)(cDataPointer)
	o.MemoryRequirements.Size = int(outData.memoryRequirements.size)
	o.MemoryRequirements.Alignment = int(outData.memoryRequirements.alignment)
	o.MemoryRequirements.MemoryTypeBits = uint32(outData.memoryRequirements.memoryTypeBits)

	return outData.pNext, nil
}
