package khr_shader_integer_dot_product

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

const (
	// ExtensionName is "VK_KHR_shader_integer_dot_product"
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VK_KHR_synchronization2.html
	ExtensionName string = C.VK_KHR_SHADER_INTEGER_DOT_PRODUCT_EXTENSION_NAME
)
