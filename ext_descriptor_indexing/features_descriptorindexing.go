package ext_descriptor_indexing

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"unsafe"
)

// PhysicalDeviceDescriptorIndexingFeatures describes descriptor indexing
// features that can be supported by an implementation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html
type PhysicalDeviceDescriptorIndexingFeatures struct {
	// ShaderInputAttachmentArrayDynamicIndexing indicates whether arrays of input attachments
	// can be indexed by dynamically uniform integer expressions in shader code
	ShaderInputAttachmentArrayDynamicIndexing bool
	// ShaderUniformTexelBufferArrayDynamicIndexing indicates whether arrays of uniform texel
	// Buffer objects can be indexed by dynamically uniform integer expressions in shader code
	ShaderUniformTexelBufferArrayDynamicIndexing bool
	// ShaderStorageTexelBufferArrayDynamicIndexing indicates whether arrays of storage texel
	// Buffer objects can be indexed by dynamically uniform integer expressions in shader code
	ShaderStorageTexelBufferArrayDynamicIndexing bool
	// ShaderUniformBufferArrayNonUniformIndexing indicates whether arrays of uniform Buffer objects
	// can be indexed by non-uniform integer expressions in shader code.
	ShaderUniformBufferArrayNonUniformIndexing bool
	// ShaderSampledImageArrayNonUniformIndexing indicates whether arrays of Sampler objects or sampled
	// Image objects can be indexed by non-uniform integer expressions in shader code
	ShaderSampledImageArrayNonUniformIndexing bool
	// ShaderStorageBufferArrayNonUniformIndexing indicates whether arrays of storage buffers
	// can be indexed by non-uniform integer expressions in shader code
	ShaderStorageBufferArrayNonUniformIndexing bool
	// ShaderStorageImageArrayNonUniformIndexing indicates whether arrays of storage Image objects can
	// be indexed by non-uniform integer expressions in shader code
	ShaderStorageImageArrayNonUniformIndexing bool
	// ShaderInputAttachmentArrayNonUniformIndexing indicates whether arrays of input attachments
	// can be indexed by non-uniform integer expressions in shader code
	ShaderInputAttachmentArrayNonUniformIndexing bool
	// ShaderUniformTexelBufferArrayNonUniformIndexing indicates whether arrays of uniform texel
	// Buffer objects can be indexed by non-uniform integer expressions in shader code
	ShaderUniformTexelBufferArrayNonUniformIndexing bool
	// ShaderStorageTexelBufferArrayNonUniformIndexing indicates whether arrays of storage texel
	// Buffer objects can be indexed by non-uniform integer expressions in shader code
	ShaderStorageTexelBufferArrayNonUniformIndexing bool
	// DescriptorBindingUniformBufferUpdateAfterBind indicates whether the implementation supports
	// updating uniform Buffer descriptors after a set is bound
	DescriptorBindingUniformBufferUpdateAfterBind bool
	// DescriptorBindingSampledImageUpdateAfterBind indicates whether the implementation supports
	// updating sampled Image descriptors after a set is bound
	DescriptorBindingSampledImageUpdateAfterBind bool
	// DescriptorBindingStorageImageUpdateAfterBind indicates whether the implementation supports
	// updating storage Image descriptors after a set is bound
	DescriptorBindingStorageImageUpdateAfterBind bool
	// DescriptorBindingStorageBufferUpdateAfterBind indicates whether the implementation
	// supports updating storage Buffer descriptors after a set is bound
	DescriptorBindingStorageBufferUpdateAfterBind bool
	// DescriptorBindingUniformTexelBufferUpdateAfterBind indicates whether the implementation
	// supports updating uniform texel Buffer descriptors after a set is bound
	DescriptorBindingUniformTexelBufferUpdateAfterBind bool
	// DescriptorBindingStorageTexelBufferUpdateAfterBind indicates whether the impelementation
	// supports updating storage texel Buffer descriptors after a set is bound
	DescriptorBindingStorageTexelBufferUpdateAfterBind bool
	// DescriptorBindingUpdateUnusedWhilePending indicates whether the implementation supports
	// updating descriptors while the set is in use
	DescriptorBindingUpdateUnusedWhilePending bool
	// DescriptorBindingPartiallyBound indicates whether the implementation supports statically
	// using a DescriptorSet binding in which some descriptors are not valid
	DescriptorBindingPartiallyBound bool
	// DescriptorBindingVariableDescriptorCount indicates whether the implementation supports
	// DescriptorSet object with a variable-sized last binding
	DescriptorBindingVariableDescriptorCount bool
	// RuntimeDescriptorArray indicates whether the implementation supports the SPIR-V
	// RuntimeDescriptorArray capability
	RuntimeDescriptorArray bool

	common.NextOptions
	common.NextOutData
}

