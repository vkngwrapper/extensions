package khr_imageless_framebuffer

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/core1_0"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	ExtensionName string = C.VK_KHR_IMAGELESS_FRAMEBUFFER_EXTENSION_NAME

	FramebufferCreateImageless core1_0.FramebufferCreateFlags = C.VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT_KHR
)

func init() {
	FramebufferCreateImageless.Register("Imageless")
}
