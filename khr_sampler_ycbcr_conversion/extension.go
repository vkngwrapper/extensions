package khr_sampler_ycbcr_conversion

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/loader"
	"github.com/vkngwrapper/extensions/v3/khr_sampler_ycbcr_conversion/loader"
)

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	driver khr_sampler_ycbcr_conversion_loader.Loader
	device core.Device
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// khr_sampler_ycbcr_conversion loaded
func CreateExtensionDriverFromCoreDriver(driver core1_0.DeviceDriver) *VulkanExtensionDriver {
	device := driver.Device()
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtensionDriver{
		driver: khr_sampler_ycbcr_conversion_loader.CreateLoaderFromCore(driver.Loader()),
	}
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver khr_sampler_ycbcr_conversion_loader.Loader, device core.Device) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		driver: driver,
		device: device,
	}
}

func (e *VulkanExtensionDriver) CreateSamplerYcbcrConversion(o SamplerYcbcrConversionCreateInfo, allocator *loader.AllocationCallbacks) (core.SamplerYcbcrConversion, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return core.SamplerYcbcrConversion{}, core1_0.VKErrorUnknown, err
	}

	var ycbcrHandle khr_sampler_ycbcr_conversion_loader.VkSamplerYcbcrConversionKHR
	res, err := e.driver.VkCreateSamplerYcbcrConversionKHR(
		e.device.Handle(),
		(*khr_sampler_ycbcr_conversion_loader.VkSamplerYcbcrConversionCreateInfoKHR)(optionPtr),
		allocator.Handle(),
		&ycbcrHandle,
	)
	if err != nil {
		return core.SamplerYcbcrConversion{}, res, err
	}

	ycbcr := core.InternalSamplerYcbcrConversion(e.device.Handle(), loader.VkSamplerYcbcrConversion(ycbcrHandle), e.device.APIVersion())

	return ycbcr, res, nil
}

func (e *VulkanExtensionDriver) DestroySamplerYcbcrConversion(conversion core.SamplerYcbcrConversion, allocator *loader.AllocationCallbacks) {
	e.driver.VkDestroySamplerYcbcrConversionKHR(conversion.DeviceHandle(), khr_sampler_ycbcr_conversion_loader.VkSamplerYcbcrConversionKHR(conversion.Handle()), allocator.Handle())
}
