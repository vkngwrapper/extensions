package khr_dynamic_rendering

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	khr_dynamic_rendering_loader "github.com/vkngwrapper/extensions/v3/khr_dynamic_rendering/loader"
)

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	driver khr_dynamic_rendering_loader.Loader
	device core1_0.Device
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// khr_dynamic_rendering loaded
func CreateExtensionDriverFromCoreDriver(coreDriver core1_0.DeviceDriver) ExtensionDriver {
	device := coreDriver.Device()
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}
	return CreateExtensionDriverFromLoader(khr_dynamic_rendering_loader.CreateLoaderFromCore(coreDriver.Loader()), device)
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver khr_dynamic_rendering_loader.Loader, device core1_0.Device) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		driver: driver,
		device: device,
	}
}

func (e *VulkanExtensionDriver) CmdBeginRendering(commandBuffer core1_0.CommandBuffer, renderingInfo RenderingInfo) error {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	renderingInfoPtr, err := common.AllocOptions(arena, renderingInfo)
	if err != nil {
		return err
	}

	e.driver.VkCmdBeginRenderingKHR(commandBuffer.Handle(), (*khr_dynamic_rendering_loader.VkRenderingInfoKHR)(renderingInfoPtr))

	return nil
}

func (e *VulkanExtensionDriver) CmdEndRendering(commandBuffer core1_0.CommandBuffer) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	e.driver.VkCmdEndRenderingKHR(commandBuffer.Handle())
}

var _ ExtensionDriver = &VulkanExtensionDriver{}
