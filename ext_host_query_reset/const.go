package ext_host_query_reset

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

const (
	// ExtensionName is "ext_host_query_reset"
	ExtensionName string = C.VK_EXT_HOST_QUERY_RESET_EXTENSION_NAME
)
