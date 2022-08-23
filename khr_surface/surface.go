package khr_surface

//go:generate mockgen -source surface.go -destination ./mocks/surface.go -package mock_surface

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	ext_driver "github.com/vkngwrapper/extensions/v2/khr_surface/driver"
	"unsafe"
)

// SurfaceCapabilities describes capabilities of a Surface
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesKHR.html
type SurfaceCapabilities struct {
	// MinImageCount is the minimum number of Image objects the specified Device supports for a khr_swapchain.Swapchain
	// created for the Surface, and will be at least one
	MinImageCount int
	// MaxImageCount is the maximum number of Image objects the specified Device supports for a
	// khr_swapchain.Swapchain created for the Surface, and will either be 0, or greater than or equal to
	// MinImageCount. A value of 0 means that there is no limit on the number of Image objects, though
	// there may be limits on memory used
	MaxImageCount int

	// CurrentExtent is the current width and height of the Surface, or the special values -1, -1 indicating
	// that the Surface size will be determined by the extent of a khr_swapchain.Swapchain targeting the
	// Surface
	CurrentExtent core1_0.Extent2D
	// MinImageExtent contains the smallest valid khr_swapchain.Swapchain extent for the Surface on the
	// specified Device
	MinImageExtent core1_0.Extent2D
	// MaxImageExtent contains the largest valid khr_swapchain.Swapchain extent for the Surface on the
	// specified Device
	MaxImageExtent core1_0.Extent2D

	// MaxImageArrayLayers is the maximum number of layers presentable Image objects can have for a
	// khr_swapchain.Swapchain created for this Device and Surface, and will be at least 1
	MaxImageArrayLayers int
	// SupportedTransforms indicates the presentation transforms supported for the Surface on the
	// specified Device. At least one bit will be set
	SupportedTransforms SurfaceTransformFlags
	// CurrentTransform indicates the Surface object's current transform relative to the presentation
	// engine's natural orientation
	CurrentTransform SurfaceTransformFlags

	// SupportedCompositeAlpha represents the alpha compositing modes supported by the presentation
	// engine for the Surface on the specified Device, and at least one will be set
	SupportedCompositeAlpha CompositeAlphaFlags
	// SupportedUsageFlags represents the ways the application can use the presentable Image objects
	// of a khr_swapchain.Swapchain created with PresentMode set to PresentModeImmediate, PresentModeMailbox,
	// PresentModeFIFO, or PresentModeFIFORelaxed for the Surface on the specified Device
	SupportedUsageFlags core1_0.ImageUsageFlags
}

// SurfaceFormat describes a supported khr_swapchain.Swapchain format-color space pair
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFormatKHR.html
type SurfaceFormat struct {
	// Format is a core1_0.Format compatible with the specified Surface
	Format core1_0.Format
	// ColorSpace is a presentation ColorSpace that is compatible with the Surface
	ColorSpace ColorSpace
}

// VulkanSurface is an implementation of the Surface interface that actually communicates
// with Vulkan. This is the default implementation. See the interface for more documentation.
type VulkanSurface struct {
	instance   driver.VkInstance
	handle     ext_driver.VkSurfaceKHR
	driver     ext_driver.Driver
	coreDriver driver.Driver

	minimumAPIVersion common.APIVersion
}

// Surface abstracts native platform surface or window objects
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html
type Surface interface {
	// Handle is the internal Vulkan object handle for this Surface
	Handle() ext_driver.VkSurfaceKHR

	// Destroy deletes this Surface and underlying structures from the device. **Warning**
	// after destruction, this object will still exist, but the Vulkan object handle
	// that backs it will be invalid. Do not call further methods on this object.
	//
	// callbacks - A set of allocation callbacks to control the memory free behavior of this command
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySurfaceKHR.html
	Destroy(callbacks *driver.AllocationCallbacks)
	// PhysicalDeviceSurfaceSupport queries if presentation of this Surface is supported on the specified PhysicalDevice
	//
	// physicalDevice - The PhysicalDevice to query for support
	//
	// queueFamilyIndex - The Queue family to be used to present the Surface
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html
	PhysicalDeviceSurfaceSupport(physicalDevice core1_0.PhysicalDevice, queueFamilyIndex int) (bool, common.VkResult, error)
	// PhysicalDeviceSurfaceCapabilities queries Surface capabilities on the specified PhysicalDevice
	//
	// device - The PhysicalDevice to query for capabilities
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html
	PhysicalDeviceSurfaceCapabilities(device core1_0.PhysicalDevice) (*SurfaceCapabilities, common.VkResult, error)
	// PhysicalDeviceSurfaceFormats queries color formats supported by Surface on the specified PhysicalDevice
	//
	// device - The PhysicalDevice to query for supported Surface formats
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceFormatsKHR.html
	PhysicalDeviceSurfaceFormats(device core1_0.PhysicalDevice) ([]SurfaceFormat, common.VkResult, error)
	// PhysicalDeviceSurfacePresentModes queries supported presentation modes on the specified PhysicalDevice
	//
	// device - The PhysicalDevice to query for supported presentation modes
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfacePresentModesKHR.html
	PhysicalDeviceSurfacePresentModes(device core1_0.PhysicalDevice) ([]PresentMode, common.VkResult, error)
}

