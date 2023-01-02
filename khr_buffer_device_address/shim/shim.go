package khr_buffer_device_address_shim

import (
	"github.com/vkngwrapper/core/v2/core1_0"
	"github.com/vkngwrapper/core/v2/core1_2"
	"github.com/vkngwrapper/extensions/v2/khr_buffer_device_address"
)

//go:generate mockgen -source shim.go -destination ../mocks/shim.go -package mock_buffer_device_address

// Shim contains all the commands for the khr_buffer_device_address extension
type Shim interface {
	// GetBufferDeviceAddress queries an address of a Buffer
	//
	// o - Specifies the Buffer to retrieve an address for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html
	GetBufferDeviceAddress(o core1_2.BufferDeviceAddressInfo) (uint64, error)
	// GetBufferOpaqueCaptureAddress queries an opaque capture address of a Buffer
	//
	// o - Specifies the Buffer to retrieve an address for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetBufferOpaqueCaptureAddress.html
	GetBufferOpaqueCaptureAddress(o core1_2.BufferDeviceAddressInfo) (uint64, error)
	// GetDeviceMemoryOpaqueCaptureAddress queries an opaque capture address of a DeviceMemory object
	//
	// o - Specifies the DeviceMemory object to retrieve an address for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetDeviceMemoryOpaqueCaptureAddress.html
	GetDeviceMemoryOpaqueCaptureAddress(o core1_2.DeviceMemoryOpaqueCaptureAddressInfo) (uint64, error)
}

// Compiler check that VulkanShim satisfies Shim
var _ Shim = &VulkanShim{}

type VulkanShim struct {
	extension khr_buffer_device_address.Extension
	device    core1_0.Device
}

func NewShim(extension khr_buffer_device_address.Extension, device core1_0.Device) *VulkanShim {
	if extension == nil {
		panic("extension cannot be nil")
	}
	if device == nil {
		panic("device cannot be nil")
	}
	return &VulkanShim{
		extension: extension,
		device:    device,
	}
}

func (s *VulkanShim) GetBufferDeviceAddress(o core1_2.BufferDeviceAddressInfo) (uint64, error) {
	return s.extension.GetBufferDeviceAddress(s.device, (khr_buffer_device_address.BufferDeviceAddressInfo)(o))
}

func (s *VulkanShim) GetBufferOpaqueCaptureAddress(o core1_2.BufferDeviceAddressInfo) (uint64, error) {
	return s.extension.GetBufferOpaqueCaptureAddress(s.device, (khr_buffer_device_address.BufferDeviceAddressInfo)(o))
}

func (s *VulkanShim) GetDeviceMemoryOpaqueCaptureAddress(o core1_2.DeviceMemoryOpaqueCaptureAddressInfo) (uint64, error) {
	return s.extension.GetDeviceMemoryOpaqueCaptureAddress(s.device, (khr_buffer_device_address.DeviceMemoryOpaqueCaptureAddressInfo)(o))
}
