package khr_multiview_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/v3/vulkan"

type VkPhysicalDeviceMultiviewFeaturesKHR C.VkPhysicalDeviceMultiviewFeaturesKHR
type VkPhysicalDeviceMultiviewPropertiesKHR C.VkPhysicalDeviceMultiviewPropertiesKHR
type VkRenderPassMultiviewCreateInfoKHR C.VkRenderPassMultiviewCreateInfoKHR
