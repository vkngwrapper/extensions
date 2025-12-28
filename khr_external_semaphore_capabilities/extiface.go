package khr_external_semaphore_capabilities

import (
	"github.com/vkngwrapper/core/v3"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_external_semaphore_capabilities

// ExtensionDriver contains all commands for the khr_external_semaphore_capabilities extension
type ExtensionDriver interface {
	// GetPhysicalDeviceExternalSemaphoreProperties queries external Semaphore capabilities
	//
	// physicalDevice - The PhysicalDevice being queried
	//
	// o - Describes the parameters that would be consumed by Device.CreateSemaphore
	//
	// outData - A pre-allocated object in which the results will be populated. It should include any
	// desired chained OutData objects.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalSemaphoreProperties.html
	GetPhysicalDeviceExternalSemaphoreProperties(physicalDevice core.PhysicalDevice, o PhysicalDeviceExternalSemaphoreInfo, outData *ExternalSemaphoreProperties) error
}
