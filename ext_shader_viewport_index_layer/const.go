package ext_shader_viewport_index_layer

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/v2/vulkan"

const (
	// ExtensionName is "VK_EXT_shader_viewport_index_layer"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_viewport_index_layer.html
	ExtensionName string = C.VK_EXT_SHADER_VIEWPORT_INDEX_LAYER_EXTENSION_NAME
)
