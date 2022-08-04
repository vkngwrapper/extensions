package khr_draw_indirect_count

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

const (
	// ExtensionName is "VK_KHR_draw_indirect_count"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_draw_indirect_count.html
	ExtensionName string = C.VK_KHR_DRAW_INDIRECT_COUNT_EXTENSION_NAME
)
