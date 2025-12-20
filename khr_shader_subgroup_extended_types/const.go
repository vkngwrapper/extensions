package khr_shader_subgroup_extended_types

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/v3/vulkan"

const (
	// ExtensionName is "VK_KHR_shader_subgroup_extended_types"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_subgroup_extended_types.html
	ExtensionName string = C.VK_KHR_SHADER_SUBGROUP_EXTENDED_TYPES_EXTENSION_NAME
)
