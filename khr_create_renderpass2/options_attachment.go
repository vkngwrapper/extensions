package khr_create_renderpass2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"unsafe"
)

// AttachmentDescription2 specifies an attachment description
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription2.html
type AttachmentDescription2 struct {
	// Flags specifies additional properties of the attachment
	Flags core1_0.AttachmentDescriptionFlags
	// Format specifies the format of the Image that will be used for the attachment
	Format core1_0.Format
	// Samples specifies the number of samples of the Image
	Samples core1_0.SampleCountFlags
	// LoadOp specifies how the contents of color and depth components of the attachment
	// are treated at the beginning of the subpass where it is first used
	LoadOp core1_0.AttachmentLoadOp
	// StoreOp specifies how the contents of color and depth components of the attachment
	// are treated at the end of the subpass where it is last used
	StoreOp core1_0.AttachmentStoreOp
	// StencilLoadOp specifies how the contents of stencil components of the attachment
	// are treated at the beginning of the subpass where it is first used
	StencilLoadOp core1_0.AttachmentLoadOp
	// StencilStoreOp specifies how the contents of the stencil components of the attachment
	// are treated at the end of the last subpass where it is used
	StencilStoreOp core1_0.AttachmentStoreOp
	// InitialLayout is the layout of the attachment Image subresource will be in when
	// a RenderPass instance begins
	InitialLayout core1_0.ImageLayout
	// FinalLayout is the layout the attachment Image subresource will be transitioned to
	// when a RenderPass instance ends
	FinalLayout core1_0.ImageLayout

	common.NextOptions
}

func (o AttachmentDescription2) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAttachmentDescription2KHR{})))
	}

	info := (*C.VkAttachmentDescription2KHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2_KHR
	info.pNext = next
	info.flags = C.VkAttachmentDescriptionFlags(o.Flags)
	info.format = C.VkFormat(o.Format)
	info.samples = C.VkSampleCountFlagBits(o.Samples)
	info.loadOp = C.VkAttachmentLoadOp(o.LoadOp)
	info.storeOp = C.VkAttachmentStoreOp(o.StoreOp)
	info.stencilLoadOp = C.VkAttachmentLoadOp(o.StencilLoadOp)
	info.stencilStoreOp = C.VkAttachmentStoreOp(o.StencilStoreOp)
	info.initialLayout = C.VkImageLayout(o.InitialLayout)
	info.finalLayout = C.VkImageLayout(o.FinalLayout)

	return preallocatedPointer, nil
}
