package khr_swapchain

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/common"
	"github.com/vkngwrapper/core/core1_0"
)

// SwapchainCreateFlags controls swapchain creation
//
// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateFlagBitsKHR.html
type SwapchainCreateFlags int32

var swapchainCreateFlagsMapping = common.NewFlagStringMapping[SwapchainCreateFlags]()

func (f SwapchainCreateFlags) Register(str string) {
	swapchainCreateFlagsMapping.Register(f, str)
}
func (f SwapchainCreateFlags) String() string {
	return swapchainCreateFlagsMapping.FlagsToString(f)
}

////

const (
	// ExtensionName is "VK_KHR_swapchain"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html
	ExtensionName string = C.VK_KHR_SWAPCHAIN_EXTENSION_NAME

	// ObjectTypeSwapchain specifies a Swapchain handle
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html
	ObjectTypeSwapchain core1_0.ObjectType = C.VK_OBJECT_TYPE_SWAPCHAIN_KHR

	// ImageLayoutPresentSrc must only be used for presenting a presentable Image for display.
	// A Swapchain object's Image must be transitioned to this layout before calling Extension.QueuePresent,
	// and must be transitioned away from this layout after calling Swapchain.AcquireNextImage
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html
	ImageLayoutPresentSrc core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_PRESENT_SRC_KHR

	// VKErrorOutOfDate indicates a Surface has changed in such a way that it is no longer cmopatible
	// with the Swapchain, and further presentation requests using the Swapchain will fail
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKErrorOutOfDate common.VkResult = C.VK_ERROR_OUT_OF_DATE_KHR
	// VKSuboptimal indicates a Swapchain no longer matches the Surface properties exactly, but can
	// still be used to present to the Surface successfully
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VKSuboptimal common.VkResult = C.VK_SUBOPTIMAL_KHR
)

func init() {
	ObjectTypeSwapchain.Register("Swapchain")

	ImageLayoutPresentSrc.Register("Present Src")

	VKErrorOutOfDate.Register("out of date")
	VKSuboptimal.Register("Suboptimal")
}
