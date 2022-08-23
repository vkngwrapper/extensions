package khr_descriptor_update_template

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	khr_descriptor_update_template_driver "github.com/vkngwrapper/extensions/v2/khr_descriptor_update_template/driver"
	"unsafe"
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

func (t *VulkanDescriptorUpdateTemplate) Handle() khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR {
	return t.descriptorTemplateHandle
}

func (t *VulkanDescriptorUpdateTemplate) Destroy(allocator *driver.AllocationCallbacks) {
	t.driver.VkDestroyDescriptorUpdateTemplateKHR(t.device, t.descriptorTemplateHandle, allocator.Handle())
}

func (t *VulkanDescriptorUpdateTemplate) UpdateDescriptorSetFromImage(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorImageInfo) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoUnsafe := arena.Malloc(C.sizeof_struct_VkDescriptorImageInfo)
	info := (*C.VkDescriptorImageInfo)(infoUnsafe)
	info.sampler = C.VkSampler(unsafe.Pointer(data.Sampler.Handle()))
	info.imageView = C.VkImageView(unsafe.Pointer(data.ImageView.Handle()))
	info.imageLayout = C.VkImageLayout(data.ImageLayout)

	t.driver.VkUpdateDescriptorSetWithTemplateKHR(
		t.device,
		descriptorSet.Handle(),
		t.descriptorTemplateHandle,
		infoUnsafe,
	)
}

func (t *VulkanDescriptorUpdateTemplate) UpdateDescriptorSetFromBuffer(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorBufferInfo) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoUnsafe := arena.Malloc(C.sizeof_struct_VkDescriptorBufferInfo)
	info := (*C.VkDescriptorBufferInfo)(infoUnsafe)
	info.buffer = C.VkBuffer(unsafe.Pointer(data.Buffer.Handle()))
	info.offset = C.VkDeviceSize(data.Offset)
	info._range = C.VkDeviceSize(data.Range)

	t.driver.VkUpdateDescriptorSetWithTemplateKHR(
		t.device,
		descriptorSet.Handle(),
		t.descriptorTemplateHandle,
		infoUnsafe,
	)
}

func (t *VulkanDescriptorUpdateTemplate) UpdateDescriptorSetFromObjectHandle(descriptorSet core1_0.DescriptorSet, data driver.VulkanHandle) {
	t.driver.VkUpdateDescriptorSetWithTemplateKHR(
		t.device,
		descriptorSet.Handle(),
		t.descriptorTemplateHandle,
		unsafe.Pointer(data),
	)
}
