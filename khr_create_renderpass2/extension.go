package khr_create_renderpass2

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/common/extensions"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	khr_create_renderpass2_driver "github.com/vkngwrapper/extensions/v3/khr_create_renderpass2/driver"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_create_renderpass2_driver.Driver
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_create_renderpass2 loaded
func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}
	return CreateExtensionFromDriver(khr_create_renderpass2_driver.CreateDriverFromCore(device.Driver()))
}

// CreateExtensionFromDriver generates an Extension from a driver.Driver object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_create_renderpass2_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) CmdBeginRenderPass2(commandBuffer core1_0.CommandBuffer, renderPassBegin core1_0.RenderPassBeginInfo, subpassBegin SubpassBeginInfo) error {
	if commandBuffer == nil {
		panic("commandBuffer cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	renderPassBeginPtr, err := common.AllocOptions(arena, renderPassBegin)
	if err != nil {
		return err
	}

	subpassBeginPtr, err := common.AllocOptions(arena, subpassBegin)
	if err != nil {
		return err
	}

	e.driver.VkCmdBeginRenderPass2KHR(
		commandBuffer.Handle(),
		(*driver.VkRenderPassBeginInfo)(renderPassBeginPtr),
		(*khr_create_renderpass2_driver.VkSubpassBeginInfoKHR)(subpassBeginPtr),
	)

	return nil
}

func (e *VulkanExtension) CmdEndRenderPass2(commandBuffer core1_0.CommandBuffer, subpassEnd SubpassEndInfo) error {
	if commandBuffer == nil {
		panic("commandBuffer cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	subpassEndPtr, err := common.AllocOptions(arena, subpassEnd)
	if err != nil {
		return err
	}

	e.driver.VkCmdEndRenderPass2KHR(
		commandBuffer.Handle(),
		(*khr_create_renderpass2_driver.VkSubpassEndInfoKHR)(subpassEndPtr),
	)

	return nil
}

func (e *VulkanExtension) CmdNextSubpass2(commandBuffer core1_0.CommandBuffer, subpassBegin SubpassBeginInfo, subpassEnd SubpassEndInfo) error {
	if commandBuffer == nil {
		panic("commandBuffer cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	subpassBeginPtr, err := common.AllocOptions(arena, subpassBegin)
	if err != nil {
		return err
	}

	subpassEndPtr, err := common.AllocOptions(arena, subpassEnd)
	if err != nil {
		return err
	}

	e.driver.VkCmdNextSubpass2KHR(
		commandBuffer.Handle(),
		(*khr_create_renderpass2_driver.VkSubpassBeginInfoKHR)(subpassBeginPtr),
		(*khr_create_renderpass2_driver.VkSubpassEndInfoKHR)(subpassEndPtr),
	)

	return nil
}

func (e *VulkanExtension) CreateRenderPass2(device core1_0.Device, allocator *driver.AllocationCallbacks, options RenderPassCreateInfo2) (core1_0.RenderPass, common.VkResult, error) {
	if device == nil {
		panic("device cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var renderPassHandle driver.VkRenderPass
	res, err := e.driver.VkCreateRenderPass2KHR(
		device.Handle(),
		(*khr_create_renderpass2_driver.VkRenderPassCreateInfo2KHR)(infoPtr),
		allocator.Handle(),
		&renderPassHandle,
	)
	if err != nil {
		return nil, res, err
	}

	renderPass := extensions.CreateRenderPassObject(
		device.Driver(),
		device.Handle(),
		renderPassHandle,
		device.APIVersion(),
	)

	return renderPass, res, nil
}
