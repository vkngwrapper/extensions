package ext_host_query_reset

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

const (
	// ExtensionName is "VK_EXT_host_query_reset"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_host_query_reset.html
	ExtensionName string = C.VK_EXT_HOST_QUERY_RESET_EXTENSION_NAME
)
