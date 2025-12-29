package khr_surface

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	khr_surface_loader "github.com/vkngwrapper/extensions/v3/khr_surface/loader"
)

// ExtensionDriver contains all commands for the khr_surface extension
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html
type ExtensionDriver interface {
	// CreateSurfaceFromHandle produces a Surface object from a surface handle received from
	// a windowing library
	CreateSurfaceFromHandle(surfaceHandle khr_surface_loader.VkSurfaceKHR) (Surface, error)

	// DestroySurface deletes a Surface and underlying structures from the device. **Warning**
	// after destruction, this object will still exist, but the Vulkan object handle
	// that backs it will be invalid. Do not call further methods on this object.
	//
	// callbacks - A set of allocation callbacks to control the memory free behavior of this command
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySurfaceKHR.html
	DestroySurface(surface Surface, callbacks *loader.AllocationCallbacks)
	// GetPhysicalDeviceSurfaceSupport queries if presentation of this Surface is supported on the specified PhysicalDevice
	//
	// physicalDevice - The PhysicalDevice to query for support
	//
	// queueFamilyIndex - The Queue family to be used to present the Surface
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html
	GetPhysicalDeviceSurfaceSupport(surface Surface, physicalDevice core1_0.PhysicalDevice, queueFamilyIndex int) (bool, common.VkResult, error)
	// GetPhysicalDeviceSurfaceCapabilities queries Surface capabilities on the specified PhysicalDevice
	//
	// device - The PhysicalDevice to query for capabilities
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html
	GetPhysicalDeviceSurfaceCapabilities(surface Surface, device core1_0.PhysicalDevice) (*SurfaceCapabilities, common.VkResult, error)
	// GetPhysicalDeviceSurfaceFormats queries color formats supported by Surface on the specified PhysicalDevice
	//
	// device - The PhysicalDevice to query for supported Surface formats
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceFormatsKHR.html
	GetPhysicalDeviceSurfaceFormats(surface Surface, device core1_0.PhysicalDevice) ([]SurfaceFormat, common.VkResult, error)
	// GetPhysicalDeviceSurfacePresentModes queries supported presentation modes on the specified PhysicalDevice
	//
	// device - The PhysicalDevice to query for supported presentation modes
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfacePresentModesKHR.html
	GetPhysicalDeviceSurfacePresentModes(surface Surface, device core1_0.PhysicalDevice) ([]PresentMode, common.VkResult, error)
}
