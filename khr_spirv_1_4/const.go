package khr_spirv_1_4

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	ExtensionName string = C.VK_KHR_SPIRV_1_4_EXTENSION_NAME
)
