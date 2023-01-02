package khr_descriptor_update_template

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	khr_descriptor_update_template_driver "github.com/vkngwrapper/extensions/v2/khr_descriptor_update_template/driver"
)

//go:generate mockgen -source iface.go -destination ./mocks/extension.go -package mock_descriptor_update_template

// DescriptorUpdateTemplate specifies a mapping from descriptor update information in host memory to
// descriptors in a DescriptorSet
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplate.html
type DescriptorUpdateTemplate interface {
	// DeviceHandle is the internal Vulkan object handle for the Device this DescriptorUpdateTemplate belongs to
	DeviceHandle() driver.VkDevice
	// Driver is the Vulkan wrapper driver used by this DescriptorUpdateTemplate
	Driver() driver.Driver
	// Handle is the internal Vulkan object handle for this DescriptorUpdateTemplate
	Handle() khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR
	// APIVersion is the maximum Vulkan API version supported by this DescriptorUpdateTemplate. If it is at least
	// Vulkan 1.2, core1_2.PromoteDescriptorUpdateTemplate can be used to promote this to a core1_2.DescriptorUpdateTemplate,
	// etc.
	APIVersion() common.APIVersion
	// Destroy destroys the DescriptorUpdateTemplate object and the underlying structures. **Warning** after
	// destruction, this object will continue to exist, but the Vulkan object handle that backs it will
	// be invalid. Do not call further methods on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkDestroyDescriptorUpdateTemplate.html
	Destroy(allocator *driver.AllocationCallbacks)

	// UpdateDescriptorSetFromImage updates the contents of a DescriptorSet object with this template and
	// an Image
	//
	// descriptorSet - The DescriptorSet to update
	//
	// data - Information and an Image used to write the descriptor
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplateKHR.html
	UpdateDescriptorSetFromImage(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorImageInfo)
	// UpdateDescriptorSetFromBuffer updates the contents of a DescriptorSet object with this template
	// and a Buffer
	//
	// descriptorSet - The DescriptorSet to update
	//
	// data - Information and a Buffer used to write the descriptor
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplateKHR.html
	UpdateDescriptorSetFromBuffer(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorBufferInfo)
	// UpdateDescriptorSetFromObjectHandle updates the contents of a DescriptorSet object with this template
	// and an arbitrary handle
	//
	// descriptorSet - The DescriptorSet to update
	//
	// data - A Vulkan object handle used to write the descriptor. Can be a BufferView handle or
	// perhaps an acceleration structure.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplateKHR.html
	UpdateDescriptorSetFromObjectHandle(descriptorSet core1_0.DescriptorSet, data driver.VulkanHandle)
}

type Extension interface {
	// CreateDescriptorUpdateTemplate creates a new DescriptorUpdateTemplate
	//
	// device - The Device to create DescriptorUpdateTemplate from
	//
	// o - Specifies the set of descriptors to update with a single call to DescriptorUpdateTemplate.UpdateDescriptorSet...
	//
	// allocator - Controls host allocation behavior
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorUpdateTemplate.html
	CreateDescriptorUpdateTemplate(device core1_0.Device, o DescriptorUpdateTemplateCreateInfo, allocator *driver.AllocationCallbacks) (DescriptorUpdateTemplate, common.VkResult, error)
}
