package khr_external_memory_capabilities

import (
	"github.com/vkngwrapper/core/v3"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_external_memory_capabilities

// Extension contains all the commands for the khr_external_memory_capabilities extension
type Extension interface {
	// PhysicalDeviceExternalBufferProperties queries external types supported by Buffer objects
	//
	// physicalDevice - The PhysicalDevice being queried
	//
	// o - Describes the parameters that would be consumed by Device.CreateBuffer
	//
	// outData - A pre-allocated object in which the results will be populated. It should include any
	// desired chained OutData objects.
	//
	// https://www.khronos.org/registry/VulkanSC/specs/1.0-extensions/man/html/vkGetPhysicalDeviceExternalBufferProperties.html
	PhysicalDeviceExternalBufferProperties(physicalDevice core.PhysicalDevice, o PhysicalDeviceExternalBufferInfo, outData *ExternalBufferProperties) error
}
