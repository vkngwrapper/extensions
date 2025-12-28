package khr_create_renderpass2

import (
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_create_renderpass2

// ExtensionDriver contains all the commands for the khr_create_renderpass2 extension
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_create_renderpass2.html
type ExtensionDriver interface {
	// CmdBeginRenderPass2 begins a new RenderPass
	//
	// commandBuffer - The CommandBuffer to begin the RenderPass in
	//
	// renderPassBegin - Specifies the RenderPass to begin an instance of, and the Framebuffer the instance
	// uses
	//
	// subpassBegin - Contains information about the subpass which is about to begin rendering
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass2.html
	CmdBeginRenderPass2(commandBuffer core.CommandBuffer, renderPassBegin core1_0.RenderPassBeginInfo, subpassBegin SubpassBeginInfo) error
	// CmdEndRenderPass2 ends the current RenderPass
	//
	// commandBuffer - The CommandBuffer to end the RenderPass in
	//
	// subpassEnd - Contains information about how the previous subpass will be ended
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdEndRenderPass2.html
	CmdEndRenderPass2(commandBuffer core.CommandBuffer, subpassEnd SubpassEndInfo) error
	// CmdNextSubpass2 transitions to the next subpass of a RenderPass
	//
	// commandBuffer - The CommandBuffer to end the RenderPass in
	//
	// subpassBegin - Contains information about the subpass which is about to begin rendering.
	//
	// subpassEnd - Contains information about how the previous subpass will be ended.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdNextSubpass2.html
	CmdNextSubpass2(commandBuffer core.CommandBuffer, subpassBegin SubpassBeginInfo, subpassEnd SubpassEndInfo) error

	// CreateRenderPass2 creates a new RenderPass object
	//
	// device - The Device to create the RenderPass from
	//
	// allocator - Controls host memory allocation behavior
	//
	// options - Describes the parameters of the RenderPass
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass2.html
	CreateRenderPass2(device core.Device, allocator *loader.AllocationCallbacks, options RenderPassCreateInfo2) (core.RenderPass, common.VkResult, error)
}
