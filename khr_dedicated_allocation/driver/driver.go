package khr_dedicated_allocation_driver

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/v2/vulkan"

type VkMemoryDedicatedAllocateInfoKHR C.VkMemoryDedicatedAllocateInfoKHR
type VkMemoryDedicatedRequirementsKHR C.VkMemoryDedicatedRequirementsKHR
