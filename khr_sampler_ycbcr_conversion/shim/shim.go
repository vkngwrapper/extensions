package khr_sampler_ycbcr_conversion_shim

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_1"
	"github.com/vkngwrapper/core/v2/driver"
	"github.com/vkngwrapper/extensions/v2/khr_sampler_ycbcr_conversion"
)

//go:generate mockgen -source shim.go -destination ../mocks/shim.go -package mock_sampler_ycbcr_conversion

// Shim contains all commands for the khr_sampler_ycbcr_conversion extension
type Shim interface {
	// CreateSamplerYcbcrConversion creates a new Y'CbCr conversion
	//
	// o - Specifies the requested sampler Y'CbCr conversion
	//
	// allocator - Controls host allocation behavior
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSamplerYcbcrConversion.html
	CreateSamplerYcbcrConversion(o core1_1.SamplerYcbcrConversionCreateInfo, allocator *driver.AllocationCallbacks) (core1_1.SamplerYcbcrConversion, common.VkResult, error)
}

// VulkanShimSamplerYcbcrConversion is a wrapper for khr_sampler_ycbcr_conversion.SamplerYcbcrConversion that
// converts the Handle() method to return the core 1.1 handle type, so that the extension SamplerYcbcrConversion
// can be made compatible with the core 1.1 SamplerYcbcrConversion for the shim
type VulkanShimSamplerYcbcrConversion struct {
	khr_sampler_ycbcr_conversion.SamplerYcbcrConversion
}

func (t *VulkanShimSamplerYcbcrConversion) Handle() driver.VkSamplerYcbcrConversion {
	return driver.VkSamplerYcbcrConversion(t.SamplerYcbcrConversion.Handle())
}

type VulkanShim struct {
	extension khr_sampler_ycbcr_conversion.Extension
	device    core1_0.Device
}

// compiler check that VulkanShim satisfies Shim
var _ Shim = &VulkanShim{}

func (s *VulkanShim) CreateSamplerYcbcrConversion(o core1_1.SamplerYcbcrConversionCreateInfo, allocator *driver.AllocationCallbacks) (core1_1.SamplerYcbcrConversion, common.VkResult, error) {
	inOptions := khr_sampler_ycbcr_conversion.SamplerYcbcrConversionCreateInfo{
		Format:                      o.Format,
		YcbcrModel:                  khr_sampler_ycbcr_conversion.SamplerYcbcrModelConversion(o.YcbcrModel),
		YcbcrRange:                  khr_sampler_ycbcr_conversion.SamplerYcbcrRange(o.YcbcrRange),
		Components:                  o.Components,
		XChromaOffset:               khr_sampler_ycbcr_conversion.ChromaLocation(o.XChromaOffset),
		YChromaOffset:               khr_sampler_ycbcr_conversion.ChromaLocation(o.YChromaOffset),
		ChromaFilter:                o.ChromaFilter,
		ForceExplicitReconstruction: o.ForceExplicitReconstruction,
		NextOptions:                 o.NextOptions,
	}

	var outSampler core1_1.SamplerYcbcrConversion
	sampler, res, err := s.extension.CreateSamplerYcbcrConversion(
		s.device,
		inOptions,
		allocator,
	)
	if sampler != nil {
		outSampler = &VulkanShimSamplerYcbcrConversion{sampler}
	}

	return outSampler, res, err
}
