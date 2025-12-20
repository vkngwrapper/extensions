package khr_buffer_device_address

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/v2/common"
	"github.com/vkngwrapper/core/v2/core1_0"
	khr_buffer_device_address_driver "github.com/vkngwrapper/extensions/v3/khr_buffer_device_address/driver"
)

// VulkanExtension is an implementation of the Extension interface that actually communicates with Vulkan. This
// is the default implementation. See the interface for more documentation.
type VulkanExtension struct {
	driver khr_buffer_device_address_driver.Driver
}

// CreateExtensionFromDevice produces an Extension object from a Device with
// khr_buffer_device_address loaded
func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}
	return CreateExtensionFromDriver(khr_buffer_device_address_driver.CreateDriverFromCore(device.Driver()))
}

// CreateExtensionFromDriver generates an Extension from a driver.Driver object- this is usually
// used in tests to build an Extension from mock drivers
func CreateExtensionFromDriver(driver khr_buffer_device_address_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) GetBufferDeviceAddress(device core1_0.Device, o BufferDeviceAddressInfo) (uint64, error) {
	if device == nil {
		panic("device cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	info, err := common.AllocOptions(arena, o)
	if err != nil {
		return 0, err
	}

	address := e.driver.VkGetBufferDeviceAddressKHR(
		device.Handle(),
		(*khr_buffer_device_address_driver.VkBufferDeviceAddressInfoKHR)(info),
	)
	return uint64(address), nil
}

func (e *VulkanExtension) GetBufferOpaqueCaptureAddress(device core1_0.Device, o BufferDeviceAddressInfo) (uint64, error) {
	if device == nil {
		panic("device cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	info, err := common.AllocOptions(arena, o)
	if err != nil {
		return 0, err
	}

	address := e.driver.VkGetBufferOpaqueCaptureAddressKHR(
		device.Handle(),
		(*khr_buffer_device_address_driver.VkBufferDeviceAddressInfoKHR)(info),
	)
	return uint64(address), nil
}

func (e *VulkanExtension) GetDeviceMemoryOpaqueCaptureAddress(device core1_0.Device, o DeviceMemoryOpaqueCaptureAddressInfo) (uint64, error) {
	if device == nil {
		panic("device cannot be nil")
	}
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	info, err := common.AllocOptions(arena, o)
	if err != nil {
		return 0, err
	}

	address := e.driver.VkGetDeviceMemoryOpaqueCaptureAddressKHR(
		device.Handle(),
		(*khr_buffer_device_address_driver.VkDeviceMemoryOpaqueCaptureAddressInfoKHR)(info),
	)
	return uint64(address), nil
}
