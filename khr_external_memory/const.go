package khr_external_memory

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/v2/common"

const (
	// ExtensionName is "VK_KHR_external_memory"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory.html
	ExtensionName string = C.VK_KHR_EXTERNAL_MEMORY_EXTENSION_NAME

	// QueueFamilyExternal represents any Queue external to the resource's current Vulkan instance,
	// as long as the Queue uses the same underlying Device group or PhysicalDevice
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QUEUE_FAMILY_EXTERNAL_KHR.html
	QueueFamilyExternal int = C.VK_QUEUE_FAMILY_EXTERNAL_KHR

	// VkErrorInvalidExternalHandle indicates an external handle is not a valid handle
	// of the specified type
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VkErrorInvalidExternalHandle common.VkResult = C.VK_ERROR_INVALID_EXTERNAL_HANDLE_KHR
)

func init() {
	VkErrorInvalidExternalHandle.Register("invalid external handle")
}
