package khr_device_group_creation

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/core1_0"

const (
	// ExtensionName is "VK_KHR_device_group_creation"
	ExtensionName string = C.VK_KHR_DEVICE_GROUP_CREATION_EXTENSION_NAME

	// MaxDeviceGroupSize is the length of a PhysicalDevice handle array
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_MAX_DEVICE_GROUP_SIZE_KHR.html
	MaxDeviceGroupSize int = C.VK_MAX_DEVICE_GROUP_SIZE_KHR

	// MemoryHeapMultiInstance specifies that ina  logical Device representing more than one
	// PhysicalDevice, there is a per-PhysicalDevice instance of the heap memory
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryHeapFlagBits.html
	MemoryHeapMultiInstance core1_0.MemoryHeapFlags = C.VK_MEMORY_HEAP_MULTI_INSTANCE_BIT_KHR
)

func init() {
	MemoryHeapMultiInstance.Register("Multi-Instance")
}
