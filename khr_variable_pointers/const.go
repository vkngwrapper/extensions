package khr_variable_pointers

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	// ExtensionName is "VK_KHR_variable_pointers"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_variable_pointers.html
	ExtensionName string = C.VK_KHR_VARIABLE_POINTERS_EXTENSION_NAME
)
