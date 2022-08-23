package khr_8bit_storage

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/v2/vulkan"

const (
	// ExtensionName is "VK_KHR_8bit_storage"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_8bit_storage.html
	ExtensionName string = C.VK_KHR_8BIT_STORAGE_EXTENSION_NAME
)
