package khr_imageless_framebuffer

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/core1_0"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	// ExtensionName is "VK_KHR_imageless_framebuffer"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_imageless_framebuffer.html
	ExtensionName string = C.VK_KHR_IMAGELESS_FRAMEBUFFER_EXTENSION_NAME

	// FramebufferCreateImageless specifies that ImageView objects are not specified, and only
	// attachment compatibility information will be provided via a FramebufferAttachmentImageInfo
	// structure
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateFlagBits.html
	FramebufferCreateImageless core1_0.FramebufferCreateFlags = C.VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT_KHR
)

func init() {
	FramebufferCreateImageless.Register("Imageless")
}
