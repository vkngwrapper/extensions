package khr_sampler_ycbcr_conversion

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
	khr_sampler_ycbcr_conversion_driver "github.com/vkngwrapper/extensions/khr_sampler_ycbcr_conversion/driver"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_sampler_ycbcr_conversion

type SamplerYcbcrConversion interface {
	Handle() khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionKHR
	Destroy(allocator *driver.AllocationCallbacks)
}

type Extension interface {
	CreateSamplerYcbcrConversion(device core1_0.Device, o SamplerYcbcrConversionCreateInfo, allocator *driver.AllocationCallbacks) (SamplerYcbcrConversion, common.VkResult, error)
}
