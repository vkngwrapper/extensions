package khr_draw_indirect_count

import (
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/extensions/v3/khr_draw_indirect_count/loader"
)

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	driver khr_draw_indirect_count_loader.Loader
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// khr_draw_indirect_count loaded
func CreateExtensionDriverFromCoreDriver(driver core1_0.DeviceDriver) ExtensionDriver {
	device := driver.Device()

	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return CreateExtensionDriverFromLoader(khr_draw_indirect_count_loader.CreateDriverFromCore(driver.Loader()))
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver khr_draw_indirect_count_loader.Loader) *VulkanExtensionDriver {
	ext := &VulkanExtensionDriver{
		driver: driver,
	}

	return ext
}

func (e *VulkanExtensionDriver) CmdDrawIndexedIndirectCount(commandBuffer core1_0.CommandBuffer, buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}
	if !buffer.Initialized() {
		panic("buffer cannot be uninitialized")
	}
	if !countBuffer.Initialized() {
		panic("countBuffer cannot be uninitialized")
	}
	e.driver.VkCmdDrawIndexedIndirectCountKHR(
		commandBuffer.Handle(),
		buffer.Handle(),
		loader.VkDeviceSize(offset),
		countBuffer.Handle(),
		loader.VkDeviceSize(countBufferOffset),
		loader.Uint32(maxDrawCount),
		loader.Uint32(stride),
	)
}

func (e *VulkanExtensionDriver) CmdDrawIndirectCount(commandBuffer core1_0.CommandBuffer, buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}
	if !buffer.Initialized() {
		panic("buffer cannot be uninitialized")
	}
	if !countBuffer.Initialized() {
		panic("countBuffer cannot be uninitialized")
	}
	e.driver.VkCmdDrawIndirectCountKHR(
		commandBuffer.Handle(),
		buffer.Handle(),
		loader.VkDeviceSize(offset),
		countBuffer.Handle(),
		loader.VkDeviceSize(countBufferOffset),
		loader.Uint32(maxDrawCount),
		loader.Uint32(stride),
	)
}
