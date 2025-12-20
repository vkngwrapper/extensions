package khr_external_fence_capabilities_shim

import (
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/extensions/v3/khr_external_fence_capabilities"
)

//go:generate mockgen -source shim.go -destination ../mocks/shim.go -package mock_external_fence_capabilities

// Shim contains all the commands for the khr_external_fence_capabilities extension
type Shim interface {
	// ExternalFenceProperties queries external Fence capabilities
	//
	// o - Describes the parameters that would be consumed by Device.CreateFence
	//
	// outData - A pre-allocated object in which the results will be populated. It should include
	// any desired chained OutData objects.
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalFenceProperties.html
	ExternalFenceProperties(o core1_1.PhysicalDeviceExternalFenceInfo, outData *core1_1.ExternalFenceProperties) error
}

// Compiler check that VulkanShim satisfies Shim
var _ Shim = &VulkanShim{}

type VulkanShim struct {
	extension      khr_external_fence_capabilities.Extension
	physicalDevice core1_0.PhysicalDevice
}

func NewShim(extension khr_external_fence_capabilities.Extension, device core1_0.PhysicalDevice) *VulkanShim {
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

func (s *VulkanShim) ExternalFenceProperties(o core1_1.PhysicalDeviceExternalFenceInfo, outData *core1_1.ExternalFenceProperties) error {
	inOptions := khr_external_fence_capabilities.PhysicalDeviceExternalFenceInfo{
		HandleType:  khr_external_fence_capabilities.ExternalFenceHandleTypeFlags(o.HandleType),
		NextOptions: o.NextOptions,
	}

	var inOutData khr_external_fence_capabilities.ExternalFenceProperties

	if outData != nil {
		inOutData.NextOutData = outData.NextOutData
	}

	err := s.extension.PhysicalDeviceExternalFenceProperties(
		s.physicalDevice,
		inOptions,
		&inOutData,
	)

	if err != nil {
		return err
	}

	if outData != nil {
		outData.ExportFromImportedHandleTypes = core1_1.ExternalFenceHandleTypeFlags(inOutData.ExportFromImportedHandleTypes)
		outData.CompatibleHandleTypes = core1_1.ExternalFenceHandleTypeFlags(inOutData.CompatibleHandleTypes)
		outData.ExternalFenceFeatures = core1_1.ExternalFenceFeatureFlags(inOutData.ExternalFenceFeatures)
	}

	return nil
}
