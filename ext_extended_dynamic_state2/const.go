package ext_extended_dynamic_state2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/v3/core1_0"

const (
	// ExtensionName is "VK_EXT_extended_dynamic_state2"
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VK_EXT_extended_dynamic_state2.html
	ExtensionName string = C.VK_EXT_EXTENDED_DYNAMIC_STATE_2_EXTENSION_NAME

	// DynamicStateDepthBiasEnable specifies that the DepthBiasEnable state in
	// PipelineRasterizationStateCreateInfo will be ignored and must be set dynamically
	// with CmdSetDepthBiasEnable before any drawing commands.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkDynamicState.html
	DynamicStateDepthBiasEnable core1_0.DynamicState = C.VK_DYNAMIC_STATE_DEPTH_BIAS_ENABLE_EXT

	// DynamicStateLogicOp specifies that the LogicOp state in
	// PipelineColorBlendStateCreateInfo will be ignored and must be set dynamically with
	// CmdSetLogicOpEXT before any drawing commands.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkDynamicState.html
	DynamicStateLogicOp core1_0.DynamicState = C.VK_DYNAMIC_STATE_LOGIC_OP_EXT

	// DynamicStatePatchControlPoints specifies that the PatchControlPoints state in
	// PipelineTessellationStateCreateInfo will be ignored and must be set dynamically
	// with CmdSetPatchControlPointsEXT before any drawing commands.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkDynamicState.html
	DynamicStatePatchControlPoints core1_0.DynamicState = C.VK_DYNAMIC_STATE_PATCH_CONTROL_POINTS_EXT

	// DynamicStatePrimitiveRestartEnable specifies that the PrimitiveRestartEnable state
	// in PipelineInputAssemblyStateCreateInfo will be ignored and must be set dynamically
	// with CmdSetPrimitiveRestartEnable before any drawing commands.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkDynamicState.html
	DynamicStatePrimitiveRestartEnable core1_0.DynamicState = C.VK_DYNAMIC_STATE_PRIMITIVE_RESTART_ENABLE_EXT

	// DynamicStateRasterizerDiscardEnable specifies that the RasterizerDiscardEnable state
	// in PipelineRasterizationStateCreateInfo will be ignored and must be set dynamically with
	// CmdSetRasterizerDiscardEnable before any drawing commands.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkDynamicState.html
	DynamicStateRasterizerDiscardEnable core1_0.DynamicState = C.VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE_EXT
)

func init() {
	DynamicStateDepthBiasEnable.Register("DepthBiasEnable")
	DynamicStateLogicOp.Register("LogicOp")
	DynamicStatePatchControlPoints.Register("PatchControlPoints")
	DynamicStatePrimitiveRestartEnable.Register("PrimitiveRestartEnable")
	DynamicStateRasterizerDiscardEnable.Register("RasterizerDiscardEnable")
}
