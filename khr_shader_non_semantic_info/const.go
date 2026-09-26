package khr_shader_non_semantic_info

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/v3/vulkan"

const (
	// ExtensionName is "VK_KHR_shader_non_semantic_info"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_viewport_index_layer.html
	ExtensionName string = C.VK_KHR_SHADER_NON_SEMANTIC_INFO_EXTENSION_NAME
)
