package khr_dedicated_allocation

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	// ExtensionName is "VK_KHR_dedicated_allocation"
	ExtensionName string = C.VK_KHR_DEDICATED_ALLOCATION_EXTENSION_NAME
)
