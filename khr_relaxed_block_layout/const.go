package khr_relaxed_block_layout

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	ExtensionName string = C.VK_KHR_RELAXED_BLOCK_LAYOUT_EXTENSION_NAME
)
