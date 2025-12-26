package khr_separate_depth_stencil_layouts_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/v3/vulkan"

type VkAttachmentDescriptionStencilLayoutKHR C.VkAttachmentDescriptionStencilLayoutKHR
type VkAttachmentReferenceStencilLayoutKHR C.VkAttachmentReferenceStencilLayoutKHR
type VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR C.VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR
