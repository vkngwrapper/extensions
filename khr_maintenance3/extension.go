package khr_maintenance3

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	khr_maintenance3_driver "github.com/vkngwrapper/extensions/v2/khr_maintenance3/driver"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_maintenance3_driver.Driver
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_maintenance3 loaded
func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: khr_maintenance3_driver.CreateDriverFromCore(device.Driver()),
	}
}

// CreateExtensionFromDriver generates an Extension from a driver.Driver object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_maintenance3_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) DescriptorSetLayoutSupport(device core1_0.Device, setLayoutOptions core1_0.DescriptorSetLayoutCreateInfo, support *DescriptorSetLayoutSupport) error {
	if device == nil {
		panic("device cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionsPtr, err := common.AllocOptions(arena, setLayoutOptions)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOutDataHeader(arena, support)
	if err != nil {
		return err
	}

	e.driver.VkGetDescriptorSetLayoutSupportKHR(device.Handle(), (*driver.VkDescriptorSetLayoutCreateInfo)(optionsPtr), (*khr_maintenance3_driver.VkDescriptorSetLayoutSupportKHR)(outDataPtr))

	return common.PopulateOutData(support, outDataPtr)
}

var _ Extension = &VulkanExtension{}
