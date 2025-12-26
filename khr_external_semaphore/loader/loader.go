package khr_external_semaphore_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/v3/vulkan"

type VkExportSemaphoreCreateInfoKHR C.VkExportSemaphoreCreateInfoKHR
