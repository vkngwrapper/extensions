package ext_extended_dynamic_state

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	ext_extended_dynamic_state_loader "github.com/vkngwrapper/extensions/v3/ext_extended_dynamic_state/loader"
)

type VulkanExtensionDriver struct {
	driver ext_extended_dynamic_state_loader.Loader
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// ext_extended_dynamic_state loaded
func CreateExtensionDriverFromCoreDriver(coreDriver core1_0.DeviceDriver) ExtensionDriver {
	device := coreDriver.Device()

	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtensionDriver{
		driver: ext_extended_dynamic_state_loader.CreateLoaderFromCore(coreDriver.Loader()),
	}
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver ext_extended_dynamic_state_loader.Loader) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		driver: driver,
	}
}

func (v *VulkanExtensionDriver) CmdBindVertexBuffers2(commandBuffer core1_0.CommandBuffer, firstBinding int, buffers []core1_0.Buffer, offsets []int, sizes []int, strides []int) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}
	if len(buffers) != len(offsets) {
		panic(fmt.Sprintf("buffers has %d elements, but offsets has %d elements", len(buffers), len(offsets)))
	}
	if len(sizes) > 0 && len(sizes) != len(buffers) {
		panic(fmt.Sprintf("buffers has %d elements, but sizes has %d elements", len(buffers), len(sizes)))
	}
	if len(strides) > 0 && len(strides) != len(buffers) {
		panic(fmt.Sprintf("buffers has %d elements, but strides has %d elements", len(buffers), len(strides)))
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	bindingCount := len(buffers)
	var bufferArrayPtr *loader.VkBuffer = nil
	var offsetArrayPtr *loader.VkDeviceSize = nil
	var sizeArrayPtr *loader.VkDeviceSize = nil
	var strideArrayPtr *loader.VkDeviceSize = nil

	if bindingCount > 0 {
		bufferArrayUnsafe := arena.Malloc(bindingCount * int(unsafe.Sizeof([1]C.VkBuffer{})))
		offsetArrayUnsafe := arena.Malloc(bindingCount * int(unsafe.Sizeof(C.VkDeviceSize(0))))

		bufferArrayPtr = (*loader.VkBuffer)(bufferArrayUnsafe)
		offsetArrayPtr = (*loader.VkDeviceSize)(offsetArrayUnsafe)

		bufferArraySlice := ([]loader.VkBuffer)(unsafe.Slice(bufferArrayPtr, bindingCount))
		offsetArraySlice := ([]loader.VkDeviceSize)(unsafe.Slice(offsetArrayPtr, bindingCount))

		for i := 0; i < bindingCount; i++ {
			bufferArraySlice[i] = buffers[i].Handle()
			offsetArraySlice[i] = loader.VkDeviceSize(offsets[i])
		}
	}

	if len(sizes) > 0 {
		sizeArrayUnsafe := arena.Malloc(bindingCount * int(unsafe.Sizeof(C.VkDeviceSize(0))))
		sizeArrayPtr = (*loader.VkDeviceSize)(sizeArrayUnsafe)
		sizeArraySlice := ([]loader.VkDeviceSize)(unsafe.Slice(sizeArrayPtr, bindingCount))

		for i := 0; i < bindingCount; i++ {
			sizeArraySlice[i] = loader.VkDeviceSize(sizes[i])
		}
	}

	if len(strides) > 0 {
		strideArrayUnsafe := arena.Malloc(bindingCount * int(unsafe.Sizeof(C.VkDeviceSize(0))))
		strideArrayPtr = (*loader.VkDeviceSize)(strideArrayUnsafe)
		strideArraySlice := ([]loader.VkDeviceSize)(unsafe.Slice(strideArrayPtr, bindingCount))

		for i := 0; i < bindingCount; i++ {
			strideArraySlice[i] = loader.VkDeviceSize(strides[i])
		}
	}

	v.driver.VkCmdBindVertexBuffers2(commandBuffer.Handle(), loader.Uint32(firstBinding), loader.Uint32(bindingCount), bufferArrayPtr, offsetArrayPtr, sizeArrayPtr, strideArrayPtr)
}

func (v *VulkanExtensionDriver) CmdSetCullMode(commandBuffer core1_0.CommandBuffer, cullMode core1_0.CullModeFlags) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	v.driver.VkCmdSetCullMode(commandBuffer.Handle(), loader.VkCullModeFlags(cullMode))
}

func (v *VulkanExtensionDriver) CmdSetDepthBoundsTestEnable(commandBuffer core1_0.CommandBuffer, depthBoundsTestEnable bool) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	enable := loader.VkBool32(0)
	if depthBoundsTestEnable {
		enable = loader.VkBool32(1)
	}
	v.driver.VkCmdSetDepthBoundsTestEnable(commandBuffer.Handle(), enable)
}

func (v *VulkanExtensionDriver) CmdSetDepthCompareOp(commandBuffer core1_0.CommandBuffer, depthCompareOp core1_0.CompareOp) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	v.driver.VkCmdSetDepthCompareOp(commandBuffer.Handle(), loader.VkCompareOp(depthCompareOp))
}

