package khr_maintenance3

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/extensions/v3/khr_maintenance3/loader"
)

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	driver khr_maintenance3_loader.Loader
	device core.Device
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// khr_maintenance3 loaded
func CreateExtensionDriverFromCoreDriver(driver core1_0.DeviceDriver) *VulkanExtensionDriver {
	device := driver.Device()
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtensionDriver{
		driver: khr_maintenance3_loader.CreateLoaderFromCore(driver.Loader()),
		device: device,
	}
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver khr_maintenance3_loader.Loader, device core.Device) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		driver: driver,
		device: device,
	}
}

func (e *VulkanExtensionDriver) DescriptorSetLayoutSupport(setLayoutOptions core1_0.DescriptorSetLayoutCreateInfo, support *DescriptorSetLayoutSupport) error {
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

	e.driver.VkGetDescriptorSetLayoutSupportKHR(e.device.Handle(), (*loader.VkDescriptorSetLayoutCreateInfo)(optionsPtr), (*khr_maintenance3_loader.VkDescriptorSetLayoutSupportKHR)(outDataPtr))

	return common.PopulateOutData(support, outDataPtr)
}

var _ ExtensionDriver = &VulkanExtensionDriver{}
