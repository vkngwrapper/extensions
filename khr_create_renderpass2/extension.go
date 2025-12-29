package khr_create_renderpass2

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/extensions/v3/khr_create_renderpass2/loader"
)

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	driver khr_create_renderpass2_loader.Loader
	device core1_0.Device
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// khr_create_renderpass2 loaded
func CreateExtensionDriverFromCoreDriver(coreDriver core1_0.DeviceDriver) ExtensionDriver {
	device := coreDriver.Device()
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}
	return CreateExtensionDriverFromLoader(khr_create_renderpass2_loader.CreateLoaderFromCore(coreDriver.Loader()), device)
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver khr_create_renderpass2_loader.Loader, device core1_0.Device) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		driver: driver,
		device: device,
	}
}

func (e *VulkanExtensionDriver) CmdBeginRenderPass2(commandBuffer core1_0.CommandBuffer, renderPassBegin core1_0.RenderPassBeginInfo, subpassBegin SubpassBeginInfo) error {
	if !commandBuffer.Initialized() {
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

func (e *VulkanExtensionDriver) CmdEndRenderPass2(commandBuffer core1_0.CommandBuffer, subpassEnd SubpassEndInfo) error {
	if !commandBuffer.Initialized() {
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

func (e *VulkanExtensionDriver) CmdNextSubpass2(commandBuffer core1_0.CommandBuffer, subpassBegin SubpassBeginInfo, subpassEnd SubpassEndInfo) error {
	if !commandBuffer.Initialized() {
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

func (e *VulkanExtensionDriver) CreateRenderPass2(allocator *loader.AllocationCallbacks, options RenderPassCreateInfo2) (core1_0.RenderPass, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, options)
	if err != nil {
		return core1_0.RenderPass{}, core1_0.VKErrorUnknown, err
	}

	var renderPassHandle loader.VkRenderPass
	res, err := e.driver.VkCreateRenderPass2KHR(
		e.device.Handle(),
		(*khr_create_renderpass2_loader.VkRenderPassCreateInfo2KHR)(infoPtr),
		allocator.Handle(),
		&renderPassHandle,
	)
	if err != nil {
		return core1_0.RenderPass{}, res, err
	}

	renderPass := core1_0.InternalRenderPass(
		e.device.Handle(),
		renderPassHandle,
		e.device.APIVersion(),
	)

	return renderPass, res, nil
}
