package khr_copy_commands2

import "github.com/vkngwrapper/core/v3/core1_0"

//go:generate go run go.uber.org/mock/mockgen -source extiface.go -destination ./mocks/extension.go -package mock_copy_commands2

// ExtensionDriver contains all the commands for the khr_copy_commands2 extension
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VK_KHR_copy_commands2.html
type ExtensionDriver interface {
	CmdBlitImage2(commandBuffer core1_0.CommandBuffer, blitImageInfo BlitImageInfo2) error

	CmdCopyBuffer2(commandBuffer core1_0.CommandBuffer, copyBufferInfo CopyBufferInfo2) error

	CmdCopyBufferToImage2(commandBuffer core1_0.CommandBuffer, copyBufferToImageInfo CopyBufferToImageInfo2) error

	CmdCopyImage2(commandBuffer core1_0.CommandBuffer, copyImageInfo CopyImageInfo2) error

	CmdCopyImageToBuffer2(commandBuffer core1_0.CommandBuffer, copyImageToBufferInfo CopyImageToBufferInfo2) error

	CmdResolveImage2(commandBuffer core1_0.CommandBuffer, resolveImageInfo ResolveImageInfo2) error
}
