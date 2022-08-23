package khr_create_renderpass2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"unsafe"
)

// SubpassDescription2 specifies a subpass description
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html
type SubpassDescription2 struct {
	// Flags specifies usage of the subpass
	Flags core1_0.SubpassDescriptionFlags
	// PipelineBindPoint specifies the Pipeline type supported for this subpass
	PipelineBindPoint core1_0.PipelineBindPoint
	// ViewMask describes which views rendering is broadcast to in this subpass, when
	// multiview is enabled
	ViewMask uint32
	// InputAttachments is a slice of AttachmentReference2 structures defining the input
	// attachments for this subpass and their layouts
	InputAttachments []AttachmentReference2
	// ColorAttachments is a slice of AttachmentReference2 structures defining the color
	// attachments for this subpass and their layouts
	ColorAttachments []AttachmentReference2
	// ResolveAttachments is a slice of AttachmentReference2 structures defining the resolve
	// attachments for this subpass and their layouts
	ResolveAttachments []AttachmentReference2
	// DepthStencilAttachment specifies the depth/stencil attachment for this subpass and
	// its layout
	DepthStencilAttachment *AttachmentReference2
	// PreserveAttachments is a slice of RenderPass attachment indices identifying attachments
	// that are not used by this subpass, but whose contents must be preserved throughout the
	// subpass
	PreserveAttachments []int

	common.NextOptions
}

func (o SubpassDescription2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSubpassDescription2KHR{})))
	}

	info := (*C.VkSubpassDescription2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2_KHR
	info.pNext = next
	info.flags = C.VkSubpassDescriptionFlags(o.Flags)
	info.pipelineBindPoint = C.VkPipelineBindPoint(o.PipelineBindPoint)
	info.viewMask = C.uint32_t(o.ViewMask)

	inputAttachmentCount := len(o.InputAttachments)
	colorAttachmentCount := len(o.ColorAttachments)
	resolveAttachmentCount := len(o.ResolveAttachments)
	preserveAttachmentCount := len(o.PreserveAttachments)

	if resolveAttachmentCount > 0 && resolveAttachmentCount != colorAttachmentCount {
		return nil, errors.Newf("in this subpass, %d color attachments are defined, but %d resolve attachments are defined- they should be equal", colorAttachmentCount, resolveAttachmentCount)
	}

	info.inputAttachmentCount = C.uint32_t(inputAttachmentCount)
	info.pInputAttachments = nil
	info.colorAttachmentCount = C.uint32_t(colorAttachmentCount)
	info.pColorAttachments = nil
	info.pResolveAttachments = nil
	info.pDepthStencilAttachment = nil
	info.preserveAttachmentCount = C.uint32_t(preserveAttachmentCount)
	info.pPreserveAttachments = nil

	var err error
	if inputAttachmentCount > 0 {
		info.pInputAttachments, err = common.AllocOptionSlice[C.VkAttachmentReference2KHR, AttachmentReference2](allocator, o.InputAttachments)
		if err != nil {
			return nil, err
		}
	}

	if colorAttachmentCount > 0 {
		info.pColorAttachments, err = common.AllocOptionSlice[C.VkAttachmentReference2KHR, AttachmentReference2](allocator, o.ColorAttachments)
		if err != nil {
			return nil, err
		}

		info.pResolveAttachments, err = common.AllocOptionSlice[C.VkAttachmentReference2KHR, AttachmentReference2](allocator, o.ResolveAttachments)
		if err != nil {
			return nil, err
		}
	}

	if o.DepthStencilAttachment != nil {
		depthStencilPtr, err := common.AllocOptions(allocator, o.DepthStencilAttachment)
		if err != nil {
			return nil, err
		}

		info.pDepthStencilAttachment = (*C.VkAttachmentReference2KHR)(depthStencilPtr)
	}

	if preserveAttachmentCount > 0 {
		attachmentsPtr := (*C.uint32_t)(allocator.Malloc(preserveAttachmentCount * int(unsafe.Sizeof(C.uint32_t(0)))))
		attachmentsSlice := unsafe.Slice(attachmentsPtr, preserveAttachmentCount)
		for i := 0; i < preserveAttachmentCount; i++ {
			attachmentsSlice[i] = C.uint32_t(o.PreserveAttachments[i])
		}
		info.pPreserveAttachments = attachmentsPtr
	}

	return preallocatedPointer, nil
}
