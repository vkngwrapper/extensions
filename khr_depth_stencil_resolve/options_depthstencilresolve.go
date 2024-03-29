package khr_depth_stencil_resolve

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/extensions/v2/khr_create_renderpass2"
	"unsafe"
)

// SubpassDescriptionDepthStencilResolve specifies depth/stencil resolve operations for
// a subpass
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescriptionDepthStencilResolve.html
type SubpassDescriptionDepthStencilResolve struct {
	// DepthResolveMode describes the depth resolve mode
	DepthResolveMode ResolveModeFlags
	// StencilResolveMode describes the stencil resolve mode
	StencilResolveMode ResolveModeFlags
	// DepthStencilResolveAttachment defines the depth/stencil resolve attachment
	// for this subpass and its layout
	DepthStencilResolveAttachment *khr_create_renderpass2.AttachmentReference2

	common.NextOptions
}

func (o SubpassDescriptionDepthStencilResolve) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassDescriptionDepthStencilResolveKHR{})))
	}

	info := (*C.VkSubpassDescriptionDepthStencilResolveKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_DEPTH_STENCIL_RESOLVE_KHR
	info.pNext = next
	info.depthResolveMode = C.VkResolveModeFlagBits(o.DepthResolveMode)
	info.stencilResolveMode = C.VkResolveModeFlagBits(o.StencilResolveMode)
	info.pDepthStencilResolveAttachment = nil

	if o.DepthStencilResolveAttachment != nil {
		attachment, err := common.AllocOptions(allocator, o.DepthStencilResolveAttachment)
		if err != nil {
			return nil, err
		}

		info.pDepthStencilResolveAttachment = (*C.VkAttachmentReference2KHR)(attachment)
	}

	return preallocatedPointer, nil
}
