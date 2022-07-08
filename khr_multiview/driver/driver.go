package khr_multiview_driver

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

type VkPhysicalDeviceMultiviewFeaturesKHR C.VkPhysicalDeviceMultiviewFeaturesKHR
type VkPhysicalDeviceMultiviewPropertiesKHR C.VkPhysicalDeviceMultiviewPropertiesKHR
type VkRenderPassMultiviewCreateInfoKHR C.VkRenderPassMultiviewCreateInfoKHR
