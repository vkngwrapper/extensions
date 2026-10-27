package ext_extended_dynamic_state

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/v3/core1_0"

const (
	// ExtensionName is "VK_EXT_extended_dynamic_state"
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VK_EXT_extended_dynamic_state.html
	ExtensionName string = C.VK_EXT_EXTENDED_DYNAMIC_STATE_EXTENSION_NAME

	// DynamicStateCullMode specifies that the cullMode state in
	// PipelineRasterizationStateCreateInfo will be ignored and must be set dynamically
	// with CmdSetCullMode before any drawing commands.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkDynamicState.html
	DynamicStateCullMode core1_0.DynamicState = C.VK_DYNAMIC_STATE_CULL_MODE_EXT
	// DynamicStateDepthBoundsTestEnable specifies that the DepthBoundsTestEnable state in
	// PipelineDepthStencilStateCreateInfo will be ignored and must be set dynamically with
	// CmdSetDepthBoundsTestEnable before any draw call.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkDynamicState.html
	DynamicStateDepthBoundsTestEnable core1_0.DynamicState = C.VK_DYNAMIC_STATE_DEPTH_BOUNDS_TEST_ENABLE_EXT
	// DynamicStateDepthCompareOp specifies that the DepthCompareOp state in
	// PipelineDepthStencilStateCreateInfo will be ignored and must be set dynamically with
	// CmdSetDepthCompareOp before any draw call.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkDynamicState.html
	DynamicStateDepthCompareOp core1_0.DynamicState = C.VK_DYNAMIC_STATE_DEPTH_COMPARE_OP_EXT
	// DynamicStateDepthTestEnable specifies that the DepthTestEnable state in
	// PipelineDepthStencilStateCreateInfo will be ignored and must be set dynamically with
	// vkCmdSetDepthTestEnable before any draw call.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkDynamicState.html
	DynamicStateDepthTestEnable core1_0.DynamicState = C.VK_DYNAMIC_STATE_DEPTH_TEST_ENABLE_EXT
	// DynamicStateDepthWriteEnable specifies that the DepthWriteEnable state in
	// PipelineDepthStencilStateCreateInfo will be ignored and must be set dynamically
	// with CmdSetDepthWriteEnable before any draw call.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkDynamicState.html
	DynamicStateDepthWriteEnable core1_0.DynamicState = C.VK_DYNAMIC_STATE_DEPTH_WRITE_ENABLE_EXT
	// DynamicStateFrontFace specifies that the FrontFace state in
	// PipelineRasterizationStateCreateInfo will be ignored and must be set dynamically
	// with CmdSetFrontFace before any drawing commands.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkDynamicState.html
	DynamicStateFrontFace core1_0.DynamicState = C.VK_DYNAMIC_STATE_FRONT_FACE_EXT
	// DynamicStatePrimitiveTopology specifies that the Topology state in
	// PipelineInputAssemblyStateCreateInfo only specifies the topology class, and the
	// specific topology order and adjacency must be set dynamically with
	// CmdSetPrimitiveTopology before any drawing commands.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkDynamicState.html
	DynamicStatePrimitiveTopology core1_0.DynamicState = C.VK_DYNAMIC_STATE_PRIMITIVE_TOPOLOGY_EXT
	// DynamicStateScissorWithCount specifies that the Scissors state in
	// PipelineViewportStateCreateInfo will be ignored and must be set dynamically with
	// CmdSetScissorWithCount before any draw call.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkDynamicState.html
	DynamicStateScissorWithCount core1_0.DynamicState = C.VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT_EXT
	// DynamicStateStencilOp specifies that the FailOp, PassOp, DepthFailOp, and CompareOp
	// states in PipelineDepthStencilStateCreateInfo for both front and back will be ignored
	// and must be set dynamically with CmdSetStencilOp before any draws are performed with
	// a pipeline state with PipelineDepthStencilStateCreateInfo member StencilTestEnable
	// set to true.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkDynamicState.html
	DynamicStateStencilOp core1_0.DynamicState = C.VK_DYNAMIC_STATE_STENCIL_OP_EXT
	// DynamicStateStencilTestEnable specifies that the StencilTestEnable state in
	// PipelineDepthStencilStateCreateInfo will be ignored and must be set dynamically with
	// CmdSetStencilTestEnable before any draw call.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkDynamicState.html
	DynamicStateStencilTestEnable core1_0.DynamicState = C.VK_DYNAMIC_STATE_STENCIL_TEST_ENABLE_EXT
	// DynamicStateVertexInputBindingStride specifies that the stride state in
	// VertexInputBindingDescription will be ignored and must be set dynamically with
	// CmdBindVertexBuffers2 before any draw call.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkDynamicState.html
	DynamicStateVertexInputBindingStride core1_0.DynamicState = C.VK_DYNAMIC_STATE_VERTEX_INPUT_BINDING_STRIDE_EXT
	// DynamicStateViewportWithCount specifies that the Viewports state in
	// PipelineViewportStateCreateInfo will be ignored and must be set dynamically with
	// CmdSetViewportWithCount before any draw call.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkDynamicState.html
	DynamicStateViewportWithCount core1_0.DynamicState = C.VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT_EXT
)

func init() {
	DynamicStateCullMode.Register("CullMode")
	DynamicStateDepthBoundsTestEnable.Register("DepthBoundsTestEnable")
	DynamicStateDepthCompareOp.Register("DepthCompareOp")
	DynamicStateDepthTestEnable.Register("DepthTestEnable")
	DynamicStateDepthWriteEnable.Register("DepthWriteEnable")
	DynamicStateFrontFace.Register("FrontFace")
	DynamicStatePrimitiveTopology.Register("PrimitiveTopology")
	DynamicStateScissorWithCount.Register("ScissorWithCount")
	DynamicStateStencilOp.Register("StencilOp")
	DynamicStateStencilTestEnable.Register("StencilTestEnable")
	DynamicStateVertexInputBindingStride.Register("VertexInputBindingStride")
	DynamicStateViewportWithCount.Register("ViewportWithCount")
}
