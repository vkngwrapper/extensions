package khr_get_memory_requirements2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
)

// BufferMemoryRequirementsInfo2 has no documentation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferMemoryRequirementsInfo2.html
type BufferMemoryRequirementsInfo2 struct {
	// Buffer is the Buffer to query
	Buffer core.Buffer

	common.NextOptions
}

func (o BufferMemoryRequirementsInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Buffer.Handle() == 0 {
		return nil, errors.New("khr_get_memory_requirements2.BufferMemoryRequirementsInfo2.Buffer cannot be uninitialized")
	}
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkBufferMemoryRequirementsInfo2KHR{})))
	}

	options := (*C.VkBufferMemoryRequirementsInfo2KHR)(preallocatedPointer)
	options.sType = C.VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2_KHR
	options.pNext = next
	options.buffer = C.VkBuffer(unsafe.Pointer(o.Buffer.Handle()))

	return preallocatedPointer, nil
}
