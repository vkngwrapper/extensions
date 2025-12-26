package khr_maintenance1

//go:generate mockgen -source extension.go -destination ./mocks/extension.go -package mock_maintenance1

import (
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/extensions/v3/khr_maintenance1/loader"
)

// Extension contains all commands for the khr_maintenance1 extension
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance1.html
type Extension interface {
	// TrimCommandPool trims a CommandPool
	//
	// commandPool - The CommandPool to trim
	//
	// flags - Reserved for future use
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkTrimCommandPool.html
	TrimCommandPool(commandPool core.CommandPool, flags CommandPoolTrimFlags)
}

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_maintenance1_loader.Loader
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_maintenance1 loaded
func CreateExtensionFromDevice(device core.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: khr_maintenance1_loader.CreateLoaderFromCore(device.Driver()),
	}
}

// CreateExtensionFromDriver generates an Extension from a loader.Loader object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_maintenance1_loader.Loader) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) TrimCommandPool(commandPool core.CommandPool, flags CommandPoolTrimFlags) {
	if commandPool.Handle() == 0 {
		panic("commandPool cannot be uninitialized")
	}
	e.driver.VkTrimCommandPoolKHR(commandPool.DeviceHandle(), commandPool.Handle(), khr_maintenance1_loader.VkCommandPoolTrimFlagsKHR(flags))
}
