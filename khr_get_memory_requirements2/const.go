package khr_get_memory_requirements2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

const (
	// ExtensionName is "VK_KHR_get_memory_requirements2"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_memory_requirements2.html
	ExtensionName string = C.VK_KHR_GET_MEMORY_REQUIREMENTS_2_EXTENSION_NAME
)
