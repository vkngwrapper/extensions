package khr_get_surface_capabilities2

import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_get_surface_capabilities2

type ExtensionDriver interface {
	GetPhysicalDeviceSurfaceCapabilities2(physicalDevice core1_0.PhysicalDevice, o PhysicalDeviceSurfaceInfo2, out *SurfaceCapabilities2) (common.VkResult, error)
	GetPhysicalDeviceSurfaceFormats2(physicalDevice core1_0.PhysicalDevice, o PhysicalDeviceSurfaceInfo2, outDataFactory func() *SurfaceFormat2) ([]*SurfaceFormat2, common.VkResult, error)
}
