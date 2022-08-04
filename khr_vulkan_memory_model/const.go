package khr_vulkan_memory_model

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	// ExtensionName is "VK_KHR_vulkan_memory_model"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_vulkan_memory_model.html
	ExtensionName string = C.VK_KHR_VULKAN_MEMORY_MODEL_EXTENSION_NAME
)
