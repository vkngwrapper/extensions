package khr_imageless_framebuffer_loader

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"
*/
import "C"
import _ "github.com/vkngwrapper/extensions/v3/vulkan"

type VkFramebufferAttachmentImageInfoKHR C.VkFramebufferAttachmentImageInfoKHR
type VkFramebufferAttachmentsCreateInfoKHR C.VkFramebufferAttachmentsCreateInfoKHR
type VkPhysicalDeviceImagelessFramebufferFeaturesKHR C.VkPhysicalDeviceImagelessFramebufferFeaturesKHR
type VkRenderPassAttachmentBeginInfoKHR C.VkRenderPassAttachmentBeginInfoKHR
