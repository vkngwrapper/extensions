package khr_maintenance4

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/v3/core1_0"

const (
	// ExtensionName is "VK_KHR_maintenance4"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance4.html
	ExtensionName string = C.VK_KHR_MAINTENANCE_4_EXTENSION_NAME

	// ImageAspectNone specifies no image aspect, or the image aspect is not applicable
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html
	ImageAspectNone core1_0.ImageAspectFlags = C.VK_IMAGE_ASPECT_NONE_KHR
)

func init() {
	ImageAspectNone.Register("None")
}
