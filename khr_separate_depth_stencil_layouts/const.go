package khr_separate_depth_stencil_layouts

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/vkngwrapper/core/core1_0"
import _ "github.com/vkngwrapper/extensions/vulkan"

const (
	// ExtensionName is "VK_KHR_separate_depth_stencil_layouts"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_separate_depth_stencil_layouts.html
	ExtensionName string = C.VK_KHR_SEPARATE_DEPTH_STENCIL_LAYOUTS_EXTENSION_NAME

	// ImageLayoutDepthAttachmentOptimal specifies a layout for the depth aspect of a depth/stencil
	// format Image allowing read and write access as a depth attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutDepthAttachmentOptimal core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL_KHR
	// ImageLayoutDepthReadOnlyOptimal specifies a layout for the depth aspect of a depth/stencil
	// format Image allowing read-only access as a depth attachment or in shaders as a sampled Image,
	// combined Image/Sampler, or input attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutDepthReadOnlyOptimal core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL_KHR
	// ImageLayoutStencilAttachmentOptimal specifies a layout for the stencil aspect of a
	// depth/stencil format Image allowing read and write access as a stencil attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutStencilAttachmentOptimal core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL_KHR
	// ImageLayoutStencilReadOnlyOptimal specifies a layout for the stencil aspect of a depth/stencil
	// format Image allowing read-only access as a stencil attachment or in shaders as a sampled
	// Image, combined Image/Sampler, or input attachment
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutStencilReadOnlyOptimal core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL_KHR
)

func init() {
	ImageLayoutDepthAttachmentOptimal.Register("Depth Attachment Optimal")
	ImageLayoutDepthReadOnlyOptimal.Register("Depth Read-Only Optimal")
	ImageLayoutStencilAttachmentOptimal.Register("Stencil Attachment Optimal")
	ImageLayoutStencilReadOnlyOptimal.Register("Stencil Read-Only Optimal")
}
