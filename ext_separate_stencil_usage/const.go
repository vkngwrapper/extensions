package ext_separate_stencil_usage

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	// ExtensionName is "ext_separate_stencil_usage"
	ExtensionName string = C.VK_EXT_SEPARATE_STENCIL_USAGE_EXTENSION_NAME
)
