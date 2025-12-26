package khr_descriptor_update_template

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
)

// DescriptorUpdateTemplateCreateInfo specifies parameters of a newly-created Descriptor Update
// Template
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateCreateInfo.html
type DescriptorUpdateTemplateCreateInfo struct {
	// Flags is reserved for future use
	Flags DescriptorUpdateTemplateCreateFlags
	// DescriptorUpdateEntries is a slice of DescriptorUpdateTemplateEntry structures describing
	// the descriptors to be updated by the DescriptorUpdateTEmplate
	DescriptorUpdateEntries []DescriptorUpdateTemplateEntry
	// TemplateType specifies the type of the DescriptorUpdateTemplate
	TemplateType DescriptorUpdateTemplateType

	// DescriptorSetLayout is the DescriptorSetLayout used to build the DescriptorUpdateTemplate
	DescriptorSetLayout core.DescriptorSetLayout

	// PipelineBindPoint indicates the type of the Pipeline that will use the descriptors
	PipelineBindPoint core1_0.PipelineBindPoint
	// PipelineLayout is a PipelineLayout object used to program the bindings
	PipelineLayout core.PipelineLayout
	// Set is the set number of the DescriptorSet in the PipelineLayout that will be updated
	Set int

	common.NextOptions
}

func (o DescriptorUpdateTemplateCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDescriptorUpdateTemplateCreateInfoKHR{})))
	}

	createInfo := (*C.VkDescriptorUpdateTemplateCreateInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO_KHR
	createInfo.pNext = next
	createInfo.flags = C.VkDescriptorUpdateTemplateCreateFlags(o.Flags)

	entryCount := len(o.DescriptorUpdateEntries)
	createInfo.descriptorUpdateEntryCount = C.uint32_t(entryCount)

	var err error
	createInfo.pDescriptorUpdateEntries, err = common.AllocSlice[C.VkDescriptorUpdateTemplateEntryKHR, DescriptorUpdateTemplateEntry](allocator, o.DescriptorUpdateEntries)
	if err != nil {
		return nil, err
	}

	createInfo.templateType = C.VkDescriptorUpdateTemplateType(o.TemplateType)
	createInfo.descriptorSetLayout = nil
	createInfo.pipelineLayout = nil

	if o.DescriptorSetLayout.Handle() != 0 {
		createInfo.descriptorSetLayout = C.VkDescriptorSetLayout(unsafe.Pointer(o.DescriptorSetLayout.Handle()))
	}

	if o.PipelineLayout.Handle() != 0 {
		createInfo.pipelineLayout = C.VkPipelineLayout(unsafe.Pointer(o.PipelineLayout.Handle()))
	}

	createInfo.pipelineBindPoint = C.VkPipelineBindPoint(o.PipelineBindPoint)
	createInfo.set = C.uint32_t(o.Set)

	return preallocatedPointer, nil
}
