package amd_device_coherent_memory

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/v3/core1_0"

const (
	// ExtensionName is "VK_AMD_device_coherent_memory"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_device_coherent_memory.html
	ExtensionName string = C.VK_AMD_DEVICE_COHERENT_MEMORY_EXTENSION_NAME

	// MemoryPropertyDeviceCoherentAMD specifies that device accesses to allocations of this
	// memory type are automatically made available and visible.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryPropertyFlagBits.html
	MemoryPropertyDeviceCoherentAMD core1_0.MemoryPropertyFlags = C.VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD
	// MemoryPropertyDeviceUncachedAMD specifies that memory allocated with this type is not
	// cached on the device. Uncached device memory is always device coherent.
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryPropertyFlagBits.html
	MemoryPropertyDeviceUncachedAMD core1_0.MemoryPropertyFlags = C.VK_MEMORY_PROPERTY_DEVICE_UNCACHED_BIT_AMD
)

func init() {
	MemoryPropertyDeviceUncachedAMD.Register("Device Uncached (AMD)")
	MemoryPropertyDeviceCoherentAMD.Register("Device Coherent (AMD)")
}
