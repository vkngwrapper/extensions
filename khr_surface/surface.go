package khr_surface

import "C"
import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	ext_driver "github.com/vkngwrapper/extensions/v3/khr_surface/loader"
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

type Surface struct {
	instance loader.VkInstance
	handle   ext_driver.VkSurfaceKHR

	apiVersion common.APIVersion
}

func (s Surface) Handle() ext_driver.VkSurfaceKHR {
	return s.handle
}

func (s Surface) InstanceHandle() loader.VkInstance {
	return s.instance
}

func (s Surface) APIVersion() common.APIVersion {
	return s.apiVersion
}
