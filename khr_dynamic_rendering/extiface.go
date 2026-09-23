package khr_dynamic_rendering

import "github.com/vkngwrapper/core/v3/core1_0"

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_dynamic_rendering

// ExtensionDriver contains all the commands for the khr_dynamic_rendering extension
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VK_KHR_dynamic_rendering.html
type ExtensionDriver interface {
	// CmdBeginRendering begins a new render pass instance
	//
	// commandBuffer - The CommandBuffer to begin rendering in
	//
	// renderingInfo - Specifies the rendering parameters and attachments
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdBeginRenderingKHR.html
	CmdBeginRendering(commandBuffer core1_0.CommandBuffer, renderingInfo RenderingInfo) error
	// CmdEndRendering ends the current render pass instance
	//
	// commandBuffer - The CommandBuffer to end rendering in
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdEndRenderingKHR.html
	CmdEndRendering(commandBuffer core1_0.CommandBuffer) error
}
