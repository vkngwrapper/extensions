package ext_extended_dynamic_state2

import (
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	ext_extended_dynamic_state2_loader "github.com/vkngwrapper/extensions/v3/ext_extended_dynamic_state2/loader"
)

type VulkanExtensionDriver struct {
	driver ext_extended_dynamic_state2_loader.Loader
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// ext_extended_dynamic_state2 loaded
func CreateExtensionDriverFromCoreDriver(coreDriver core1_0.DeviceDriver) ExtensionDriver {
	device := coreDriver.Device()

	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtensionDriver{
		driver: ext_extended_dynamic_state2_loader.CreateLoaderFromCore(coreDriver.Loader()),
	}
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver ext_extended_dynamic_state2_loader.Loader) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		driver: driver,
	}
}

func (v *VulkanExtensionDriver) CmdSetDepthBiasEnable(commandBuffer core1_0.CommandBuffer, depthBiasEnable bool) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	enable := loader.VkBool32(0)
	if depthBiasEnable {
		enable = loader.VkBool32(1)
	}
	v.driver.VkCmdSetDepthBiasEnableEXT(commandBuffer.Handle(), enable)
}

func (v *VulkanExtensionDriver) CmdSetLogicOp(commandBuffer core1_0.CommandBuffer, logicOp core1_0.LogicOp) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	v.driver.VkCmdSetLogicOpEXT(commandBuffer.Handle(), loader.VkLogicOp(logicOp))
}

func (v *VulkanExtensionDriver) CmdSetPatchControlPoints(commandBuffer core1_0.CommandBuffer, patchControlPoints int) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	v.driver.VkCmdSetPatchControlPointsEXT(commandBuffer.Handle(), loader.Uint32(patchControlPoints))
}

func (v *VulkanExtensionDriver) CmdSetPrimitiveRestartEnable(commandBuffer core1_0.CommandBuffer, primitiveRestartEnable bool) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	enable := loader.VkBool32(0)
	if primitiveRestartEnable {
		enable = loader.VkBool32(1)
	}

	v.driver.VkCmdSetPrimitiveRestartEnableEXT(commandBuffer.Handle(), enable)
}

func (v *VulkanExtensionDriver) CmdSetRasterizerDiscardEnable(commandBuffer core1_0.CommandBuffer, rasterizerDiscardEnable bool) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	enable := loader.VkBool32(0)
	if rasterizerDiscardEnable {
		enable = loader.VkBool32(1)
	}

	v.driver.VkCmdSetRasterizerDiscardEnableEXT(commandBuffer.Handle(), enable)
}
