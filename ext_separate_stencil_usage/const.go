package ext_separate_stencil_usage

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	// ExtensionName is "VK_EXT_separate_stencil_usage"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_separate_stencil_usage.html
	ExtensionName string = C.VK_EXT_SEPARATE_STENCIL_USAGE_EXTENSION_NAME
)
