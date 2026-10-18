package khr_synchronization2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"

const uint64_t ACCESS2_NONE = VK_ACCESS_2_NONE_KHR;
const uint64_t ACCESS2_INDIRECT_COMMAND_READ = VK_ACCESS_2_INDIRECT_COMMAND_READ_BIT_KHR;
const uint64_t ACCESS2_INDEX_READ = VK_ACCESS_2_INDEX_READ_BIT_KHR;
const uint64_t ACCESS2_VERTEX_ATTRIBUTE_READ = VK_ACCESS_2_VERTEX_ATTRIBUTE_READ_BIT_KHR;
const uint64_t ACCESS2_UNIFORM_READ = VK_ACCESS_2_UNIFORM_READ_BIT_KHR;
const uint64_t ACCESS2_INPUT_ATTACHMENT_READ = VK_ACCESS_2_INPUT_ATTACHMENT_READ_BIT_KHR;
const uint64_t ACCESS2_SHADER_READ = VK_ACCESS_2_SHADER_READ_BIT_KHR;
const uint64_t ACCESS2_SHADER_WRITE = VK_ACCESS_2_SHADER_WRITE_BIT_KHR;
const uint64_t ACCESS2_COLOR_ATTACHMENT_READ = VK_ACCESS_2_COLOR_ATTACHMENT_READ_BIT_KHR;
const uint64_t ACCESS2_COLOR_ATTACHMENT_WRITE = VK_ACCESS_2_COLOR_ATTACHMENT_WRITE_BIT_KHR;
const uint64_t ACCESS2_DEPTH_STENCIL_ATTACHMENT_READ = VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_READ_BIT_KHR;
const uint64_t ACCESS2_DEPTH_STENCIL_ATTACHMENT_WRITE = VK_ACCESS_2_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT_KHR;
const uint64_t ACCESS2_TRANSFER_READ = VK_ACCESS_2_TRANSFER_READ_BIT_KHR;
const uint64_t ACCESS2_TRANSFER_WRITE = VK_ACCESS_2_TRANSFER_WRITE_BIT_KHR;
const uint64_t ACCESS2_HOST_READ = VK_ACCESS_2_HOST_READ_BIT_KHR;
const uint64_t ACCESS2_HOST_WRITE = VK_ACCESS_2_HOST_WRITE_BIT_KHR;
const uint64_t ACCESS2_MEMORY_READ = VK_ACCESS_2_MEMORY_READ_BIT_KHR;
const uint64_t ACCESS2_MEMORY_WRITE = VK_ACCESS_2_MEMORY_WRITE_BIT_KHR;
const uint64_t ACCESS2_SHADER_SAMPLED_READ = VK_ACCESS_2_SHADER_SAMPLED_READ_BIT_KHR;
const uint64_t ACCESS2_SHADER_STORAGE_READ = VK_ACCESS_2_SHADER_STORAGE_READ_BIT_KHR;
const uint64_t ACCESS2_SHADER_STORAGE_WRITE = VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT_KHR;
const uint64_t ACCESS2_TRANSFORM_FEEDBACK_WRITE = VK_ACCESS_2_TRANSFORM_FEEDBACK_WRITE_BIT_EXT;
const uint64_t ACCESS2_TRANSFORM_FEEDBACK_COUNTER_READ = VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_READ_BIT_EXT;
const uint64_t ACCESS2_TRANSFORM_FEEDBACK_COUNTER_WRITE = VK_ACCESS_2_TRANSFORM_FEEDBACK_COUNTER_WRITE_BIT_EXT;
const uint64_t ACCESS2_CONDITIONAL_RENDERING_READ = VK_ACCESS_2_CONDITIONAL_RENDERING_READ_BIT_EXT;
const uint64_t ACCESS2_COMMAND_PREPROCESS_READ_NV = VK_ACCESS_2_COMMAND_PREPROCESS_READ_BIT_NV;
const uint64_t ACCESS2_COMMAND_PREPROCESS_WRITE_NV = VK_ACCESS_2_COMMAND_PREPROCESS_WRITE_BIT_NV;
const uint64_t ACCESS2_COMMAND_PREPROCESS_READ = VK_ACCESS_2_COMMAND_PREPROCESS_READ_BIT_EXT;
const uint64_t ACCESS2_COMMAND_PREPROCESS_WRITE = VK_ACCESS_2_COMMAND_PREPROCESS_WRITE_BIT_EXT;
const uint64_t ACCESS2_FRAGMENT_SHADING_RATE_ATTACHMENT_READ = VK_ACCESS_2_FRAGMENT_SHADING_RATE_ATTACHMENT_READ_BIT_KHR;
const uint64_t ACCESS2_SHADING_RATE_IMAGE_READ_NV = VK_ACCESS_2_SHADING_RATE_IMAGE_READ_BIT_NV;
const uint64_t ACCESS2_ACCELERATION_STRUCTURE_READ = VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_KHR;
const uint64_t ACCESS2_ACCELERATION_STRUCTURE_WRITE = VK_ACCESS_2_ACCELERATION_STRUCTURE_WRITE_BIT_KHR;
const uint64_t ACCESS2_ACCELERATION_STRUCTURE_READ_NV = VK_ACCESS_2_ACCELERATION_STRUCTURE_READ_BIT_NV;
const uint64_t ACCESS2_ACCELERATION_STRUCTURE_WRITE_NV = VK_ACCESS_2_ACCELERATION_STRUCTURE_WRITE_BIT_NV;
const uint64_t ACCESS2_FRAGMENT_DENSITY_MAP_READ = VK_ACCESS_2_FRAGMENT_DENSITY_MAP_READ_BIT_EXT;
const uint64_t ACCESS2_COLOR_ATTACHMENT_READ_NONCOHERENT = VK_ACCESS_2_COLOR_ATTACHMENT_READ_NONCOHERENT_BIT_EXT;

