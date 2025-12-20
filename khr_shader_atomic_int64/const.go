package khr_shader_atomic_int64

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/v3/vulkan"

const (
	// ExtensionName is "VK_KHR_shader_atomic_int64"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_atomic_int64.html
	ExtensionName string = C.VK_KHR_SHADER_ATOMIC_INT64_EXTENSION_NAME
)
