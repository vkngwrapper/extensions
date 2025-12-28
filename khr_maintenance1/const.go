package khr_maintenance1

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/vkngwrapper/core/v3/common"
	"github.com/vkngwrapper/core/v3/core1_0"
	"github.com/vkngwrapper/core/v3/core1_1"
)

type CommandPoolTrimFlags = core1_1.CommandPoolTrimFlags

////

const (
	// ExtensionName is "VK_KHR_maintenance1"
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance1.html
	ExtensionName string = C.VK_KHR_MAINTENANCE1_EXTENSION_NAME

	// FormatFeatureTransferDst specifies that an Image can be used as a destination Image for copy
	// commands and clear commands
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureTransferDst core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_TRANSFER_DST_BIT_KHR
	// FormatFeatureTransferSrc specifies that an Image can be used as a source Image for copy commands
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html
	FormatFeatureTransferSrc core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_TRANSFER_SRC_BIT_KHR

	// ImageCreate2DArrayCompatible specifies that the Image can be used to create an ImageView of
	// type core1_0.ImageViewType2D or core1_0.ImageViewType2DArray
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html
	ImageCreate2DArrayCompatible core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT_KHR

	// VkErrorOutOfPoolMemory indicates a pool memory allocation has failed
	//
	// https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html
	VkErrorOutOfPoolMemory common.VkResult = C.VK_ERROR_OUT_OF_POOL_MEMORY_KHR
)

func init() {
	FormatFeatureTransferDst.Register("Transfer Destination")
	FormatFeatureTransferSrc.Register("Transfer Source")

	ImageCreate2DArrayCompatible.Register("2D Array Compatible")

	VkErrorOutOfPoolMemory.Register("out of pool memory")
}
