package ext_shader_viewport_index_layer

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	// ExtensionName is "ext_shader_viewport_index_layer"
	ExtensionName string = C.VK_EXT_SHADER_VIEWPORT_INDEX_LAYER_EXTENSION_NAME
)
