package ext_host_query_reset

import (
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	ext_host_query_reset_driver "github.com/vkngwrapper/extensions/v3/ext_host_query_reset/loader"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver ext_host_query_reset_driver.Loader
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// ext_host_query_reset loaded
func CreateExtensionFromDevice(device core.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: ext_host_query_reset_driver.CreateLoaderFromCore(device.Driver()),
	}
}

// CreateExtensionFromDriver generates an Extension from a loader.Loader object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver ext_host_query_reset_driver.Loader) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) ResetQueryPool(queryPool core.QueryPool, firstQuery, queryCount int) {
	if queryPool.Handle() == 0 {
		panic("queryPool cannot be uninitialized")
	}
	e.driver.VkResetQueryPoolEXT(queryPool.DeviceHandle(), queryPool.Handle(), loader.Uint32(firstQuery), loader.Uint32(queryCount))
}
