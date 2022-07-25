package khr_16bit_storage

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	// ExtensionName is "khr_16bit_storage"
	ExtensionName string = C.VK_KHR_16BIT_STORAGE_EXTENSION_NAME
)
