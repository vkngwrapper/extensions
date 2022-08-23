package khr_external_semaphore_capabilities

import "github.com/vkngwrapper/core/v2/core1_0"

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_external_semaphore_capabilities

// Extension contains all commands for the khr_external_semaphore_capabilities extension
type Extension interface {
	// PhysicalDeviceExternalSemaphoreProperties queries external Semaphore capabilities
	//
	// physicalDevice - The PhysicalDevice being queried
	//
	// o - Describes the parameters that would be consumed by Device.CreateSemaphore
	//
	// outData - A pre-allocated object in which the results will be populated. It should include any
	// desired chained OutData objects.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalSemaphoreProperties.html
	PhysicalDeviceExternalSemaphoreProperties(physicalDevice core1_0.PhysicalDevice, o PhysicalDeviceExternalSemaphoreInfo, outData *ExternalSemaphoreProperties) error
}
