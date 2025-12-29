package khr_surface

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"unsafe"

	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/extensions/v3/khr_surface/loader"
)

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from an Insstance with
// khr_surface loaded
func CreateExtensionDriverFromCoreDriver(coreDriver core1_0.CoreInstanceDriver) ExtensionDriver {
	instance := coreDriver.Instance()

	if !instance.IsInstanceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtensionDriver{
		driver:   khr_surface_loader.CreateDriverFromCore(coreDriver.Loader()),
		instance: instance,
	}
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver khr_surface_loader.Loader, instance core.Instance) *VulkanExtensionDriver {
	if !instance.Initialized() {
		panic("instance cannot be uninitialized")
	}
	return &VulkanExtensionDriver{
		driver:   driver,
		instance: instance,
	}
}

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	driver   khr_surface_loader.Loader
	instance core.Instance
}

func (e *VulkanExtensionDriver) CreateSurfaceFromHandle(surfaceHandle khr_surface_loader.VkSurfaceKHR) (Surface, error) {
	instanceHandle := e.instance.Handle()
	apiVersion := e.instance.APIVersion()

	surface := Surface{
		handle:     surfaceHandle,
		instance:   instanceHandle,
		apiVersion: apiVersion,
	}
	return surface, nil
}

func (e *VulkanExtensionDriver) DestroySurface(surface Surface, callbacks *loader.AllocationCallbacks) {
	if !surface.Initialized() {
		panic("surface cannot be uninitialized")
	}
	e.driver.VkDestroySurfaceKHR(surface.InstanceHandle(), surface.Handle(), callbacks.Handle())
}

func (e *VulkanExtensionDriver) GetPhysicalDeviceSurfaceSupport(surface Surface, physicalDevice core.PhysicalDevice, queueFamilyIndex int) (bool, common.VkResult, error) {
	if !surface.Initialized() {
		panic("surface cannot be uninitialized")
	}
	if !physicalDevice.Initialized() {
		panic("physicalDevice cannot be uninitialized")
	}
	var canPresent loader.VkBool32

	res, err := e.driver.VkGetPhysicalDeviceSurfaceSupportKHR(physicalDevice.Handle(), loader.Uint32(queueFamilyIndex), surface.Handle(), &canPresent)

	return canPresent != C.VK_FALSE, res, err
}

func (e *VulkanExtensionDriver) GetPhysicalDeviceSurfaceCapabilities(surface Surface, device core.PhysicalDevice) (*SurfaceCapabilities, common.VkResult, error) {
	if !surface.Initialized() {
		panic("surface cannot be uninitialized")
	}
	if !device.Initialized() {
		panic("device cannot be uninitialized")
	}
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	capabilitiesPtr := allocator.Malloc(int(unsafe.Sizeof([1]C.VkSurfaceCapabilitiesKHR{})))
	cCapabilities := (*C.VkSurfaceCapabilitiesKHR)(capabilitiesPtr)

	res, err := e.driver.VkGetPhysicalDeviceSurfaceCapabilitiesKHR(device.Handle(), surface.Handle(), (*khr_surface_loader.VkSurfaceCapabilitiesKHR)(unsafe.Pointer(cCapabilities)))
	if err != nil {
		return nil, res, err
	}

	return &SurfaceCapabilities{
		MinImageCount: int(cCapabilities.minImageCount),
		MaxImageCount: int(cCapabilities.maxImageCount),
		CurrentExtent: core1_0.Extent2D{
			Width:  int(cCapabilities.currentExtent.width),
			Height: int(cCapabilities.currentExtent.height),
		},
		MinImageExtent: core1_0.Extent2D{
			Width:  int(cCapabilities.minImageExtent.width),
			Height: int(cCapabilities.minImageExtent.height),
		},
		MaxImageExtent: core1_0.Extent2D{
			Width:  int(cCapabilities.maxImageExtent.width),
			Height: int(cCapabilities.maxImageExtent.height),
		},
		MaxImageArrayLayers: int(cCapabilities.maxImageArrayLayers),

		SupportedTransforms: SurfaceTransformFlags(cCapabilities.supportedTransforms),
		CurrentTransform:    SurfaceTransformFlags(cCapabilities.currentTransform),

		SupportedCompositeAlpha: CompositeAlphaFlags(cCapabilities.supportedCompositeAlpha),
		SupportedUsageFlags:     core1_0.ImageUsageFlags(cCapabilities.supportedUsageFlags),
	}, res, nil
}

