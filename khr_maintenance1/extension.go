package khr_maintenance1

//go:generate mockgen -source extension.go -destination ./mocks/extension.go -package mock_maintenance1

import (
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/extensions/khr_maintenance1/driver"
)

type Extension interface {
	TrimCommandPool(commandPool core1_0.CommandPool, flags CommandPoolTrimFlags)
}

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_maintenance1_driver.Driver
}

func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: khr_maintenance1_driver.CreateDriverFromCore(device.Driver()),
	}
}

// CreateExtensionFromDriver generates an Extension from a driver.Driver object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_maintenance1_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) TrimCommandPool(commandPool core1_0.CommandPool, flags CommandPoolTrimFlags) {
	e.driver.VkTrimCommandPoolKHR(commandPool.DeviceHandle(), commandPool.Handle(), khr_maintenance1_driver.VkCommandPoolTrimFlagsKHR(flags))
}
