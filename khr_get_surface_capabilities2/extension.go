package khr_get_surface_capabilities2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	khr_get_surface_capabilities2_loader "github.com/vkngwrapper/extensions/v3/khr_get_surface_capabilities2/loader"
)

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	loader khr_get_surface_capabilities2_loader.Loader
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from an Instance with
// khr_get_physical_device_properties2 loaded
func CreateExtensionDriverFromCoreDriver(coreDriver core1_0.CoreInstanceDriver) ExtensionDriver {
	instance := coreDriver.Instance()

	if !instance.IsInstanceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtensionDriver{
		loader: khr_get_surface_capabilities2_loader.CreateLoaderFromCore(coreDriver.Loader()),
	}
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver khr_get_surface_capabilities2_loader.Loader) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		loader: driver,
	}
}

func (e *VulkanExtensionDriver) GetPhysicalDeviceSurfaceCapabilities2(physicalDevice core1_0.PhysicalDevice, o PhysicalDeviceSurfaceInfo2, out *SurfaceCapabilities2) (common.VkResult, error) {
	if !physicalDevice.Initialized() {
		panic("physicalDevice cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionData, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	outData, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	res, err := e.loader.VkGetPhysicalDeviceSurfaceCapabilities2KHR(physicalDevice.Handle(),
		(*khr_get_surface_capabilities2_loader.VkPhysicalDeviceSurfaceInfo2KHR)(optionData),
		(*khr_get_surface_capabilities2_loader.VkSurfaceCapabilities2KHR)(outData))
	if err != nil {
		return res, err
	}

	err = common.PopulateOutData(out, outData)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return res, nil
}

func (e *VulkanExtensionDriver) GetPhysicalDeviceSurfaceFormats2(physicalDevice core1_0.PhysicalDevice, o PhysicalDeviceSurfaceInfo2, outDataFactory func() *SurfaceFormat2) ([]*SurfaceFormat2, common.VkResult, error) {
	if !physicalDevice.Initialized() {
		panic("physicalDevice cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionData, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	outDataCountPtr := (*loader.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	res, err := e.loader.VkGetPhysicalDeviceSurfaceFormats2KHR(physicalDevice.Handle(),
		(*khr_get_surface_capabilities2_loader.VkPhysicalDeviceSurfaceInfo2KHR)(optionData),
		(*loader.Uint32)(outDataCountPtr),
		nil)
	if err != nil {
		return nil, res, err
	}

	outDataCount := int(*outDataCountPtr)
	if outDataCount == 0 {
		return nil, core1_0.VKSuccess, nil
	}

	out := common.InitSlice[SurfaceFormat2](outDataCount, outDataFactory)
	outData, err := common.AllocOutDataHeaderSlice[C.VkSurfaceFormat2KHR, *SurfaceFormat2](arena, out)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	res, err = e.loader.VkGetPhysicalDeviceSurfaceFormats2KHR(physicalDevice.Handle(),
		(*khr_get_surface_capabilities2_loader.VkPhysicalDeviceSurfaceInfo2KHR)(optionData),
		(*loader.Uint32)(outDataCountPtr),
		(*khr_get_surface_capabilities2_loader.VkSurfaceFormat2KHR)(unsafe.Pointer(outData)))
	if err != nil {
		return nil, res, err
	}

	err = common.PopulateOutDataSlice[C.VkSurfaceFormat2KHR, *SurfaceFormat2](out, unsafe.Pointer(outData))
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}
	return out, res, nil
}
