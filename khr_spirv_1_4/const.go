package khr_spirv_1_4

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/v3/vulkan"

const (
	// ExtensionName is "VK_KHR_spirv_1_4"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_spirv_1_4.html
	ExtensionName string = C.VK_KHR_SPIRV_1_4_EXTENSION_NAME
)
