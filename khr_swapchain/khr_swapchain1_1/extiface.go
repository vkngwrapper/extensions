package khr_swapchain1_1

import (
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/extensions/v3/khr_surface"
)

//go:generate mockgen -source extiface.go -destination ../mocks/extension1_1.go -package mock_swapchain -mock_names Extension=MockExtension1_1

// Extension contains all the core1.1-only commands for the khr_swapchain extension
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html
type Extension interface {
	// AcquireNextImage2 retrieves the index of the next available presentable Image
	//
	// device - The Device which owns the requested khr_swapchain.Swapchain
	//
	// o - Contains parameters of the acquire operation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImage2KHR.html
	AcquireNextImage2(device core.Device, o AcquireNextImageInfo) (int, common.VkResult, error)
	// DeviceGroupPresentCapabilities queries Device group present capabilities for a surface
	//
	// device - The Device being queried
	//
	// outData - A pre-allocated object in which the capabilities will be populated. It should include any desired
	// chained OutData objects
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupPresentCapabilitiesKHR.html
	DeviceGroupPresentCapabilities(device core.Device, outData *DeviceGroupPresentCapabilities) (common.VkResult, error)
	// DeviceGroupSurfacePresentModes queries present capabilities for a khr_surface.Surface
	//
	// device - The Device being queried
	//
	// surface - The Surface whose present capabilities are being requested
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupSurfacePresentModesKHR.html
	DeviceGroupSurfacePresentModes(device core.Device, surface khr_surface.Surface) (DeviceGroupPresentModeFlags, common.VkResult, error)
	// PhysicalDevicePresentRectangles queries present rectangles for a khr_surface.Surface on a PhysicalDevice
	//
	// physicalDevice - The PhysicalDevice being queried
	//
	// surface - The Surface whose present rectangles are being requested
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDevicePresentRectanglesKHR.html
	PhysicalDevicePresentRectangles(physicalDevice core.PhysicalDevice, surface khr_surface.Surface) ([]core1_0.Rect2D, common.VkResult, error)
}
