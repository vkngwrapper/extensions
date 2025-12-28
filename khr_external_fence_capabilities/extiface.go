package khr_external_fence_capabilities

import (
	"github.com/vkngwrapper/core/v3"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_external_fence_capabilities

// ExtensionDriver contains all the commands for the khr_external_fence_capabilities extension
type ExtensionDriver interface {
	// GetPhysicalDeviceExternalFenceProperties queries external Fence capabilities
	//
	// physicalDevice - The PhysicalDevice to retrieve capabilities for
	//
	// o - Describes the parameters that would be consumed by Device.CreateFence
	//
	// outData - A pre-allocated object in which the results will be populated. It should include
	// any desired chained OutData objects.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalFenceProperties.html
	GetPhysicalDeviceExternalFenceProperties(physicalDevice core.PhysicalDevice, o PhysicalDeviceExternalFenceInfo, outData *ExternalFenceProperties) error
}