const uint64_t PIPELINE2_NONE = VK_PIPELINE_STAGE_2_NONE_KHR;
const uint64_t PIPELINE2_TOP_OF_PIPE = VK_PIPELINE_STAGE_2_TOP_OF_PIPE_BIT_KHR;
const uint64_t PIPELINE2_DRAW_INDIRECT = VK_PIPELINE_STAGE_2_DRAW_INDIRECT_BIT_KHR;
const uint64_t PIPELINE2_VERTEX_SHADER = VK_PIPELINE_STAGE_2_VERTEX_SHADER_BIT_KHR;
const uint64_t PIPELINE2_VERTEX_INPUT = VK_PIPELINE_STAGE_2_VERTEX_INPUT_BIT_KHR;
const uint64_t PIPELINE2_TESSELLATION_CONTROL_SHADER = VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT_KHR;
const uint64_t PIPELINE2_TESSELLATION_EVALUATION_SHADER = VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT_KHR;
const uint64_t PIPELINE2_GEOMETRY_SHADER = VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT_KHR;
const uint64_t PIPELINE2_FRAGMENT_SHADER = VK_PIPELINE_STAGE_2_FRAGMENT_SHADER_BIT_KHR;
const uint64_t PIPELINE2_EARLY_FRAGMENT_TESTS = VK_PIPELINE_STAGE_2_EARLY_FRAGMENT_TESTS_BIT_KHR;
const uint64_t PIPELINE2_LATE_FRAGMENT_TESTS = VK_PIPELINE_STAGE_2_LATE_FRAGMENT_TESTS_BIT_KHR;
const uint64_t PIPELINE2_COLOR_ATTACHMENT_OUTPUT = VK_PIPELINE_STAGE_2_COLOR_ATTACHMENT_OUTPUT_BIT_KHR;
const uint64_t PIPELINE2_COMPUTE_SHADER = VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT_KHR;
const uint64_t PIPELINE2_ALL_TRANSFER = VK_PIPELINE_STAGE_2_ALL_TRANSFER_BIT_KHR;
const uint64_t PIPELINE2_BOTTOM_OF_PIPE = VK_PIPELINE_STAGE_2_BOTTOM_OF_PIPE_BIT_KHR;
const uint64_t PIPELINE2_HOST = VK_PIPELINE_STAGE_2_HOST_BIT_KHR;
const uint64_t PIPELINE2_ALL_GRAPHICS = VK_PIPELINE_STAGE_2_ALL_GRAPHICS_BIT_KHR;
const uint64_t PIPELINE2_ALL_COMMANDS = VK_PIPELINE_STAGE_2_ALL_COMMANDS_BIT_KHR;
const uint64_t PIPELINE2_COPY = VK_PIPELINE_STAGE_2_COPY_BIT_KHR;
const uint64_t PIPELINE2_RESOLVE = VK_PIPELINE_STAGE_2_RESOLVE_BIT_KHR;
const uint64_t PIPELINE2_BLIT = VK_PIPELINE_STAGE_2_BLIT_BIT_KHR;
const uint64_t PIPELINE2_CLEAR = VK_PIPELINE_STAGE_2_CLEAR_BIT_KHR;
const uint64_t PIPELINE2_INDEX_INPUT = VK_PIPELINE_STAGE_2_INDEX_INPUT_BIT_KHR;
const uint64_t PIPELINE2_VERTEX_ATTRIBUTE_INPUT = VK_PIPELINE_STAGE_2_VERTEX_ATTRIBUTE_INPUT_BIT_KHR;
const uint64_t PIPELINE2_PRE_RASTERIZATION_SHADERS = VK_PIPELINE_STAGE_2_PRE_RASTERIZATION_SHADERS_BIT_KHR;
const uint64_t PIPELINE2_TRANSFORM_FEEDBACK = VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT;
const uint64_t PIPELINE2_CONDITIONAL_RENDERING = VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT;
const uint64_t PIPELINE2_COMMAND_PREPROCESS_NV = VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_NV;
const uint64_t PIPELINE2_COMMAND_PREPROCESS = VK_PIPELINE_STAGE_2_COMMAND_PREPROCESS_BIT_EXT;
const uint64_t PIPELINE2_FRAGMENT_SHADING_RATE_ATTACHMENT = VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR;
const uint64_t PIPELINE2_ACCELERATION_STRUCTURE_BUILD = VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR;
const uint64_t PIPELINE2_RAY_TRACING_SHADER = VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR;
const uint64_t PIPELINE2_RAY_TRACING_SHADER_NV = VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_NV;
const uint64_t PIPELINE2_ACCELERATION_STRUCTURE_BUILD_NV = VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_NV;
const uint64_t PIPELINE2_FRAGMENT_DENSITY_PROCESS = VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT;
const uint64_t PIPELINE2_TASK_SHADER_NV = VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_NV;
const uint64_t PIPELINE2_MESH_SHADER_NV = VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_NV;
const uint64_t PIPELINE2_TASK_SHADER = VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT;
const uint64_t PIPELINE2_MESH_SHADER = VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT;
*/
import "C"
import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
)

