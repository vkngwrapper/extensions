package ext_descriptor_indexing

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"unsafe"
)

// PhysicalDeviceDescriptorIndexingProperties describes descriptor indexing properties
// that can be supported by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingProperties.html
type PhysicalDeviceDescriptorIndexingProperties struct {
	// MaxUpdateAfterBindDescriptorsInAllPools is the maximum number of descriptors (summed over
	// all descriptor types) that can be created across all pools that are created with
	// DescriptorPoolCreateUpdateAfterBind
	MaxUpdateAfterBindDescriptorsInAllPools int
	// ShaderUniformBufferArrayNonUniformIndexingNative is a boolean value indicating whether
	// uniform Buffer descriptors natively support nonuniform indexing
	ShaderUniformBufferArrayNonUniformIndexingNative bool
	// ShaderSampledImageArrayNonUniformIndexingNative is a boolean value indicating whether
	// Sampler and Image descriptors natively support nonuniform indexing
	ShaderSampledImageArrayNonUniformIndexingNative bool
	// ShaderStorageBufferArrayNonUniformIndexingNative is a boolean value indicating whether
	// storage Buffer descriptors natively support nonuniform indexing
	ShaderStorageBufferArrayNonUniformIndexingNative bool
	// ShaderStorageImageArrayNonUniformIndexingNative is a boolean value indicating whether storage
	// Image descriptors natively support nonuniform indexing
	ShaderStorageImageArrayNonUniformIndexingNative bool
	// ShaderInputAttachmentArrayNonUniformIndexingNative is a boolean value indicating whether
	// input attachment descriptors natively support nonuniform indexing
	ShaderInputAttachmentArrayNonUniformIndexingNative bool
	// RobustBufferAccessUpdateAfterBind is a boolean value indicating whether RobustBufferAccess
	// can be enabled in a Device simultaneously with DescriptorBindingUniformBufferUpdateAfterBind,
	// DescriptorBindingStorageBufferUpdateAfterBind,
	// DescriptorBindingUniformTexelBufferUpdateAfterBind, and/or
	// DescriptorBindingStorageTexelBufferUpdateAfterBind
	RobustBufferAccessUpdateAfterBind bool
	// QuadDivergentImplicitLod is a boolean value indicating whether implicit level of detail
	// calculations for Image operations have well-defined results when the Image and/or Sampler
	// objects used for the instruction are not uniform within a quad
	QuadDivergentImplicitLod bool

	// MaxPerStageDescriptorUpdateAfterBindSamplers is similar to <axPerStageDescriptorSamplers
	// but counts descriptors from descriptor sets created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageDescriptorUpdateAfterBindSamplers int
	// MaxPerStageDescriptorUpdateAfterBindUniformBuffers is similar to
	// MaxPerStageDescriptorUniformBuffers but counts descriptors from DescriptorSet objects
	// created with or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageDescriptorUpdateAfterBindUniformBuffers int
	// MaxPerStageDescriptorUpdateAfterBindStorageBuffers is similar to
	// MaxPerStageDescriptorStorageBuffers but counts descriptors from DescriptorSet created with
	// or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageDescriptorUpdateAfterBindStorageBuffers int
	// MaxPerStageDescriptorUpdateAfterBindSampledImages is similar to
	// MaxPerStageDescriptorSampledImages but counts descriptors from DescriptorSets created with
	// or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageDescriptorUpdateAfterBindSampledImages int
	// MaxPerStageDescriptorUpdateAfterBindStorageImages is similar to
	// MaxPerStageDescriptorStorageImages but counts descriptors from DescriptorSet objects created
	// with or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageDescriptorUpdateAfterBindStorageImages int
	// MaxPerStageDescriptorUpdateAfterBindInputAttachments  is similar to
	// MaxPerStageDescriptorInputAttachments but counts descriptors from DescriptorSet objects
	// created with or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageDescriptorUpdateAfterBindInputAttachments int
	// MaxPerStageUpdateAfterBindResources is similar to MaxPerStageResources but counts
	// descriptors from DescriptorSet objects created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxPerStageUpdateAfterBindResources int

	// MaxDescriptorSetUpdateAfterBindSamplers is similar to MaxDescriptorSetSamplers but counts
	// descriptors from DescriptorSet created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindSamplers int
	// MaxDescriptorSetUpdateAfterBindUniformBuffers is similar to MaxDescriptorSetUniformBuffers
	// but counts descriptors from DescriptorSet objects created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindUniformBuffers int
	// MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic is similar to
	// MaxDescriptorSetUniformBuffersDynamic but counts descriptors from DescriptorSet objects
	// created with or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic int
	// MaxDescriptorSetUpdateAfterBindStorageBuffers is similar to MaxDescriptorSetStorageBuffers
	// but counts descriptors from DescriptorSet objects created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindStorageBuffers int
	// MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic is similar to
	// MaxDescriptorSetStorageBuffersDynamic but counts descriptors from DescriptorSet objects
	// created with or without DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic int
	// MaxDescriptorSetUpdateAfterBindSampledImages is similar to MaxDescriptorSetSampledImages
	// but counts descriptors from DescriptorSet objects created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindSampledImages int
	// MaxDescriptorSetUpdateAfterBindStorageImages is similar to MaxDescriptorSetStorageImages
	// but counts descriptors from DescriptorSet objects created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindStorageImages int
	// MaxDescriptorSetUpdateAfterBindInputAttachments is similar to MaxDescriptorSetInputAttachments
	// but counts descriptors from DescriptorSet objects created with or without
	// DescriptorSetLayoutCreateUpdateAfterBindPool
	MaxDescriptorSetUpdateAfterBindInputAttachments int

	common.NextOutData
}

