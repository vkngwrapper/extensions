package ext_memory_priority

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

// MemoryPriorityAllocateInfo specifies memory priority for a new allocation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryPriorityAllocateInfoEXT.html
type MemoryPriorityAllocateInfo struct {
	Priority float32

	common.NextOptions
}

func (o MemoryPriorityAllocateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(C.sizeof_struct_VkMemoryPriorityAllocateInfoEXT))
	}

	info := (*C.VkMemoryPriorityAllocateInfoEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_MEMORY_PRIORITY_ALLOCATE_INFO_EXT
	info.pNext = next
	info.priority = C.float(o.Priority)

	return preallocatedPointer, nil
}