func (v *VulkanExtensionDriver) CmdSetDepthTestEnable(commandBuffer core1_0.CommandBuffer, depthTestEnable bool) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	enable := loader.VkBool32(0)
	if depthTestEnable {
		enable = loader.VkBool32(1)
	}
	v.driver.VkCmdSetDepthTestEnable(commandBuffer.Handle(), enable)
}

func (v *VulkanExtensionDriver) CmdSetDepthWriteEnable(commandBuffer core1_0.CommandBuffer, depthWriteEnable bool) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	enable := loader.VkBool32(0)
	if depthWriteEnable {
		enable = loader.VkBool32(1)
	}
	v.driver.VkCmdSetDepthWriteEnable(commandBuffer.Handle(), enable)
}

func (v *VulkanExtensionDriver) CmdSetFrontFace(commandBuffer core1_0.CommandBuffer, frontFace core1_0.FrontFace) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	v.driver.VkCmdSetFrontFace(commandBuffer.Handle(), loader.VkFrontFace(frontFace))
}

func (v *VulkanExtensionDriver) CmdSetPrimitiveTopology(commandBuffer core1_0.CommandBuffer, primitiveTopology core1_0.PrimitiveTopology) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	v.driver.VkCmdSetPrimitiveTopology(commandBuffer.Handle(), loader.VkPrimitiveTopology(primitiveTopology))
}

func (v *VulkanExtensionDriver) CmdSetScissorWithCount(commandBuffer core1_0.CommandBuffer, scissors []core1_0.Rect2D) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	scissorCount := len(scissors)
	var scissorPtr *C.VkRect2D = nil

	if scissorCount > 0 {
		scissorUnsafe := arena.Malloc(scissorCount * int(unsafe.Sizeof(C.VkRect2D{})))
		scissorPtr = (*C.VkRect2D)(scissorUnsafe)
		scissorSlice := ([]C.VkRect2D)(unsafe.Slice(scissorPtr, scissorCount))

		for i := 0; i < scissorCount; i++ {
			scissorSlice[i].offset.x = C.int32_t(scissors[i].Offset.X)
			scissorSlice[i].offset.y = C.int32_t(scissors[i].Offset.Y)
			scissorSlice[i].extent.width = C.uint32_t(scissors[i].Extent.Width)
			scissorSlice[i].extent.height = C.uint32_t(scissors[i].Extent.Height)
		}
	}

	v.driver.VkCmdSetScissorWithCount(commandBuffer.Handle(), loader.Uint32(scissorCount), (*loader.VkRect2D)(unsafe.Pointer(scissorPtr)))
}

func (v *VulkanExtensionDriver) CmdSetStencilOp(commandBuffer core1_0.CommandBuffer, faceMask core1_0.StencilFaceFlags, failOp core1_0.StencilOp, passOp core1_0.StencilOp, depthFailOp core1_0.StencilOp, compareOp core1_0.CompareOp) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	v.driver.VkCmdSetStencilOp(commandBuffer.Handle(), loader.VkStencilFaceFlags(faceMask), loader.VkStencilOp(failOp), loader.VkStencilOp(passOp), loader.VkStencilOp(depthFailOp), loader.VkCompareOp(compareOp))
}

func (v *VulkanExtensionDriver) CmdSetStencilTestEnable(commandBuffer core1_0.CommandBuffer, stencilTestEnable bool) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	enable := loader.VkBool32(0)
	if stencilTestEnable {
		enable = loader.VkBool32(1)
	}

	v.driver.VkCmdSetStencilTestEnable(commandBuffer.Handle(), enable)
}

func (v *VulkanExtensionDriver) CmdSetViewportWithCount(commandBuffer core1_0.CommandBuffer, viewports []core1_0.Viewport) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	viewportCount := len(viewports)
	var viewportPtr *C.VkViewport = nil

	if viewportCount > 0 {
		viewportUnsafe := arena.Malloc(viewportCount * int(unsafe.Sizeof(C.VkViewport{})))
		viewportPtr = (*C.VkViewport)(viewportUnsafe)
		viewportSlice := ([]C.VkViewport)(unsafe.Slice(viewportPtr, viewportCount))

		for i := 0; i < viewportCount; i++ {
			viewportSlice[i].height = C.float(viewports[i].Height)
			viewportSlice[i].width = C.float(viewports[i].Width)
			viewportSlice[i].x = C.float(viewports[i].X)
			viewportSlice[i].y = C.float(viewports[i].Y)
			viewportSlice[i].minDepth = C.float(viewports[i].MinDepth)
			viewportSlice[i].maxDepth = C.float(viewports[i].MaxDepth)
		}
	}

	v.driver.VkCmdSetViewportWithCount(commandBuffer.Handle(), loader.Uint32(viewportCount), (*loader.VkViewport)(unsafe.Pointer(viewportPtr)))
}
