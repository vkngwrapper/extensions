package khr_copy_commands2

import (
	"errors"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	khr_copy_commands2_loader "github.com/vkngwrapper/extensions/v3/khr_copy_commands2/loader"
)

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	driver khr_copy_commands2_loader.Loader
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// ext_host_query_reset loaded
func CreateExtensionDriverFromCoreDriver(coreDriver core1_0.DeviceDriver) ExtensionDriver {
	device := coreDriver.Device()

	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtensionDriver{
		driver: khr_copy_commands2_loader.CreateLoaderFromCore(coreDriver.Loader()),
	}
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver khr_copy_commands2_loader.Loader) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		driver: driver,
	}
}

func (e *VulkanExtensionDriver) CmdBlitImage2(commandBuffer core1_0.CommandBuffer, blitImageInfo BlitImageInfo2) error {
	if !commandBuffer.Initialized() {
		return errors.New("commandBuffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	blitInfo, err := common.AllocOptions(arena, blitImageInfo)
	if err != nil {
		return err
	}

	e.driver.VkCmdBlitImage2KHR(commandBuffer.Handle(), (*khr_copy_commands2_loader.VkBlitImageInfo2KHR)(blitInfo))
	return nil
}

func (e *VulkanExtensionDriver) CmdCopyBuffer2(commandBuffer core1_0.CommandBuffer, copyBufferInfo CopyBufferInfo2) error {
	if !commandBuffer.Initialized() {
		return errors.New("commandBuffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	copyInfo, err := common.AllocOptions(arena, copyBufferInfo)
	if err != nil {
		return err
	}

	e.driver.VkCmdCopyBuffer2KHR(commandBuffer.Handle(), (*khr_copy_commands2_loader.VkCopyBufferInfo2KHR)(copyInfo))
	return nil
}

func (e *VulkanExtensionDriver) CmdCopyBufferToImage2(commandBuffer core1_0.CommandBuffer, copyBufferToImageInfo CopyBufferToImageInfo2) error {
	if !commandBuffer.Initialized() {
		return errors.New("commandBuffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	copyInfo, err := common.AllocOptions(arena, copyBufferToImageInfo)
	if err != nil {
		return err
	}

	e.driver.VkCmdCopyBufferToImage2KHR(commandBuffer.Handle(), (*khr_copy_commands2_loader.VkCopyBufferToImageInfo2KHR)(copyInfo))
	return nil
}

func (e *VulkanExtensionDriver) CmdCopyImage2(commandBuffer core1_0.CommandBuffer, copyImageInfo CopyImageInfo2) error {
	if !commandBuffer.Initialized() {
		return errors.New("commandBuffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	copyInfo, err := common.AllocOptions(arena, copyImageInfo)
	if err != nil {
		return err
	}

	e.driver.VkCmdCopyImage2KHR(commandBuffer.Handle(), (*khr_copy_commands2_loader.VkCopyImageInfo2KHR)(copyInfo))
	return nil
}

func (e *VulkanExtensionDriver) CmdCopyImageToBuffer2(commandBuffer core1_0.CommandBuffer, copyImageToBufferInfo CopyImageToBufferInfo2) error {
	if !commandBuffer.Initialized() {
		return errors.New("commandBuffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	copyInfo, err := common.AllocOptions(arena, copyImageToBufferInfo)
	if err != nil {
		return err
	}

	e.driver.VkCmdCopyImageToBuffer2KHR(commandBuffer.Handle(), (*khr_copy_commands2_loader.VkCopyImageToBufferInfo2KHR)(copyInfo))
	return nil
}

func (e *VulkanExtensionDriver) CmdResolveImage2(commandBuffer core1_0.CommandBuffer, resolveImageInfo ResolveImageInfo2) error {
	if !commandBuffer.Initialized() {
		return errors.New("commandBuffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	resolveInfo, err := common.AllocOptions(arena, resolveImageInfo)
	if err != nil {
		return err
	}

	e.driver.VkCmdResolveImage2KHR(commandBuffer.Handle(), (*khr_copy_commands2_loader.VkResolveImageInfo2KHR)(resolveInfo))
	return nil
}
