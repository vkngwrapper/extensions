package khr_device_group

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	khr_device_group_driver "github.com/vkngwrapper/extensions/v3/khr_device_group/loader"
	"github.com/vkngwrapper/extensions/v3/khr_surface"
	"github.com/vkngwrapper/extensions/v3/khr_swapchain"
)

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	driver        khr_device_group_driver.Loader
	device        core1_0.Device
	withSurface   *VulkanExtensionDriverWithKHRSurface
	withSwapchain *VulkanExtensionDriverWithKHRSwapchain
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// khr_device_group loaded
func CreateExtensionDriverFromCoreDriver(coreDriver core1_0.DeviceDriver, instance core1_0.Instance) ExtensionDriver {
	device := coreDriver.Device()

	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	surfaceInteraction := instance.IsInstanceExtensionActive(khr_surface.ExtensionName)
	swapchainInteraction := device.IsDeviceExtensionActive(khr_swapchain.ExtensionName)

	return CreateExtensionDriverFromLoader(khr_device_group_driver.CreateLoaderFromCore(coreDriver.Loader()), device, surfaceInteraction, swapchainInteraction)
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver khr_device_group_driver.Loader, device core1_0.Device, khrSurfaceInteraction bool, khrSwapchainInteraction bool) *VulkanExtensionDriver {
	ext := &VulkanExtensionDriver{
		driver: driver,
		device: device,
	}

	if khrSurfaceInteraction {
		ext.withSurface = &VulkanExtensionDriverWithKHRSurface{driver: driver, device: device}
	}

	if khrSwapchainInteraction {
		ext.withSwapchain = &VulkanExtensionDriverWithKHRSwapchain{driver: driver, device: device}
	}

	return ext
}

// WithKHRSurface returns nil if khr_surface is not currently active, or
func (v *VulkanExtensionDriver) WithKHRSurface() ExtensionDriverWithKHRSurface {
	return v.withSurface
}

func (v *VulkanExtensionDriver) WithKHRSwapchain() ExtensionDriverWithKHRSwapchain {
	return v.withSwapchain
}

func (v *VulkanExtensionDriver) CmdDispatchBase(commandBuffer core1_0.CommandBuffer, baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ int) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}
	v.driver.VkCmdDispatchBaseKHR(commandBuffer.Handle(),
		loader.Uint32(baseGroupX),
		loader.Uint32(baseGroupY),
		loader.Uint32(baseGroupZ),
		loader.Uint32(groupCountX),
		loader.Uint32(groupCountY),
		loader.Uint32(groupCountZ))
}

func (v *VulkanExtensionDriver) CmdSetDeviceMask(commandBuffer core1_0.CommandBuffer, deviceMask uint32) {
	if !commandBuffer.Initialized() {
		panic("commandBuffer cannot be uninitialized")
	}
	v.driver.VkCmdSetDeviceMaskKHR(commandBuffer.Handle(), loader.Uint32(deviceMask))
}

func (v *VulkanExtensionDriver) GetDeviceGroupPeerMemoryFeatures(heapIndex, localDeviceIndex, remoteDeviceIndex int) PeerMemoryFeatureFlags {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	featuresPtr := (*khr_device_group_driver.VkPeerMemoryFeatureFlagsKHR)(arena.Malloc(int(unsafe.Sizeof(C.VkPeerMemoryFeatureFlagsKHR(0)))))

	v.driver.VkGetDeviceGroupPeerMemoryFeaturesKHR(
		v.device.Handle(),
		loader.Uint32(heapIndex),
		loader.Uint32(localDeviceIndex),
		loader.Uint32(remoteDeviceIndex),
		featuresPtr,
	)

	return PeerMemoryFeatureFlags(*featuresPtr)
}

type VulkanExtensionDriverWithKHRSurface struct {
	driver khr_device_group_driver.Loader
	device core1_0.Device
}

func (v *VulkanExtensionDriverWithKHRSurface) GetDeviceGroupPresentCapabilities(outData *DeviceGroupPresentCapabilities) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOutDataHeader(arena, outData)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	res, err := v.driver.VkGetDeviceGroupPresentCapabilitiesKHR(
		v.device.Handle(),
		(*khr_device_group_driver.VkDeviceGroupPresentCapabilitiesKHR)(optionPtr),
	)
	if err != nil {
		return res, err
	}

	err = common.PopulateOutData(outData, optionPtr)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return res, nil
}

