package ext_host_query_reset

import (
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/driver"
	ext_host_query_reset_driver "github.com/vkngwrapper/extensions/v3/ext_host_query_reset/driver"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver ext_host_query_reset_driver.Driver
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// ext_host_query_reset loaded
func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: ext_host_query_reset_driver.CreateDriverFromCore(device.Driver()),
	}
}

// CreateExtensionFromDriver generates an Extension from a driver.Driver object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver ext_host_query_reset_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) ResetQueryPool(queryPool core1_0.QueryPool, firstQuery, queryCount int) {
	if queryPool == nil {
		panic("queryPool cannot be nil")
	}
	e.driver.VkResetQueryPoolEXT(queryPool.DeviceHandle(), queryPool.Handle(), driver.Uint32(firstQuery), driver.Uint32(queryCount))
}
