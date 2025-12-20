package khr_external_memory_capabilities_shim

import (
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/extensions/v3/khr_external_memory_capabilities"
)

//go:generate mockgen -source shim.go -destination ../mocks/shim.go -package mock_external_memory_capabilities

// Shim contains all the commands for the khr_external_memory_capabilities extension
type Shim interface {
	// ExternalBufferProperties queries external types supported by Buffer objects
	//
	// o - Describes the parameters that would be consumed by Device.CreateBuffer
	//
	// outData - A pre-allocated object in which the results will be populated. It should include any
	// desired chained OutData objects.
	//
	// https://www.khronos.org/registry/VulkanSC/specs/1.0-extensions/man/html/vkGetPhysicalDeviceExternalBufferProperties.html
	ExternalBufferProperties(o core1_1.PhysicalDeviceExternalBufferInfo, outData *core1_1.ExternalBufferProperties) error
}

// Compiler check that VulkanShim satisfies Shim
var _ Shim = &VulkanShim{}

type VulkanShim struct {
	extension      khr_external_memory_capabilities.Extension
	physicalDevice core1_0.PhysicalDevice
}

func NewShim(extension khr_external_memory_capabilities.Extension, device core1_0.PhysicalDevice) *VulkanShim {
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

func (s *VulkanShim) ExternalBufferProperties(o core1_1.PhysicalDeviceExternalBufferInfo, outData *core1_1.ExternalBufferProperties) error {
	inOptions := khr_external_memory_capabilities.PhysicalDeviceExternalBufferInfo{
		Flags:       o.Flags,
		Usage:       o.Usage,
		HandleType:  khr_external_memory_capabilities.ExternalMemoryHandleTypeFlags(o.HandleType),
		NextOptions: o.NextOptions,
	}
	var inOutData khr_external_memory_capabilities.ExternalBufferProperties

	if outData != nil {
		inOutData.NextOutData = outData.NextOutData
	}

	err := s.extension.PhysicalDeviceExternalBufferProperties(
		s.physicalDevice,
		inOptions,
		&inOutData,
	)
	if err != nil {
		return err
	}

	if outData != nil {
		outData.ExternalMemoryProperties.ExternalMemoryFeatures = core1_1.ExternalMemoryFeatureFlags(
			inOutData.ExternalMemoryProperties.ExternalMemoryFeatures,
		)
		outData.ExternalMemoryProperties.ExportFromImportedHandleTypes = core1_1.ExternalMemoryHandleTypeFlags(
			inOutData.ExternalMemoryProperties.ExportFromImportedHandleTypes,
		)
		outData.ExternalMemoryProperties.CompatibleHandleTypes = core1_1.ExternalMemoryHandleTypeFlags(
			inOutData.ExternalMemoryProperties.CompatibleHandleTypes,
		)
	}

	return nil
}
