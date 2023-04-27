package ext_memory_budget

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

const (
	// ExtensionName is "VK_EXT_memory_budget"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_memory_budget.html
	ExtensionName string = C.VK_EXT_MEMORY_BUDGET_EXTENSION_NAME
)