// AccessFlags2 is a bitmask indicating the descriptor type used, or how a fixed-
// function stage accesses memory.
//
// https://docs.vulkan.org/spec/latest/chapters/synchronization.html#VkAccessFlagBits2
type AccessFlags2 uint64

var accessFlags2Mapping = common.NewFlagStringMapping[AccessFlags2]()

func (f AccessFlags2) Register(str string) {
	accessFlags2Mapping.Register(f, str)
}

func (f AccessFlags2) String() string {
	return accessFlags2Mapping.FlagsToString(f)
}

// SubmitFlags is a bitmask specifying behavior of a submission.
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkSubmitFlagBits.html
type SubmitFlags int32

var submitFlagsMapping = common.NewFlagStringMapping[SubmitFlags]()

func (f SubmitFlags) Register(str string) {
	submitFlagsMapping.Register(f, str)
}

func (f SubmitFlags) String() string {
	return submitFlagsMapping.FlagsToString(f)
}

// PipelineStage2 is a bitmask specifying stages of execution
//
// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
type PipelineStageFlags2 uint64

var pipelineStageFlags2Mapping = common.NewFlagStringMapping[PipelineStageFlags2]()

func (f PipelineStageFlags2) Register(str string) {
	pipelineStageFlags2Mapping.Register(f, str)
}

func (f PipelineStageFlags2) String() string {
	return pipelineStageFlags2Mapping.FlagsToString(f)
}

const (
	// ExtensionName is "VK_KHR_synchronization2"
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VK_KHR_synchronization2.html
	ExtensionName string = C.VK_KHR_SYNCHRONIZATION_2_EXTENSION_NAME

	// AccessNone specifies no accesses
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits.html
	AccessNone core1_0.AccessFlags = C.VK_ACCESS_NONE_KHR

	// EventCreateDeviceOnly specifies that host event commands will not be used with this event.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkEventCreateFlagBits.html
	EventCreateDeviceOnly core1_0.EventCreateFlags = C.VK_EVENT_CREATE_DEVICE_ONLY_BIT_KHR

	// ImageLayoutAttachmentOptimal specifies a layout that must only be used with attachment
	// accesses in the graphics pipeline.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkImageLayout.html
	ImageLayoutAttachmentOptimal core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR

	// ImageLayoutReadOnlyOptimal specifies a layout allowing read only access as an attachment,
	// or in shaders as a sampled image, combined image/sampler, or input attachment.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkImageLayout.html
	ImageLayoutReadOnlyOptimal core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR

	// PipelineStageNone specifies no stages of execution.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkImageLayout.html
	PipelineStageNone core1_0.PipelineStageFlags = C.VK_PIPELINE_STAGE_NONE_KHR

	// SubmitProtected specifies that this batch is a protected submission.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkSubmitFlagBits.html
	SubmitProtected SubmitFlags = C.VK_SUBMIT_PROTECTED_BIT_KHR
)

