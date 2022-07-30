package khr_image_format_list

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	// ExtensionName is "VK_KHR_image_format_list"
	ExtensionName string = C.VK_KHR_IMAGE_FORMAT_LIST_EXTENSION_NAME
)
