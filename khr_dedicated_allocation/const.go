package khr_dedicated_allocation

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/v2/vulkan"

const (
	// ExtensionName is "VK_KHR_dedicated_allocation"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dedicated_allocation.html
	ExtensionName string = C.VK_KHR_DEDICATED_ALLOCATION_EXTENSION_NAME
)
