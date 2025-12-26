package khr_sampler_ycbcr_conversion

import (
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/loader"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_sampler_ycbcr_conversion

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
	CreateSamplerYcbcrConversion(device core.Device, o SamplerYcbcrConversionCreateInfo, allocator *loader.AllocationCallbacks) (core.SamplerYcbcrConversion, common.VkResult, error)

	// DestroySamplerYcbcrConversion destroys the SamplerYcbcrConversion object and the underlying structures. **Warning** after
	// destruction, this object will continue to exist, but the Vulkan object handle that backs it will
	// be invalid. Do not call further methods on this object.
	//
	// callbacks - Controls host memory deallocation
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySamplerYcbcrConversion.html
	DestroySamplerYcbcrConversion(conversion core.SamplerYcbcrConversion, allocator *loader.AllocationCallbacks)
}
