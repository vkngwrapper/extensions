package khr_separate_depth_stencil_layouts_driver

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/vulkan"

type VkAttachmentDescriptionStencilLayoutKHR C.VkAttachmentDescriptionStencilLayoutKHR
type VkAttachmentReferenceStencilLayoutKHR C.VkAttachmentReferenceStencilLayoutKHR
type VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR C.VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR
