package khr_depth_stencil_resolve_driver

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/v2/vulkan"

type VkPhysicalDeviceDepthStencilResolvePropertiesKHR C.VkPhysicalDeviceDepthStencilResolvePropertiesKHR
type VkSubpassDescriptionDepthStencilResolveKHR C.VkSubpassDescriptionDepthStencilResolveKHR
type VkResolveModeFlagsKHR C.VkResolveModeFlagsKHR
