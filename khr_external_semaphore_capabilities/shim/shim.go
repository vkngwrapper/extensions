package khr_external_semaphore_capabilities_shim

import (
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/extensions/v3/khr_external_semaphore_capabilities"
)

//go:generate mockgen -source shim.go -destination ../mocks/shim.go -package mock_external_semaphore_capabilities

// Shim contains all commands for the khr_external_semaphore_capabilities extension
type Shim interface {
	// ExternalSemaphoreProperties queries external Semaphore capabilities
	//
	// o - Describes the parameters that would be consumed by Device.CreateSemaphore
	//
	// outData - A pre-allocated object in which the results will be populated. It should include any
	// desired chained OutData objects.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalSemaphoreProperties.html
	ExternalSemaphoreProperties(o core1_1.PhysicalDeviceExternalSemaphoreInfo, outData *core1_1.ExternalSemaphoreProperties) error
}

// Compiler check that VulkanShim satisfies Shim
var _ Shim = &VulkanShim{}

type VulkanShim struct {
	extension      khr_external_semaphore_capabilities.Extension
	physicalDevice core1_0.PhysicalDevice
}

func NewShim(extension khr_external_semaphore_capabilities.Extension, device core1_0.PhysicalDevice) *VulkanShim {
	if extension == nil {
		panic("extension cannot be nil")
	}
	if device == nil {
		panic("device cannot be nil")
	}

	return &VulkanShim{
		extension:      extension,
		physicalDevice: device,
	}
}

func (s *VulkanShim) ExternalSemaphoreProperties(o core1_1.PhysicalDeviceExternalSemaphoreInfo, outData *core1_1.ExternalSemaphoreProperties) error {

	inOptions := khr_external_semaphore_capabilities.PhysicalDeviceExternalSemaphoreInfo{
		HandleType:  khr_external_semaphore_capabilities.ExternalSemaphoreHandleTypeFlags(o.HandleType),
		NextOptions: o.NextOptions,
	}

	var inOutData khr_external_semaphore_capabilities.ExternalSemaphoreProperties

	if outData != nil {
		inOutData.NextOutData = outData.NextOutData
	}

	err := s.extension.PhysicalDeviceExternalSemaphoreProperties(s.physicalDevice, inOptions, &inOutData)
	if err != nil {
		return err
	}

	if outData != nil {
		outData.ExportFromImportedHandleTypes = core1_1.ExternalSemaphoreHandleTypeFlags(inOutData.ExportFromImportedHandleTypes)
		outData.CompatibleHandleTypes = core1_1.ExternalSemaphoreHandleTypeFlags(inOutData.CompatibleHandleTypes)
		outData.ExternalSemaphoreFeatures = core1_1.ExternalSemaphoreFeatureFlags(inOutData.ExternalSemaphoreFeatures)
	}

	return nil
}
