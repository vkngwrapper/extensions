package khr_portability_subset

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
#include "../vulkan/vulkan_beta.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/v3/vulkan"

const (
	// ExtensionName is "VK_KHR_portability_subset"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_portability_subset.html
	ExtensionName string = C.VK_KHR_PORTABILITY_SUBSET_EXTENSION_NAME
)
