package khr_swapchain1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
	"github.com/vkngwrapper/extensions/khr_surface"
	"github.com/vkngwrapper/extensions/khr_swapchain"
	khr_swapchain_driver "github.com/vkngwrapper/extensions/khr_swapchain/driver"
	"unsafe"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	khr_swapchain.Extension

	driver khr_swapchain_driver.Driver
}

// PromoteExtension accepts a khr_swapchain.Extension object from core 1.0. If provided an Extension
// that supports at least core 1.1, it will return a core1_1.Extension. Otherwise, it will return nil.
func PromoteExtension(extension khr_swapchain.Extension) *VulkanExtension {
	if extension == nil || !extension.APIVersion().IsAtLeast(common.Vulkan1_1) {
		return nil
	}

	return &VulkanExtension{
		Extension: extension,

		driver: extension.Driver(),
	}
}

// CreateExtensionFromDriver generates an Extension from a driver.Driver object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_swapchain_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		Extension: khr_swapchain.CreateExtensionFromDriver(driver),
		driver:    driver,
	}
}

func (v *VulkanExtension) AcquireNextImage2(device core1_0.Device, o AcquireNextImageInfo) (int, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return -1, core1_0.VKErrorUnknown, err
	}

	indexPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := v.driver.VkAcquireNextImage2KHR(
		device.Handle(),
		(*khr_swapchain_driver.VkAcquireNextImageInfoKHR)(optionPtr),
		indexPtr,
	)
	if err != nil {
		return -1, res, err
	}

	return int(*indexPtr), res, nil
}

func (v *VulkanExtension) DeviceGroupPresentCapabilities(device core1_0.Device, outData *DeviceGroupPresentCapabilities) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOutDataHeader(arena, outData)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	res, err := v.driver.VkGetDeviceGroupPresentCapabilitiesKHR(
		device.Handle(),
		(*khr_swapchain_driver.VkDeviceGroupPresentCapabilitiesKHR)(optionPtr),
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

func (v *VulkanExtension) DeviceGroupSurfacePresentModes(device core1_0.Device, surface khr_surface.Surface) (DeviceGroupPresentModeFlags, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	flagsPtr := (*khr_swapchain_driver.VkDeviceGroupPresentModeFlagsKHR)(arena.Malloc(int(unsafe.Sizeof(C.VkDeviceGroupPresentModeFlagsKHR(0)))))

	res, err := v.driver.VkGetDeviceGroupSurfacePresentModesKHR(
		device.Handle(),
		surface.Handle(),
		flagsPtr,
	)
	if err != nil {
		return 0, res, err
	}

	return DeviceGroupPresentModeFlags(*flagsPtr), res, nil
}

func (v *VulkanExtension) attemptGetPhysicalDevicePresentRectangles(physicalDevice core1_0.PhysicalDevice, surface khr_surface.Surface) ([]core1_0.Rect2D, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	countPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

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
		(*driver.Uint32)(countPtr),
		(*driver.VkRect2D)(rectsPtr),
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

func (v *VulkanExtension) PhysicalDevicePresentRectangles(physicalDevice core1_0.PhysicalDevice, surface khr_surface.Surface) ([]core1_0.Rect2D, common.VkResult, error) {
	var outData []core1_0.Rect2D
	var result common.VkResult
	var err error

	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		outData, result, err = v.attemptGetPhysicalDevicePresentRectangles(physicalDevice, surface)
	}
	return outData, result, err
}
