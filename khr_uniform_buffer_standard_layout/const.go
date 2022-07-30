package khr_uniform_buffer_standard_layout

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	// ExtensionName is "VK_KHR_uniform_buffer_standard_layout"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_uniform_buffer_standard_layout.html
	ExtensionName string = C.VK_KHR_UNIFORM_BUFFER_STANDARD_LAYOUT_EXTENSION_NAME
)
