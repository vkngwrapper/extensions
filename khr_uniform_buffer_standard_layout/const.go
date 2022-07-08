package khr_uniform_buffer_standard_layout

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	ExtensionName string = C.VK_KHR_UNIFORM_BUFFER_STANDARD_LAYOUT_EXTENSION_NAME
)
