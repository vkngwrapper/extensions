package khr_create_renderpass2

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/extensions/v3/khr_create_renderpass2/loader"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_create_renderpass2_loader.Loader
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_create_renderpass2 loaded
func CreateExtensionFromDevice(device core.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}
	return CreateExtensionFromDriver(khr_create_renderpass2_loader.CreateLoaderFromCore(device.Driver()))
}

// CreateExtensionFromDriver generates an Extension from a loader.Loader object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_create_renderpass2_loader.Loader) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) CmdBeginRenderPass2(commandBuffer core.CommandBuffer, renderPassBegin core1_0.RenderPassBeginInfo, subpassBegin SubpassBeginInfo) error {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
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
		(*loader.VkRenderPassBeginInfo)(renderPassBeginPtr),
		(*khr_create_renderpass2_loader.VkSubpassBeginInfoKHR)(subpassBeginPtr),
	)

	return nil
}

func (e *VulkanExtension) CmdEndRenderPass2(commandBuffer core.CommandBuffer, subpassEnd SubpassEndInfo) error {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	subpassEndPtr, err := common.AllocOptions(arena, subpassEnd)
	if err != nil {
		return err
	}

	e.driver.VkCmdEndRenderPass2KHR(
		commandBuffer.Handle(),
		(*khr_create_renderpass2_loader.VkSubpassEndInfoKHR)(subpassEndPtr),
	)

	return nil
}

func (e *VulkanExtension) CmdNextSubpass2(commandBuffer core.CommandBuffer, subpassBegin SubpassBeginInfo, subpassEnd SubpassEndInfo) error {
	if commandBuffer.Handle() == 0 {
		panic("commandBuffer cannot be uninitialized")
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
		(*khr_create_renderpass2_loader.VkSubpassBeginInfoKHR)(subpassBeginPtr),
		(*khr_create_renderpass2_loader.VkSubpassEndInfoKHR)(subpassEndPtr),
	)

	return nil
}

func (e *VulkanExtension) CreateRenderPass2(device core.Device, allocator *loader.AllocationCallbacks, options RenderPassCreateInfo2) (core.RenderPass, common.VkResult, error) {
	if device.Handle() == 0 {
		panic("device cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, options)
	if err != nil {
		return core.RenderPass{}, core1_0.VKErrorUnknown, err
	}

	var renderPassHandle loader.VkRenderPass
	res, err := e.driver.VkCreateRenderPass2KHR(
		device.Handle(),
		(*khr_create_renderpass2_loader.VkRenderPassCreateInfo2KHR)(infoPtr),
		allocator.Handle(),
		&renderPassHandle,
	)
	if err != nil {
		return core.RenderPass{}, res, err
	}

	renderPass := core.InternalRenderPass(
		device.Handle(),
		renderPassHandle,
		device.APIVersion(),
	)

	return renderPass, res, nil
}