func (e *VulkanExtensionDriver) attemptFormats(surface Surface, device core.PhysicalDevice) ([]SurfaceFormat, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	formatCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	formatCount := (*loader.Uint32)(formatCountPtr)

	res, err := e.driver.VkGetPhysicalDeviceSurfaceFormatsKHR(device.Handle(), surface.Handle(), formatCount, nil)
	if err != nil {
		return nil, res, err
	}

	count := int(*formatCount)

	if count == 0 {
		return nil, res, nil
	}

	formatsPtr := allocator.Malloc(count * int(unsafe.Sizeof([1]C.VkSurfaceFormatKHR{})))

	res, err = e.driver.VkGetPhysicalDeviceSurfaceFormatsKHR(device.Handle(), surface.Handle(), formatCount, (*khr_surface_loader.VkSurfaceFormatKHR)(formatsPtr))
	if err != nil || res == core1_0.VKIncomplete {
		return nil, res, err
	}

	formatSlice := ([]C.VkSurfaceFormatKHR)(unsafe.Slice((*C.VkSurfaceFormatKHR)(formatsPtr), count))
	var result []SurfaceFormat
	for i := 0; i < count; i++ {
		result = append(result, SurfaceFormat{
			Format:     core1_0.Format(formatSlice[i].format),
			ColorSpace: ColorSpace(formatSlice[i].colorSpace),
		})
	}

	return result, res, nil
}

func (e *VulkanExtensionDriver) GetPhysicalDeviceSurfaceFormats(surface Surface, device core.PhysicalDevice) ([]SurfaceFormat, common.VkResult, error) {
	if !surface.Initialized() {
		panic("surface cannot be uninitialized")
	}
	if !device.Initialized() {
		panic("device cannot be uninitialized")
	}
	var formats []SurfaceFormat
	var result common.VkResult
	var err error
	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		formats, result, err = e.attemptFormats(surface, device)
	}

	return formats, result, err
}

func (e *VulkanExtensionDriver) attemptPresentModes(surface Surface, device core.PhysicalDevice) ([]PresentMode, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	modeCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	modeCount := (*loader.Uint32)(modeCountPtr)

	res, err := e.driver.VkGetPhysicalDeviceSurfacePresentModesKHR(device.Handle(), surface.Handle(), modeCount, nil)
	if err != nil {
		return nil, res, err
	}

	count := int(*modeCount)
	if count == 0 {
		return nil, res, nil
	}

	modesPtr := allocator.Malloc(count * int(unsafe.Sizeof(C.VkPresentModeKHR(0))))
	presentModes := (*khr_surface_loader.VkPresentModeKHR)(modesPtr)

	res, err = e.driver.VkGetPhysicalDeviceSurfacePresentModesKHR(device.Handle(), surface.Handle(), modeCount, presentModes)
	if err != nil || res == core1_0.VKIncomplete {
		return nil, res, err
	}

	presentModeSlice := ([]khr_surface_loader.VkPresentModeKHR)(unsafe.Slice(presentModes, count))
	var result []PresentMode
	for i := 0; i < count; i++ {
		result = append(result, PresentMode(presentModeSlice[i]))
	}

	return result, res, nil
}

func (e *VulkanExtensionDriver) GetPhysicalDeviceSurfacePresentModes(surface Surface, device core.PhysicalDevice) ([]PresentMode, common.VkResult, error) {
	if !surface.Initialized() {
		panic("surface cannot be uninitialized")
	}
	if !device.Initialized() {
		panic("device cannot be uninitialized")
	}
	var presentModes []PresentMode
	var result common.VkResult
	var err error
	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		presentModes, result, err = e.attemptPresentModes(surface, device)
	}

	return presentModes, result, err
}
