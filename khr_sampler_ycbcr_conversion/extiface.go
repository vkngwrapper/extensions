package khr_sampler_ycbcr_conversion

import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/driver"
	khr_sampler_ycbcr_conversion_driver "github.com/vkngwrapper/extensions/khr_sampler_ycbcr_conversion/driver"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_sampler_ycbcr_conversion

// SamplerYcbcrConversion is an opaque representation of a device-specific sampler YCbCr conversion
// description.
//
// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html
type SamplerYcbcrConversion interface {
	// Handle is the internal Vulkan object handle for this SamplerYcbcrConversion
	Handle() khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionKHR
	// Destroy destroys the SamplerYcbcrConversion object and the underlying structures. **Warning** after
	// destruction, this object will continue to exist, but the Vulkan object handle that backs it will
	// be invalid. Do not call further methods on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySamplerYcbcrConversion.html
	Destroy(allocator *driver.AllocationCallbacks)
}

// Extension contains all commands for the khr_sampler_ycbcr_conversion extension
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_sampler_ycbcr_conversion.html
type Extension interface {
	// CreateSamplerYcbcrConversion creates a new Y'CbCr conversion
	//
	// device - The Device which will own the new SamplerYcbcrConversion object
	//
	// o - Specifies the requested sampler Y'CbCr conversion
	//
	// allocator - Controls host allocation behavior
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSamplerYcbcrConversion.html
	CreateSamplerYcbcrConversion(device core1_0.Device, o SamplerYcbcrConversionCreateInfo, allocator *driver.AllocationCallbacks) (SamplerYcbcrConversion, common.VkResult, error)
}
