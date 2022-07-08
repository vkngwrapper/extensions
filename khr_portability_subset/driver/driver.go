package khr_portability_subset_driver

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
#include "../../vulkan/vulkan_beta.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

type VkPhysicalDevicePortabilitySubsetFeaturesKHR C.VkPhysicalDevicePortabilitySubsetFeaturesKHR
type VkPhysicalDevicePortabilitySubsetPropertiesKHR C.VkPhysicalDevicePortabilitySubsetPropertiesKHR
