package khr_create_renderpass2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"unsafe"
)

// RenderPassCreateInfo2 specifies parameters of a newly-created RenderPass
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo2.html
type RenderPassCreateInfo2 struct {
	// Flags is reserved for future use
	Flags core1_0.RenderPassCreateFlags

	// Attachments is a slice of AttachmentDescription2 structures describing the attachments
	// used by the RenderPass
	Attachments []AttachmentDescription2
	// Subpasses is a slice of SubpassDescription2 structures describing each subpass
	Subpasses []SubpassDescription2
	// Dependencies is a slice of SubpassDependency2 structures describing dependencies
	// between pairs of subpasses
	Dependencies []SubpassDependency2

	// CorrelatedViewMasks is a slice of view masks indicating sets of views that may be
	// more efficient to render concurrently
	CorrelatedViewMasks []uint32

	common.NextOptions
}

func (o RenderPassCreateInfo2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkRenderPassCreateInfo2KHR{})))
	}

	info := (*C.VkRenderPassCreateInfo2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO_2_KHR
	info.pNext = next
	info.flags = C.VkRenderPassCreateFlags(o.Flags)

	attachmentCount := len(o.Attachments)
	subpassCount := len(o.Subpasses)
	dependencyCount := len(o.Dependencies)
	viewMaskCount := len(o.CorrelatedViewMasks)

	info.attachmentCount = C.uint32_t(attachmentCount)
	info.pAttachments = nil
	info.subpassCount = C.uint32_t(subpassCount)
	info.pSubpasses = nil
	info.dependencyCount = C.uint32_t(dependencyCount)
	info.pDependencies = nil
	info.correlatedViewMaskCount = C.uint32_t(viewMaskCount)
	info.pCorrelatedViewMasks = nil

	var err error
	if attachmentCount > 0 {
		info.pAttachments, err = common.AllocOptionSlice[C.VkAttachmentDescription2KHR, AttachmentDescription2](allocator, o.Attachments)
		if err != nil {
			return nil, err
		}
	}

	if subpassCount > 0 {
		info.pSubpasses, err = common.AllocOptionSlice[C.VkSubpassDescription2KHR, SubpassDescription2](allocator, o.Subpasses)
		if err != nil {
			return nil, err
		}
	}

	if dependencyCount > 0 {
		info.pDependencies, err = common.AllocOptionSlice[C.VkSubpassDependency2KHR, SubpassDependency2](allocator, o.Dependencies)
		if err != nil {
			return nil, err
		}
	}

	if viewMaskCount > 0 {
		viewMaskPtr := (*C.uint32_t)(allocator.Malloc(viewMaskCount * int(unsafe.Sizeof(C.uint32_t(0)))))
		viewMaskSlice := unsafe.Slice(viewMaskPtr, viewMaskCount)
		for i := 0; i < viewMaskCount; i++ {
			viewMaskSlice[i] = C.uint32_t(o.CorrelatedViewMasks[i])
		}
		info.pCorrelatedViewMasks = viewMaskPtr
	}

	return preallocatedPointer, nil
}
