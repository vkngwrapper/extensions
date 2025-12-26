package khr_external_memory_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/v3/vulkan"

type VkExternalMemoryBufferCreateInfoKHR C.VkExternalMemoryBufferCreateInfoKHR
type VkExternalMemoryImageCreateInfoKHR C.VkExternalMemoryImageCreateInfoKHR
type VkExportMemoryAllocateInfoKHR C.VkExportMemoryAllocateInfoKHR
