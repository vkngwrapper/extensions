//go:build windows

package ext_full_screen_exclusive

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	ext_full_screen_exclusive_loader "github.com/vkngwrapper/extensions/v3/ext_full_screen_exclusive/loader"
	"github.com/vkngwrapper/extensions/v3/khr_device_group"
	khr_device_group_loader "github.com/vkngwrapper/extensions/v3/khr_device_group/loader"
	"github.com/vkngwrapper/extensions/v3/khr_get_physical_device_properties2"
	"github.com/vkngwrapper/extensions/v3/khr_get_surface_capabilities2"
	khr_get_surface_capabilities2_loader "github.com/vkngwrapper/extensions/v3/khr_get_surface_capabilities2/loader"
	"github.com/vkngwrapper/extensions/v3/khr_surface"
	khr_surface_loader "github.com/vkngwrapper/extensions/v3/khr_surface/loader"
	"github.com/vkngwrapper/extensions/v3/khr_swapchain"
)

type VulkanExtensionDriverCombined struct {
	VulkanExtensionDriver
	VulkanExtensionDriverWithDeviceGroups
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// khr_device_group loaded
func CreateExtensionDriverFromCoreDriver(coreDriver core1_0.DeviceDriver, instance core1_0.Instance) ExtensionDriver {
	device := coreDriver.Device()

	if !device.IsDeviceExtensionActive(ExtensionName) ||
		!device.IsDeviceExtensionActive(khr_swapchain.ExtensionName) ||
		!instance.IsInstanceExtensionActive(khr_surface.ExtensionName) ||
		!instance.IsInstanceExtensionActive(khr_get_surface_capabilities2.ExtensionName) {
		return nil
	}

	if !instance.APIVersion().IsAtLeast(common.Vulkan1_1) &&
		!instance.IsInstanceExtensionActive(khr_get_physical_device_properties2.ExtensionName) {
		return nil
	}

	deviceGroupInteraction := device.IsDeviceExtensionActive(khr_device_group.ExtensionName)

	return CreateExtensionDriverFromLoader(ext_full_screen_exclusive_loader.CreateLoaderFromCore(coreDriver.Loader()), device, deviceGroupInteraction)
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(loader ext_full_screen_exclusive_loader.Loader, device core1_0.Device, deviceGroupInteraction bool) ExtensionDriver {

	if deviceGroupInteraction {
		return &VulkanExtensionDriverCombined{
			VulkanExtensionDriver{
				loader: loader,
				device: device,
			},
			VulkanExtensionDriverWithDeviceGroups{
				loader: loader,
				device: device,
			},
		}
	}

	return &VulkanExtensionDriver{
		loader: loader,
		device: device,
	}
}

type VulkanExtensionDriver struct {
	loader ext_full_screen_exclusive_loader.Loader
	device core1_0.Device
}

func (v *VulkanExtensionDriver) AcquireFullScreenExclusiveMode(swapchain khr_swapchain.Swapchain) (common.VkResult, error) {
	if !swapchain.Initialized() {
		return core1_0.VKErrorUnknown, errors.New("swapchain cannot be uninitialized")
	}

	return v.loader.VkAcquireFullScreenExclusiveModeEXT(
		v.device.Handle(),
		swapchain.Handle(),
	)
}

func (v *VulkanExtensionDriver) GetPhysicalDeviceSurfacePresentModes2(physicalDevice core1_0.PhysicalDevice, o khr_get_surface_capabilities2.PhysicalDeviceSurfaceInfo2) ([]khr_surface.PresentMode, common.VkResult, error) {
	if !physicalDevice.Initialized() {
		return nil, core1_0.VKErrorUnknown, errors.New("physicalDevice cannot be uninitialized")
	}

	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionData, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	outDataCountPtr := (*loader.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := v.loader.VkGetPhysicalDeviceSurfacePresentModes2EXT(
		physicalDevice.Handle(),
		(*khr_get_surface_capabilities2_loader.VkPhysicalDeviceSurfaceInfo2KHR)(optionData),
		outDataCountPtr,
		nil,
	)
	if err != nil {
		return nil, res, err
	}

	count := int(*outDataCountPtr)
	if count == 0 {
		return nil, core1_0.VKSuccess, nil
	}

	retValSlice := make([]khr_surface.PresentMode, count)
	outDataPtr := arena.Malloc(int(unsafe.Sizeof(khr_surface_loader.VkPresentModeKHR(0))) * count)

	res, err = v.loader.VkGetPhysicalDeviceSurfacePresentModes2EXT(
		physicalDevice.Handle(),
		(*khr_get_surface_capabilities2_loader.VkPhysicalDeviceSurfaceInfo2KHR)(optionData),
		outDataCountPtr,
		(*khr_surface_loader.VkPresentModeKHR)(outDataPtr),
	)
	if err != nil {
		return nil, res, err
	}

	outDataSlice := unsafe.Slice((*C.VkPresentModeKHR)(outDataPtr), count)
	for retValIndex := range retValSlice {
		retValSlice[retValIndex] = khr_surface.PresentMode(outDataSlice[retValIndex])
	}

	return retValSlice, res, nil
}

func (v *VulkanExtensionDriver) ReleaseFullScreenExclusiveMode(swapchain khr_swapchain.Swapchain) (common.VkResult, error) {
	if !swapchain.Initialized() {
		return core1_0.VKErrorUnknown, errors.New("swapchain cannot be uninitialized")
	}

	return v.loader.VkReleaseFullScreenExclusiveModeEXT(v.device.Handle(), swapchain.Handle())
}

type VulkanExtensionDriverWithDeviceGroups struct {
	loader ext_full_screen_exclusive_loader.Loader

	device core1_0.Device
}

func (v *VulkanExtensionDriverWithDeviceGroups) GetDeviceGroupSurfacePresentModes2(o khr_get_surface_capabilities2.PhysicalDeviceSurfaceInfo2, out *khr_device_group.DeviceGroupPresentModeFlags) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionData, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	outDataPtr := arena.Malloc(int(unsafe.Sizeof(khr_device_group_loader.VkDeviceGroupPresentModeFlagsKHR(0))))
	flagsPtr := (*khr_device_group_loader.VkDeviceGroupPresentModeFlagsKHR)(outDataPtr)

	res, err := v.loader.VkGetDeviceGroupSurfacePresentModes2EXT(
		v.device.Handle(),
		(*khr_get_surface_capabilities2_loader.VkPhysicalDeviceSurfaceInfo2KHR)(optionData),
		flagsPtr,
	)
	if err != nil {
		return res, err
	}

	*out = khr_device_group.DeviceGroupPresentModeFlags(*flagsPtr)
	return res, nil
}
