package khr_shader_draw_parameters

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/v2/vulkan"

const (
	// ExtensionName is "VK_KHR_shader_draw_parameters"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_draw_parameters.html
	ExtensionName string = C.VK_KHR_SHADER_DRAW_PARAMETERS_EXTENSION_NAME
)
