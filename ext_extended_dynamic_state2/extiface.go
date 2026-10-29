package ext_extended_dynamic_state2

import "github.com/vkngwrapper/core/v3/core1_0"

//go:generate go run go.uber.org/mock/mockgen -source extiface.go -destination ./mocks/extension.go -package mock_extended_dynamic_state2

// ExtensionDriver contains all commands for the ext_extended_dynamic_state2
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VK_EXT_extended_dynamic_state2.html
type ExtensionDriver interface {
	// CmdSetDepthBiasEnable controls whether to bias fragment depth values dynamically for
	// a command buffer
	//
	// commandBuffer - the command buffer into which the command will be recorded
	//
	// depthBiasEnable - whether to bias fragment depth values
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdSetDepthBiasEnable.html
	CmdSetDepthBiasEnable(commandBuffer core1_0.CommandBuffer, depthBiasEnable bool)

	// CmdSetLogicOp selects which logical operation to apply for blend state dynamically
	// for a command buffer
	//
	// commandBuffer - the command buffer into which the command will be recorded
	//
	// logicOp - specifies the logical operation to apply for blend state.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdSetLogicOpEXT.html
	CmdSetLogicOp(commandBuffer core1_0.CommandBuffer, logicOp core1_0.LogicOp)

	// CmdSetPatchControlPoints specifies the number of control points per patch dynamically
	// for a command buffer
	//
	// commandBuffer - the command buffer into which the command will be recorded
	//
	// patchControlPoints - specifies the number of control points per patch
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdSetPatchControlPointsEXT.html
	CmdSetPatchControlPoints(commandBuffer core1_0.CommandBuffer, patchControlPoints int)

	// CmdSetPrimitiveRestartEnable sets primitive assembly restart state dynamically for a
	// command buffer
	//
	// commandBuffer - the command buffer into which the command will be recorded
	//
	// primitiveRestartEnable - controls whether a special vertex index value is treated as
	// restarting the assembly of primitives. It behaves in the same way as
	// PipelineInputAssemblyStateCreateInfo.PrimitiveRestartEnable.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdSetPrimitiveRestartEnable.html
	CmdSetPrimitiveRestartEnable(commandBuffer core1_0.CommandBuffer, primitiveRestartEnable bool)

	// CmdSetRasterizerDiscardEnable controls whether primitives are discarded before the
	// rasterization stage dynamically for a command buffer
	//
	// commandBuffer - the command buffer into which the command will be recorded
	//
	// rasterizerDiscardEnable - controls whether primitives are discarded immediately before
	// the rasterization stage
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/vkCmdSetRasterizerDiscardEnable.html
	CmdSetRasterizerDiscardEnable(commandBuffer core1_0.CommandBuffer, rasterizerDiscardEnable bool)
}
