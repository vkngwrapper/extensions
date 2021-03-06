package khr_multiview

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/core1_0"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	ExtensionName string = C.VK_KHR_MULTIVIEW_EXTENSION_NAME

	DependencyViewLocal core1_0.DependencyFlags = C.VK_DEPENDENCY_VIEW_LOCAL_BIT_KHR
)

func init() {
	DependencyViewLocal.Register("View Local")
}
