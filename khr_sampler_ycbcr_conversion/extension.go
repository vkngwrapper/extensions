package khr_sampler_ycbcr_conversion

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/extensions/v3/khr_sampler_ycbcr_conversion/loader"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_sampler_ycbcr_conversion_loader.Loader
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_sampler_ycbcr_conversion loaded
func CreateExtensionFromDevice(device core.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: khr_sampler_ycbcr_conversion_loader.CreateLoaderFromCore(device.Driver()),
	}
}

// CreateExtensionFromDriver generates an Extension from a loader.Loader object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_sampler_ycbcr_conversion_loader.Loader) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) CreateSamplerYcbcrConversion(device core.Device, o SamplerYcbcrConversionCreateInfo, allocator *loader.AllocationCallbacks) (core.SamplerYcbcrConversion, common.VkResult, error) {
	if device.Handle() == 0 {
		panic("device cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.SamplerYcbcrConversion{}, core1_0.VKErrorUnknown, err
	}

	var ycbcrHandle khr_sampler_ycbcr_conversion_loader.VkSamplerYcbcrConversionKHR
	res, err := e.driver.VkCreateSamplerYcbcrConversionKHR(
		device.Handle(),
		(*khr_sampler_ycbcr_conversion_loader.VkSamplerYcbcrConversionCreateInfoKHR)(optionPtr),
		allocator.Handle(),
		&ycbcrHandle,
	)
	if err != nil {
		return core.SamplerYcbcrConversion{}, res, err
	}

	ycbcr := core.InternalSamplerYcbcrConversion(device.Handle(), loader.VkSamplerYcbcrConversion(ycbcrHandle), device.APIVersion())

	return ycbcr, res, nil
}

func (e *VulkanExtension) DestroySamplerYcbcrConversion(conversion core.SamplerYcbcrConversion, allocator *loader.AllocationCallbacks) {
	e.driver.VkDestroySamplerYcbcrConversionKHR(conversion.DeviceHandle(), khr_sampler_ycbcr_conversion_loader.VkSamplerYcbcrConversionKHR(conversion.Handle()), allocator.Handle())
}