var (
	// Access2None specifies no accesses
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2None = AccessFlags2(C.ACCESS2_NONE)

	// Access2IndirectCommandRead specifies read access to command data read from indirect
	// buffers as part of an indirect build, trace, drawing or dispatch command. Such access
	// occurs in the PipelineStage2DrawIndirect pipeline stage. It also specifies read access to
	// command data read from indirect buffers as part of a copy command with access
	// occurring in the PipelineStage2CopyIndirect pipeline stage
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2IndirectCommandRead = AccessFlags2(C.ACCESS2_INDIRECT_COMMAND_READ)

	// Access2IndexRead specifies read access to an index buffer as part of an indexed
	// drawing command, bound by CmdBindIndexBuffer2 and CmdBindIndexBuffer. Such access
	// occurs in the PipelineStage2IndexInput pipeline stage.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2IndexRead = AccessFlags2(C.ACCESS2_INDEX_READ)

	// Access2VertexAttributeRead specifies read access to a vertex buffer as part of a
	// drawing command, bound by CmdBindVertexBuffers. Such access occurs in the
	// PipelineStage2VertexAttributeInput pipeline stage.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2VertexAttributeRead = AccessFlags2(C.ACCESS2_VERTEX_ATTRIBUTE_READ)

	// Access2UniformRead specifies read access to a uniform buffer in any shader pipeline stage.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2UniformRead = AccessFlags2(C.ACCESS2_UNIFORM_READ)

	// Access2InputAttachmentRead specifies read access to an input attachment within a
	// render pass during subpass shading or fragment shading. Such access occurs in the
	// PipelineStage2SubpassShaderHuawei or PipelineStage2FragmentShader pipeline stage.
	Access2InputAttachmentRead = AccessFlags2(C.ACCESS2_INPUT_ATTACHMENT_READ)

	// Access2ShaderRead is equivalent to (Access2ShaderSampledRead|Access2ShaderStorageRead|
	// Access2ShaderBindingTableRead|Access2ShaderTileAttachmentReadQcom)
	Access2ShaderRead = AccessFlags2(C.ACCESS2_SHADER_READ)

	// Access2ShaderWrite is equivalent to Access2ShaderStorageWrite
	Access2ShaderWrite = AccessFlags2(C.ACCESS2_SHADER_WRITE)

	// Access2ColorAttachmentRead specifies read access to a color attachment, such as via
	// blending (other than advanced blend operations), logic operations or certain render
	// pass load operations in the PipelineStage2ColorAttachmentOutput pipeline stage or
	// via fragment shader tile image reads in the PipelineStage2FragmentShader pipeline stage.
	Access2ColorAttachmentRead = AccessFlags2(C.ACCESS2_COLOR_ATTACHMENT_READ)

	// Access2ColorAttachmentWrite specifies write access to a color attachment during a
	// render pass or via certain render pass load, store, and multisample resolve operations.
	// This includes multisample resolve operations for depth/stencil resolve attachments.
	// Such access occurs in the PipelineStage2ColorAttachmentOutput pipeline stage.
	Access2ColorAttachmentWrite = AccessFlags2(C.ACCESS2_COLOR_ATTACHMENT_WRITE)

	// Access2DepthStencilAttachmentRead specifies read access to a depth/stencil attachment,
	// via depth or stencil operations or certain render pass load operations in the
	// PipelineStage2EarlyFragmentTests or PipelineStage2LateFragmentTests pipeline stages
	// or via fragment shader tile image reads in the PipelineStage2FragmentShader pipeline stage.
	Access2DepthStencilAttachmentRead = AccessFlags2(C.ACCESS2_DEPTH_STENCIL_ATTACHMENT_READ)

	// Access2DepthStencilAttachmentWrite specifies write access to a depth/stencil
	// attachment, via depth or stencil operations or certain render pass load and store
	// operations. Such access occurs in the PipelineStage2EarlyFragmentTests or
	// PipelineStage2LateFragmentTests pipeline stages.
	Access2DepthStencilAttachmentWrite = AccessFlags2(C.ACCESS2_DEPTH_STENCIL_ATTACHMENT_WRITE)

	// Access2TransferRead specifies read access to an image or buffer in a copy operation.
	// Such access occurs in the PipelineStage2Copy, PipelineStage2Blit, or PipelineStage2Resolve
	// pipeline stages.
	Access2TransferRead = AccessFlags2(C.ACCESS2_TRANSFER_READ)

	// Access2TransferWrite specifies write access to an image or buffer in a clear or
	// copy operation. Such access occurs in the PipelineStage2Copy, PipelineStage2Blit,
	// PipelineStage2Clear, or PipelineStage2Resolve pipeline stages.
	Access2TransferWrite = AccessFlags2(C.ACCESS2_TRANSFER_WRITE)

	// Access2HostRead specifies read access by a host operation. Accesses of this type
	// are not performed through a resource, but directly on memory. Such access occurs
	// in the PipelineStage2Host pipeline stage.
	Access2HostRead = AccessFlags2(C.ACCESS2_HOST_READ)

	// Access2HostWrite specifies write access by a host operation. Accesses of this type
	// are not performed through a resource, but directly on memory. Such access occurs in
	// the PipelineStage2Host pipeline stage.
	Access2HostWrite = AccessFlags2(C.ACCESS2_HOST_WRITE)

	// Access2MemoryRead specifies all read accesses. It is always valid in any access
	// mask, and is treated as equivalent to setting all READ access flags that are valid
	// where it is used.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2MemoryRead = AccessFlags2(C.ACCESS2_MEMORY_READ)

	// Access2MemoryWrite specifies all write accesses. It is always valid in any access
	// mask, and is treated as equivalent to setting all WRITE access flags that are valid
	// where it is used.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2MemoryWrite = AccessFlags2(C.ACCESS2_MEMORY_WRITE)

	// Access2ShaderSampledRead specifies read access to a uniform texel buffer or sampled
	// image in any shader pipeline stage.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2ShaderSampledRead = AccessFlags2(C.ACCESS2_SHADER_SAMPLED_READ)

	// Access2ShaderStorageRead specifies read access to a storage buffer, physical
	// storage buffer, storage texel buffer, or storage image in any shader pipeline stage.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2ShaderStorageRead = AccessFlags2(C.ACCESS2_SHADER_STORAGE_READ)

	// Access2ShaderStorageWrite specifies write access to a storage buffer, physical
	// storage buffer, storage texel buffer, or storage image in any shader pipeline stage.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2ShaderStorageWrite = AccessFlags2(C.ACCESS2_SHADER_STORAGE_WRITE)

	// Access2TransformFeedbackWrite specifies write access to a transform feedback buffer
	// made when transform feedback is active. Such access occurs in the
	// PipelineStage2TransformFeedback pipeline stage.
	//
	// Requires ext_transform_feedback.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2TransformFeedbackWrite = AccessFlags2(C.ACCESS2_TRANSFORM_FEEDBACK_WRITE)

	// Access2TransformFeedbackCounterRead specifies read access to a transform feedback
	// counter buffer which is read when CmdBeginTransformFeedback executes. Such access
	// occurs in the PipelineStage2TransformFeedback pipeline stage.
	//
	// Requires ext_transform_feedback.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2TransformFeedbackCounterRead = AccessFlags2(C.ACCESS2_TRANSFORM_FEEDBACK_COUNTER_READ)

	// Access2TransformFeedbackCounterWrite specifies write access to a transform feedback
	// counter buffer which is written when CmdEndTransformFeedback executes. Such access
	// occurs in the PipelineStage2TransformFeedback pipeline stage.
	//
	// Requires ext_transform_feedback.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2TransformFeedbackCounterWrite = AccessFlags2(C.ACCESS2_TRANSFORM_FEEDBACK_COUNTER_WRITE)

	// Access2ConditionalRenderingRead specifies read access to a predicate as part of
	// conditional rendering. Such access occurs in the PipelineStage2ConditionalRendering pipeline
	// stage.
	//
	// Requires ext_conditional_rendering.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2ConditionalRenderingRead = AccessFlags2(C.ACCESS2_CONDITIONAL_RENDERING_READ)

	// Access2CommandPreprocessReadNV specifies reads from buffer inputs to
	// CmdPreprocessGeneratedCommandsNV. Such access occurs in the
	// PipelineStage2CommandPreprocessNV pipeline stage.
	//
	// Requires nv_device_generated_commands.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2CommandPreprocessReadNV = AccessFlags2(C.ACCESS2_COMMAND_PREPROCESS_READ_NV)

	// Access2CommandPreprocessWriteNV specifies writes to the target command buffer
	// preprocess outputs. Such access occurs in the PipelineStage2CommandPreprocessNV pipeline
	// stage.
	//
	// Requires nv_device_generated_commands.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2CommandPreprocessWriteNV = AccessFlags2(C.ACCESS2_COMMAND_PREPROCESS_WRITE_NV)

	// Access2CommandPreprocessRead specifies reads from buffer inputs to
	// CmdPreprocessGeneratedCommands. Such access occurs in the PipelineStage2CommandPreprocess
	// pipeline stage.
	//
	// Requires ext_device_generated_commands.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2CommandPreprocessRead = AccessFlags2(C.ACCESS2_COMMAND_PREPROCESS_READ)

	// Access2CommandPreprocessWrite specifies writes to the target command buffer preprocess
	// outputs. Such access occurs in the PipelineStage2CommandPreprocess pipeline stage.
	//
	// Requires ext_device_generated_commands.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2CommandPreprocessWrite = AccessFlags2(C.ACCESS2_COMMAND_PREPROCESS_WRITE)

	// Access2FragmentShadingRateAttachmentRead specifies read access to a fragment shading
	// rate attachment during rasterization. Such access occurs in the
	// PipelineStage2FragmentShadingRateAttachment pipeline stage.
	//
	// Requires khr_fragment_shading_rate
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2FragmentShadingRateAttachmentRead = AccessFlags2(C.ACCESS2_FRAGMENT_SHADING_RATE_ATTACHMENT_READ)

	// Access2ShadingRateImageReadNV specifies read access to a shading rate image during
	// rasterization. Such access occurs in the PipelineStage2ShadingRateImageNV pipeline
	// stage. It is equivalent to Access2FragmentShadingRateAttachmentRead.
	//
	// Requires nv_shading_rate_image
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2ShadingRateImageReadNV = AccessFlags2(C.ACCESS2_SHADING_RATE_IMAGE_READ_NV)

	// Access2AccelerationStructureRead specifies read access to an acceleration structure as
	// part of a trace, build, or copy command, or to an acceleration structure scratch buffer
	// as part of a build command. Such access occurs in the PipelineStage2RayTracingShader
	// pipeline stage or PipelineStage2AccelerationStructureBuild pipeline stage.
	//
	// Requires khr_acceleration_structure.
	//
	// https://docs.vulkan.org/spec/latest/chapters/synchronization.html#VkAccessFlagBits2KHR
	Access2AccelerationStructureRead = AccessFlags2(C.ACCESS2_ACCELERATION_STRUCTURE_READ)

	// Access2AccelerationStructureWrite specifies write access to an acceleration structure
	// or acceleration structure scratch buffer as part of a build or copy command. Such access
	// occurs in the PipelineStage2AccelerationStructure pipeline stage.
	//
	// Requires khr_acceleration_structure.
	//
	// https://docs.vulkan.org/spec/latest/chapters/synchronization.html#VkAccessFlagBits2KHR
	Access2AccelerationStructureWrite = AccessFlags2(C.ACCESS2_ACCELERATION_STRUCTURE_WRITE)

	// Access2AccelerationStructureReadNV specifies read access to an acceleration structure
	// as part of a trace, build, or copy command, or to an acceleration structure scratch
	// buffer as part of a build command. Such access occurs in the PipelineStage2RayTracingShader
	// pipeline stage or PipelineStage2AccelerationStructureBuild pipeline stage.
	//
	// Requires nv_ray_tracing.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2AccelerationStructureReadNV = AccessFlags2(C.ACCESS2_ACCELERATION_STRUCTURE_READ_NV)

	// Access2AccelerationStructureWriteNV specifies write access to an acceleration structure
	// or acceleration structure scratch buffer as part of a build or copy command. Such access
	// occurs in the PipelineStage2AccelerationStructureBuild pipeline stage.
	//
	// Requires nv_ray_tracing.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2AccelerationStructureWriteNV = AccessFlags2(C.ACCESS2_ACCELERATION_STRUCTURE_WRITE_NV)

	// Access2FragmentDensityMapRead specifies read access to a fragment density map attachment
	// during dynamic fragment density map operations. Such access occurs in the
	// PipelineStage2FragmentDensityProcess pipeline stage.
	//
	// Requires ext_fragment_density_map
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2FragmentDensityMapRead = AccessFlags2(C.ACCESS2_FRAGMENT_DENSITY_MAP_READ)

	// Access2ColorAttachmentReadNoncoherent specifies read access to color attachments,
	// including advanced blend operations. Such access occurs in the
	// PipelineStage2ColorAttachmentOutput pipeline stage.
	//
	// Requires ext_blend_operation_advanced.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkAccessFlagBits2.html
	Access2ColorAttachmentReadNoncoherent = AccessFlags2(C.ACCESS2_COLOR_ATTACHMENT_READ_NONCOHERENT)
)

