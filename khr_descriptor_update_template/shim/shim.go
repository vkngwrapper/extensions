package khr_descriptor_update_template_shim

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
	"github.com/vkngwrapper/extensions/v2/khr_descriptor_update_template"
)

//go:generate mockgen -source shim.go -destination ../mocks/shim.go -package mock_descriptor_update_template

type Shim interface {
	// CreateDescriptorUpdateTemplate creates a new DescriptorUpdateTemplate
	//
	// o - Specifies the set of descriptors to update with a single call to DescriptorUpdateTemplate.UpdateDescriptorSet...
	//
	// allocator - Controls host allocation behavior
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorUpdateTemplate.html
	CreateDescriptorUpdateTemplate(o core1_1.DescriptorUpdateTemplateCreateInfo, allocator *driver.AllocationCallbacks) (core1_1.DescriptorUpdateTemplate, common.VkResult, error)
}

// Compiler check that VulkanShim satisfies Shim
var _ Shim = &VulkanShim{}

type VulkanShim struct {
	extension khr_descriptor_update_template.Extension
	device    core1_0.Device
}

func NewShim(extension khr_descriptor_update_template.Extension, device core1_0.Device) *VulkanShim {
	if extension == nil {
		panic("extension cannot be nil")
	}
	if device == nil {
		panic("device cannot be nil")
	}
	return &VulkanShim{
		extension: extension,
		device:    device,
	}
}

// VulkanShimDescriptorUpdateTemplate is a wrapper for khr_descriptor_update_template.DescriptorUpdateTemplate that
// converts the Handle() method to return the core 1.1 handle type, so that the extension DescriptorUpdateTemplate
// can be made compatible with the core 1.1 DescriptorUpdateTemplate for the shim
type VulkanShimDescriptorUpdateTemplate struct {
	khr_descriptor_update_template.DescriptorUpdateTemplate
}

func (t *VulkanShimDescriptorUpdateTemplate) Handle() driver.VkDescriptorUpdateTemplate {
	return driver.VkDescriptorUpdateTemplate(t.DescriptorUpdateTemplate.Handle())
}

func (s *VulkanShim) CreateDescriptorUpdateTemplate(o core1_1.DescriptorUpdateTemplateCreateInfo, allocator *driver.AllocationCallbacks) (core1_1.DescriptorUpdateTemplate, common.VkResult, error) {
	inOptions := khr_descriptor_update_template.DescriptorUpdateTemplateCreateInfo{
		Flags:                   khr_descriptor_update_template.DescriptorUpdateTemplateCreateFlags(o.Flags),
		DescriptorUpdateEntries: make([]khr_descriptor_update_template.DescriptorUpdateTemplateEntry, 0, len(o.DescriptorUpdateEntries)),
		TemplateType:            khr_descriptor_update_template.DescriptorUpdateTemplateType(o.TemplateType),
		DescriptorSetLayout:     o.DescriptorSetLayout,
		PipelineBindPoint:       o.PipelineBindPoint,
		PipelineLayout:          o.PipelineLayout,
		Set:                     o.Set,
		NextOptions:             o.NextOptions,
	}

	for _, entry := range o.DescriptorUpdateEntries {
		inOptions.DescriptorUpdateEntries = append(inOptions.DescriptorUpdateEntries, khr_descriptor_update_template.DescriptorUpdateTemplateEntry(entry))
	}

	template, res, err := s.extension.CreateDescriptorUpdateTemplate(
		s.device,
		inOptions,
		allocator,
	)

	var outTemplate *VulkanShimDescriptorUpdateTemplate
	if template != nil {
		outTemplate = &VulkanShimDescriptorUpdateTemplate{template}
	}
	return outTemplate, res, err
}