func (v *VulkanExtensionDriverWithKHRSurface) GetDeviceGroupSurfacePresentModes(surface khr_surface.Surface) (DeviceGroupPresentModeFlags, common.VkResult, error) {
	if !surface.Initialized() {
		panic("surface cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	flagsPtr := (*khr_device_group_driver.VkDeviceGroupPresentModeFlagsKHR)(arena.Malloc(int(unsafe.Sizeof(C.VkDeviceGroupPresentModeFlagsKHR(0)))))

	res, err := v.driver.VkGetDeviceGroupSurfacePresentModesKHR(
		v.device.Handle(),
		surface.Handle(),
		flagsPtr,
	)
	if err != nil {
		return 0, res, err
	}

	return DeviceGroupPresentModeFlags(*flagsPtr), res, nil
}

func (v *VulkanExtensionDriverWithKHRSurface) attemptGetPhysicalDevicePresentRectangles(physicalDevice core1_0.PhysicalDevice, surface khr_surface.Surface) ([]core1_0.Rect2D, common.VkResult, error) {
	if !physicalDevice.Initialized() {
		panic("physicalDevice cannot be uninitialized")
	}
	if !surface.Initialized() {
		panic("surface cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	countPtr := (*loader.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := v.driver.VkGetPhysicalDevicePresentRectanglesKHR(
		physicalDevice.Handle(),
		surface.Handle(),
		countPtr,
		nil,
	)
	if err != nil {
		return nil, res, err
	}

	count := int(*countPtr)
	if count == 0 {
		return nil, res, nil
	}

	rectsPtr := arena.Malloc(count * C.sizeof_struct_VkRect2D)
	res, err = v.driver.VkGetPhysicalDevicePresentRectanglesKHR(
		physicalDevice.Handle(),
		surface.Handle(),
		(*loader.Uint32)(countPtr),
		(*loader.VkRect2D)(rectsPtr),
	)
	if res != core1_0.VKSuccess {
		return nil, res, err
	}

	rectsSlice := ([]C.VkRect2D)(unsafe.Slice((*C.VkRect2D)(rectsPtr), count))
	outRects := make([]core1_0.Rect2D, count)
	for i := 0; i < count; i++ {
		outRects[i].Offset.X = int(rectsSlice[i].offset.x)
		outRects[i].Offset.Y = int(rectsSlice[i].offset.y)
		outRects[i].Extent.Width = int(rectsSlice[i].extent.width)
		outRects[i].Extent.Height = int(rectsSlice[i].extent.height)
	}

	return outRects, res, nil
}

func (v *VulkanExtensionDriverWithKHRSurface) GetPhysicalDevicePresentRectangles(physicalDevice core1_0.PhysicalDevice, surface khr_surface.Surface) ([]core1_0.Rect2D, common.VkResult, error) {
	if !physicalDevice.Initialized() {
		return nil, core1_0.VKErrorUnknown, fmt.Errorf("physicalDevice cannot be uninitialized")
	}
	var outData []core1_0.Rect2D
	var result common.VkResult
	var err error

	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		outData, result, err = v.attemptGetPhysicalDevicePresentRectangles(physicalDevice, surface)
	}
	return outData, result, err
}

type VulkanExtensionDriverWithKHRSwapchain struct {
	driver khr_device_group_driver.Loader
	device core1_0.Device
}

func (v *VulkanExtensionDriverWithKHRSwapchain) AcquireNextImage2(o AcquireNextImageInfo) (int, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return -1, core1_0.VKErrorUnknown, err
	}

	indexPtr := (*loader.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := v.driver.VkAcquireNextImage2KHR(
		v.device.Handle(),
		(*khr_device_group_driver.VkAcquireNextImageInfoKHR)(optionPtr),
		indexPtr,
	)
	if err != nil {
		return -1, res, err
	}

	return int(*indexPtr), res, nil
}