var (
	// PipelineStage2None specifies no accesses.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2None = PipelineStageFlags2(C.PIPELINE2_NONE)

	// PipelineStage2TopOfPipe is equivalent to PipelineStage2AllCommands with AccessFlags2
	// set to 0 when specified in the second synchronization scope, but equivalent to
	// PipelineStage2None in the first scope.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2TopOfPipe = PipelineStageFlags2(C.PIPELINE2_TOP_OF_PIPE)

	// PipelineStage2DrawIndirect specifies the stage of the pipeline where indirect command
	// parameters are consumed. This stage also includes reading commands written by
	// CmdPreprocessGeneratedCommands. This stage also includes reading commands written
	// by CmdPreprocessGeneratedCommands.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2DrawIndirect = PipelineStageFlags2(C.PIPELINE2_DRAW_INDIRECT)

	// PipelineStage2VertexInput is equivalent to
	// (PipelineStage2IndexInput|PipelineStage2VertexAttributeInput)
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2VertexInput = PipelineStageFlags2(C.PIPELINE2_VERTEX_INPUT)

	// PipelineStage2VertexShader specifies the vertex shader stage
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2VertexShader = PipelineStageFlags2(C.PIPELINE2_VERTEX_SHADER)

	// PipelineStage2TessellationControlShader specifies the tessellation control shader stage.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2TessellationControlShader = PipelineStageFlags2(C.PIPELINE2_TESSELLATION_CONTROL_SHADER)

	// PipelineStage2TessellationEvaluationShader specifies the tessellation evaluation shader stage.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2TessellationEvaluationShader = PipelineStageFlags2(C.PIPELINE2_TESSELLATION_EVALUATION_SHADER)

	// PipelineStage2GeometryShader specifies the geometry shader stage.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2GeometryShader = PipelineStageFlags2(C.PIPELINE2_GEOMETRY_SHADER)

	// PipelineStage2FragmentShader specifies the fragment shader stage.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2FragmentShader = PipelineStageFlags2(C.PIPELINE2_FRAGMENT_SHADER)

	// PipelineStage2EarlyFragmentTests specifies the stage of the pipeline where early
	// fragment tests (depth and stencil tests before fragment shading) are performed. This
	// stage also includes render pass load operations for framebuffer attachments with a
	// depth/stencil format.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2EarlyFragmentTests = PipelineStageFlags2(C.PIPELINE2_EARLY_FRAGMENT_TESTS)

	// PipelineStage2LateFragmentTests specifies the stage of the pipeline where late
	// fragment tests (depth and stencil tests after fragment shading) are performed. This
	// stage also includes render pass store operations for framebuffer attachments with a
	// depth/stencil format.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2LateFragmentTests = PipelineStageFlags2(C.PIPELINE2_LATE_FRAGMENT_TESTS)

	// PipelineStage2ColorAttachmentOutput specifies the stage of the pipeline where final
	// color values are output from the pipeline. This stage includes blending, logic operations,
	// render pass load and store operations for color attachments, render pass multisample
	// resolve operations, and CmdClearAttachments.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2ColorAttachmentOutput = PipelineStageFlags2(C.PIPELINE2_COLOR_ATTACHMENT_OUTPUT)

	// PipelineStage2ComputeShader specifies the compute shader stage.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2ComputeShader = PipelineStageFlags2(C.PIPELINE2_COMPUTE_SHADER)

	// PipelineStage2AllTransfer is equivalent to (PipelineStage2Copy|PipelineStage2Blit|
	// PipelineStage2Resolve|PipelineStage2Clear|PipelineStage2AccelerationStructureCopy)
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2AllTransfer = PipelineStageFlags2(C.PIPELINE2_ALL_TRANSFER)

	// PipelineStage2BottomOfPipe is equivalent to PipelineStage2AllCommands with AccessFlags2
	// set to 0 when specified in the first synchronization scope, but equivalent to
	// PipelineStage2None in the second scope.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2BottomOfPipe = PipelineStageFlags2(C.PIPELINE2_BOTTOM_OF_PIPE)

	// PipelineStage2Host specifies a pseudo-stage indicating execution on the host of
	// reads/writes of device memory. This stage is not invoked by any commands recorded in
	// a command buffer.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2Host = PipelineStageFlags2(C.PIPELINE2_HOST)

	// PipelineStage2AllGraphics specifies the execution of all graphics pipeline stages,
	// and is equivalent to (PipelineStage2DrawIndirect|PipelineStage2TaskShader|
	// PipelineStage2MeshShader|PipelineStage2VertexInput|PipelineStage2VertexShader|
	// PipelineStage2TessellationControlShader|PipelineStage2TessellationEvaluationShader|
	// PipelineStage2GeometryShader|PipelineStage2FragmentShader|PipelineStage2EarlyFragmentTests|
	// PipelineStage2LateFragmentTests|PipelineStage2ColorAttachmentOutput|PipelineStage2ConditionalRendering|
	// PipelineStage2TransformFeedback|PipelineStage2FragmentShadingRateAttachment|
	// PipelineStage2FragmentDensityProcess)
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2AllGraphics = PipelineStageFlags2(C.PIPELINE2_ALL_GRAPHICS)

	// PipelineStage2AllCommands specifies all operations performed by all commands supported on the queue
	// it is used with.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2AllCommands = PipelineStageFlags2(C.PIPELINE2_ALL_COMMANDS)

	// PipelineStage2Copy specifies the execution of all copy commands, including
	// CmdCopyQueryPoolResults, with the exception of CmdBlitImage2 and CmdBlitImage.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2Copy = PipelineStageFlags2(C.PIPELINE2_COPY)

	// PipelineStage2Resolve specifies the execution of CmdResolveImage2 and CmdResolveImage.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2Resolve = PipelineStageFlags2(C.PIPELINE2_RESOLVE)

	// PipelineStage2Blit specifies the execution of CmdBlitImage2 and CmdBlitImage.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2Blit = PipelineStageFlags2(C.PIPELINE2_BLIT)

	// PipelineStage2Clear specifies the execution of clear commands, with the exception of
	// CmdClearAttachments.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2Clear = PipelineStageFlags2(C.PIPELINE2_CLEAR)

	// PipelineStage2IndexInput specifies the stage of the pipeline where index buffers
	// are consumed.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2IndexInput = PipelineStageFlags2(C.PIPELINE2_INDEX_INPUT)

	// PipelineStage2VertexAttributeInput specifies the stage of the pipeline where vertex
	// buffers are consumed.
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2VertexAttributeInput = PipelineStageFlags2(C.PIPELINE2_VERTEX_ATTRIBUTE_INPUT)

	// PipelineStage2PreRasterizationShaders is equivalent to (PipelineStage2VertexShader|
	// PipelineStage2TessellationControlShader|PipelineStage2TessellationEvaluationShader|
	// PipelineStage2GeometryShader|PipelineStage2TaskShader|PipelineStage2MeshShader)
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2PreRasterizationShaders = PipelineStageFlags2(C.PIPELINE2_PRE_RASTERIZATION_SHADERS)

	// PipelineStage2TransformFeedback specifies the stage of the pipeline where vertex attribute
	// output values are written to the transform feedback buffers.
	//
	// Requires ext_transform_feedback
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2TransformFeedback = PipelineStageFlags2(C.PIPELINE2_TRANSFORM_FEEDBACK)

	// PipelineStage2ConditionalRendering specifies the stage of the pipeline where the
	// predicate of conditional rendering is consumed.
	//
	// Requires ext_conditional_rendering
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2ConditionalRendering = PipelineStageFlags2(C.PIPELINE2_CONDITIONAL_RENDERING)

	// Requires nv_device_generated_commands
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2CommandPreprocessNV = PipelineStageFlags2(C.PIPELINE2_COMMAND_PREPROCESS_NV)

	// PipelineStage2CommandPreprocess specifies the stage of the pipeline where device-side
	// generation of commands via CmdPreprocessGeneratedCommandsNV is handled.
	//
	// Requires ext_device_generated_commands
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2CommandPreprocess = PipelineStageFlags2(C.PIPELINE2_COMMAND_PREPROCESS)

	// PipelineStage2FragmentShadingRateAttachment specifies the stage of the pipeline where
	// the fragment shading rate attachment or shading rate image is read to determine the
	// fragment shading rate for portions of a rasterized primitive.
	//
	// Requires khr_fragment_shading_rate
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2FragmentShadingRateAttachment = PipelineStageFlags2(C.PIPELINE2_FRAGMENT_SHADING_RATE_ATTACHMENT)

	// PipelineStage2AccelerationStructureBuild specifies the execution of acceleration structure commands or acceleration structure copy commands.
	//
	// Requires khr_acceleration_structure
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2AccelerationStructureBuild = PipelineStageFlags2(C.PIPELINE2_ACCELERATION_STRUCTURE_BUILD)

	// PipelineStage2RayTracingShader specifies the execution of the ray tracing shader stages.
	//
	// Requires khr_ray_tracing_pipeline
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2RayTracingShader = PipelineStageFlags2(C.PIPELINE2_RAY_TRACING_SHADER)

	// PipelineStage2RayTracingShaderNV specifies the execution of the ray tracing shader stages.
	//
	// Requires nv_ray_tracing
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2RayTracingShaderNV = PipelineStageFlags2(C.PIPELINE2_RAY_TRACING_SHADER_NV)

	// PipelineStage2AccelerationStructureBuildNV specifies the execution of acceleration
	// structure commands or acceleration structure copy commands.
	//
	// Requires nv_ray_tracing
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2AccelerationStructureBuildNV = PipelineStageFlags2(C.PIPELINE2_ACCELERATION_STRUCTURE_BUILD_NV)

	// PipelineStage2FragmentDensityProcess specifies the stage of the pipeline where the
	// fragment density map is read to generate the fragment areas.
	//
	// Requires ext_fragment_density_map
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2FragmentDensityProcess = PipelineStageFlags2(C.PIPELINE2_FRAGMENT_DENSITY_PROCESS)

	// PipelineStage2TaskShaderNV specifies the task shader stage
	//
	// Requires nv_mesh_shader
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2TaskShaderNV = PipelineStageFlags2(C.PIPELINE2_TASK_SHADER_NV)

	// PipelineStage2MeshShaderNV specifies the mesh shader stage
	//
	// Requires nv_mesh_shader
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2MeshShaderNV = PipelineStageFlags2(C.PIPELINE2_MESH_SHADER_NV)

	// PipelineStage2TaskShader specifies the task shader stage
	//
	// Requires ext_mesh_shader
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2TaskShader = PipelineStageFlags2(C.PIPELINE2_TASK_SHADER)

	// PipelineStage2MeshShader specifies the mesh shader stage
	//
	// Requires ext_mesh_shader
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VkPipelineStageFlagBits2.html
	PipelineStage2MeshShader = PipelineStageFlags2(C.PIPELINE2_MESH_SHADER)
)

