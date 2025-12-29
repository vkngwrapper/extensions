package khr_dedicated_allocation

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
)

// MemoryDedicatedAllocateInfo specifies a dedicated memory allocation resource
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html
type MemoryDedicatedAllocateInfo struct {
	// Image is nil or the Image object which this memory will be bound to
	Image core1_0.Image
	// Buffer is nil or the Buffer object this memory will be bound to
	Buffer core1_0.Buffer

	common.NextOptions
}

func (o MemoryDedicatedAllocateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Image.Initialized() && o.Buffer.Initialized() {
		return nil, errors.New("both Image and Buffer fields are set in MemoryDedicatedAllocateInfo- only one must be set")
	} else if !o.Image.Initialized() && !o.Buffer.Initialized() {
		return nil, errors.New("neither Image nor Buffer fields are set in MemoryDedicatedAllocateInfo- one must be set")
	}

	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkMemoryDedicatedAllocateInfoKHR{})))
	}

	createInfo := (*C.VkMemoryDedicatedAllocateInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO_KHR
	createInfo.pNext = next
	createInfo.image = nil
	createInfo.buffer = nil

	if o.Image.Initialized() {
		createInfo.image = C.VkImage(unsafe.Pointer(o.Image.Handle()))
	} else if o.Buffer.Initialized() {
		createInfo.buffer = C.VkBuffer(unsafe.Pointer(o.Buffer.Handle()))
	}

	return preallocatedPointer, nil
}
