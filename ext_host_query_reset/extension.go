package ext_host_query_reset

import (
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	ext_host_query_reset_driver "github.com/vkngwrapper/extensions/v3/ext_host_query_reset/loader"
)

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	driver ext_host_query_reset_driver.Loader
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// ext_host_query_reset loaded
func CreateExtensionDriverFromCoreDriver(coreDriver core1_0.DeviceDriver) ExtensionDriver {
	device := coreDriver.Device()

	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtensionDriver{
		driver: ext_host_query_reset_driver.CreateLoaderFromCore(coreDriver.Loader()),
	}
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver ext_host_query_reset_driver.Loader) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		driver: driver,
	}
}

func (e *VulkanExtensionDriver) ResetQueryPool(queryPool core.QueryPool, firstQuery, queryCount int) {
	if queryPool.Handle() == 0 {
		panic("queryPool cannot be uninitialized")
	}
	e.driver.VkResetQueryPoolEXT(queryPool.DeviceHandle(), queryPool.Handle(), loader.Uint32(firstQuery), loader.Uint32(queryCount))
}
