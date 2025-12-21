package khr_draw_indirect_count_shim

import (
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/extensions/v3/khr_draw_indirect_count"
)

//go:generate mockgen -source shim.go -destination ../mocks/shim.go -package mock_draw_indirect_count

// Shim contains all the commands for the khr_draw_indirect_count extension
type Shim interface {
	// CmdDrawIndexedIndirectCount draws with indirect parameters, indexed vertices, and draw count
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
	CmdDrawIndexedIndirectCount(buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int)
	// CmdDrawIndirectCount draws primitives with indirect parameters and draw count
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
	CmdDrawIndirectCount(buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int)
}

// Compiler check that VulkanShim satisfies Shim
var _ Shim = &VulkanShim{}

type VulkanShim struct {
	extension     khr_draw_indirect_count.Extension
	commandBuffer core1_0.CommandBuffer
}

func NewShim(extension khr_draw_indirect_count.Extension, buffer core1_0.CommandBuffer) *VulkanShim {
	if extension == nil {
		panic("extension cannot be nil")
	}
	if buffer == nil {
		panic("buffer cannot be nil")
	}

	return &VulkanShim{
		extension:     extension,
		commandBuffer: buffer,
	}
}

func (s *VulkanShim) CmdDrawIndexedIndirectCount(buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int) {
	s.extension.CmdDrawIndexedIndirectCount(s.commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride)
}

func (s *VulkanShim) CmdDrawIndirectCount(buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int) {
	s.extension.CmdDrawIndirectCount(s.commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride)
}
