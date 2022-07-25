package khr_external_memory_capabilities

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	khr_external_memory_capabilities_driver "github.com/vkngwrapper/extensions/khr_external_memory_capabilities/driver"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_external_memory_capabilities_driver.Driver
}

func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return CreateExtensionFromDriver(khr_external_memory_capabilities_driver.CreateDriverFromCore(device.Driver()))
}

// CreateExtensionFromDriver generates an Extension from a driver.Driver object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_external_memory_capabilities_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) PhysicalDeviceExternalBufferProperties(physicalDevice core1_0.PhysicalDevice, o PhysicalDeviceExternalBufferInfo, outData *ExternalBufferProperties) error {
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
		(*khr_external_memory_capabilities_driver.VkPhysicalDeviceExternalBufferInfoKHR)(optionsPtr),
		(*khr_external_memory_capabilities_driver.VkExternalBufferPropertiesKHR)(outDataPtr))

	return common.PopulateOutData(outData, outDataPtr)
}
