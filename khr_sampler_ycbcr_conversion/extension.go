package khr_sampler_ycbcr_conversion

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/driver"
	khr_sampler_ycbcr_conversion_driver "github.com/vkngwrapper/extensions/v3/khr_sampler_ycbcr_conversion/driver"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_sampler_ycbcr_conversion_driver.Driver
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_sampler_ycbcr_conversion loaded
func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: khr_sampler_ycbcr_conversion_driver.CreateDriverFromCore(device.Driver()),
	}
}

// CreateExtensionFromDriver generates an Extension from a driver.Driver object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_sampler_ycbcr_conversion_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) CreateSamplerYcbcrConversion(device core1_0.Device, o SamplerYcbcrConversionCreateInfo, allocator *driver.AllocationCallbacks) (SamplerYcbcrConversion, common.VkResult, error) {
	if device == nil {
		panic("device cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var ycbcrHandle khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionKHR
	res, err := e.driver.VkCreateSamplerYcbcrConversionKHR(
		device.Handle(),
		(*khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionCreateInfoKHR)(optionPtr),
		allocator.Handle(),
		&ycbcrHandle,
	)
	if err != nil {
		return nil, res, err
	}

	ycbcr := device.Driver().ObjectStore().GetOrCreate(driver.VulkanHandle(ycbcrHandle), driver.Core1_1,
		func() any {
			return &VulkanSamplerYcbcrConversion{
				coreDriver:        device.Driver(),
				driver:            e.driver,
				device:            device.Handle(),
				ycbcrHandle:       ycbcrHandle,
				maximumAPIVersion: device.APIVersion(),
			}
		}).(*VulkanSamplerYcbcrConversion)

	return ycbcr, res, nil
}
