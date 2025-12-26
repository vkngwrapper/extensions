package khr_buffer_device_address

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v3"
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/extensions/v3/khr_buffer_device_address/loader"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_buffer_device_address_loader.Loader
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_buffer_device_address loaded
func CreateExtensionFromDevice(device core.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}
	return CreateExtensionFromDriver(khr_buffer_device_address_loader.CreateLoaderFromCore(device.Driver()))
}

// CreateExtensionFromDriver generates an Extension from a loader.Loader object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_buffer_device_address_loader.Loader) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) GetBufferDeviceAddress(device core.Device, o BufferDeviceAddressInfo) (uint64, error) {
	if device.Handle() == 0 {
		panic("device cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	info, err := common.AllocOptions(arena, o)
	if err != nil {
		return 0, err
	}

	address := e.driver.VkGetBufferDeviceAddressKHR(
		device.Handle(),
		(*khr_buffer_device_address_loader.VkBufferDeviceAddressInfoKHR)(info),
	)
	return uint64(address), nil
}

func (e *VulkanExtension) GetBufferOpaqueCaptureAddress(device core.Device, o BufferDeviceAddressInfo) (uint64, error) {
	if device.Handle() == 0 {
		panic("device cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	info, err := common.AllocOptions(arena, o)
	if err != nil {
		return 0, err
	}

	address := e.driver.VkGetBufferOpaqueCaptureAddressKHR(
		device.Handle(),
		(*khr_buffer_device_address_loader.VkBufferDeviceAddressInfoKHR)(info),
	)
	return uint64(address), nil
}

func (e *VulkanExtension) GetDeviceMemoryOpaqueCaptureAddress(device core.Device, o DeviceMemoryOpaqueCaptureAddressInfo) (uint64, error) {
	if device.Handle() == 0 {
		panic("device cannot be uninitialized")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	info, err := common.AllocOptions(arena, o)
	if err != nil {
		return 0, err
	}

	address := e.driver.VkGetDeviceMemoryOpaqueCaptureAddressKHR(
		device.Handle(),
		(*khr_buffer_device_address_loader.VkDeviceMemoryOpaqueCaptureAddressInfoKHR)(info),
	)
	return uint64(address), nil
}
