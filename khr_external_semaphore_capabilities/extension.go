package khr_external_semaphore_capabilities

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/extensions/v3/khr_external_semaphore_capabilities/loader"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_external_semaphore_capabilities_loader.Loader
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_external_semaphore_capabilities loaded
func CreateExtensionFromDevice(device core.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return CreateExtensionFromDriver(khr_external_semaphore_capabilities_loader.CreateLoaderFromCore(device.Driver()))
}

// CreateExtensionFromDriver generates an Extension from a loader.Loader object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_external_semaphore_capabilities_loader.Loader) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) PhysicalDeviceExternalSemaphoreProperties(physicalDevice core.PhysicalDevice, o PhysicalDeviceExternalSemaphoreInfo, outData *ExternalSemaphoreProperties) error {
	if physicalDevice.Handle() == 0 {
		panic("physicalDevice cannot be uninitialized")
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
		(*khr_external_semaphore_capabilities_loader.VkPhysicalDeviceExternalSemaphoreInfoKHR)(infoPtr),
		(*khr_external_semaphore_capabilities_loader.VkExternalSemaphorePropertiesKHR)(outDataPtr),
	)

	return common.PopulateOutData(outData, outDataPtr)
}
