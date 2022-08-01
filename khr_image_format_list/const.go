package khr_image_format_list

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	// ExtensionName is "VK_KHR_image_format_list"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_image_format_list.html
	ExtensionName string = C.VK_KHR_IMAGE_FORMAT_LIST_EXTENSION_NAME
)
