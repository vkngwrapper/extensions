package khr_descriptor_update_template

import (
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
)

//go:generate mockgen -source iface.go -destination ./mocks/extension.go -package mock_descriptor_update_template

type ExtensionDriver interface {
	// CreateDescriptorUpdateTemplate creates a new DescriptorUpdateTemplate
	//
	// device - The Device to create DescriptorUpdateTemplate from
	//
	// o - Specifies the set of descriptors to update with a single call to DescriptorUpdateTemplate.UpdateDescriptorSet...
	//
	// allocator - Controls host allocation behavior
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorUpdateTemplate.html
	CreateDescriptorUpdateTemplate(o DescriptorUpdateTemplateCreateInfo, allocator *loader.AllocationCallbacks) (core.DescriptorUpdateTemplate, common.VkResult, error)
	// DestroyDescriptorUpdateTemplate destroys a DescriptorUpdateTemplate object and the underlying structures. **Warning** after
	// destruction, the object will continue to exist, but the Vulkan object handle that backs it will
	// be invalid. Do not call further methods with the DescriptorUpdateTemplate.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyDescriptorUpdateTemplate.html
	DestroyDescriptorUpdateTemplate(template core.DescriptorUpdateTemplate, allocator *loader.AllocationCallbacks)

	// UpdateDescriptorSetWithTemplateFromImage updates the contents of a DescriptorSet object with this template and
	// an Image
	//
	// descriptorSet - The DescriptorSet to update
	//
	// data - Information and an Image used to write the descriptor
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplateKHR.html
	UpdateDescriptorSetWithTemplateFromImage(descriptorSet core.DescriptorSet, template core.DescriptorUpdateTemplate, data core1_0.DescriptorImageInfo)
	// UpdateDescriptorSetWithTemplateFromBuffer updates the contents of a DescriptorSet object with this template
	// and a Buffer
	//
	// descriptorSet - The DescriptorSet to update
	//
	// data - Information and a Buffer used to write the descriptor
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplateKHR.html
	UpdateDescriptorSetWithTemplateFromBuffer(descriptorSet core.DescriptorSet, template core.DescriptorUpdateTemplate, data core1_0.DescriptorBufferInfo)
	// UpdateDescriptorSetWithTemplateFromObjectHandle updates the contents of a DescriptorSet object with this template
	// and an arbitrary handle
	//
	// descriptorSet - The DescriptorSet to update
	//
	// data - A Vulkan object handle used to write the descriptor. Can be a BufferView handle or
	// perhaps an acceleration structure.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplateKHR.html
	UpdateDescriptorSetWithTemplateFromObjectHandle(descriptorSet core.DescriptorSet, template core.DescriptorUpdateTemplate, data loader.VulkanHandle)
}
