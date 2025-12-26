package khr_maintenance3

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/extensions/v3/khr_maintenance3/loader"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_maintenance3_loader.Loader
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_maintenance3 loaded
func CreateExtensionFromDevice(device core.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: khr_maintenance3_loader.CreateLoaderFromCore(device.Driver()),
	}
}

// CreateExtensionFromDriver generates an Extension from a loader.Loader object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_maintenance3_loader.Loader) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) DescriptorSetLayoutSupport(device core.Device, setLayoutOptions core1_0.DescriptorSetLayoutCreateInfo, support *DescriptorSetLayoutSupport) error {
	if device.Handle() == 0 {
		panic("device cannot be uninitialized")
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

	e.driver.VkGetDescriptorSetLayoutSupportKHR(device.Handle(), (*loader.VkDescriptorSetLayoutCreateInfo)(optionsPtr), (*khr_maintenance3_loader.VkDescriptorSetLayoutSupportKHR)(outDataPtr))

	return common.PopulateOutData(support, outDataPtr)
}

var _ Extension = &VulkanExtension{}
