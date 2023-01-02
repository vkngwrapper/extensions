package khr_create_renderpass2_shim

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_2"
	"github.com/vkngwrapper/core/v2/driver"
	"github.com/vkngwrapper/extensions/v2/khr_create_renderpass2"
)

//go:generate mockgen -source shim.go -destination ../mocks/shim.go -package mock_create_renderpass2

// CommandBufferShim contains all the CommandBuffer commands for the khr_create_renderpass2 extension
type CommandBufferShim interface {
	// CmdBeginRenderPass2 begins a new RenderPass
	//
	// renderPassBegin - Specifies the RenderPass to begin an instance of, and the Framebuffer the instance
	// uses
	//
	// subpassBegin - Contains information about the subpass which is about to begin rendering
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass2.html
	CmdBeginRenderPass2(renderPassBegin core1_0.RenderPassBeginInfo, subpassBegin core1_2.SubpassBeginInfo) error
	// CmdEndRenderPass2 ends the current RenderPass
	//
	// subpassEnd - Contains information about how the previous subpass will be ended
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdEndRenderPass2.html
	CmdEndRenderPass2(subpassEnd core1_2.SubpassEndInfo) error
	// CmdNextSubpass2 transitions to the next subpass of a RenderPass
	//
	// subpassBegin - Contains information about the subpass which is about to begin rendering.
	//
	// subpassEnd - Contains information about how the previous subpass will be ended.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdNextSubpass2.html
	CmdNextSubpass2(subpassBegin core1_2.SubpassBeginInfo, subpassEnd core1_2.SubpassEndInfo) error
}

type VulkanCommandBufferShim struct {
	extension     khr_create_renderpass2.Extension
	commandBuffer core1_0.CommandBuffer
}

func NewCommandBufferShim(extension khr_create_renderpass2.Extension, commandBuffer core1_0.CommandBuffer) *VulkanCommandBufferShim {
	if extension == nil {
		panic("extension cannot be nil")
	}
	if commandBuffer == nil {
		panic("commandBuffer cannot be nil")
	}
	return &VulkanCommandBufferShim{
		extension:     extension,
		commandBuffer: commandBuffer,
	}
}

func (s *VulkanCommandBufferShim) CmdBeginRenderPass2(renderPassBegin core1_0.RenderPassBeginInfo, subpassBegin core1_2.SubpassBeginInfo) error {
	return s.extension.CmdBeginRenderPass2(s.commandBuffer, renderPassBegin, (khr_create_renderpass2.SubpassBeginInfo)(subpassBegin))
}

func (s *VulkanCommandBufferShim) CmdEndRenderPass2(subpassEnd core1_2.SubpassEndInfo) error {
	return s.extension.CmdEndRenderPass2(s.commandBuffer, (khr_create_renderpass2.SubpassEndInfo)(subpassEnd))
}

func (s *VulkanCommandBufferShim) CmdNextSubpass2(subpassBegin core1_2.SubpassBeginInfo, subpassEnd core1_2.SubpassEndInfo) error {
	return s.extension.CmdNextSubpass2(
		s.commandBuffer,
		khr_create_renderpass2.SubpassBeginInfo(subpassBegin),
		khr_create_renderpass2.SubpassEndInfo(subpassEnd),
	)
}

// DeviceShim contains all the Device commands for the khr_create_renderpass2 extension
type DeviceShim interface {
	// CreateRenderPass2 creates a new RenderPass object
	//
	// allocator - Controls host memory allocation behavior
	//
	// options - Describes the parameters of the RenderPass
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass2.html
	CreateRenderPass2(allocator *driver.AllocationCallbacks, options core1_2.RenderPassCreateInfo2) (core1_0.RenderPass, common.VkResult, error)
}

type VulkanDeviceShim struct {
	extension khr_create_renderpass2.Extension
	device    core1_0.Device
}

func NewDeviceShim(extension khr_create_renderpass2.Extension, device core1_0.Device) *VulkanDeviceShim {
	if extension == nil {
		panic("extension cannot be nil")
	}
	if device == nil {
		panic("device cannot be nil")
	}

	return &VulkanDeviceShim{
		extension: extension,
		device:    device,
	}
}

func (s *VulkanDeviceShim) CreateRenderPass2(allocator *driver.AllocationCallbacks, options core1_2.RenderPassCreateInfo2) (core1_0.RenderPass, common.VkResult, error) {
	// The one thing we can't autoconvert is slices of two equivalent types- sadly, there's a lot of that here.
	inOptions := khr_create_renderpass2.RenderPassCreateInfo2{
		Flags:               options.Flags,
		CorrelatedViewMasks: options.CorrelatedViewMasks,
		Attachments:         make([]khr_create_renderpass2.AttachmentDescription2, 0, len(options.Attachments)),
		Subpasses:           make([]khr_create_renderpass2.SubpassDescription2, 0, len(options.Subpasses)),
		Dependencies:        make([]khr_create_renderpass2.SubpassDependency2, 0, len(options.Dependencies)),
		NextOptions:         options.NextOptions,
	}

	for _, attachment := range options.Attachments {
		inOptions.Attachments = append(inOptions.Attachments, khr_create_renderpass2.AttachmentDescription2(attachment))
	}

	for _, subpass := range options.Subpasses {
		inSubpass := khr_create_renderpass2.SubpassDescription2{
			Flags:                  subpass.Flags,
			PipelineBindPoint:      subpass.PipelineBindPoint,
			ViewMask:               subpass.ViewMask,
			InputAttachments:       make([]khr_create_renderpass2.AttachmentReference2, 0, len(subpass.InputAttachments)),
			ColorAttachments:       make([]khr_create_renderpass2.AttachmentReference2, 0, len(subpass.ColorAttachments)),
			ResolveAttachments:     make([]khr_create_renderpass2.AttachmentReference2, 0, len(subpass.ResolveAttachments)),
			DepthStencilAttachment: (*khr_create_renderpass2.AttachmentReference2)(subpass.DepthStencilAttachment),
			PreserveAttachments:    subpass.PreserveAttachments,
			NextOptions:            subpass.NextOptions,
		}

		for _, input := range subpass.InputAttachments {
			inSubpass.InputAttachments = append(inSubpass.InputAttachments, khr_create_renderpass2.AttachmentReference2(input))
		}
		for _, color := range subpass.ColorAttachments {
			inSubpass.ColorAttachments = append(inSubpass.ColorAttachments, khr_create_renderpass2.AttachmentReference2(color))
		}
		for _, resolve := range subpass.ResolveAttachments {
			inSubpass.ResolveAttachments = append(inSubpass.ResolveAttachments, khr_create_renderpass2.AttachmentReference2(resolve))
		}
		inOptions.Subpasses = append(inOptions.Subpasses, inSubpass)
	}

	for _, dependency := range options.Dependencies {
		inOptions.Dependencies = append(inOptions.Dependencies, khr_create_renderpass2.SubpassDependency2(dependency))
	}

	return s.extension.CreateRenderPass2(
		s.device,
		allocator,
		inOptions,
	)
}
