package khr_shader_terminate_invocation

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

const (
	// ExtensionName is "VK_KHR_shader_terminate_invocation"
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VK_KHR_shader_terminate_invocation.html
	ExtensionName string = C.VK_KHR_SHADER_TERMINATE_INVOCATION_EXTENSION_NAME
)
