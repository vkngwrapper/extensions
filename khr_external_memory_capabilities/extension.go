package khr_external_memory_capabilities

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/extensions/v3/khr_external_memory_capabilities/loader"
)

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	driver khr_external_memory_capabilities_loader.Loader
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// khr_external_memory_capabilities loaded
func CreateExtensionDriverFromCoreDriver(driver core1_0.DeviceDriver) *VulkanExtensionDriver {
	device := driver.Device()
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return CreateExtensionDriverFromLoader(khr_external_memory_capabilities_loader.CreateLoaderFromCore(driver.Loader()))
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver khr_external_memory_capabilities_loader.Loader) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		driver: driver,
	}
}

func (e *VulkanExtensionDriver) GetPhysicalDeviceExternalBufferProperties(physicalDevice core.PhysicalDevice, o PhysicalDeviceExternalBufferInfo, outData *ExternalBufferProperties) error {
	if physicalDevice.Handle() == 0 {
		panic("physicalDevice cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionsPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOutDataHeader(arena, outData)
	if err != nil {
		return err
	}

	e.driver.VkGetPhysicalDeviceExternalBufferPropertiesKHR(physicalDevice.Handle(),
		(*khr_external_memory_capabilities_loader.VkPhysicalDeviceExternalBufferInfoKHR)(optionsPtr),
		(*khr_external_memory_capabilities_loader.VkExternalBufferPropertiesKHR)(outDataPtr))

	return common.PopulateOutData(outData, outDataPtr)
}
