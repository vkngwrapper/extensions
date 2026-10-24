package khr_zero_initialize_workgroup_memory

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

const (
	// ExtensionName is "VK_KHR_zero_initialize_workgroup_memory"
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VK_KHR_zero_initialize_workgroup_memory.html
	ExtensionName string = C.VK_KHR_ZERO_INITIALIZE_WORKGROUP_MEMORY_EXTENSION_NAME
)
