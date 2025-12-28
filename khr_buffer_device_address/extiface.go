package khr_buffer_device_address

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_buffer_device_address

// ExtensionDriver contains all the commands for the khr_buffer_device_address extension
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_buffer_device_address.html
type ExtensionDriver interface {
	// GetBufferDeviceAddress queries an address of a Buffer
	//
	// device - The Device that owns the Buffer
	//
	// o - Specifies the Buffer to retrieve an address for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html
	GetBufferDeviceAddress(o BufferDeviceAddressInfo) (uint64, error)
	// GetBufferOpaqueCaptureAddress queries an opaque capture address of a Buffer
	//
	// device - The Device that owns the Buffer
	//
	// o - Specifies the Buffer to retrieve an address for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetBufferOpaqueCaptureAddress.html
	GetBufferOpaqueCaptureAddress(o BufferDeviceAddressInfo) (uint64, error)
	// GetDeviceMemoryOpaqueCaptureAddress queries an opaque capture address of a DeviceMemory object
	//
	// device - The Device that owns the DeviceMemory
	//
	// o - Specifies the DeviceMemory object to retrieve an address for
	//
	// https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/vkGetDeviceMemoryOpaqueCaptureAddress.html
	GetDeviceMemoryOpaqueCaptureAddress(o DeviceMemoryOpaqueCaptureAddressInfo) (uint64, error)
}
