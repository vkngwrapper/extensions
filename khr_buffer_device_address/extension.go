package khr_buffer_device_address

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/extensions/v3/khr_buffer_device_address/loader"
)

// VulkanExtensionDriver is an implementation of the ExtensionDriver interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtensionDriver struct {
	driver khr_buffer_device_address_loader.Loader
	device core.Device
}

// CreateExtensionDriverFromCoreDriver produces an ExtensionDriver object from a Device with
// khr_buffer_device_address loaded
func CreateExtensionDriverFromCoreDriver(coreDriver core1_0.DeviceDriver) *VulkanExtensionDriver {
	device := coreDriver.Device()
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}
	return CreateExtensionDriverFromLoader(khr_buffer_device_address_loader.CreateLoaderFromCore(coreDriver.Loader()), device)
}

// CreateExtensionDriverFromLoader generates an ExtensionDriver from a loader.Loader object- this is usually
// used in tests to build an ExtensionDriver from mock drivers
func CreateExtensionDriverFromLoader(driver khr_buffer_device_address_loader.Loader, device core.Device) *VulkanExtensionDriver {
	return &VulkanExtensionDriver{
		driver: driver,
		device: device,
	}
}

func (e *VulkanExtensionDriver) GetBufferDeviceAddress(o BufferDeviceAddressInfo) (uint64, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	info, err := common.AllocOptions(arena, o)
	if err != nil {
		return 0, err
	}

	address := e.driver.VkGetBufferDeviceAddressKHR(
		e.device.Handle(),
		(*khr_buffer_device_address_loader.VkBufferDeviceAddressInfoKHR)(info),
	)
	return uint64(address), nil
}

func (e *VulkanExtensionDriver) GetBufferOpaqueCaptureAddress(o BufferDeviceAddressInfo) (uint64, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	info, err := common.AllocOptions(arena, o)
	if err != nil {
		return 0, err
	}

	address := e.driver.VkGetBufferOpaqueCaptureAddressKHR(
		e.device.Handle(),
		(*khr_buffer_device_address_loader.VkBufferDeviceAddressInfoKHR)(info),
	)
	return uint64(address), nil
}

func (e *VulkanExtensionDriver) GetDeviceMemoryOpaqueCaptureAddress(o DeviceMemoryOpaqueCaptureAddressInfo) (uint64, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	info, err := common.AllocOptions(arena, o)
	if err != nil {
		return 0, err
	}

	address := e.driver.VkGetDeviceMemoryOpaqueCaptureAddressKHR(
		e.device.Handle(),
		(*khr_buffer_device_address_loader.VkDeviceMemoryOpaqueCaptureAddressInfoKHR)(info),
	)
	return uint64(address), nil
}
