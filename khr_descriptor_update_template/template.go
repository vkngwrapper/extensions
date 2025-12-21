package khr_descriptor_update_template

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	khr_descriptor_update_template_driver "github.com/vkngwrapper/extensions/v3/khr_descriptor_update_template/driver"
)

// VulkanDescriptorUpdateTemplate is an implementation of the DescriptorUpdateTemplate interface that actually communicates
// with Vulkan. This is the default implementation. See the interface for more documentation.
type VulkanDescriptorUpdateTemplate struct {
	coreDriver               driver.Driver
	driver                   khr_descriptor_update_template_driver.Driver
	device                   driver.VkDevice
	descriptorTemplateHandle khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR

	maximumAPIVersion common.APIVersion
}

func (t *VulkanDescriptorUpdateTemplate) DeviceHandle() driver.VkDevice {
	return t.device
}

func (t *VulkanDescriptorUpdateTemplate) Driver() driver.Driver {
	return t.coreDriver
}

func (t *VulkanDescriptorUpdateTemplate) APIVersion() common.APIVersion {
	return t.maximumAPIVersion
}

func (t *VulkanDescriptorUpdateTemplate) Handle() khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR {
	return t.descriptorTemplateHandle
}

func (t *VulkanDescriptorUpdateTemplate) Destroy(allocator *driver.AllocationCallbacks) {
	t.driver.VkDestroyDescriptorUpdateTemplateKHR(t.device, t.descriptorTemplateHandle, allocator.Handle())
}

func (t *VulkanDescriptorUpdateTemplate) UpdateDescriptorSetFromImage(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorImageInfo) {
	if descriptorSet == nil {
		panic("descriptorSet cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoUnsafe := arena.Malloc(C.sizeof_struct_VkDescriptorImageInfo)
	info := (*C.VkDescriptorImageInfo)(infoUnsafe)
	info.imageView = nil
	info.sampler = nil
	info.imageLayout = C.VkImageLayout(data.ImageLayout)

	if data.Sampler != nil {
		info.sampler = C.VkSampler(unsafe.Pointer(data.Sampler.Handle()))
	}

	if data.ImageView != nil {
		info.imageView = C.VkImageView(unsafe.Pointer(data.ImageView.Handle()))
	}

	t.driver.VkUpdateDescriptorSetWithTemplateKHR(
		t.device,
		descriptorSet.Handle(),
		t.descriptorTemplateHandle,
		infoUnsafe,
	)
}

func (t *VulkanDescriptorUpdateTemplate) UpdateDescriptorSetFromBuffer(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorBufferInfo) {
	if descriptorSet == nil {
		panic("descriptorSet cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoUnsafe := arena.Malloc(C.sizeof_struct_VkDescriptorBufferInfo)
	info := (*C.VkDescriptorBufferInfo)(infoUnsafe)
	info.buffer = nil
	info.offset = C.VkDeviceSize(data.Offset)
	info._range = C.VkDeviceSize(data.Range)

	if data.Buffer != nil {
		info.buffer = C.VkBuffer(unsafe.Pointer(data.Buffer.Handle()))
	}

	t.driver.VkUpdateDescriptorSetWithTemplateKHR(
		t.device,
		descriptorSet.Handle(),
		t.descriptorTemplateHandle,
		infoUnsafe,
	)
}

func (t *VulkanDescriptorUpdateTemplate) UpdateDescriptorSetFromObjectHandle(descriptorSet core1_0.DescriptorSet, data driver.VulkanHandle) {
	if descriptorSet == nil {
		panic("descriptorSet cannot be nil")
	}
	t.driver.VkUpdateDescriptorSetWithTemplateKHR(
		t.device,
		descriptorSet.Handle(),
		t.descriptorTemplateHandle,
		unsafe.Pointer(data),
	)
}
