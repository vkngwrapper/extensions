package khr_descriptor_update_template

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
	khr_descriptor_update_template_driver "github.com/vkngwrapper/extensions/khr_descriptor_update_template/driver"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_descriptor_update_template_driver.Driver
}

func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}
	return CreateExtensionFromDriver(khr_descriptor_update_template_driver.CreateDriverFromCore(device.Driver()))
}

// CreateExtensionFromDriver generates an Extension from a driver.Driver object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_descriptor_update_template_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) CreateDescriptorUpdateTemplate(device core1_0.Device, o DescriptorUpdateTemplateCreateInfo, allocator *driver.AllocationCallbacks) (DescriptorUpdateTemplate, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfoPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var templateHandle khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR
	res, err := e.driver.VkCreateDescriptorUpdateTemplateKHR(device.Handle(),
		(*khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateCreateInfoKHR)(createInfoPtr),
		allocator.Handle(),
		&templateHandle,
	)
	if err != nil {
		return nil, res, err
	}

	descriptorTemplate := device.Driver().ObjectStore().GetOrCreate(driver.VulkanHandle(templateHandle), driver.Core1_1,
		func() any {
			template := &vulkanDescriptorUpdateTemplate{
				driver:                   e.driver,
				coreDriver:               device.Driver(),
				device:                   device.Handle(),
				descriptorTemplateHandle: templateHandle,
				maximumAPIVersion:        device.APIVersion(),
			}

			return template
		}).(*vulkanDescriptorUpdateTemplate)
	device.Driver().ObjectStore().SetParent(driver.VulkanHandle(device.Handle()), driver.VulkanHandle(templateHandle))

	return descriptorTemplate, res, nil
}
