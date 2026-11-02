package ext_shader_demote_to_helper_invocation

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

const (
	// ExtensionName is "VK_EXT_shader_demote_to_helper_invocation"
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VK_EXT_shader_demote_to_helper_invocation.html
	ExtensionName string = C.VK_EXT_SHADER_DEMOTE_TO_HELPER_INVOCATION_EXTENSION_NAME
)
