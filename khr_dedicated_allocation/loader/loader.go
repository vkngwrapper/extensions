package khr_dedicated_allocation_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/v3/vulkan"

type VkMemoryDedicatedAllocateInfoKHR C.VkMemoryDedicatedAllocateInfoKHR
type VkMemoryDedicatedRequirementsKHR C.VkMemoryDedicatedRequirementsKHR
