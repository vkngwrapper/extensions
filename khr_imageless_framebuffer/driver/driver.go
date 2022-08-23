package khr_imageless_framebuffer_driver

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/v2/vulkan"

type VkFramebufferAttachmentImageInfoKHR C.VkFramebufferAttachmentImageInfoKHR
type VkFramebufferAttachmentsCreateInfoKHR C.VkFramebufferAttachmentsCreateInfoKHR
type VkPhysicalDeviceImagelessFramebufferFeaturesKHR C.VkPhysicalDeviceImagelessFramebufferFeaturesKHR
type VkRenderPassAttachmentBeginInfoKHR C.VkRenderPassAttachmentBeginInfoKHR
