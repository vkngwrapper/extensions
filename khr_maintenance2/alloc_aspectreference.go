package khr_maintenance2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/core1_0"
	"unsafe"
)

// InputAttachmentAspectReference specifies a subpass/input attachment pair and an aspect mask that
// can be read
type InputAttachmentAspectReference struct {
	// Subpass is an index into RenderPassCreateInfo.Subpasses
	Subpass int
	// InputAttachmentIndex is an index into the InputAttachments of the specified subpass
	InputAttachmentIndex int
	// AspectMask is a mask of which aspect(s) can be accessed within the specified subpass
	AspectMask core1_0.ImageAspectFlags
}

func (ref InputAttachmentAspectReference) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkInputAttachmentAspectReferenceKHR{})))
	}

	val := (*C.VkInputAttachmentAspectReferenceKHR)(preallocatedPointer)
	val.subpass = C.uint32_t(ref.Subpass)
	val.inputAttachmentIndex = C.uint32_t(ref.InputAttachmentIndex)
	val.aspectMask = C.VkImageAspectFlags(ref.AspectMask)

	return preallocatedPointer, nil
}
