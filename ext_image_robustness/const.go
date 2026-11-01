package ext_image_robustness

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"

const (
	// ExtensionName is "VK_EXT_image_robustness"
	//
	// https://docs.vulkan.org/refpages/latest/refpages/source/VK_EXT_image_robustness.html
	ExtensionName string = C.VK_EXT_IMAGE_ROBUSTNESS_EXTENSION_NAME
)
