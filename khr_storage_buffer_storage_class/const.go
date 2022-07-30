package khr_storage_buffer_storage_class

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	// ExtensionName is "VK_KHR_storage_buffer_storage_class"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_storage_buffer_storage_class.html
	ExtensionName string = C.VK_KHR_STORAGE_BUFFER_STORAGE_CLASS_EXTENSION_NAME
)