func (o *PhysicalDeviceDescriptorIndexingFeatures) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceDescriptorIndexingFeaturesEXT{})))
	}

	info := (*C.VkPhysicalDeviceDescriptorIndexingFeaturesEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES_EXT
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceDescriptorIndexingFeatures) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceDescriptorIndexingFeaturesEXT)(cDataPointer)

	o.ShaderInputAttachmentArrayDynamicIndexing = info.shaderInputAttachmentArrayDynamicIndexing != C.VkBool32(0)
	o.ShaderUniformTexelBufferArrayDynamicIndexing = info.shaderUniformTexelBufferArrayDynamicIndexing != C.VkBool32(0)
	o.ShaderStorageTexelBufferArrayDynamicIndexing = info.shaderStorageTexelBufferArrayDynamicIndexing != C.VkBool32(0)
	o.ShaderUniformBufferArrayNonUniformIndexing = info.shaderUniformBufferArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderSampledImageArrayNonUniformIndexing = info.shaderSampledImageArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderStorageBufferArrayNonUniformIndexing = info.shaderStorageBufferArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderStorageImageArrayNonUniformIndexing = info.shaderStorageImageArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderInputAttachmentArrayNonUniformIndexing = info.shaderInputAttachmentArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderUniformTexelBufferArrayNonUniformIndexing = info.shaderUniformTexelBufferArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderStorageTexelBufferArrayNonUniformIndexing = info.shaderStorageTexelBufferArrayNonUniformIndexing != C.VkBool32(0)
	o.DescriptorBindingUniformBufferUpdateAfterBind = info.descriptorBindingUniformBufferUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingSampledImageUpdateAfterBind = info.descriptorBindingSampledImageUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingStorageImageUpdateAfterBind = info.descriptorBindingStorageImageUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingStorageBufferUpdateAfterBind = info.descriptorBindingStorageBufferUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingUniformTexelBufferUpdateAfterBind = info.descriptorBindingUniformTexelBufferUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingStorageTexelBufferUpdateAfterBind = info.descriptorBindingStorageTexelBufferUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingUpdateUnusedWhilePending = info.descriptorBindingUpdateUnusedWhilePending != C.VkBool32(0)
	o.DescriptorBindingPartiallyBound = info.descriptorBindingPartiallyBound != C.VkBool32(0)
	o.DescriptorBindingVariableDescriptorCount = info.descriptorBindingVariableDescriptorCount != C.VkBool32(0)
	o.RuntimeDescriptorArray = info.runtimeDescriptorArray != C.VkBool32(0)

	return info.pNext, nil
}

