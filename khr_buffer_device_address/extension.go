package khr_buffer_device_address

import (
	"github.com/CannibalVox/cgoparam"
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	khr_buffer_device_address_driver "github.com/vkngwrapper/extensions/khr_buffer_device_address/driver"
)

type VulkanExtension struct {
	driver khr_buffer_device_address_driver.Driver
}

func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}
	return CreateExtensionFromDriver(khr_buffer_device_address_driver.CreateDriverFromCore(device.Driver()))
}

func CreateExtensionFromDriver(driver khr_buffer_device_address_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) GetBufferDeviceAddress(device core1_0.Device, o BufferDeviceAddressInfo) (uint64, error) {
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

func (e *VulkanExtension) GetDeviceMemoryOpaqueCaptureAddress(device core1_0.Device, o DeviceMemoryOpaqueAddressInfo) (uint64, error) {
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
