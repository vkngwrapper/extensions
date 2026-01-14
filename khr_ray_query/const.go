package khr_ray_query

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

const (
	// ExtensionName is "VK_KHR_ray_query"
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VK_KHR_ray_query.html
	ExtensionName string = C.VK_KHR_RAY_QUERY_EXTENSION_NAME
)
