package khr_draw_indirect_count

import (
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	khr_draw_indirect_count_driver "github.com/vkngwrapper/extensions/v2/khr_draw_indirect_count/driver"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_draw_indirect_count_driver.Driver
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_draw_indirect_count loaded
func CreateExtensionFromDevice(device core1_0.Device, instance core1_0.Instance) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return CreateExtensionFromDriver(khr_draw_indirect_count_driver.CreateDriverFromCore(device.Driver()))
}

// CreateExtensionFromDriver generates an Extension from a driver.Driver object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_draw_indirect_count_driver.Driver) *VulkanExtension {
	ext := &VulkanExtension{
		driver: driver,
	}

	return ext
}

func (e *VulkanExtension) CmdDrawIndexedIndirectCount(commandBuffer core1_0.CommandBuffer, buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int) {
	e.driver.VkCmdDrawIndexedIndirectCountKHR(
		commandBuffer.Handle(),
		buffer.Handle(),
		driver.VkDeviceSize(offset),
		countBuffer.Handle(),
		driver.VkDeviceSize(countBufferOffset),
		driver.Uint32(maxDrawCount),
		driver.Uint32(stride),
	)
}

func (e *VulkanExtension) CmdDrawIndirectCount(commandBuffer core1_0.CommandBuffer, buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int) {
	e.driver.VkCmdDrawIndirectCountKHR(
		commandBuffer.Handle(),
		buffer.Handle(),
		driver.VkDeviceSize(offset),
		countBuffer.Handle(),
		driver.VkDeviceSize(countBufferOffset),
		driver.Uint32(maxDrawCount),
		driver.Uint32(stride),
	)
}
