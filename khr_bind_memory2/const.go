package khr_bind_memory2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/v2/core1_0"

const (
	// ExtensionName is "VK_KHR_bind_memory2"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_bind_memory2.html
	ExtensionName string = C.VK_KHR_BIND_MEMORY_2_EXTENSION_NAME

	// ImageCreateAlias specifies that two Image objects created with the same creation parameters
	// and aliased to the same memory can interpret the contents of the memory consistently each
	// other
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html
	ImageCreateAlias core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_ALIAS_BIT_KHR
)

func init() {
	ImageCreateAlias.Register("Alias")
}
