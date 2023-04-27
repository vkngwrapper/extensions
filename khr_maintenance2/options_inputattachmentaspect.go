package khr_maintenance2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/pkg/errors"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

// RenderPassInputAttachmentAspectCreateInfo specifies, for a given subpass/input attachment
// pair, which aspect can be read
type RenderPassInputAttachmentAspectCreateInfo struct {
	// AspectReferences is a slice of InputAttachmentAspectReference structures containing
	// a mask describing which aspect(s) can be accessed for a given input attachment within a
	// given subpass
	AspectReferences []InputAttachmentAspectReference

	common.NextOptions
}

func (o RenderPassInputAttachmentAspectCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkRenderPassInputAttachmentAspectCreateInfoKHR{})))
	}

	createInfo := (*C.VkRenderPassInputAttachmentAspectCreateInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO_KHR
	createInfo.pNext = next

	count := len(o.AspectReferences)
	if count < 1 {
		return nil, errors.New("options RenderPassInputAttachmentAspectCreateInfo must include at least 1 entry in AspectReferences")
	}

	createInfo.aspectReferenceCount = C.uint32_t(count)
	references, err := common.AllocSlice[C.VkInputAttachmentAspectReference, InputAttachmentAspectReference](allocator, o.AspectReferences)
	if err != nil {
		return nil, err
	}
	createInfo.pAspectReferences = references

	return preallocatedPointer, nil
}