func init() {
	AccessNone.Register("None")

	EventCreateDeviceOnly.Register("DeviceOnly")

	ImageLayoutAttachmentOptimal.Register("AttachmentOptimal")
	ImageLayoutReadOnlyOptimal.Register("ReadOnlyOptimal")

	PipelineStageNone.Register("None")

	SubmitProtected.Register("Protected")

	Access2None.Register("None")
	Access2IndirectCommandRead.Register("IndirectCommandRead")
	Access2IndexRead.Register("IndexRead")
	Access2VertexAttributeRead.Register("VertexAttributeRead")
	Access2UniformRead.Register("UniformRead")
	Access2InputAttachmentRead.Register("InputAttachmentRead")
	Access2ShaderRead.Register("ShaderRead")
	Access2ShaderWrite.Register("ShaderWrite")
	Access2ColorAttachmentRead.Register("ColorAttachmentRead")
	Access2ColorAttachmentWrite.Register("ColorAttachmentWrite")
	Access2DepthStencilAttachmentRead.Register("DepthStencilAttachmentRead")
	Access2DepthStencilAttachmentWrite.Register("DepthStencilAttachmentWrite")
	Access2TransferRead.Register("TransferRead")
	Access2TransferWrite.Register("TransferWrite")
	Access2HostRead.Register("HostRead")
	Access2HostWrite.Register("HostWrite")
	Access2MemoryRead.Register("MemoryRead")
	Access2MemoryWrite.Register("MemoryWrite")
	Access2ShaderSampledRead.Register("ShaderSampledRead")
	Access2ShaderStorageRead.Register("ShaderStorageRead")
	Access2ShaderStorageWrite.Register("ShaderStorageWrite")
	Access2TransformFeedbackWrite.Register("TransformFeedbackWrite")
	Access2TransformFeedbackCounterRead.Register("TransformFeedbackCounterRead")
	Access2TransformFeedbackCounterWrite.Register("TransformFeedbackCounterWrite")
	Access2ConditionalRenderingRead.Register("ConditionalRenderingRead")
	Access2CommandPreprocessReadNV.Register("CommandPreprocessReadNV")
	Access2CommandPreprocessWriteNV.Register("CommandPreprocessWriteNV")
	Access2CommandPreprocessRead.Register("CommandPreprocessRead")
	Access2CommandPreprocessWrite.Register("CommandPreprocessWrite")
	Access2FragmentShadingRateAttachmentRead.Register("FragmentShadingRateAttachmentRead")
	Access2ShadingRateImageReadNV.Register("ShadingRateImageReadNV")
	Access2AccelerationStructureRead.Register("AccelerationStructureRead")
	Access2AccelerationStructureWrite.Register("AccelerationStructureWrite")
	Access2AccelerationStructureReadNV.Register("AccelerationStructureReadNV")
	Access2AccelerationStructureWriteNV.Register("AccelerationStructureWriteNV")
	Access2FragmentDensityMapRead.Register("FragmentDensityMapRead")
	Access2ColorAttachmentReadNoncoherent.Register("ColorAttachmentReadNoncoherent")

	PipelineStage2None.Register("None")
	PipelineStage2TopOfPipe.Register("TopOfPipe")
	PipelineStage2DrawIndirect.Register("DrawIndirect")
	PipelineStage2VertexInput.Register("VertexInput")
	PipelineStage2VertexShader.Register("VertexShader")
	PipelineStage2TessellationControlShader.Register("TessellationControlShader")
	PipelineStage2TessellationEvaluationShader.Register("TessellationEvaluationShader")
	PipelineStage2GeometryShader.Register("GeometryShader")
	PipelineStage2FragmentShader.Register("FragmentShader")
	PipelineStage2EarlyFragmentTests.Register("EarlyFragmentTests")
	PipelineStage2LateFragmentTests.Register("LateFragmentTests")
	PipelineStage2ColorAttachmentOutput.Register("ColorAttachmentOutput")
	PipelineStage2ComputeShader.Register("ComputeShader")
	PipelineStage2AllTransfer.Register("AllTransfer")
	PipelineStage2BottomOfPipe.Register("BottomOfPipe")
	PipelineStage2Host.Register("Host")
	PipelineStage2AllGraphics.Register("AllGraphics")
	PipelineStage2AllCommands.Register("AllCommands")
	PipelineStage2Copy.Register("Copy")
	PipelineStage2Resolve.Register("Resolve")
	PipelineStage2Blit.Register("Blit")
	PipelineStage2Clear.Register("Clear")
	PipelineStage2IndexInput.Register("IndexInput")
	PipelineStage2VertexAttributeInput.Register("VertexAttributeInput")
	PipelineStage2PreRasterizationShaders.Register("PreRasterizationShaders")
	PipelineStage2TransformFeedback.Register("TransformFeedback")
	PipelineStage2ConditionalRendering.Register("ConditionalRendering")
	PipelineStage2CommandPreprocess.Register("CommandPreprocess")
	PipelineStage2CommandPreprocessNV.Register("CommandPreprocessNV")
	PipelineStage2FragmentShadingRateAttachment.Register("FragmentShadingRateAttachment")
	PipelineStage2AccelerationStructureBuild.Register("AccelerationStructureBuild")
	PipelineStage2RayTracingShader.Register("RayTracingShader")
	PipelineStage2RayTracingShaderNV.Register("RayTracingShaderNV")
	PipelineStage2AccelerationStructureBuildNV.Register("AccelerationStructureBuildNV")
	PipelineStage2FragmentDensityProcess.Register("FragmentDensityProcess")
	PipelineStage2TaskShaderNV.Register("TaskShaderNV")
	PipelineStage2MeshShaderNV.Register("MeshShaderNV")
	PipelineStage2TaskShader.Register("TaskShader")
	PipelineStage2MeshShader.Register("MeshShader")
}
