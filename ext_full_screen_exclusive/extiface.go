package ext_full_screen_exclusive

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/extensions/v3/khr_device_group"
	"github.com/vkngwrapper/extensions/v3/khr_get_surface_capabilities2"
	"github.com/vkngwrapper/extensions/v3/khr_surface"
	"github.com/vkngwrapper/extensions/v3/khr_swapchain"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_full_screen_exclusive

type ExtensionDriver interface {
	AcquireFullScreenExclusiveMode(swapchain khr_swapchain.Swapchain) (common.VkResult, error)
	GetPhysicalDeviceSurfacePresentModes2(physicalDevice core1_0.PhysicalDevice, o khr_get_surface_capabilities2.PhysicalDeviceSurfaceInfo2) ([]khr_surface.PresentMode, common.VkResult, error)
	ReleaseFullScreenExclusiveMode(swapchain khr_swapchain.Swapchain) (common.VkResult, error)
}

type ExtensionDriverWithDeviceGroups interface {
	GetDeviceGroupSurfacePresentModes2(o khr_get_surface_capabilities2.PhysicalDeviceSurfaceInfo2, out *khr_device_group.DeviceGroupPresentModeFlags) (common.VkResult, error)
}
