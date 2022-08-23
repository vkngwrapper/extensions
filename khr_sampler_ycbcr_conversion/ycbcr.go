package khr_sampler_ycbcr_conversion

import (
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/driver"
	khr_sampler_ycbcr_conversion_driver "github.com/vkngwrapper/extensions/v2/khr_sampler_ycbcr_conversion/driver"
)

// VulkanSamplerYcbcrConversion is an implementation of the SamplerYcbcrConversion interface that actually communicates
// with Vulkan. This is the default implementation. See the interface for more documentation.
type VulkanSamplerYcbcrConversion struct {
	coreDriver  driver.Driver
	driver      khr_sampler_ycbcr_conversion_driver.Driver
	device      driver.VkDevice
	ycbcrHandle khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionKHR

	maximumAPIVersion common.APIVersion
}

func (y *VulkanSamplerYcbcrConversion) Handle() khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionKHR {
	return y.ycbcrHandle
}

func (y *VulkanSamplerYcbcrConversion) Destroy(allocator *driver.AllocationCallbacks) {
	y.driver.VkDestroySamplerYcbcrConversionKHR(y.device, y.ycbcrHandle, allocator.Handle())
}
