package khr_draw_indirect_count

import (
	"github.com/vkngwrapper/core/v3"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_draw_indirect_count

// ExtensionDriver contains all the commands for the khr_draw_indirect_count extension
type ExtensionDriver interface {
	// CmdDrawIndexedIndirectCount draws with indirect parameters, indexed vertices, and draw count
	//
	// commandBuffer - The CommandBuffer to queue the draw to
	//
	// buffer - The Buffer containing draw parameters
	//
	// offset - The byte offset into buffer where parameters begin
	//
	// countBuffer - The Buffer containing the draw count
	//
	// countBufferOffset - The byte offset into countBuffer where the draw count begins
	//
	// maxDrawCount - Specifies the maximum number of draws that will be executed.
	//
	// stride - The byte stride between successive sets of draw parameters
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexedIndirectCount.html
	CmdDrawIndexedIndirectCount(commandBuffer core.CommandBuffer, buffer core.Buffer, offset uint64, countBuffer core.Buffer, countBufferOffset uint64, maxDrawCount, stride int)
	// CmdDrawIndirectCount draws primitives with indirect parameters and draw count
	//
	// commandBuffer - The CommandBuffer to queue the draw to
	//
	// buffer - The Buffer containing draw parameters
	//
	// offset - The byte offset into buffer where parameters begin
	//
	// countBuffer - The Buffer containing the draw count
	//
	// countBufferOffset - The byte offset into countBuffer where the draw count begins
	//
	// maxDrawCount - Specifies the maximum number of draws that will be executed.
	//
	// stride - The byte stride between successive sets of draw parameters
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndirectCount.html
	CmdDrawIndirectCount(commandBuffer core.CommandBuffer, buffer core.Buffer, offset uint64, countBuffer core.Buffer, countBufferOffset uint64, maxDrawCount, stride int)
}
