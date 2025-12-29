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
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/extensions/v3/khr_descriptor_update_template/loader"
)

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	driver khr_descriptor_update_template_loader.Loader
	device core.Device
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// khr_descriptor_update_template loaded
func CreateExtensionDriverFromCoreDriver(coreDriver core1_0.DeviceDriver) ExtensionDriver {
	device := coreDriver.Device()
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}
	return CreateExtensionDriverFromLoader(khr_descriptor_update_template_loader.CreateLoaderFromCore(coreDriver.Loader()), device)
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver khr_descriptor_update_template_loader.Loader, device core.Device) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		driver: driver,
		device: device,
	}
}

func (e *VulkanExtensionDriver) CreateDescriptorUpdateTemplate(o DescriptorUpdateTemplateCreateInfo, allocator *loader.AllocationCallbacks) (core.DescriptorUpdateTemplate, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfoPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.DescriptorUpdateTemplate{}, core1_0.VKErrorUnknown, err
	}

	var templateHandle khr_descriptor_update_template_loader.VkDescriptorUpdateTemplateKHR
	res, err := e.driver.VkCreateDescriptorUpdateTemplateKHR(e.device.Handle(),
		(*khr_descriptor_update_template_loader.VkDescriptorUpdateTemplateCreateInfoKHR)(createInfoPtr),
		allocator.Handle(),
		&templateHandle,
	)
	if err != nil {
		return core.DescriptorUpdateTemplate{}, res, err
	}

	descriptorTemplate := core.InternalDescriptorUpdateTemplate(
		e.device.Handle(),
		loader.VkDescriptorUpdateTemplate(templateHandle),
		e.device.APIVersion(),
	)

	return descriptorTemplate, res, nil
}

func (t *VulkanExtensionDriver) DestroyDescriptorUpdateTemplate(template core.DescriptorUpdateTemplate, allocator *loader.AllocationCallbacks) {
	if template.Handle() == 0 {
		panic("template cannot be uninitialized")
	}
	t.driver.VkDestroyDescriptorUpdateTemplateKHR(template.DeviceHandle(), khr_descriptor_update_template_loader.VkDescriptorUpdateTemplateKHR(template.Handle()), allocator.Handle())
}

func (t *VulkanExtensionDriver) UpdateDescriptorSetWithTemplateFromImage(descriptorSet core.DescriptorSet, template core.DescriptorUpdateTemplate, data core1_0.DescriptorImageInfo) {
	if descriptorSet.Handle() == 0 {
		panic("descriptorSet cannot be uninitialized")
	}
	if template.Handle() == 0 {
		panic("template cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoUnsafe := arena.Malloc(C.sizeof_struct_VkDescriptorImageInfo)
	info := (*C.VkDescriptorImageInfo)(infoUnsafe)
	info.imageView = nil
	info.sampler = nil
	info.imageLayout = C.VkImageLayout(data.ImageLayout)

	if data.Sampler.Handle() != 0 {
		info.sampler = C.VkSampler(unsafe.Pointer(data.Sampler.Handle()))
	}

	if data.ImageView.Handle() != 0 {
		info.imageView = C.VkImageView(unsafe.Pointer(data.ImageView.Handle()))
	}

	t.driver.VkUpdateDescriptorSetWithTemplateKHR(
		descriptorSet.DeviceHandle(),
		descriptorSet.Handle(),
		khr_descriptor_update_template_loader.VkDescriptorUpdateTemplateKHR(template.Handle()),
		infoUnsafe,
	)
}

func (t *VulkanExtensionDriver) UpdateDescriptorSetWithTemplateFromBuffer(descriptorSet core.DescriptorSet, template core.DescriptorUpdateTemplate, data core1_0.DescriptorBufferInfo) {
	if descriptorSet.Handle() == 0 {
		panic("descriptorSet cannot be uninitialized")
	}
	if template.Handle() == 0 {
		panic("template cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoUnsafe := arena.Malloc(C.sizeof_struct_VkDescriptorBufferInfo)
	info := (*C.VkDescriptorBufferInfo)(infoUnsafe)
	info.buffer = nil
	info.offset = C.VkDeviceSize(data.Offset)
	info._range = C.VkDeviceSize(data.Range)

	if data.Buffer.Handle() != 0 {
		info.buffer = C.VkBuffer(unsafe.Pointer(data.Buffer.Handle()))
	}

	t.driver.VkUpdateDescriptorSetWithTemplateKHR(
		descriptorSet.DeviceHandle(),
		descriptorSet.Handle(),
		khr_descriptor_update_template_loader.VkDescriptorUpdateTemplateKHR(template.Handle()),
		infoUnsafe,
	)
}

func (t *VulkanExtensionDriver) UpdateDescriptorSetWithTemplateFromObjectHandle(descriptorSet core.DescriptorSet, template core.DescriptorUpdateTemplate, data loader.VulkanHandle) {
	if descriptorSet.Handle() == 0 {
		panic("descriptorSet cannot be uninitialized")
	}
	if template.Handle() == 0 {
		panic("template cannot be uninitialized")
	}
	t.driver.VkUpdateDescriptorSetWithTemplateKHR(
		descriptorSet.DeviceHandle(),
		descriptorSet.Handle(),
		khr_descriptor_update_template_loader.VkDescriptorUpdateTemplateKHR(template.Handle()),
		unsafe.Pointer(data),
	)
}
