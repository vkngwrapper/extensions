package khr_buffer_device_address

import "github.com/vkngwrapper/core/core1_0"

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_buffer_device_address

// Extension contains all the commands for the khr_buffer_device_address extension
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_buffer_device_address.html
type Extension interface {
	// GetBufferDeviceAddress queries an address of a Buffer
	//
	// device - The Device that owns the Buffer
	//
	// o - Specifies the Buffer to retrieve an address for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html
	GetBufferDeviceAddress(device core1_0.Device, o BufferDeviceAddressInfo) (uint64, error)
	// GetBufferOpaqueCaptureAddress queries an opaque capture address of a Buffer
	//
	// device - The Device that owns the Buffer
	//
	// o - Specifies the Buffer to retrieve an address for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetBufferOpaqueCaptureAddress.html
	GetBufferOpaqueCaptureAddress(device core1_0.Device, o BufferDeviceAddressInfo) (uint64, error)
	// GetDeviceMemoryOpaqueCaptureAddress queries an opaque capture address of a DeviceMemory object
	//
	// device - The Device that owns the DeviceMemory
	//
	// o - Specifies the DeviceMemory object to retrieve an address for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetDeviceMemoryOpaqueCaptureAddress.html
	GetDeviceMemoryOpaqueCaptureAddress(device core1_0.Device, o DeviceMemoryOpaqueCaptureAddressInfo) (uint64, error)
}
