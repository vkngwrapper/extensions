package khr_multiview

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/core1_0"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	// ExtensionName is "VK_KHR_multiview"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_multiview.html
	ExtensionName string = C.VK_KHR_MULTIVIEW_EXTENSION_NAME

	// DependencyViewLocal specifies that a subpass has more than one view
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlagBits.html
	DependencyViewLocal core1_0.DependencyFlags = C.VK_DEPENDENCY_VIEW_LOCAL_BIT_KHR
)

func init() {
	DependencyViewLocal.Register("View Local")
}
