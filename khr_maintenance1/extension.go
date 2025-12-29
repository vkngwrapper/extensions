package khr_maintenance1

//go:generate mockgen -source extension.go -destination ./mocks/extension.go -package mock_maintenance1

import (
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/extensions/v3/khr_maintenance1/loader"
)

// ExtensionDriver contains all commands for the khr_maintenance1 extension
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance1.html
type ExtensionDriver interface {
	// TrimCommandPool trims a CommandPool
	//
	// commandPool - The CommandPool to trim
	//
	// flags - Reserved for future use
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkTrimCommandPool.html
	TrimCommandPool(commandPool core1_0.CommandPool, flags CommandPoolTrimFlags)
}

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	driver khr_maintenance1_loader.Loader
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// khr_maintenance1 loaded
func CreateExtensionDriverFromCoreDriver(driver core1_0.DeviceDriver) ExtensionDriver {
	device := driver.Device()
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtensionDriver{
		driver: khr_maintenance1_loader.CreateLoaderFromCore(driver.Loader()),
	}
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver khr_maintenance1_loader.Loader) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		driver: driver,
	}
}

func (e *VulkanExtensionDriver) TrimCommandPool(commandPool core1_0.CommandPool, flags CommandPoolTrimFlags) {
	if !commandPool.Initialized() {
		panic("commandPool cannot be uninitialized")
	}
	e.driver.VkTrimCommandPoolKHR(commandPool.DeviceHandle(), commandPool.Handle(), khr_maintenance1_loader.VkCommandPoolTrimFlagsKHR(flags))
}