func (o PhysicalDeviceDescriptorIndexingFeatures) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceDescriptorIndexingFeaturesEXT{})))
	}

	info := (*C.VkPhysicalDeviceDescriptorIndexingFeaturesEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES_EXT
	info.pNext = next
	info.shaderInputAttachmentArrayDynamicIndexing = C.VkBool32(0)
	info.shaderUniformTexelBufferArrayDynamicIndexing = C.VkBool32(0)
	info.shaderStorageTexelBufferArrayDynamicIndexing = C.VkBool32(0)
	info.shaderUniformBufferArrayNonUniformIndexing = C.VkBool32(0)
	info.shaderSampledImageArrayNonUniformIndexing = C.VkBool32(0)
	info.shaderStorageBufferArrayNonUniformIndexing = C.VkBool32(0)
	info.shaderStorageImageArrayNonUniformIndexing = C.VkBool32(0)
	info.shaderInputAttachmentArrayNonUniformIndexing = C.VkBool32(0)
	info.shaderUniformTexelBufferArrayNonUniformIndexing = C.VkBool32(0)
	info.shaderStorageTexelBufferArrayNonUniformIndexing = C.VkBool32(0)
	info.descriptorBindingUniformBufferUpdateAfterBind = C.VkBool32(0)
	info.descriptorBindingSampledImageUpdateAfterBind = C.VkBool32(0)
	info.descriptorBindingStorageImageUpdateAfterBind = C.VkBool32(0)
	info.descriptorBindingStorageBufferUpdateAfterBind = C.VkBool32(0)
	info.descriptorBindingUniformTexelBufferUpdateAfterBind = C.VkBool32(0)
	info.descriptorBindingStorageTexelBufferUpdateAfterBind = C.VkBool32(0)
	info.descriptorBindingUpdateUnusedWhilePending = C.VkBool32(0)
	info.descriptorBindingPartiallyBound = C.VkBool32(0)
	info.descriptorBindingVariableDescriptorCount = C.VkBool32(0)
	info.runtimeDescriptorArray = C.VkBool32(0)

	if o.ShaderInputAttachmentArrayDynamicIndexing {
		info.shaderInputAttachmentArrayDynamicIndexing = C.VkBool32(1)
	}

	if o.ShaderUniformTexelBufferArrayDynamicIndexing {
		info.shaderUniformTexelBufferArrayDynamicIndexing = C.VkBool32(1)
	}

	if o.ShaderStorageTexelBufferArrayDynamicIndexing {
		info.shaderStorageTexelBufferArrayDynamicIndexing = C.VkBool32(1)
	}

	if o.ShaderUniformBufferArrayNonUniformIndexing {
		info.shaderUniformBufferArrayNonUniformIndexing = C.VkBool32(1)
	}

	if o.ShaderSampledImageArrayNonUniformIndexing {
		info.shaderSampledImageArrayNonUniformIndexing = C.VkBool32(1)
	}

	if o.ShaderStorageBufferArrayNonUniformIndexing {
		info.shaderStorageBufferArrayNonUniformIndexing = C.VkBool32(1)
	}

	if o.ShaderStorageImageArrayNonUniformIndexing {
		info.shaderStorageImageArrayNonUniformIndexing = C.VkBool32(1)
	}

	if o.ShaderInputAttachmentArrayNonUniformIndexing {
		info.shaderInputAttachmentArrayNonUniformIndexing = C.VkBool32(1)
	}

	if o.ShaderUniformTexelBufferArrayNonUniformIndexing {
		info.shaderUniformTexelBufferArrayNonUniformIndexing = C.VkBool32(1)
	}

	if o.ShaderStorageTexelBufferArrayNonUniformIndexing {
		info.shaderStorageTexelBufferArrayNonUniformIndexing = C.VkBool32(1)
	}

	if o.DescriptorBindingUniformBufferUpdateAfterBind {
		info.descriptorBindingUniformBufferUpdateAfterBind = C.VkBool32(1)
	}

	if o.DescriptorBindingSampledImageUpdateAfterBind {
		info.descriptorBindingSampledImageUpdateAfterBind = C.VkBool32(1)
	}

	if o.DescriptorBindingStorageImageUpdateAfterBind {
		info.descriptorBindingStorageImageUpdateAfterBind = C.VkBool32(1)
	}

	if o.DescriptorBindingStorageBufferUpdateAfterBind {
		info.descriptorBindingStorageBufferUpdateAfterBind = C.VkBool32(1)
	}

	if o.DescriptorBindingUniformTexelBufferUpdateAfterBind {
		info.descriptorBindingUniformTexelBufferUpdateAfterBind = C.VkBool32(1)
	}

	if o.DescriptorBindingStorageTexelBufferUpdateAfterBind {
		info.descriptorBindingStorageTexelBufferUpdateAfterBind = C.VkBool32(1)
	}

	if o.DescriptorBindingUpdateUnusedWhilePending {
		info.descriptorBindingUpdateUnusedWhilePending = C.VkBool32(1)
	}

	if o.DescriptorBindingPartiallyBound {
		info.descriptorBindingPartiallyBound = C.VkBool32(1)
	}

	if o.DescriptorBindingVariableDescriptorCount {
		info.descriptorBindingVariableDescriptorCount = C.VkBool32(1)
	}

	if o.RuntimeDescriptorArray {
		info.runtimeDescriptorArray = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}
