package khr_buffer_device_address

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
	"github.com/vkngwrapper/core/core1_1"
)

const (
	// ExtensionName is "VK_KHR_buffer_device_address"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_buffer_device_address.html
	ExtensionName string = C.VK_KHR_BUFFER_DEVICE_ADDRESS_EXTENSION_NAME

	// BufferCreateDeviceAddressCaptureReplay specifies that the Buffer object's address can
	// be saved and reused on a subsequent run (e.g. for trace capture and replay)
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateFlagBits.html
	BufferCreateDeviceAddressCaptureReplay core1_0.BufferCreateFlags = C.VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR

	// BufferUsageShaderDeviceAddress specifies that the Buffer can be used to retrieve a
	// Buffer device address via Device.GetBufferDeviceAddress and use that address to
	// access the Buffer object's memory from a shader
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html
	BufferUsageShaderDeviceAddress core1_0.BufferUsageFlags = C.VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT_KHR

	// MemoryAllocateDeviceAddress specifies that the memory can be attached to a Buffer object
	// created with BufferUsageShaderDeviceAddress set in Usage, and that the DeviceMemory object
	// can be used to retrieve an opaque address via Device.GetDeviceMemoryOpaqueCaptureAddress
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagBits.html
	MemoryAllocateDeviceAddress core1_1.MemoryAllocateFlags = C.VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT_KHR
	// MemoryAllocateDeviceAddressCaptureReplay specifies that the memory's address can be saved
	// and reused on a subsequent run (e.g. for trace capture and replay)
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagBits.html
	MemoryAllocateDeviceAddressCaptureReplay core1_1.MemoryAllocateFlags = C.VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR

	// VkErrorInvalidOpaqueCaptureAddress indicates a Buffer creation or memory allocation failed
	// because the requested address is not available
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VkErrorInvalidOpaqueCaptureAddress common.VkResult = C.VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR
)

func init() {
	BufferCreateDeviceAddressCaptureReplay.Register("Device Address (Capture/Replay)")

	BufferUsageShaderDeviceAddress.Register("Shader Device Address")

	MemoryAllocateDeviceAddress.Register("Device Address")
	MemoryAllocateDeviceAddressCaptureReplay.Register("Device Address (Capture/Replay)")

	VkErrorInvalidOpaqueCaptureAddress.Register("invalid opaque capture address")
}
