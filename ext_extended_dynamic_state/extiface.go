package ext_extended_dynamic_state

import "github.com/vkngwrapper/core/v3/core1_0"

//go:generate go run go.uber.org/mock/mockgen -source extiface.go -destination ./mocks/extension.go -package mock_extended_dynamic_state

// ExtensionDriver contains all commands for the ext_extended_dynamic_state
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VK_EXT_extended_dynamic_state.html
type ExtensionDriver interface {
	// CmdBindVertexBuffers2 binds vertex buffers to a command buffer and dynamically sets
	// strides
	//
	// commandBuffer - the command buffer into which the command is recorded
	//
	// firstBinding - the index of the first vertex input binding whose state is updated by the command
	//
	// buffers - a slice of Buffer objects
	//
	// offsets - a slice of buffer offsets
	//
	// sizes - nil or a slice of the size in bytes of vertex data bound from buffers
	//
	// strides - nil or a slice of buffer strides
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdBindVertexBuffers2.html
	CmdBindVertexBuffers2(commandBuffer core1_0.CommandBuffer, firstBinding int, buffers []core1_0.Buffer, offsets []int, sizes []int, strides []int)

	// CmdSetCullMode sets cull mode dynamically for a command buffer
	//
	// commandBuffer - the command buffer into which the command is recorded
	//
	// cullMode - the cull mode property to use for drawing
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdSetCullMode.html
	CmdSetCullMode(commandBuffer core1_0.CommandBuffer, cullMode core1_0.CullModeFlags)

	// CmdSetDepthBoundsTestEnable sets depth bounds test enable dynamically for a command buffer
	//
	// commandBuffer - the command buffer into which the command is recorded
	//
	// depthBoundsTestEnable - specifies if the depth bounds test is enabled
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdSetDepthBoundsTestEnable.html
	CmdSetDepthBoundsTestEnable(commandBuffer core1_0.CommandBuffer, depthBoundsTestEnable bool)

	// CmdSetDepthCompareOp sets depth comparison operator dynamically for a command buffer
	//
	// commandBuffer - the command buffer into which the command is recorded
	//
	// depthCompareOp - a CompareOp value specifying the comparison operator used for the
	// Depth Comparison step of the depth test
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdSetDepthCompareOp.html
	CmdSetDepthCompareOp(commandBuffer core1_0.CommandBuffer, depthCompareOp core1_0.CompareOp)

	// CmdSetDepthTestEnable sets depth test enable dynamically for a command buffer
	//
	// commandBuffer - the command buffer into which the command is recorded
	//
	// depthTestEnable - specifies if the depth test is enabled
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdSetDepthTestEnable.html
	CmdSetDepthTestEnable(commandBuffer core1_0.CommandBuffer, depthTestEnable bool)

	// CmdSetDepthWriteEnable sets depth write enable dynamically for a command buffer
	//
	// commandBuffer - the command buffer into which the command is recorded
	//
	// depthWriteEnable - specifies if depth writes are enabled
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdSetDepthWriteEnable.html
	CmdSetDepthWriteEnable(commandBuffer core1_0.CommandBuffer, depthWriteEnable bool)

	// CmdSetFrontFace sets front face orientation dynamically for a command buffer
	//
	// commandBuffer - the command buffer into which the command is recorded
	//
	// frontFace - a FrontFace value specifying the front-facing triangle orientation to
	// be used for culling
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdSetFrontFace.html
	CmdSetFrontFace(commandBuffer core1_0.CommandBuffer, frontFace core1_0.FrontFace)

	// CmdSetPrimitiveTopology sets primitive topology state dynamically for a command buffer
	//
	// commandBuffer - the command buffer into which the command is recorded
	//
	// primitiveTopology - specifies the primitive topology to use for drawing
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdSetPrimitiveTopology.html
	CmdSetPrimitiveTopology(commandBuffer core1_0.CommandBuffer, primitiveTopology core1_0.PrimitiveTopology)

	// CmdSetScissorWithCount sets the scissor rectangular bounds
	// dynamically for a command buffer
	//
	// commandBuffer - the command buffer into which the command is recorded
	//
	// scissors - specifies the scissors to use for drawing
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdSetScissorWithCount.html
	CmdSetScissorWithCount(commandBuffer core1_0.CommandBuffer, scissors []core1_0.Rect2D)

	// CmdSetStencilOp sets stencil operation dynamically for a command buffer
	//
	// commandBuffer - the command buffer into which the command is recorded
	//
	// faceMask - a bitmask of StencilFaceFlags specifying the set of stencil state for which
	// to update the stencil operation
	//
	// failOp - a StencilOp value specifying the actions performed on samples that fail the
	// stencil test
	//
	// passOp - a StencilOp value specifying the actions performed on samples that pass
	// both the depth and stencil tests
	//
	// depthFailOp - a StencilOp value specifying the actions performed on samples that pass
	// the stencil test and fail the depth test
	//
	// compareOp - a CompareOp value specifying the comparison operator used in the stencil test
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdSetStencilOp.html
	CmdSetStencilOp(commandBuffer core1_0.CommandBuffer, faceMask core1_0.StencilFaceFlags, failOp core1_0.StencilOp, passOp core1_0.StencilOp, depthFailOp core1_0.StencilOp, compareOp core1_0.CompareOp)

	// CmdSetStencilTestEnable sets stencil test enable dynamically for a command buffer
	//
	// commandBuffer - the command buffer into which the command is recorded
	//
	// stencilTestEnable - specifies if the stencil test is enabled
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdSetStencilTestEnable.html
	CmdSetStencilTestEnable(commandBuffer core1_0.CommandBuffer, stencilTestEnable bool)

	// CmdSetViewportWithCount sets the viewports dynamically for a command buffer
	//
	// commandBuffer - the command buffer into which the command is recorded
	//
	// viewports - the viewports to use for drawing
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdSetViewportWithCount.html
	CmdSetViewportWithCount(commandBuffer core1_0.CommandBuffer, viewports []core1_0.Viewport)
}
