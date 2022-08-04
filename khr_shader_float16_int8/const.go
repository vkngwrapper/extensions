package khr_shader_float16_int8

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	// ExtensionName is "VK_KHR_shader_float16_int8"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_float16_int8.html
	ExtensionName string = C.VK_KHR_SHADER_FLOAT16_INT8_EXTENSION_NAME
)
