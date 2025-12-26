package khr_draw_indirect_count

import (
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/extensions/v3/khr_draw_indirect_count/loader"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_draw_indirect_count_loader.Loader
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_draw_indirect_count loaded
func CreateExtensionFromDevice(device core.Device, instance core.Instance) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return CreateExtensionFromDriver(khr_draw_indirect_count_loader.CreateDriverFromCore(device.Driver()))
}

// CreateExtensionFromDriver generates an Extension from a loader.Loader object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_draw_indirect_count_loader.Loader) *VulkanExtension {
	ext := &VulkanExtension{
		driver: driver,
	}

	return ext
}

func (e *VulkanExtension) CmdDrawIndexedIndirectCount(commandBuffer core.CommandBuffer, buffer core.Buffer, offset uint64, countBuffer core.Buffer, countBufferOffset uint64, maxDrawCount, stride int) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}
	if buffer.Handle() == 0 {
		panic("buffer cannot be uninitialized")
	}
	if countBuffer.Handle() == 0 {
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

func (e *VulkanExtension) CmdDrawIndirectCount(commandBuffer core.CommandBuffer, buffer core.Buffer, offset uint64, countBuffer core.Buffer, countBufferOffset uint64, maxDrawCount, stride int) {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}
	if buffer.Handle() == 0 {
		panic("buffer cannot be uninitialized")
	}
	if countBuffer.Handle() == 0 {
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
