package ext_scalar_block_layout

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	ExtensionName string = C.VK_EXT_SCALAR_BLOCK_LAYOUT_EXTENSION_NAME
)
