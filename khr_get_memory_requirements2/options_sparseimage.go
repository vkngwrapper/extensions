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

// ImageSparseMemoryRequirementsInfo2 has no documentation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSparseMemoryRequirementsInfo2.html
type ImageSparseMemoryRequirementsInfo2 struct {
	// Image is the Image to query
	Image core.Image

	common.NextOptions
}

func (o ImageSparseMemoryRequirementsInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Image.Handle() == 0 {
		return nil, errors.New("khr_get_memory_requirements2.ImageSparseMemoryRequirementsInfo2.Image cannot be uninitialized")
	}
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkImageSparseMemoryRequirementsInfo2KHR{})))
	}

	options := (*C.VkImageSparseMemoryRequirementsInfo2KHR)(preallocatedPointer)
	options.sType = C.VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2_KHR
	options.pNext = next
	options.image = C.VkImage(unsafe.Pointer(o.Image.Handle()))

	return preallocatedPointer, nil
}
