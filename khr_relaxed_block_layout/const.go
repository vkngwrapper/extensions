package khr_relaxed_block_layout

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/v2/vulkan"

const (
	// ExtensionName is "VK_KHR_relaxed_block_layout"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_relaxed_block_layout.html
	ExtensionName string = C.VK_KHR_RELAXED_BLOCK_LAYOUT_EXTENSION_NAME
)
