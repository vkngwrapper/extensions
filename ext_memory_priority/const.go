package ext_memory_priority

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

const (
	// ExtensionName is "VK_EXT_memory_priority"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_memory_priority.html
	ExtensionName string = C.VK_EXT_MEMORY_PRIORITY_EXTENSION_NAME
)