func (s *VulkanSurface) Handle() ext_driver.VkSurfaceKHR {
	return s.handle
}

func (s *VulkanSurface) Destroy(callbacks *driver.AllocationCallbacks) {
	s.driver.VkDestroySurfaceKHR(s.instance, s.handle, callbacks.Handle())
	s.coreDriver.ObjectStore().Delete(driver.VulkanHandle(s.handle))
}

func (s *VulkanSurface) PhysicalDeviceSurfaceSupport(physicalDevice core1_0.PhysicalDevice, queueFamilyIndex int) (bool, common.VkResult, error) {
	var canPresent driver.VkBool32

	res, err := s.driver.VkGetPhysicalDeviceSurfaceSupportKHR(physicalDevice.Handle(), driver.Uint32(queueFamilyIndex), s.handle, &canPresent)

	return canPresent != C.VK_FALSE, res, err
}

func (s *VulkanSurface) PhysicalDeviceSurfaceCapabilities(device core1_0.PhysicalDevice) (*SurfaceCapabilities, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	capabilitiesPtr := allocator.Malloc(int(unsafe.Sizeof([1]C.VkSurfaceCapabilitiesKHR{})))
	cCapabilities := (*C.VkSurfaceCapabilitiesKHR)(capabilitiesPtr)

	res, err := s.driver.VkGetPhysicalDeviceSurfaceCapabilitiesKHR(device.Handle(), s.handle, (*ext_driver.VkSurfaceCapabilitiesKHR)(unsafe.Pointer(cCapabilities)))
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

func (s *VulkanSurface) attemptFormats(device core1_0.PhysicalDevice) ([]SurfaceFormat, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	formatCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	formatCount := (*driver.Uint32)(formatCountPtr)

	res, err := s.driver.VkGetPhysicalDeviceSurfaceFormatsKHR(device.Handle(), s.handle, formatCount, nil)
	if err != nil {
		return nil, res, err
	}

	count := int(*formatCount)

	if count == 0 {
		return nil, res, nil
	}

	formatsPtr := allocator.Malloc(count * int(unsafe.Sizeof([1]C.VkSurfaceFormatKHR{})))

	res, err = s.driver.VkGetPhysicalDeviceSurfaceFormatsKHR(device.Handle(), s.handle, formatCount, (*ext_driver.VkSurfaceFormatKHR)(formatsPtr))
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

func (s *VulkanSurface) PhysicalDeviceSurfaceFormats(device core1_0.PhysicalDevice) ([]SurfaceFormat, common.VkResult, error) {
	var formats []SurfaceFormat
	var result common.VkResult
	var err error
	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		formats, result, err = s.attemptFormats(device)
	}

	return formats, result, err
}

func (s *VulkanSurface) attemptPresentModes(device core1_0.PhysicalDevice) ([]PresentMode, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	modeCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	modeCount := (*driver.Uint32)(modeCountPtr)

	res, err := s.driver.VkGetPhysicalDeviceSurfacePresentModesKHR(device.Handle(), s.handle, modeCount, nil)
	if err != nil {
		return nil, res, err
	}

	count := int(*modeCount)
	if count == 0 {
		return nil, res, nil
	}

	modesPtr := allocator.Malloc(count * int(unsafe.Sizeof(C.VkPresentModeKHR(0))))
	presentModes := (*ext_driver.VkPresentModeKHR)(modesPtr)

	res, err = s.driver.VkGetPhysicalDeviceSurfacePresentModesKHR(device.Handle(), s.handle, modeCount, presentModes)
	if err != nil || res == core1_0.VKIncomplete {
		return nil, res, err
	}

	presentModeSlice := ([]ext_driver.VkPresentModeKHR)(unsafe.Slice(presentModes, count))
	var result []PresentMode
	for i := 0; i < count; i++ {
		result = append(result, PresentMode(presentModeSlice[i]))
	}

	return result, res, nil
}

func (s *VulkanSurface) PhysicalDeviceSurfacePresentModes(device core1_0.PhysicalDevice) ([]PresentMode, common.VkResult, error) {
	var presentModes []PresentMode
	var result common.VkResult
	var err error
	for doWhile := true; doWhile; doWhile = (result == core1_0.VKIncomplete) {
		presentModes, result, err = s.attemptPresentModes(device)
	}

	return presentModes, result, err
}
