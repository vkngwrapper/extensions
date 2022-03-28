package khr_get_memory_requirements2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type SparseImageRequirementsOptions struct {
	Image core1_0.Image

	common.HaveNext
}

func (o SparseImageRequirementsOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageSparseMemoryRequirementsInfo2KHR{})))
	}

	options := (*C.VkImageSparseMemoryRequirementsInfo2KHR)(preallocatedPointer)
	options.sType = C.VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2_KHR
	options.pNext = next
	options.image = C.VkImage(unsafe.Pointer(o.Image.Handle()))

	return preallocatedPointer, nil
}

func (o SparseImageRequirementsOptions) PopulateOutData(cPointer unsafe.Pointer) (next unsafe.Pointer, err error) {
	options := (*C.VkImageSparseMemoryRequirementsInfo2KHR)(cPointer)
	return options.pNext, nil
}