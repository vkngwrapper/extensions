package khr_variable_pointers

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	ExtensionName string = C.VK_KHR_VARIABLE_POINTERS_EXTENSION_NAME
)