func (o *PhysicalDeviceDescriptorIndexingProperties) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceDescriptorIndexingPropertiesEXT{})))
	}

	info := (*C.VkPhysicalDeviceDescriptorIndexingPropertiesEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES_EXT
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceDescriptorIndexingProperties) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceDescriptorIndexingPropertiesEXT)(cDataPointer)

	o.MaxUpdateAfterBindDescriptorsInAllPools = int(info.maxUpdateAfterBindDescriptorsInAllPools)
	o.ShaderUniformBufferArrayNonUniformIndexingNative = info.shaderUniformBufferArrayNonUniformIndexingNative != C.VkBool32(0)
	o.ShaderSampledImageArrayNonUniformIndexingNative = info.shaderSampledImageArrayNonUniformIndexingNative != C.VkBool32(0)
	o.ShaderStorageBufferArrayNonUniformIndexingNative = info.shaderStorageBufferArrayNonUniformIndexingNative != C.VkBool32(0)
	o.ShaderStorageImageArrayNonUniformIndexingNative = info.shaderStorageImageArrayNonUniformIndexingNative != C.VkBool32(0)
	o.ShaderInputAttachmentArrayNonUniformIndexingNative = info.shaderInputAttachmentArrayNonUniformIndexingNative != C.VkBool32(0)
	o.RobustBufferAccessUpdateAfterBind = info.robustBufferAccessUpdateAfterBind != C.VkBool32(0)
	o.QuadDivergentImplicitLod = info.quadDivergentImplicitLod != C.VkBool32(0)

	o.MaxPerStageDescriptorUpdateAfterBindSamplers = int(info.maxPerStageDescriptorUpdateAfterBindSamplers)
	o.MaxPerStageDescriptorUpdateAfterBindUniformBuffers = int(info.maxPerStageDescriptorUpdateAfterBindUniformBuffers)
	o.MaxPerStageDescriptorUpdateAfterBindStorageBuffers = int(info.maxPerStageDescriptorUpdateAfterBindStorageBuffers)
	o.MaxPerStageDescriptorUpdateAfterBindSampledImages = int(info.maxPerStageDescriptorUpdateAfterBindSampledImages)
	o.MaxPerStageDescriptorUpdateAfterBindStorageImages = int(info.maxPerStageDescriptorUpdateAfterBindStorageImages)
	o.MaxPerStageDescriptorUpdateAfterBindInputAttachments = int(info.maxPerStageDescriptorUpdateAfterBindInputAttachments)
	o.MaxPerStageUpdateAfterBindResources = int(info.maxPerStageUpdateAfterBindResources)

	o.MaxDescriptorSetUpdateAfterBindSamplers = int(info.maxDescriptorSetUpdateAfterBindSamplers)
	o.MaxDescriptorSetUpdateAfterBindUniformBuffers = int(info.maxDescriptorSetUpdateAfterBindUniformBuffers)
	o.MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic = int(info.maxDescriptorSetUpdateAfterBindUniformBuffersDynamic)
	o.MaxDescriptorSetUpdateAfterBindStorageBuffers = int(info.maxDescriptorSetUpdateAfterBindStorageBuffers)
	o.MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic = int(info.maxDescriptorSetUpdateAfterBindStorageBuffersDynamic)
	o.MaxDescriptorSetUpdateAfterBindSampledImages = int(info.maxDescriptorSetUpdateAfterBindSampledImages)
	o.MaxDescriptorSetUpdateAfterBindStorageImages = int(info.maxDescriptorSetUpdateAfterBindStorageImages)
	o.MaxDescriptorSetUpdateAfterBindInputAttachments = int(info.maxDescriptorSetUpdateAfterBindInputAttachments)

	return info.pNext, nil
}
