package khr_8bit_storage

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	// ExtensionName is "VK_KHR_8bit_storage"
	ExtensionName string = C.VK_KHR_8BIT_STORAGE_EXTENSION_NAME
)
