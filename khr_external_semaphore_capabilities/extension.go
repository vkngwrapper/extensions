package khr_external_semaphore_capabilities

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	khr_external_semaphore_capabilities_driver "github.com/vkngwrapper/extensions/v2/khr_external_semaphore_capabilities/driver"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_external_semaphore_capabilities_driver.Driver
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_external_semaphore_capabilities loaded
func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return CreateExtensionFromDriver(khr_external_semaphore_capabilities_driver.CreateDriverFromCore(device.Driver()))
}

// CreateExtensionFromDriver generates an Extension from a driver.Driver object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_external_semaphore_capabilities_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) PhysicalDeviceExternalSemaphoreProperties(physicalDevice core1_0.PhysicalDevice, o PhysicalDeviceExternalSemaphoreInfo, outData *ExternalSemaphoreProperties) error {
	if physicalDevice == nil {
		panic("physicalDevice cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOutDataHeader(arena, outData)
	if err != nil {
		return err
	}

	e.driver.VkGetPhysicalDeviceExternalSemaphorePropertiesKHR(
		physicalDevice.Handle(),
		(*khr_external_semaphore_capabilities_driver.VkPhysicalDeviceExternalSemaphoreInfoKHR)(infoPtr),
		(*khr_external_semaphore_capabilities_driver.VkExternalSemaphorePropertiesKHR)(outDataPtr),
	)

	return common.PopulateOutData(outData, outDataPtr)
}
